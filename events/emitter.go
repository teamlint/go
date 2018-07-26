// Package events 提供简单的事件管理器
package events

import (
	"log"
	"reflect"
	"sync"
)

const (
	// Version 版本号
	Version = "0.1"
	// DefaultMaxListeners 每个事件最大监听器数量
	DefaultMaxListeners = 0
	// EnableWarning 允许警告信息,如果监听器数量达到设定值输出警告信息,默认false
	EnableWarning = false
)

// Listener 事件监听器
type Listener func(...interface{})

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
}

type emitter struct {
	mu           sync.Mutex
	maxListeners int
	events       map[interface{}][]Listener
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
	e.events = make(map[interface{}][]Listener)
	e.maxListeners = DefaultMaxListeners
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

// 默认 Emitter 接口实现
//==============================================================================
// Emit 触发事件
func (e *emitter) Emit(evt interface{}, arguments ...interface{}) Emitter {
	var (
		listeners []Listener
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
		go func(l Listener) {
			defer wg.Done()
			l(arguments...)
		}(fn)
	}
	wg.Wait()
	return e
}

// EmitSync 同步触发事件
func (e *emitter) EmitSync(evt interface{}, arguments ...interface{}) Emitter {

	var (
		listeners []Listener
		ok        bool
	)
	e.mu.Lock()
	if listeners, ok = e.events[evt]; !ok {
		e.mu.Unlock()
		return e
	}
	e.mu.Unlock()

	for _, l := range listeners {
		l(arguments...)
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

	if e.maxListeners > 0 && len(e.events[evt]) == e.maxListeners {
		if EnableWarning {
			log.Printf("Warning: event `%v` has exceeded the maximum number of listeners of %d.\n", evt, e.maxListeners)
		}
	}

	e.events[evt] = append(e.events[evt], listener)
	return e
}

// Once 绑定一次性事件
func (e *emitter) Once(evt interface{}, listener Listener) Emitter {
	var run Listener

	run = func(arguments ...interface{}) {
		defer e.Off(evt, run)
		listener(arguments...)
	}

	e.On(evt, run)
	return e
}

// Off 移除指定事件监听器
func (e *emitter) Off(evt interface{}, listener Listener) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	newListeners := []Listener{}
	if listeners, ok := e.events[evt]; ok {
		listenerPointer := reflect.ValueOf(listener).Pointer()
		for _, l := range listeners {
			itemPointer := reflect.ValueOf(l).Pointer()
			if itemPointer != listenerPointer {
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

	e.events = map[interface{}][]Listener{}
}

// SetMaxListeners 设置事件最大监听器数量
func (e *emitter) SetMaxListeners(n int) Emitter {
	e.mu.Lock()
	defer e.mu.Unlock()

	if n < 0 {
		if EnableWarning {
			log.Printf("warning: MaxListeners must be positive number, tried to set: %d", n)
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
		return listeners
	}

	return nil
}

func (e *emitter) Len() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.events)
}
