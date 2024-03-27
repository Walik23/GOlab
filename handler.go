package main

import (
	"io"
)

// ComputeHandler обробляє обчислення виразу
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute обчислює вираз та записує результат
func (h *ComputeHandler) Compute() error {
	// ... (Реалізація обчислення виразу)
}
