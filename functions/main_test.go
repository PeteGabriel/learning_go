package main

import (
	"testing"
)

func TestGetAvg(t *testing.T) {
	elms := []float64 { 12.3, 15.1, 16.7, 12.3, 18.0}
	r := GetAvg(elms)
	if r != 14.879999999999999 {
		t.Error(r)
	}
}
