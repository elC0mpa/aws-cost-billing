package main

import (
	awsconfig "aws-billing/service/aws_config"
	awscostexplorer "aws-billing/service/costexplorer"
	"aws-billing/service/flag"
	awssts "aws-billing/service/sts"
	"aws-billing/utils"
	"context"
)

func main() {
	flagService := flag.NewService()
	flags, err := flagService.GetParsedFlags()
	if err != nil {
		panic(err)
	}

	cfgService := awsconfig.NewService()
	awsCfg, err := cfgService.GetAWSCfg(context.Background(), flags.Region)
	if err != nil {
		panic(err)
	}

	costService := awscostexplorer.NewService(awsCfg)
	stsService := awssts.NewService(awsCfg)

	currentMonthData, err := costService.GetCurrentMonthCostsByService(context.Background())
	if err != nil {
		panic(err)
	}

	lastMonthData, err := costService.GetLastMonthCostsByService(context.Background())
	if err != nil {
		panic(err)
	}

	currentTotalCost, err := costService.GetCurrentMonthTotalCosts(context.Background())
	if err != nil {
		panic(err)
	}

	lastTotalCost, err := costService.GetLastMonthTotalCosts(context.Background())
	if err != nil {
		panic(err)
	}

	stsResult, err := stsService.GetCallerIdentity(context.Background())
	if err != nil {
		panic(err)
	}

	utils.DrawTable(*stsResult.Account, *lastTotalCost, *currentTotalCost, lastMonthData, currentMonthData, "UnblendedCost")
}
