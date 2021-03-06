package main

import (
	"fmt"

	"github.com/jrallison/go-workers"
)

// Hello is a job that provides encouragement after someone submits "Hello World".
// The job receives the uuid of a submission and submits a comment from rikki-
// to the conversation on exercism.
type Hello struct {
	exercism *Exercism
	comment  []byte
}

// NewHello configures a Hello job to talk to the exercism API.
func NewHello(exercism *Exercism, dir string) (*Hello, error) {
	b, err := read(fmt.Sprintf("%s/hello/hello.md", dir))
	if err != nil {
		return nil, err
	}
	return &Hello{
		exercism: exercism,
		comment:  b,
	}, nil
}

func (hello *Hello) process(msg *workers.Msg) {
	args := msg.Args()
	uuid, err := args.GetIndex(0).String()
	if err != nil {
		lgr.Printf("unable to determine submission uuid - %s\n", err)
		return
	}

	if args.GetIndex(1).MustInt(1) > 1 {
		return
	}

	if err := hello.exercism.SubmitComment(hello.comment, uuid); err != nil {
		lgr.Printf("%s\n", err)
	}
}
