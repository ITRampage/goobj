package inc

import (
	"fmt"
)

type mKind = int

const (
	_ mKind = iota
	incMsg
	printMsg
	desctructMsg
)

type Obj struct {
	counter  int64
	messages chan mKind
}

func receive(obj *Obj) {
	for m := range obj.messages {
		switch m {
		case incMsg:
			obj.counter++
		case printMsg:
			fmt.Printf("\n\t\tcounter: %d\n", obj.counter)
		case desctructMsg:
			fmt.Printf("\n\tfinal counter: %d\n\n", obj.counter)
			return
		}
	}
}
func New(init int64) *Obj {
	obj := &Obj{
		counter:  init,
		messages: make(chan mKind),
	}
	go receive(obj)
	return obj
}

func (o *Obj) Inc() {
	o.messages <- incMsg
}
func (o *Obj) Print() {
	o.messages <- printMsg
}
func (o *Obj) Descruct() {
	o.messages <- desctructMsg
}
