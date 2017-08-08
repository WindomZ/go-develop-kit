package queue

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var queue Queue

func TestNew(t *testing.T) {
	queue = New(10)
	assert.NotEmpty(t, queue)
}

func TestPool_Capacity(t *testing.T) {
	assert.Equal(t, 10, queue.Capacity())
}

func TestPool_Push(t *testing.T) {
	assert.Equal(t, 1, queue.Push(1))
	assert.Equal(t, 2, queue.Push(2))
	assert.Equal(t, 3, queue.Push(3))
}

func TestPool_PushSlice(t *testing.T) {
	assert.Equal(t, 9, queue.PushSlice([]interface{}{4, 5, 6, 7, 8, 9}))

	assert.Equal(t, 9, queue.PushSlice([]interface{}{}))
	assert.Equal(t, 10, queue.PushSlice([]interface{}{0}))

	assert.Equal(t, -1, queue.Push(2))
	assert.Equal(t, -1, queue.PushSlice([]interface{}{2}))
}

func TestPool_Size(t *testing.T) {
	assert.Equal(t, 10, queue.Size())
}

func TestPool_IsIdle(t *testing.T) {
	assert.False(t, queue.IsIdle())
}

func TestPool_Pull(t *testing.T) {
	assert.Equal(t, 1, queue.Pull())
	for i := 2; i <= 7; i++ {
		assert.Equal(t, i, queue.Pull())
	}
	assert.Equal(t, 8, queue.Pull())
}

func TestPool_PullSync(t *testing.T) {
	assert.Equal(t, 9, queue.PullSync())
	assert.Equal(t, 0, queue.PullSync())
	assert.Empty(t, queue.PullSync())
}

func TestPool_Exchange(t *testing.T) {
	assert.True(t, queue.Exchange(5))
	assert.Equal(t, 0, queue.Size())
	assert.Equal(t, -1, queue.PushSlice([]interface{}{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 5, queue.PushSlice([]interface{}{1, 2, 3, 4, 5}))
	assert.True(t, queue.Exchange(10))
	assert.Equal(t, -1, queue.PushSlice([]interface{}{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 10, queue.PushSlice([]interface{}{1, 2, 3, 4, 5}))
	assert.False(t, queue.Exchange(5))
}

func TestPool_Free(t *testing.T) {
	queue.Free()
	assert.Equal(t, 0, queue.Size())
}

func Benchmark_Pull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PushSlice([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		for j := 0; j < 10; j++ {
			queue.Pull()
		}
	}
}

func Benchmark_PullSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PushSlice([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		for j := 0; j < 10; j++ {
			queue.PullSync()
		}
	}
}
