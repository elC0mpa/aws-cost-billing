package flag

import (
	"aws-billing/model"
	"flag"
)

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetParsedFlags() (model.Flags, error) {
	region := flag.String("region", "us-east-1", "AWS region")
	return model.Flags{
		Region: *region,
	}, nil
}
