package domestic

import (
	"factory-design-pattern/plan"
	"factory-design-pattern/pojo"
)

type domestic struct {
	pojo.Plan
}

func DomesticCost(units int) (cost plan.Plan) {
	cost = &domestic{
		Plan: pojo.Plan{
			Rate:    2.50,
			Ammount: 2.50 * float64(units),
		},
	}
	return
}
