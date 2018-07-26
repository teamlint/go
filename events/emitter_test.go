package events

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testEvents = map[interface{}][]Listener{
	"user_created": []Listener{
		func(payload ...interface{}) {
			fmt.Printf("A new User just created!\n")
		},
		func(payload ...interface{}) {
			fmt.Printf("A new User just created, *from second event listener\n")
		},
	},
	"user_joined": []Listener{
		func(payload ...interface{}) {
			user := payload[0].(string)
			room := payload[1].(string)
			fmt.Printf("%s joined to room: %s\n", user, room)
		},
	},
	"user_left": []Listener{func(payload ...interface{}) {
		user := payload[0].(string)
		room := payload[1].(string)
		fmt.Printf("%s left from the room: %s\n", user, room)
	}},
}

func createUser(user string) {
	EmitSync("user_created", user)
}

func joinUserTo(user string, room string) {
	EmitSync("user_joined", user, room)
}

func leaveFromRoom(user string, room string) {
	EmitSync("user_left", user, room)
}

func ExampleEvents() {
	// regiter our events to the default event emmiter
	for evt, listeners := range testEvents {
		for _, l := range listeners {
			On(evt, l)
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

	// Output:
	// A new User just created!
	// A new User just created, *from second event listener
	// user1 joined to room: room1
	// user1 left from the room: room1
}

func TestEvents(t *testing.T) {
	assert := assert.New(t)
	e := New()
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

	// e.On("evt", func(name string, age int) {
	// 	assert.Equal("john", name)
	// 	assert.Equal(32, age)
	// })
	e.On("evt", func(name ...interface{}) {
		assert.Equal("john", name[0].(string))
	})
	// type Model struct {
	// 	Username   string
	// 	Age        int
	// 	IsApproved bool
	// 	CreatedAt  time.Time
	// }
	// user:=Model{Username:"john",Age:32,IsApproved:true,CreatedAt:time.Now()}
	e.Trigger("evt", "john", 32, time.Now)
}

func TestEventsOnce(t *testing.T) {
	// on default
	Clear()

	var count = 0
	Once("my_event", func(payload ...interface{}) {
		if count > 0 {
			t.Fatalf("Once's listener fired more than one time! count: %d", count)
		}
		count++
	})

	if l := ListenerCount("my_event"); l != 1 {
		t.Fatalf("Real  event's listeners should be: %d but has: %d", 1, l)
	}

	if l := len(Listeners("my_event")); l != 1 {
		t.Fatalf("Real  event's listeners (from Listeners) should be: %d but has: %d", 1, l)
	}

	for i := 0; i < 10; i++ {
		Emit("my_event")
	}

	if l := ListenerCount("my_event"); l > 0 {
		t.Fatalf("Real event's listeners length count should be: %d but has: %d", 0, l)
	}

	if l := len(Listeners("my_event")); l > 0 {
		t.Fatalf("Real event's listeners length count ( from Listeners) should be: %d but has: %d", 0, l)
	}

}

func TestRemoveListener(t *testing.T) {
	// on default
	assert := assert.New(t)
	e := New()

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
}
