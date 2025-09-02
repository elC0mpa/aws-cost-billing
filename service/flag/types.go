package flag

import "github.com/elC0mpa/aws-billing/model"


type service struct {
}

type FlagService interface {
	GetParsedFlags() (model.Flags, error)
}
