package orchestrator

import (
	"aws-billing/model"
	awscostexplorer "aws-billing/service/costexplorer"
	awssts "aws-billing/service/sts"
)

type service struct {
	stsService awssts.STSService
	costService awscostexplorer.CostService
}

type OrchestratorService interface {
	Orchestrate(model.Flags) (error)
}
