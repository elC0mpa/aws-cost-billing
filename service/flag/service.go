package flag

import (
	"aws-billing/model"
	"flag"
)

func NewService() *service {
	return &service{}
}

func (s *service) GetParsedFlags() (model.Flags, error) {
	region := flag.String("region", "us-east-1", "AWS region")
	profile := flag.String("profile", "", "AWS profile configuration")
	trend := flag.Bool("trend", false, "Display a trend report for the last 6 months")

	flag.Parse()

	return model.Flags{
		Region:  *region,
		Profile: *profile,
		Trend:   *trend,
	}, nil
}
