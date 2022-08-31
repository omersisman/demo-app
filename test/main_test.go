package test

import (
	"strings"
	"testing"
)

func TestDemoApp(t *testing.T) {
	text := "text"

	if strings.ToUpper(text) != "TEXT" {
		t.Error("Demo test failed!")
	}
}
