package main

import "testing"

func TestRun(t *testing.T) {
	db, err := run()

	if db != nil {

	}

	if err != nil {
		t.Errorf("failed run")
	}
}
