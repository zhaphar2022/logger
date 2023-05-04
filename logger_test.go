package logger

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

var now = time.Now().Format("2006/01/02 15:04:05")

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	Info("test message")

	got := buf.String()
	want := now + " INFO - test message\n"

	if got != want {
		t.Errorf("Info() = %q, want %q", got, want)
	}
}

func TestWarning(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	Warning("test message")

	got := buf.String()
	want := now + " WARN - test message\n"

	if got != want {
		t.Errorf("Warning() = %q, want %q", got, want)
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	Error("test message")

	got := buf.String()
	want := now + " ERROR - test message\n"

	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestLoggingDisabled(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	Info("test message")
	Warning("test message")
	Error("test message")
}

func TestLoggingEnabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	Info("test message")
	Warning("test message")
	Error("test message")

	got := buf.String()
	want := now + " INFO - test message\n" + now + " WARN - test message\n" + now + " ERROR - test message\n"

	if got != want {
		t.Errorf("Output() = %q, want %q", got, want)
	}
}
