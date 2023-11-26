package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestEnumerateSubDomains(t *testing.T) {
	os.Setenv("DOMAIN", "foobar.com")
	defer os.Unsetenv("DOMAIN")

	result, err := EnumerateSubDomains()
	if err != nil {
		t.Errorf("EnumerateSubDomains returned an error: %v", err)
	}

	if result == "" {
		t.Errorf("EnumerateSubDomains returned empty result")
	}

	arr1 := regexp.MustCompile("\n").Split(result, -1)
	fmt.Println("Enumerate: ", len(arr1))
}
