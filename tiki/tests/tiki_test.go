package tiki

import (
	"testing"
	"hthngoc/tikitiki"
	"strings"
	"fmt"
)

func Test_access(t *testing.T) {
	want := "Hello, world."
	if got := tiki.Access(); got != want {
		t.Errorf("accessTiki() = %q, want %q", got, want)
	}
}


func Test_get(t *testing.T) {
	want := "yamaha"
	got := tiki.Get(1)
	fmt.Println(strings.ContainsAny(got.Url, want))
	if strings.Contains(got.Url, want) == false {
		t.Errorf("GET(yamaha) = %v, want %q", got, want)
	}
}

