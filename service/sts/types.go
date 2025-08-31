package awssts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type service struct {
	client *sts.Client
}

type STSService interface {
	GetCallerIdentity(ctx context.Context) (*sts.GetCallerIdentityOutput, error)
}
