// Package events 提供简单的事件管理器
package events

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
)

const (
	// Version 版本号
	Version = "0.2"
	// DefaultMaxListeners 每个事件最大监听器数量
	DefaultMaxListeners = 0
)

// EnableWarning 允许警告信息,如果监听器数量达到设定值输出警告信息,默认false
var EnableWarning = false

// ErrNoneFunction 不是有效的函数
var ErrNoneFunction = errors.New("kind of value for listener is not function")

// Listener 事件监听器
type Listener = interface{}

// Emitter 事件管理接口
type Emitter interface {
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
}

type emitter struct {
	mu           sync.Mutex
	maxListeners int
	events       map[interface{}][]reflect.Value
	errors       map[interface{}][]error
}

var (
	_              Emitter = &emitter{}
	defaultEmitter         = New()
)

// 包级方法
//==============================================================================

// New 创建新的事件管理器
func New() Emitter {
	e := new(emitter)
	e.events = make(map[interface{}][]reflect.Value)
	e.maxListeners = DefaultMaxListeners
	e.errors = make(map[interface{}][]error)
	return e
}

// Emit 触发事件
func Emit(evt interface{}, arguments ...interface{}) Emitter {
	return defaultEmitter.Emit(evt, arguments...)
}

// EmitSync 同步触发事件
func EmitSync(evt interface{}, arguments ...interface{}) Emitter {
	return defaultEmitter.EmitSync(evt, arguments...)
}

// Trigger 触发事件
func Trigger(evt interface{}, arguments ...interface{}) Emitter {
	return defaultEmitter.Emit(evt, arguments...)
}

// TriggerSync 同步触发事件
func TriggerSync(evt interface{}, arguments ...interface{}) Emitter {
	return defaultEmitter.EmitSync(evt, arguments...)
}

// On 绑定事件
func On(evt interface{}, listener Listener) Emitter {
	return defaultEmitter.On(evt, listener)
}

// Once 绑定一次性事件
func Once(evt interface{}, listener Listener) Emitter {
	return defaultEmitter.Once(evt, listener)
}

// Off 移除指定事件监听器
func Off(evt interface{}, listener Listener) Emitter {
	return defaultEmitter.Off(evt, listener)
}

// OffAll 移除指定事件所有监听器
func OffAll(evt interface{}) Emitter {
	return defaultEmitter.OffAll(evt)
}

// Clear 清除所有事件及监听器
func Clear() {
	defaultEmitter.Clear()
}

// SetMaxListeners 设置事件监听器最大数量,0 无限制
func SetMaxListeners(n int) Emitter {
	return defaultEmitter.SetMaxListeners(n)
}

// MaxListeners 获取事件监听器最大数量
func MaxListeners() int {
	return defaultEmitter.MaxListeners()
}

// Events 获取所有事件列表
func Events() []interface{} {
	return defaultEmitter.Events()
}

// ListenerCount 获取指定事件监听器数量
func ListenerCount(evt interface{}) int {
	return defaultEmitter.ListenerCount(evt)
}

// Listeners 获取指定事件监听器列表
func Listeners(evt interface{}) []Listener {
	return defaultEmitter.Listeners(evt)
}

// Len 获取注册事件数量
func Len() int {
	return defaultEmitter.Len()
}

// Errors 获取事件管理器错误
func Errors(evt ...interface{}) []error {
	return defaultEmitter.Errors(evt...)
}

// Error 获取事件管理器最后一个错误
func Error(evt ...interface{}) error {
	return defaultEmitter.Error(evt...)
}

// ClearErrors 清除事件管理器错误
func ClearErrors(evt ...interface{}) {
	defaultEmitter.ClearErrors(evt...)
}

// 默认 Emitter 接口实现
//==============================================================================

// Emit 触发事件
func (e *emitter) Emit(evt interface{}, arguments ...interface{}) Emitter {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("[events] emit error=%v\n", err)
		}
	}()

	e.ClearErrors(evt)
	var (
		listeners []reflect.Value
		ok        bool
	)
	e.mu.Lock()
	if listeners, ok = e.events[evt]; !ok {
		e.mu.Unlock()
		return e
	}
	e.mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(len(listeners))

	for _, fn := range listeners {
		go func(l reflect.Value) {
			defer wg.Done()
			e.callFunc(evt, l, arguments...)
		}(fn)
	}
	wg.Wait()
	return e
}

