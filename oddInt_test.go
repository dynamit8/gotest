package main

import (
	"testing"
)

func TestOdd7(t *testing.T) {
	result := getOddInt([]int{7})
	expected := 7
	if result != expected {
		t.Errorf("getOddInt([]int{7}) expected %v, but got %v", expected, result)
	}
}

func TestOdd0(t *testing.T) {
	result := getOddInt([]int{0})
	expected := 0
	if result != expected {
		t.Errorf("getOddInt([]int{0}) expected %v, but got %v", expected, result)
	}
}

func TestOdd112(t *testing.T) {
	result := getOddInt([]int{1, 1, 2})
	expected := 2
	if result != expected {
		t.Errorf("getOddInt([]int{1,1,2}) expected %v, but got %v", expected, result)
	}
}

func TestOdd01010(t *testing.T) {
	result := getOddInt([]int{0, 1, 0, 1, 0})
	expected := 0
	if result != expected {
		t.Errorf("getOddInt([]int{0,1,0,1,0}) expected %v, but got %v", expected, result)
	}
}

func TestOddLong(t *testing.T) {
	result := getOddInt([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
	expected := 4
	if result != expected {
		t.Errorf("getOddInt([]int{0,1,0,1,0}) expected %v, but got %v", expected, result)
	}
}
