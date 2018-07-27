package events

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/teamlint/gox/events"
)

var testEvents = map[interface{}][]events.Listener{
	"user_created": []events.Listener{
		func(payload ...interface{}) {
			fmt.Printf("A new User just created!\n")
		},
		func(payload ...interface{}) {
			fmt.Printf("A new User just created, *from second event listener\n")
		},
	},
	"user_joined": []events.Listener{
		func(payload ...interface{}) {
			user := payload[0].(string)
			room := payload[1].(string)
			fmt.Printf("%s joined to room: %s\n", user, room)
		},
	},
	"user_left": []events.Listener{func(payload ...interface{}) {
		user := payload[0].(string)
		room := payload[1].(string)
		fmt.Printf("%s left from the room: %s\n", user, room)
	}},
}

func createUser(user string) {
	events.EmitSync("user_created", user)
}

func joinUserTo(user string, room string) {
	events.EmitSync("user_joined", user, room)
}

func leaveFromRoom(user string, room string) {
	events.EmitSync("user_left", user, room)
}

func Example() {
	// regiter our events to the default event emmiter
	for evt, listeners := range testEvents {
		for _, l := range listeners {
			events.On(evt, l)
		}
	}
	// On("user_joined", func(payload string) {
	// 	fmt.Printf("single parameter:%v\n", payload)
	// })

	user := "user1"
	room := "room1"

	createUser(user)
	joinUserTo(user, room)
	leaveFromRoom(user, room)
	events.TriggerSync("evt", "gox")
	// Output:
	// A new User just created!
	// A new User just created, *from second event listener
	// user1 joined to room: room1
	// user1 left from the room: room1
}

func TestEvents(t *testing.T) {
	assert := assert.New(t)
	e := events.New()
	expectedPayload := "this is my payload"

	e.On("my_event", func(payload ...interface{}) {
		if len(payload) <= 0 {
			t.Fatal("Expected payload but got nothing")
		}

		if s, ok := payload[0].(string); !ok {
			t.Fatalf("Payload is not the correct type, got: %#v", payload[0])
		} else if s != expectedPayload {
			t.Fatalf("Eexpected %s, got: %s", expectedPayload, s)
		}
	})

	e.Emit("my_event", expectedPayload)
	if e.Len() != 1 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 1)
	}

	if e.Len() != 1 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.ListenerCount("my_event"), 1)
	}

	e.OffAll("my_event")
	if e.Len() != 0 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 0)
	}

	if e.Len() != 0 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.ListenerCount("my_event"), 0)
	}

	now := time.Now()
	e.On("evt", func(name string, age int, created time.Time) {
		assert.Equal("john", name)
		assert.Equal(32, age)
		assert.Equal(now, created)
		t.Logf("name:%s\n", name)
		t.Logf("age:%d\n", age)
		t.Logf("created:%v\n", created)
		t.Logf("---------------------------------\n")
	})
	e.On("evt", func(args ...interface{}) {
		name := args[0].(string)
		age := args[1].(int)
		created := args[2].(time.Time)
		assert.Equal("john", args[0].(string))
		t.Logf("name:%s\n", name)
		t.Logf("age:%d\n", age)
		t.Logf("created:%v\n", created)
		t.Logf("---------------------------------\n")
	})
	e.On("evt", func(name string, age int) {
		assert.Equal("john", name)
		assert.Equal(32, age)
		t.Logf("name:%s\n", name)
		t.Logf("age:%d\n", age)
		t.Logf("---------------------------------\n")
	})
	type Address struct {
		UserID   int
		Location string
	}
	type Model struct {
		ID         int
		Username   string
		Age        int
		Money      *float64
		IsApproved bool
		CreatedAt  time.Time
		DeletedAt  *time.Time
		Addresses  []*Address
	}
	deleted := now.Add(30 * time.Minute)
	m := 1024.356
	user := Model{ID: 1001, Username: "john", Money: &m, Age: 32, IsApproved: true, CreatedAt: time.Now(), DeletedAt: &deleted}
	addresses := make([]*Address, 0)
	add1 := Address{UserID: 1, Location: "beijing of china"}
	add2 := Address{UserID: 2, Location: "hebei of china"}
	addresses = append(addresses, &add1, &add2)

	user.Addresses = addresses
	e.On("evt", func(name string, age int, a time.Time, user Model) {
		assert.Equal("john", name)
		assert.Equal(32, age)
		t.Logf("name:%s\n", name)
		t.Logf("age:%d\n", age)
		t.Logf("user.money:%v\n", *user.Money)
		t.Logf("user:%+v\n", user)
		for k, add := range user.Addresses {
			t.Logf("user.addresses[%v]:%+v\n", k, add)
		}
		t.Logf("---------------------------------\n")
	})

	for i, v := range e.Listeners("evt") {
		t.Logf("Listeners[%v]:%v\n", i, v)
	}
	e.Trigger("evt", "john", 32, now, user)
	t.Logf("----------[event end]--------------------\n")
	assert.Equal(4, e.ListenerCount("evt"))
	e.Clear()
	assert.Equal(0, e.Len())
}

func TestPanic(t *testing.T) {
	// assert := assert.New(t)
	e := events.New()

	e.On("evt", func() {
		t.Logf("none parameter function")
		panic("panic error")
	})
	e.On("evt", func() {
		t.Logf("22222222")
	})
	e.On("evt", func() {
		t.Logf("33333333")
	})
	e.Trigger("evt")
	e.Clear()

}
func TestEventsOnce(t *testing.T) {
	// on default
	events.Clear()

	var count = 0
	events.Once("my_event", func(payload ...interface{}) {
		if count > 0 {
			t.Fatalf("Once's listener fired more than one time! count: %d", count)
		}
		count++
	})

	if l := events.ListenerCount("my_event"); l != 1 {
		t.Fatalf("Real  event's listeners should be: %d but has: %d", 1, l)
	}

	if l := len(events.Listeners("my_event")); l != 1 {
		t.Fatalf("Real  event's listeners (from Listeners) should be: %d but has: %d", 1, l)
	}

	for i := 0; i < 10; i++ {
		events.Emit("my_event")
	}

	if l := events.ListenerCount("my_event"); l > 0 {
		t.Fatalf("Real event's listeners length count should be: %d but has: %d", 0, l)
	}

	if l := len(events.Listeners("my_event")); l > 0 {
		t.Fatalf("Real event's listeners length count ( from Listeners) should be: %d but has: %d", 0, l)
	}

}

func TestRemoveListener(t *testing.T) {
	assert := assert.New(t)
	e := events.New()

	var count = 0
	listener := func(payload ...interface{}) {
		if count > 1 {
			t.Fatal("Event listener should be removed")
		}

		count++
	}

	e.On("my_event", listener)
	e.On("my_event", func(payload ...interface{}) {})
	e.On("another_event", func(payload ...interface{}) {})

	e.Emit("my_event")

	assert.Equal(1, e.Off("my_event", listener).ListenerCount("my_event"))
	assert.Equal(0, e.Off("foo_bar", listener).ListenerCount("foo_bar"))

	for i, v := range e.Events() {
		t.Logf("Event[%v]:%v\n", i, v)
	}
	assert.Equal(2, len(e.Events()))

	assert.Equal(2, e.Len())
	assert.Equal(1, e.ListenerCount("my_event"))

	e.Emit("my_event")
	e.Clear()
}
