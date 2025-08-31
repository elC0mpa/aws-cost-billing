package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAWSCfg(ctx context.Context, region, profile string) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile(profile))
}
