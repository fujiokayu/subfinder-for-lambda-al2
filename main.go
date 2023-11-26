package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func EnumerateSubDomains() (string, error) {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		return "", errors.New("environment variable 'DOMAIN' not set")
	}

	subfinderOpts := &runner.Options{
		Threads:            10, // Thread controls the number of threads to use for active enumerations
		Timeout:            30, // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 30, // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		All:                true,
	}

	// disable timestamps in logs / configure logger
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		return "failed to create subfinder runner: ", err
	}

	output := &bytes.Buffer{}

	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		return "failed to enumerate single domain: ", err
	}

	return output.String(), err
}

func HandleRequest(ctx context.Context) {
	_ = godotenv.Load() // If you use .env files for local test

	result, err := EnumerateSubDomains()
	if err != nil {
		log.Fatalf(result, err)
	}

	arr1 := regexp.MustCompile("\n").Split(result, -1)
	fmt.Println(len(arr1) - 1)
	for _, s := range arr1 {
		// invoke some functions for each domain
		fmt.Printf("%s\n", s)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
