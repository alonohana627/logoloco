package sink_test

import (
	"bufio"
	s "logoloco/sink"
	"os"
	"strings"
	"testing"
)

const fileName = "./__fixtures__/testfile.txt"

func TestNewFileSink(t *testing.T) {
	sink, err := s.NewFileSink(fileName)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if sink == nil {
		t.Fatal("Expected sink to be not nil")
	}
}

func TestWrite(t *testing.T) {
	message := "test message"

	sink, _ := s.NewFileSink(fileName)
	err := sink.Write(message)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if !strings.Contains(lines[len(lines)-1], message) {
		t.Errorf("Expected file to contain '%v', but it did not", message)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
		return
	}
}
