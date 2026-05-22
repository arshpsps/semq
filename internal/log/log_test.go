package log

import "testing"

func TestAppendAssignsOffsets(t *testing.T) {
	l := NewLog()

	offset0 := l.Append([]byte("hello"))
	offset1 := l.Append([]byte("world"))

	if offset0 != 0 {
		t.Fatalf("expected offset 0, got %d", offset0)
	}

	if offset1 != 1 {
		t.Fatalf("expected offset 1, got %d", offset1)
	}
}

func TestReadFromOffset(t *testing.T) {
	l := NewLog()

	l.Append([]byte("hello"))
	l.Append([]byte("world"))

	msgs := l.Read(1, 10)

	if len(msgs) != 1 {
		t.Fatalf("expected 1 message, got %d", len(msgs))
	}

	if string(msgs[0].Payload) != "world" {
		t.Fatalf("expected world, got %s", msgs[0].Payload)
	}
}

func TestReadPastEndReturnsEmpty(t *testing.T) {
	l := NewLog()

	l.Append([]byte("hello"))

	msgs := l.Read(10, 10)

	if len(msgs) != 0 {
		t.Fatalf("expected 0 messages, got %d", len(msgs))
	}
}

