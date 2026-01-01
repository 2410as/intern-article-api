package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 5

	actual := Add(2,3)

	if actual != expectsd {
		t.Errorf("期待値は %d ですが、実際は %d でした", expected, actual) 
	}
}