// EmitSync 同步触发事件
func (e *emitter) EmitSync(evt interface{}, arguments ...interface{}) Emitter {
	var (
		listeners []reflect.Value
		ok        bool
	)
	e.mu.Lock()
	if listeners, ok = e.events[evt]; !ok {
		e.mu.Unlock()
		return e
	}
	e.mu.Unlock()

	for _, fn := range listeners {
		e.callFunc(evt, fn, arguments...)
	}

	return e
}

// Trigger 触发事件
func (e *emitter) Trigger(evt interface{}, arguments ...interface{}) Emitter {
	return e.Emit(evt, arguments...)
}

// TriggerSync 同步触发事件
func (e *emitter) TriggerSync(evt interface{}, arguments ...interface{}) Emitter {
	return e.EmitSync(evt, arguments...)
}

// On 绑定事件监听器
func (e *emitter) On(evt interface{}, listener Listener) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	fn := getFunc(listener)
	if e.maxListeners > 0 && len(e.events[evt]) == e.maxListeners {
		if EnableWarning {
			log.Printf("[events] warning: event `%v` has exceeded the maximum number of listeners of %d.\n", evt, e.maxListeners)
		}
	}

	e.events[evt] = append(e.events[evt], fn)
	return e
}

// Once 绑定一次性事件
func (e *emitter) Once(evt interface{}, listener Listener) Emitter {
	fn := getFunc(listener)

	var run Listener
	run = func(arguments ...interface{}) {
		defer e.Off(evt, run)
		e.callFunc(evt, fn, arguments...)
	}

	e.On(evt, run)
	return e
}

// Off 移除指定事件监听器
func (e *emitter) Off(evt interface{}, listener Listener) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	fn := reflect.ValueOf(listener)
	newListeners := []reflect.Value{}
	if listeners, ok := e.events[evt]; ok {
		for _, l := range listeners {
			if l.Pointer() != fn.Pointer() {
				newListeners = append(newListeners, l)
			}
		}

		e.events[evt] = newListeners
	}

	return e
}

// OffAll 移除指定事件所有监听器
func (e *emitter) OffAll(evt interface{}) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.events[evt]; ok {
		delete(e.events, evt)
	}

	return e
}

// Clear 清除所有事件及监听器
func (e *emitter) Clear() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.events = map[interface{}][]reflect.Value{}
}

// SetMaxListeners 设置事件最大监听器数量
func (e *emitter) SetMaxListeners(n int) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	if n < 0 {
		if EnableWarning {
			log.Printf("[events] warning: MaxListeners must be positive number, tried to set: %d", n)
		}
	}
	e.maxListeners = n
	return e
}

// MaxListeners 获取事件最大监听器数量
func (e *emitter) MaxListeners() (max int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.maxListeners
}

// Events 获取事件列表
func (e *emitter) Events() []interface{} {
	events := make([]interface{}, e.Len())
	i := 0
	for k := range e.events {
		events[i] = k
		i++
	}
	return events
}

// ListenerCount 获取事件监听器数量
func (e *emitter) ListenerCount(evt interface{}) (count int) {
	e.mu.Lock()
	if listeners, ok := e.events[evt]; ok {
		count = len(listeners)
	}
	e.mu.Unlock()
	return
}

// Listeners 获取指定事件监听器列表
func (e *emitter) Listeners(evt interface{}) []Listener {
	e.mu.Lock()
	defer e.mu.Unlock()

	if listeners, ok := e.events[evt]; ok {
		var returnListeners []interface{}
		for _, listener := range listeners {
			returnListeners = append(returnListeners, listener.Interface())
		}
		return returnListeners
	}

	return nil
}

// Len 获取事件数量
func (e *emitter) Len() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.events)
}

