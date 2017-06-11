package foobar

import "testing"

func TestAdd(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.Errorf("exptect 3, got %v\n", s)
	}
}

func TestAdd2(t *testing.T) {
	s := Add(3, 4)
	if s != 7 {
		t.Errorf("exptect 7, got %v\n", s)
	}
}
