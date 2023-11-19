package main

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	if err := Run(context.Background()); err != nil {
		t.Error(err)
	}
}
