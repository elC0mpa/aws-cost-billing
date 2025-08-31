package utils

import "github.com/common-nighthawk/go-figure"

func DrawBanner() {
	myFigure := figure.NewColorFigure("AWS Billing", "isometric3", "yellow", false)
	myFigure.Print()
}
