package main

import (
	"errors"
	"fmt"
	"log"

	e "github.com/teamlint/gox/events"
)

var (
	ErrPost1   = errors.New("post1")
	ErrPost2   = errors.New("post2")
	ErrPostVar = errors.New("postvar")
)

func main() {
	// events.EnableWarning = true
	// e := events.New()
	// e := events.Default()

	// e.On("evt", func() {

	// 	p("1111111")
	// })
	log.Printf("emiter=%p\n", e.Default())
	e.On("post", func() {
		p("post_handler func() noarg")
	})
	// e.On("post", func() {
	// 	p("post func2 load")
	// 	// panic("post func2 error")
	// })
	// e.On("post", func() error {
	// 	p("post 1 load,return error")
	// 	return ErrPost1
	// })
	e.Subscribe("post", func(msg string) error {
		log.Printf("emiter=%p\n", e.Default())
		p("post_handler(string) arg=%v", msg)
		return nil
		// if e.Error() == ErrPost1 {
		// 	p("post 1 error->post2 process")
		// 	return nil
		// }
		// p("post 1 error(nil)->post2 process")
		// return ErrPost2
	})
	e.On("post", func(msg string, foo string) error {
		p("post_handler(string,string) arg1=%v,arg2=%v", msg, foo)
		return nil
	})
	e.On("post", func(msg string, args ...interface{}) error {
		p("post_handler variable argument.msg=%v", msg)
		p("post_handler variable argument")
		if len(args) > 0 {
			for k, v := range args {
				p("\targ[%v] = %v", k, v)
			}
		}
		// return ErrPostVar
		return nil
	})
	e.On("post", func(args ...interface{}) error {
		p("post_handler variable argument")
		if len(args) > 0 {
			for k, v := range args {
				p("\targ[%v] = %v", k, v)
			}
		}
		// return ErrPostVar
		return nil
	})
	// e.On("evt", func() {
	// 	p("22222222")
	// 	panic("222 error")
	// })
	// e.On("evt", func() {
	// 	p("33333333")
	// })
	// e.On("evt", func() {
	// 	p("5555555")
	// 	panic("5555 error")
	// })
	p("evt listener count: %v", e.ListenerCount("post"))
	// e.Emit("evt")
	// p("1 event all errors: %+v", e.Errors())
	// p("1 event last error: %+v", e.Error())
	// p("1 event post errors: %+v", e.Errors("post"))
	// p("1 event last post error: %+v", e.Error("post"))
	// e.On("evt", func() {
	// 	p("4444444")
	// })
	// e.ClearErrors()
	// e.ClearErrors("evt")
	e.EmitSync("post")
	e.EmitSync("post", "gox")
	e.Publish("post", "a1", "a2", 333, 4444)
	e.PublishSync("post", "a1a", 2222, 333, 4444)
	// e.OffAll("evt")
	// p("2 event all errors: %+v", e.Errors())
	// p("2 event last error: %+v", e.Error())
	// p("2 event post errors: %+v", e.Errors("post"))
	// p("2 event last post error: %+v", e.Error("post"))
	e.Clear()
	p("main exit")
}
func p(msg string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf("[out] "+msg+"\n", args...)
	} else {
		fmt.Println("[out] " + msg)
	}
}
