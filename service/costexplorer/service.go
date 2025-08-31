package awscostexplorer

import (
	"aws-billing/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func NewService(awsconfig aws.Config) *service {
	client := costexplorer.NewFromConfig(awsconfig)
	return &service{
		client: client,
	}
}

func (s *service) GetCurrentMonthCostsByService(ctx context.Context) (*model.CostInfo, error) {
	return s.GetMonthCostsByService(ctx, time.Now())
}

func (s *service) GetLastMonthCostsByService(ctx context.Context) (*model.CostInfo, error) {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return s.GetMonthCostsByService(ctx, oneMonthAgo)
}

func (s *service) GetMonthCostsByService(ctx context.Context, endDate time.Time) (*model.CostInfo, error) {
	firstOfMonth := s.getFirstDayOfMonth(endDate)
	firstOfMonthStr := firstOfMonth.Format("2006-01-02")
	costsAggregation := "UnblendedCost"

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: aws.String(firstOfMonthStr),
			End:   aws.String(endDate.Format("2006-01-02")),
		},
		Metrics: []string{costsAggregation},
		GroupBy: []types.GroupDefinition{
			{
				Key:  aws.String("SERVICE"),
				Type: types.GroupDefinitionTypeDimension,
			},
		},
	}

	output, err := s.client.GetCostAndUsage(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.CostInfo{
		CostGroup:    s.filterGroups(output.ResultsByTime[0].Groups, costsAggregation),
		DateInterval: *output.ResultsByTime[0].TimePeriod,
	}, nil
}

func (s *service) GetCurrentMonthTotalCosts(ctx context.Context) (*string, error) {
	return s.GetMonthTotalCosts(ctx, time.Now())
}

func (s *service) GetLastMonthTotalCosts(ctx context.Context) (*string, error) {
	return s.GetMonthTotalCosts(ctx, time.Now().AddDate(0, -1, 0))
}

func (s *service) GetMonthTotalCosts(ctx context.Context, endDate time.Time) (*string, error) {
	firstOfMonth := s.getFirstDayOfMonth(endDate)
	firstOfMonthStr := firstOfMonth.Format("2006-01-02")
	costsAggregation := "UnblendedCost"

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: aws.String(firstOfMonthStr),
			End:   aws.String(endDate.Format("2006-01-02")),
		},
		Metrics: []string{costsAggregation},
	}

	output, err := s.client.GetCostAndUsage(ctx, input)
	if err != nil {
		return nil, err
	}

	totalInfo := output.ResultsByTime[0].Total[costsAggregation]
	amount, err := strconv.ParseFloat(*totalInfo.Amount, 64)
	if err != nil || amount == 0 {
		panic("Could not parse total amount")
	}

	total := fmt.Sprintf("%.2f %s", amount, *totalInfo.Unit)
	return &total, nil
}

func (s *service) getFirstDayOfMonth(month time.Time) time.Time {
	return time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
}

func (s *service) filterGroups(results []types.Group, costsAggregation string) model.CostGroup {
	filtered := make([]types.Group, 0, len(results))

	for _, g := range results {
		amountStr := ""
		if metric, ok := g.Metrics[costsAggregation]; ok && metric.Amount != nil {
			amountStr = *metric.Amount
		}
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount == 0 {
			continue
		}
		filtered = append(filtered, g)
	}

	costGroups := make(map[string]struct {
		Amount float64
		Unit   string
	})

	for _, v := range filtered {
		amount, _ := strconv.ParseFloat(*v.Metrics[costsAggregation].Amount, 64)
		costGroups[v.Keys[0]] = struct {
			Amount float64
			Unit   string
		}{
			Amount: amount,
			Unit:   *v.Metrics[costsAggregation].Unit,
		}
	}

	return costGroups
}
