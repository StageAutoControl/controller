package sample

import (
	"io"
	"sync"
)

// Buffer is a storage for read samples that is thread safe by mutex. This is kinda FIFO Queue.
type Buffer struct {
	m    sync.Mutex
	data []float64
}

// NewBuffer returns a new Buffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

// Write adds a sample to the buffer
func (b *Buffer) Write(data float64) {
	b.m.Lock()
	defer b.m.Unlock()

	b.data = append(b.data, data)
}

// WriteAll writes all given samples to the Buffer
func (b *Buffer) WriteAll(data []float64) {
	b.m.Lock()
	defer b.m.Unlock()

	b.data = append(b.data, data...)
}

// Read returns the first sample from the buffer and deletes it.
func (b *Buffer) Read() (data float64, err error) {
	b.m.Lock()
	defer b.m.Unlock()

	if len(b.data) == 0 {
		return 0, io.EOF
	}

	data = b.data[0]
	b.data = b.data[1:]

	return data, nil
}
