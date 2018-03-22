package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	kcl "github.com/sul-dlss-labs/go-kcl"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	s := kcl.NewLocalStore()
	config := kcl.Config{
		Limit:        1000,
		Interval:     time.Millisecond * 1000,
		IteratorType: kcl.IteratorTypeLatest,
	}
	k, err := kcl.NewStream(sess, os.Getenv("AWS_KINESIS_ENDPOINT"), os.Getenv("AWS_KINESIS_STREAM"), s, config)
	if err != nil {
		panic(err)
	}

	err = k.Listen(handler)
	if err != nil {
		panic(err)
	}
}

func handler(records []*kinesis.Record) {
	for _, r := range records {
		fmt.Printf("%s", fmt.Sprintf("%s", r.Data[:]))
	}
}
