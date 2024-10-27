package main

import (
	"testing"
)

func TestValidDatePassingValidDate(t *testing.T) {
	var response, _ = validDate(1)
	if response != true {
		t.Error("Expected", true, "Got", response)
	}

	response, _ = validDate(13)
	if response != true {
		t.Error("Expected true; got", response)
	}

	response, _ = validDate(31)
	if response != true {
		t.Error("Expected true for date 31; got", response)
	}

}

func TestValidDatePassingInvalidDate(t *testing.T) {
	var response, _ = validDate(-1)
	if response != false {
		t.Error("Expected false; got", response)
	}

	response, _ = validDate(0)
	if response != false {
		t.Error("Expected false; got", response)
	}

	response, _ = validDate(32)
	if response != false {
		t.Error("Expected false; got", response)
	}

}
