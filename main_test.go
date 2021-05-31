package main

import "testing"

func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "miao~~" {
		t.Errorf("Cat say %s", saying)
	}
}
