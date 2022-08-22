package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	"github.com/NpoolPlatform/build-chain/pkg/config"
	servicename "github.com/NpoolPlatform/build-chain/pkg/service-name"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/version"

	banner "github.com/common-nighthawk/go-figure"
	cli "github.com/urfave/cli/v2"
)

const (
	serviceName = "Build Chain"
	usageText   = "Build Chain Service"
)

var (
	chainType string
	logDir    string
	endpoint  string
)

func main() {
	logDir, _ = config.LookupEnv("ENV_LOG_DIR")
	chainType, _ = config.LookupEnv("ENV_CHAIN_TYPE")
	endpoint, _ = config.LookupEnv("ENV_ENDPOINT")

	commands := cli.Commands{runCmd, crawlCmd}

	description := fmt.Sprintf(
		"%v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName,
		serviceName,
	)
	banner.NewColorFigure(serviceName, "", "green", true).Print()
	vesion, err := version.GetVersion()
	if err != nil {
		log.Fatalf("fail to get version, %v", err)
	}

	app := &cli.App{
		Name:        serviceName,
		Version:     vesion,
		Description: description,
		Usage:       usageText,
		Commands:    commands,
	}
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", servicename.ServiceName, err)
		return
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("fail to run %v: %v", serviceName, err)
	}
}
