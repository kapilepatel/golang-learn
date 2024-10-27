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

func TestBulkValidDate(t *testing.T) {
	type testDate struct {
		date   int
		answer bool
	}

	testDates := []testDate{
		testDate{
			date:   -1,
			answer: false,
		},
		testDate{
			date:   0,
			answer: false,
		},
		testDate{
			date:   1,
			answer: true,
		},
		testDate{
			date:   31,
			answer: true,
		},
		testDate{
			date:   32,
			answer: false,
		},
	}

	for _, v := range testDates {
		response, _ := validDate(v.date)
		if response != v.answer {
			t.Error("Expected", v.answer, "got", response)
		}
	}
}

func BenchmarkValidDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validDate(0)

	}
}
