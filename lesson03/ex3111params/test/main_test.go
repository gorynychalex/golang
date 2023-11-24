package test

import (
	"testing"
	"example/functions"
)

func TestAdd(t *testing.T){
	
	// Arrange
	x, y := 5, 6
	expected := 11
	
	// Act
	result := functions.Add(x,y)
	
	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
	}


}


func TestMutiply(t *testing.T){
	// Arrange
	x, y := 5, 6
	expected := 30
	
	// Act
	result := functions.Multiply(x,y)
	
	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
	}

}