package broker

import "testing"

func TestProduceAndConsume(t *testing.T) {
	b := NewBroker()

	offset0, err := b.Produce("orders", []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	offset1, err := b.Produce("orders", []byte("world"))
	if err != nil {
		t.Fatal(err)
	}

	if offset0 != 0 {
		t.Fatalf("expected offset 0, got %d", offset0)
	}

	if offset1 != 1 {
		t.Fatalf("expected offset 1, got %d", offset1)
	}

	msgs, err := b.Consume("orders", 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	if len(msgs) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(msgs))
	}

	if string(msgs[0].Payload) != "hello" {
		t.Fatalf("expected hello, got %s", msgs[0].Payload)
	}

	if string(msgs[1].Payload) != "world" {
		t.Fatalf("expected world, got %s", msgs[1].Payload)
	}
}

func TestConsumeFromOffset(t *testing.T) {
	b := NewBroker()

	b.Produce("orders", []byte("hello"))
	b.Produce("orders", []byte("world"))

	msgs, err := b.Consume("orders", 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	if len(msgs) != 1 {
		t.Fatalf("expected 1 message, got %d", len(msgs))
	}

	if string(msgs[0].Payload) != "world" {
		t.Fatalf("expected world, got %s", msgs[0].Payload)
	}
}

func TestConsumePastEndReturnsEmpty(t *testing.T) {
	b := NewBroker()

	b.Produce("orders", []byte("hello"))

	msgs, err := b.Consume("orders", 10, 10)
	if err != nil {
		t.Fatal(err)
	}

	if len(msgs) != 0 {
		t.Fatalf("expected 0 messages, got %d", len(msgs))
	}
}

func TestConsumeMissingTopic(t *testing.T) {
	b := NewBroker()

	_, err := b.Consume("missing", 0, 10)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "topic not found" {
		t.Fatalf("expected topic not found, got %v", err)
	}
}

func TestProduceEmptyTopic(t *testing.T) {
	b := NewBroker()

	_, err := b.Produce("   ", []byte("hello"))

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "no topic name" {
		t.Fatalf("expected no topic name, got %v", err)
	}
}

func TestConsumeEmptyTopic(t *testing.T) {
	b := NewBroker()

	_, err := b.Consume("   ", 0, 10)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "no topic name" {
		t.Fatalf("expected no topic name, got %v", err)
	}
}

func TestConsumeNegativeOffset(t *testing.T) {
	b := NewBroker()

	b.Produce("orders", []byte("hello"))

	_, err := b.Consume("orders", -1, 10)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "invalid offset" {
		t.Fatalf("expected invalid offset, got %v", err)
	}
}

func TestConsumeInvalidMax(t *testing.T) {
	b := NewBroker()

	b.Produce("orders", []byte("hello"))

	_, err := b.Consume("orders", 0, 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "invalid max" {
		t.Fatalf("expected invalid max, got %v", err)
	}
}
