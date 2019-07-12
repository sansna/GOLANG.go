package main

import "testing"

func TestS1(t *testing.T) {
	if s1("1234") != 4 {
		t.Error(`s1("1234") != 4`)
	}
	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestZ(t *testing.T) {
	a,b := z("123","345")
	if a!="345" || b!="123" {
		t.Error(`z("123","345") != "345", "123"`)
	}
}
