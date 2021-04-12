package action

import "testing"

func TestA(t *testing.T) {
	if A() != "A"{
		t.Fatal("failed!")
	}
}

func TestB(t *testing.T) {
	if B() != "B"{
		t.Fatal()
	}
}