package main

import (
	"factory-design-pattern/factory"
	"testing"
)

func TestDomestic(t *testing.T) {
	expected := 2000.0000
	actual, err := factory.GetPlan("domestic", 800)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if actual.GetAmmount() != expected {
		t.Errorf("Expected output to be %f, but got %f", expected, actual.GetAmmount())

	}

}

func TestCommercial(t *testing.T) {
	expected := 6400.0000
	actual, err := factory.GetPlan("commercial", 800)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if actual.GetAmmount() != expected {
		t.Errorf("Expected output to be %f, but got %f", expected, actual.GetAmmount())

	}
}

func TestInstituitional(t *testing.T) {
	expected := 4480.0000
	actual, err := factory.GetPlan("instituitional", 800)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if actual.GetAmmount() != expected {
		t.Errorf("Expected output to be %f, but got %f", expected, actual.GetAmmount())

	}

}
