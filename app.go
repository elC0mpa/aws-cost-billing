package main

import (
	awsconfig "aws-billing/service/aws_config"
	awscostexplorer "aws-billing/service/costexplorer"
	"aws-billing/service/flag"
	"aws-billing/service/orchestrator"
	awssts "aws-billing/service/sts"
	"aws-billing/utils"
	"context"
)

func main() {
	utils.DrawBanner()
	utils.StartSpinner()

	flagService := flag.NewService()
	flags, err := flagService.GetParsedFlags()
	if err != nil {
		panic(err)
	}

	cfgService := awsconfig.NewService()
	awsCfg, err := cfgService.GetAWSCfg(context.Background(), flags.Region, flags.Profile)
	if err != nil {
		panic(err)
	}

	costService := awscostexplorer.NewService(awsCfg)
	stsService := awssts.NewService(awsCfg)

	orchestratorService := orchestrator.NewService(stsService, costService)

	err = orchestratorService.Orchestrate(flags)
	if err != nil {
		panic(err)
	}
}
