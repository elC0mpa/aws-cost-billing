package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type service struct{}

type ConfigService interface {
	GetAWSCfg(ctx context.Context, region, profile string) (aws.Config, error)
}
