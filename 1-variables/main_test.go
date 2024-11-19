package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestPrintTypes(t *testing.T) {
	// Перехватываем stdout для проверки вывода
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printTypes(
		1,
		2.2,
		"Hello",
		true,
		complex64(1+2i))

	// Восстанавливаем stdout и читаем результат
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := []string{
		"значение = 1, тип = int",
		"значение = 2.20, тип = float64",
		"значение = Hello, тип = string",
		"значение = true, тип = bool",
		"значение = (1+2i), тип = complex64",
	}

	outputLines := strings.Split(output, "\n")

	for i, exp := range expected {
		if !strings.Contains(outputLines[i], exp) {
			t.Errorf("Ожидалось %s, получено %s", exp, output)
		}
	}
}

func TestConvertToString(t *testing.T) {
	expected := "123trueHello(1+2i)"
	result := convertToString(12, 3, true, "Hello", complex64(1+2i))

	if strings.Compare(result, expected) != 0 {
		t.Errorf("Ожидалось %s, получено %s", expected, result)
	}
}

func TestConvertToRuneSlice(t *testing.T) {
	expected := []rune("Not rune string")
	result := convertToRuneSlice("Not rune string")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestHashRuneSliceWithSalt(t *testing.T) {
	expected := "af4e04569d12a0b47833a7e5b51ddb1a44ed0c20adffd94ef95e2d1ea99bb47f"
	result := hashRuneSliceWithSalt([]rune("Not rune string"), "")

	if strings.Compare(result, expected) != 0 {
		t.Errorf("Ожидалось %s, получено %s", expected, result)
	}
}
