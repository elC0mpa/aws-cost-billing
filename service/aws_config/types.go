package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Service struct {
}

type ConfigService interface {
	GetAWSCfg(ctx context.Context, region string) (aws.Config, error)
}
