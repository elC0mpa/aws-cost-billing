package awssts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type Service struct {
	client *sts.Client
}

type ConfigService interface {
	GetCallerIdentity(ctx context.Context) (*sts.GetCallerIdentityOutput, error)
}
