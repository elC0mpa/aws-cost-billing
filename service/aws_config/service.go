package awsconfig

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAWSCfg(ctx context.Context, region, profile string) (aws.Config, error) {
	fmt.Println("Received profile:", profile)
	return config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile(profile))
}
