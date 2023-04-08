package main

import "testing"

func TestShirtFactory_Create(t *testing.T) {
	shirtFactory := &ShirtFactory{}
	shirt := shirtFactory.Create()

	if shirt.GetName() != "Basic shirt" {
		t.Errorf("Shirt name is incorrect, got: %s, want: %s.", shirt.GetName(), "Basic shirt")
	}

	if shirt.GetPrice() != 25.0 {
		t.Errorf("Shirt price is incorrect, got: %f, want: %f.", shirt.GetPrice(), 25.0)
	}
}

func TestPantsFactory_Create(t *testing.T) {
	pantsFactory := &PantsFactory{}
	pants := pantsFactory.Create()

	if pants.GetName() != "Basic pants" {
		t.Errorf("Pants name is incorrect, got: %s, want: %s.", pants.GetName(), "Basic pants")
	}

	if pants.GetPrice() != 35.0 {
		t.Errorf("Pants price is incorrect, got: %f, want: %f.", pants.GetPrice(), 35.0)
	}
}
