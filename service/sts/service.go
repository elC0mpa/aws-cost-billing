package awssts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func NewService(awsconfig aws.Config) *Service {
	client := sts.NewFromConfig(awsconfig)
	return &Service{
		client: client,
	}
}

func (s *Service) GetCallerIdentity(ctx context.Context) (*sts.GetCallerIdentityOutput, error) {
	input := &sts.GetCallerIdentityInput{}

	return s.client.GetCallerIdentity(ctx, input)
}
