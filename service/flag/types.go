package flag

import "aws-billing/model"

type service struct {
}

type FlagService interface {
	GetParsedFlags() (model.Flags, error)
}
