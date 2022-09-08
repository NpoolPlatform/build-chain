package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"

	"github.com/NpoolPlatform/build-chain/api/v1"
	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/config"
	"github.com/NpoolPlatform/build-chain/pkg/db"
	res "github.com/NpoolPlatform/build-chain/resource"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Build Chain daemon",
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		err := config.Init("./", serviceName)
		if err != nil {
			panic(fmt.Sprintf("fail to init config %v: %v", serviceName, err))
		}

		err = logger.Init(logger.DebugLevel, fmt.Sprintf("%v/%v.log", logDir, serviceName))
		if err != nil {
			panic(fmt.Errorf("fail to init logger: %v", err))
		}

		// TODO: elegent set or get env
		config.SetENV(&config.ENVInfo{
			LogDir:      logDir,
			EthEndpoint: ethEndpoint,
		})
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "eth-endpoint",
			Aliases:     []string{"ee"},
			Usage:       "ethereum`s endpoint",
			EnvVars:     []string{"ENV_ETH_ENDPOINT"},
			Required:    false,
			Value:       "",
			Destination: &ethEndpoint,
		},
		&cli.StringFlag{
			Name:        "log dir",
			Aliases:     []string{"l"},
			Usage:       "log fir",
			EnvVars:     []string{"ENV_LOG_DIR"},
			Required:    false,
			Value:       "./",
			Destination: &logDir,
		},
	},
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		coins.Init()

		go func() {
			runGRPCServer(config.GetInt(config.KeyGRPCPort))
		}()

		runHTTPServer(config.GetInt(config.KeyHTTPPort), config.GetInt(config.KeyGRPCPort))
		return nil
	},
}

func runHTTPServer(httpPort, grpcPort int) {
	grpcEndpoint := fmt.Sprintf(":%v", grpcPort)
	httpEndpoint := fmt.Sprintf(":%v", httpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// register to gatway
	err := api.RegisterGateway(mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service endpoint: %v", err)
	}

	http.Handle("/v1/", mux)
	pages, err := fs.Sub(res.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}

	http.Handle("/", http.FileServer(http.FS(pages)))
	err = http.ListenAndServe(httpEndpoint, nil)
	if err != nil {
		log.Fatalf("failed to setup HTTP pages: %v", err)
	}
}

func runGRPCServer(grpcPort int) {
	endpoint := fmt.Sprintf(":%v", grpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.Register(server)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
