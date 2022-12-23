package commercial

import (
	"factory-design-pattern/plan"
	"factory-design-pattern/pojo"
)

type commercial struct {
	pojo.Plan
}

func CommercialCost(units int) (cost plan.Plan) {
	cost = &commercial{
		Plan: pojo.Plan{
			Rate:    8.00,
			Ammount: 8.00 * float64(units),
		},
	}
	return
}
