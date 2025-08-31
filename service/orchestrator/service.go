package orchestrator

import (
	"aws-billing/model"
	awscostexplorer "aws-billing/service/costexplorer"
	awssts "aws-billing/service/sts"
	"aws-billing/utils"
	"context"
)

func NewService(stsService awssts.STSService, costService awscostexplorer.CostService) *service {
	return &service{
		stsService:  stsService,
		costService: costService,
	}
}

func (s *service) Orchestrate(flags model.Flags) error {
	var err error

	switch flags.Trend {
	case false:
		err = s.defaultWorkflow()
	}

	return err
}

func (s *service) defaultWorkflow() error {
	currentMonthData, err := s.costService.GetCurrentMonthCostsByService(context.Background())
	if err != nil {
		return err
	}

	lastMonthData, err := s.costService.GetLastMonthCostsByService(context.Background())
	if err != nil {
		return err
	}

	currentTotalCost, err := s.costService.GetCurrentMonthTotalCosts(context.Background())
	if err != nil {
		return err
	}

	lastTotalCost, err := s.costService.GetLastMonthTotalCosts(context.Background())
	if err != nil {
		return err
	}

	stsResult, err := s.stsService.GetCallerIdentity(context.Background())
	if err != nil {
		return err
	}

	utils.StopSpinner()

	utils.DrawTable(*stsResult.Account, *lastTotalCost, *currentTotalCost, lastMonthData, currentMonthData, "UnblendedCost")
	return nil
}

