package model

import "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"

type CostInfo struct {
	types.DateInterval
	CostGroup
}

type CostGroup map[string]struct {
	Amount float64
	Unit   string
}

type ServiceCost struct {
	Name   string
	Amount float64
	Unit   string
}
