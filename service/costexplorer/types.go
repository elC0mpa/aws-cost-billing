package awscostexplorer

import (
	"aws-billing/model"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

type service struct {
	client *costexplorer.Client
}

type CostService interface {
	GetCurrentMonthCostsByService(ctx context.Context) (*model.CostInfo, error)
	GetLastMonthCostsByService(ctx context.Context) (*model.CostInfo, error)
	GetMonthCostsByService(ctx context.Context, endDate time.Time) (*model.CostInfo, error)
	GetCurrentMonthTotalCosts(ctx context.Context) (*string, error)
	GetLastMonthTotalCosts(ctx context.Context) (*string, error)
}

