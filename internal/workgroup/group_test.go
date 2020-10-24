package workgroup_test

import (
	"errors"
	"io"
	"testing"

	"commerceiq.ai/ticketing/internal/workgroup"
)

func TestGroupRunWithNoRegisteredFunctions(t *testing.T) {
	var g workgroup.Group
	got := g.Run()
	assert(t, nil, got)
}

func TestGroupFirstReturnValueIsReturnedToRunsCaller(t *testing.T) {
	var g workgroup.Group
	wait := make(chan int)
	g.Add(func(<-chan struct{}) error {
		<-wait
		return io.EOF
	})

	g.Add(func(stop <-chan struct{}) error {
		<-stop
		return errors.New("stopped")
	})

	result := make(chan error)
	go func() {
		result <- g.Run()
	}()
	close(wait)
	assert(t, io.EOF, <-result)
}

func assert(t *testing.T, want, got error) {
	t.Helper()
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
