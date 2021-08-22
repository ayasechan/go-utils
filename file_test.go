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

func TestGetFileStem(t *testing.T) {
	data := map[string]string{
		"./foo":         "foo",
		"../../foo.txt": "foo",
	}
	for k, v := range data {
		r := GetFileStem(k)
		if r != v {
			t.Fatal(r)
		}
	}
}
func TestSetFileStem(t *testing.T) {
	stem := "bar"
	data := map[string]string{
		"./foo":         "bar",
		"../../foo.txt": "../../bar.txt",
	}
	for k, v := range data {
		r := SetFileStem(k, stem)
		if r != v {
			t.Fatal(r)
		}
	}
}
func TestSetFileSuffix(t *testing.T) {
	suffix := ".txt"
	data := map[string]string{
		"./foo":         "foo.txt",
		"../../foo.txt": "../../foo.txt",
		".././/foo.py":  "../foo.txt",
	}
	for k, v := range data {
		r := SetFileSuffix(k, suffix)
		if r != v {
			t.Fatalf("src: %s suffix: %s", k, r)
		}
	}
}
func TestSetFileName(t *testing.T) {
	name := "bar"
	data := map[string]string{
		"./foo":         "bar",
		"../../foo.txt": "../../bar",
		".././/foo.py":  "../bar",
	}
	for k, v := range data {
		r := SetFileName(k, name)
		if r != v {
			t.Fatalf("src: %s dst: %s", k, r)
		}
	}
}
