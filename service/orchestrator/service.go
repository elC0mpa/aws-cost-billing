package orchestrator

import (
	"context"

	"github.com/elC0mpa/aws-billing/model"
	awscostexplorer "github.com/elC0mpa/aws-billing/service/costexplorer"
	awssts "github.com/elC0mpa/aws-billing/service/sts"
	"github.com/elC0mpa/aws-billing/utils"
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
	case true:
		err = s.trendWorkflow()
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

func (s *service) trendWorkflow() error {
	costInfo, err := s.costService.GetLastSixMonthsCosts(context.Background())
	if err != nil {
		return err
	}

	stsResult, err := s.stsService.GetCallerIdentity(context.Background())
	if err != nil {
		return err
	}

	utils.StopSpinner()

	utils.DrawTrendChart(*stsResult.Account, costInfo)

	return nil
}
