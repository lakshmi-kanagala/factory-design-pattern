package instituitional

import (
	"factory-design-pattern/plan"
	"factory-design-pattern/pojo"
)

type instituitional struct {
	pojo.Plan
}

func InstituitionalCost(units int) (cost plan.Plan) {
	cost = &instituitional{
		Plan: pojo.Plan{
			Rate:    5.60,
			Ammount: 5.60 * float64(units),
		},
	}
	return
}
