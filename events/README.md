## 简单事件管理器

### 安装
```
go get -u github.com/teamlint/gox/events
```

### 方法
```go
// Emit 触发事件
Emit(evt interface{}, arguments ...interface{}) Emitter
// EmitSync 同步触发事件
EmitSync(evt interface{}, arguments ...interface{}) Emitter
// Trigger Emit 别名
Trigger(evt interface{}, arguments ...interface{}) Emitter
// TriggerSync EmitSync 同步触发事件别名
TriggerSync(evt interface{}, arguments ...interface{}) Emitter
// On 注册事件监听器
On(evt interface{}, listener Listener) Emitter
// Once 注册一次性事件监听器
Once(evt interface{}, listener Listener) Emitter
// Off 移除指定事件监听器
Off(evt interface{}, listener Listener) Emitter
// Off 移除指定事件所有监听器
OffAll(evt interface{}) Emitter
// Clear 清除所有事件及监听器
Clear()
// SetMaxListeners 设置事件监听器最大数量,0 无限制
SetMaxListeners(n int) Emitter
// MaxListeners 获取事件监听器最大数量
MaxListeners() int
// Events 获取所有事件列表
Events() []interface{}
// ListenerCount 获取指定事件监听器数量
ListenerCount(evt interface{}) int
// Listeners 获取指定事件监听器列表
Listeners(evt interface{}) []Listener
// Len 获取注册事件数量
Len() int
// Errors 获取事件管理器错误
Errors(evt ...interface{}) []error
// Error 获取事件管理器最后一个错误
Error(evt ...interface{}) error
// ClearErrors 清除事件管理器错误
ClearErrors(evt ...interface{})
```

### 使用案例
```go
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

    // On 绑定事件
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
    // 同步事件调用
	e.EmitSync("post", "gox")
    // 移除事件绑定
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
```

### 版权
> Copyright (c) 2018 [venjiang](https://github.com/venjiang) 