// Errors 获取事件管理器错误
func (e *emitter) Errors(evts ...interface{}) []error {
	e.mu.Lock()
	defer e.mu.Unlock()

	var errs []error
	evtLen := len(evts)
	if evtLen > 0 {
		for i := 0; i < evtLen; i++ {
			errs = append(errs, e.errors[evts[i]]...)
		}
	} else {
		for k := range e.events {
			if item, ok := e.errors[k]; ok {
				errs = append(errs, item...)
			}
		}
	}

	return errs
}

// Error 获取事件管理器最后一个错误
func (e *emitter) Error(evt ...interface{}) error {
	errs := e.Errors()
	elen := len(errs)
	if elen > 0 {
		return errs[elen-1]
	}
	return nil
}

// ClearErrors 清除事件管理器错误
func (e *emitter) ClearErrors(evts ...interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	evtLen := len(evts)
	if evtLen > 0 {
		for i := 0; i < evtLen; i++ {
			e.errors[evts[i]] = nil
		}
		return
	}
	e.errors = map[interface{}][]error{}
}

// getFunc 获取监听器函数
func getFunc(listener interface{}) reflect.Value {
	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		if EnableWarning {
			log.Printf("[events] error: %v\n", ErrNoneFunction.Error())
		}
		panic(ErrNoneFunction)
	}

	return fn
}

// callFunc 调用函数
func (e *emitter) callFunc(evt interface{}, fn reflect.Value, arguments ...interface{}) {
	defer e.recovery(evt)

	// var values []reflect.Value = make([]reflect.Value, len(arguments))
	// for i, v := range arguments {
	// 	if v == nil {
	// 		values[i] = reflect.New(fn.Type().In(i)).Elem()
	// 	} else {
	// 		values[i] = reflect.ValueOf(arguments[i])
	// 	}
	// }
	var values map[int]reflect.Value = make(map[int]reflect.Value, len(arguments))
	for i, v := range arguments {
		if v == nil {
			values[i] = reflect.New(fn.Type().In(i)).Elem()
		} else {
			values[i] = reflect.ValueOf(arguments[i])
		}
	}

	fnArgNum := fn.Type().NumIn()
	fnOutNum := fn.Type().NumOut()
	if EnableWarning {
		log.Printf("[events] handler=%v, args.count=%v", fn.Type(), fnArgNum)
	}

	// func actual argumnets
	var fnArguments []reflect.Value = make([]reflect.Value, fnArgNum)

	log.Printf("arg.values=%v\n", values)
	// variable args
	if fn.Type().IsVariadic() {
		variableArgs := make([]reflect.Value, len(values))
		for i, v := range values {
			variableArgs[i] = v
		}
		if len(variableArgs) > 0 {
			fnArguments = variableArgs
		}
		for i, fa := range fnArguments {
			if !fa.IsValid() {
				fnArguments[i] = reflect.New(fn.Type().In(i)).Elem()
			}
		}
	} else {
		for i := 0; i < fnArgNum; i++ {
			if v, ok := values[i]; ok {
				fnArguments[i] = v
			} else {
				fnArguments[i] = reflect.New(fn.Type().In(i)).Elem()
			}
		}
	}
	log.Printf("fn.argumnets=%v\n", fnArguments)
	// if fn.Type().IsVariadic() {
	// 	fnArguments = values
	// 	// fn.Call(values)
	// } else {
	// 	fnArguments = values[:fnArgNum]
	// }
	// only support return one error type parameter
	if fnOutNum > 0 {
		out := fn.Call(fnArguments)[0]
		if !out.IsNil() {
			if err, ok := out.Interface().(error); ok {
				e.addError(evt, err)
			}
		}
		return
	}
	fn.Call(fnArguments)
}
func (e *emitter) recovery(evt interface{}) {
	if r := recover(); r != nil {
		if EnableWarning {
			log.Printf("[events] recover '%v' event, %v\n", evt, r)
		}
		err := errors.New(fmt.Sprintf("%v", r))
		e.addError(evt, err)
	}
}
func (e *emitter) addError(evt interface{}, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	errs := e.errors[evt]
	e.errors[evt] = append(errs, err)
}
