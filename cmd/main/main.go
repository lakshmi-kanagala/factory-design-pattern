package main

import (
	"factory-design-pattern/factory"
	"factory-design-pattern/plan"
	"fmt"
)

func main() {
	domesticDetails, _ := factory.GetPlan("domestic", 800)
	//musket, _ := getGun("musket")

	printDetails(domesticDetails)
	//printDetails(musket)
}

func printDetails(plan plan.Plan) {
	fmt.Printf("Ammount: %f", plan.GetAmmount())
	fmt.Println()
	fmt.Printf("Rate: %f", plan.GetRate())
	fmt.Println()
}
