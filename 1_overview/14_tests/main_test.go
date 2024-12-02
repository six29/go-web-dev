package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-values", 100.0, 10.0, 10.0, false},
	{"invalid-values", 100.0, 0, 0, true},
	{"expect-five", 100.0, 20.0, 5, false},
	{"expect-fraction", 4.5, 3.0, 1.5, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {

		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one.")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one.", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f and got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(100.0, 10.0)
// 	if err != nil {
// 		t.Error("Got an error that we should not have.")
// 	}
// }

// func TestDiveByZero(t *testing.T) {
// 	_, err := divide(100.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have.")
// 	}
// }
