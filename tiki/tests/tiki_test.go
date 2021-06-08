package tiki

import (
	"testing"
	"hthngoc/tikitiki"
	"strings"
)

func Test_access(t *testing.T) {
    want := "Hello, world."
    if got := tiki.Access(); got != want {
        t.Errorf("accessTiki() = %q, want %q", got, want)
    }
}


func Test_get(t *testing.T) {
    want := "yamaha"
    if got := tiki.Get(1); strings.Contains(got.Url, want) {
        t.Errorf("GET(yamaha) = %v, want %q", got, want)
    }
}

