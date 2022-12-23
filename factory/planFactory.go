package factory

import (
	"factory-design-pattern/commercial"
	"factory-design-pattern/domestic"
	"factory-design-pattern/instituitional"
	"factory-design-pattern/plan"
	"fmt"
)

func GetPlan(planName string, units int) (plan.Plan, error) {
	if planName == "commercial" {
		return commercial.CommercialCost(units), nil
	}
	if planName == "domestic" {
		return domestic.DomesticCost(units), nil
	}
	if planName == "instituitional" {
		return instituitional.InstituitionalCost(units), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
