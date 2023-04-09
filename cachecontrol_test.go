package cachecontrol_test

import (
	"testing"

	"github.com/0bl/cachecontrol"
)

func TestHeader(t *testing.T) {
	cc := cachecontrol.New().MaxAge(60).NoCache().Public().Header()
	expected := "no-cache, public, max-age=60"
	if cc != expected {
		t.Errorf("expected %q, but got %q", expected, cc)
	}
}
