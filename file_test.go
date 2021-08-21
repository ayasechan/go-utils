package goutils

import (
	"log"
	"testing"
)

func TestAsSafeFileName(t *testing.T) {
	m := map[string]string{
		`\fo* `: "fo",
	}

	for k, v := range m {
		r := AsSafeFileName(k, "")
		if r != v {
			log.Fatal(r)
		}
	}
}
