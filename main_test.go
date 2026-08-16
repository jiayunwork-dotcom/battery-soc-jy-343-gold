package main

import "testing"

func TestParseCurrentLog_RejectsGarbage(t *testing.T) {
	_, err := parseCurrentLog([]byte("dt,current\n1.0,10\n1.0,abc\n"))
	if err == nil {
		t.Fatal("expected error for non-numeric current on a data row")
	}
}
