package orchestrator

import (
	"github.com/elC0mpa/aws-billing/model"
	awscostexplorer "github.com/elC0mpa/aws-billing/service/costexplorer"
	awssts "github.com/elC0mpa/aws-billing/service/sts"
)

type service struct {
	stsService awssts.STSService
	costService awscostexplorer.CostService
}

type OrchestratorService interface {
	Orchestrate(model.Flags) (error)
}
