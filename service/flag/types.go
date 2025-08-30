package flag

import "aws-billing/model"

type Service struct {
}

type FlagService interface {
	GetParsedFlags() (model.Flags, error)
}
