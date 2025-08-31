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
	profile := flag.String("profile", "", "AWS profile configuration")

	flag.Parse()

	return model.Flags{
		Region:  *region,
		Profile: *profile,
	}, nil
}
