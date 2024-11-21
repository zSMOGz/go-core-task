package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

var input = []int{1, 2, 3, 4, 5}

func TestPrintSlice(t *testing.T) {
	expected := []string{"1", "2", "3", "4", "5"}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printSlice(&input)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	outputLines := strings.Split(output, "\n")

	for i, exp := range expected {
		if !strings.Contains(outputLines[i], exp) {
			t.Errorf("Ожидалось %s, получено %s", exp, output)
		}
	}
}

func TestSliceExample(t *testing.T) {
	expected := []int{2, 4}

	result := sliceExample(&input)

	for i, exp := range expected {
		if (*result)[i] != exp {
			t.Errorf("Ожидалось %d, получено %d", exp, (*result)[i])
		}
	}
}

func TestAddElements(t *testing.T) {
	appendElement := 6
	expected := []int{1, 2, 3, 4, 5, 6}

	result := addElements(&input, appendElement)

	for i, exp := range expected {
		if (*result)[i] != exp {
			t.Errorf("Ожидалось %d, получено %d", exp, (*result)[i])
		}
	}
}

func TestCopySlice(t *testing.T) {
	result := copySlice(&input)

	for i, val := range input {
		if (*result)[i] != val {
			t.Errorf("Ожидалось %d, получено %d", val, (*result)[i])
		}
	}

	(*result)[0] = 999
	if input[0] == (*result)[0] {
		t.Error("Копия не является независимой от оригинала")
	}
}

func TestRemoveElement(t *testing.T) {
	indexToRemove := 2
	expected := []int{1, 2, 4, 5}

	result := removeElement(&input, indexToRemove)

	if len(*result) != len(expected) {
		t.Errorf("Ожидалась длина %d, получена %d", len(expected), len(*result))
	}

	for i, exp := range expected {
		if (*result)[i] != exp {
			t.Errorf("Ожидалось %d, получено %d", exp, (*result)[i])
		}
	}
}
