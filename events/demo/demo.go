package main

import (
	"errors"
	"fmt"

	"github.com/teamlint/gox/events"
)

var (
	ErrPost1   = errors.New("post1")
	ErrPost2   = errors.New("post2")
	ErrPostVar = errors.New("postvar")
)

func main() {
	events.EnableWarning = true
	e := events.New()

	e.On("evt", func() {
		p("1111111")
	})
	e.On("post", func() {
		p("post func1 load")
	})
	e.On("post", func() {
		p("post func2 load")
		panic("post func2 error")
	})
	e.On("post", func() error {
		p("post 1 load,return error")
		return ErrPost1
	})
	e.On("post", func(msg string) error {
		if e.Error() == ErrPost1 {
			p("post 1 error->post2 process")
			return nil
		}
		p("post 1 error(nil)->post2 process")
		return ErrPost2
	})
	e.On("post", func(args ...interface{}) error {
		p("variable argument")
		return ErrPostVar
	})
	e.On("evt", func() {
		p("22222222")
		panic("222 error")
	})
	e.On("evt", func() {
		p("33333333")
	})
	e.On("evt", func() {
		p("5555555")
		panic("5555 error")
	})
	p("evt listener count: %v", e.ListenerCount("evt"))
	e.Emit("evt")
	// e.Emit("post")
	p("1 event all errors: %+v", e.Errors())
	p("1 event last error: %+v", e.Error())
	p("1 event post errors: %+v", e.Errors("post"))
	p("1 event last post error: %+v", e.Error("post"))
	e.On("evt", func() {
		p("4444444")
	})
	// e.ClearErrors()
	// e.ClearErrors("evt")
	// e.Emit("post", "gox")
	e.EmitSync("post", "gox")
	e.OffAll("evt")
	p("2 event all errors: %+v", e.Errors())
	p("2 event last error: %+v", e.Error())
	p("2 event post errors: %+v", e.Errors("post"))
	p("2 event last post error: %+v", e.Error("post"))
	e.Clear()
	p("main exit")
}
func p(msg string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(msg+"\n", args...)
	} else {
		fmt.Println(msg)
	}
}
