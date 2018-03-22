package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	kcl "github.com/sul-dlss-labs/go-kcl"
	"github.com/sul-dlss-labs/refritos/aws_session"
	"github.com/sul-dlss-labs/refritos/streaming"
)

var depositStream streaming.KinesisStream
var updateStream streaming.KinesisStream

func main() {
	sess := aws_session.Connect(true)
	depositStream = streaming.KinesisStream{
		Connection: streaming.Connect(sess, os.Getenv("AWS_KINESIS_ENDPOINT")),
		StreamName: "deposit",
	}
	updateStream = streaming.KinesisStream{
		Connection: streaming.Connect(sess, os.Getenv("AWS_KINESIS_ENDPOINT")),
		StreamName: "update",
	}

	s := kcl.NewLocalStore()
	config := kcl.Config{
		Limit:        1000,
		Interval:     time.Millisecond * 1000,
		IteratorType: kcl.IteratorTypeLatest,
	}
	k, err := kcl.NewStream(sess, os.Getenv("AWS_KINESIS_ENDPOINT"), "taco", s, config)
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
		dataString := fmt.Sprintf("%s", r.Data[:])
		fmt.Printf("%s", dataString)
		depositStream.SendMessage(dataString)
		updateStream.SendMessage(dataString)
	}
}
