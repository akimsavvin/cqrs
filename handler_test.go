package cqrs

import "testing"

type MyQuery struct {
	Name string
}

type MyQueryHandler struct{}

func (h *MyQueryHandler) Handle(query MyQuery) (string, error) {
	return query.Name, nil
}

type MyCommand struct {
	Age int
}

type MyCommandHandler struct{}

func (h *MyCommandHandler) Handle(command MyCommand) (int, error) {
	return command.Age, nil
}

func TestCQRS(t *testing.T) {
	// Register handlers
	RegisterHandler[MyQuery, string](&MyQueryHandler{})
	RegisterHandler[MyCommand, int](&MyCommandHandler{})

	const (
		Name = "Akim"
		Age  = 16
	)

	// Query
	queryRes := HandleEvent[string](MyQuery{Name})

	if queryRes.Payload() != Name {
		t.Errorf("expected %v payload, got %v", Name, queryRes.Payload())
	}

	if queryRes.Err() != nil {
		t.Errorf("got unexpected error for query: %v", queryRes.Err())
	}

	// Command
	cmdRes := HandleEvent[int](MyCommand{Age})

	if cmdRes.Payload() != Age {
		t.Errorf("expected %v payload, got %v", Age, queryRes.Payload())
	}

	if cmdRes.Err() != nil {
		t.Errorf("got unexpected error for command: %v", cmdRes.Err())
	}
}
