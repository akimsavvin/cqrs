package cqrs

import (
	"reflect"
)

var handlers = make(map[reflect.Type]reflect.Value)

type Event any

type Result[TPayload any] struct {
	payload TPayload
	err     error
}

func (r *Result[TPayload]) Payload() TPayload {
	return r.payload
}

func (r *Result[TPayload]) Err() error {
	return r.err
}

type Handler[TEvent, TPayload any] interface {
	Handle(event TEvent) (TPayload, error)
}

func RegisterHandler[TEvent, TResult any](handler Handler[TEvent, TResult]) {
	eventType := reflect.TypeOf(handler.Handle).In(0)
	handlers[eventType] = reflect.ValueOf(handler)
}

func HandleEvent[TResult any, TEvent Event](event TEvent) *Result[TResult] {
	handler, ok := handlers[reflect.TypeOf(event)].Interface().(Handler[TEvent, TResult])

	if !ok {
		panic("handler is not a type of Handler[TEvent, TResult]")
	}

	payload, err := handler.Handle(event)

	return &Result[TResult]{
		payload: payload,
		err:     err,
	}
}
