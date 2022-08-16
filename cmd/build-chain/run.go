package main

import (
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
	Before: func(ctx *cli.Context) error {
		// TODO: elegent set or get env
		config.SetENV(&config.ENVInfo{
			LogDir:    logDir,
			ChainType: chainType,
			Endpoint:  endpoint,
		})
		logger.Init(logger.DebugLevel, logDir)
		return nil
	},
	Flags: []cli.Flag{
		// proxy address
		&cli.StringFlag{
			Name:        "chainType",
			Aliases:     []string{"c"},
			Usage:       "chain type",
			EnvVars:     []string{"ENV_CHAIN_TYPE"},
			Required:    true,
			Value:       "",
			Destination: &chainType,
		},
		&cli.StringFlag{
			Name:        "endpoint",
			Aliases:     []string{"e"},
			Usage:       "chain`s endpoint",
			EnvVars:     []string{"ENV_ENDPOINT"},
			Required:    false,
			Value:       "",
			Destination: &endpoint,
		},
	},
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		go coins.Run()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err := api.RegisterGateway(mux, ":12315", opts)
		if err != nil {
			log.Fatalf("Fail to register gRPC gateway service endpoint: %v", err)
		}

		go func() {
			http.Handle("/v1/", mux)
			pages, err := fs.Sub(res.ResPages, "pages")
			if err != nil {
				log.Fatalf("failed to load pages: %v", err)
			}

			http.Handle("/", http.FileServer(http.FS(pages)))
			err = http.ListenAndServe(":12317", nil)
			if err != nil {
				log.Fatalf("failed to setup HTTP pages: %v", err)
			}
		}()

		lis, err := net.Listen("tcp", ":12315")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		server := grpc.NewServer()
		api.Register(server)
		reflection.Register(server)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		return nil
	},
}
