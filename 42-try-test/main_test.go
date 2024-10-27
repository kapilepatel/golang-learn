package main

import (
	"testing"
)

func TestValidDatePassingValidDate(t *testing.T) {
	var d, _ = validDate(1)
	if d != true {
		t.Errorf("Expected true; got", d)
	}

	d, _ = validDate(13)
	if d != true {
		t.Errorf("Expected true; got", d)
	}

	d, _ = validDate(31)
	if d != true {
		t.Errorf("Expected true; got", d)
	}

}

func TestValidDatePassingInvalidDate(t *testing.T) {
	var d, _ = validDate(-1)
	if d != false {
		t.Errorf("Expected false; got", d)
	}

	d, _ = validDate(0)
	if d != false {
		t.Errorf("Expected false; got", d)
	}

	d, _ = validDate(32)
	if d != false {
		t.Errorf("Expected false; got", d)
	}

}
