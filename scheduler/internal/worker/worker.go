package main
package worker

import "context"

type Worker struct{}

func New() *Worker {
	return &Worker{}
}

func (w *Worker) Run(_ context.Context) error {
	return nil
}
