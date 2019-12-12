package computer

import (
	"reflect"
	"testing"
)

func TestComputer(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	compareArrays := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Load", func(t *testing.T) {
		computer := New()
		computer.Load(program)

		got := computer.memory
		want := program
		compareArrays(t, got, want)
	})

	t.Run("Execute", func(t *testing.T) {
		computer := New()

		computer.Load(program)

		got := computer.Execute()
		want := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
		compareArrays(t, got, want)
	})

	t.Run("Add", func(t *testing.T) {
		computer := New()

		computer.Load([]int{1, 1, 2, 1, 99})

		got := computer.Execute()
		want := []int{1, 3, 2, 1, 99}
		compareArrays(t, got, want)
	})

	t.Run("Multiply", func(t *testing.T) {
		computer := New()

		computer.Load([]int{2, 2, 3, 3, 99})

		got := computer.Execute()
		want := []int{2, 2, 3, 9, 99}
		compareArrays(t, got, want)
	})
}
