package server

import (
	"errors"
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLog()

	// Append
	expected := Record{Value: []byte("test"), Offset: 0}
	offset, err := log.Append(expected)
	if err != nil {
		t.Errorf("Append returned error: %v", err)
	}

	// Read
	record, err := log.Read(offset)
	if err != nil {
		t.Errorf("Read returned error: %v", err)
	}
	if !reflect.DeepEqual(record, expected) {
		t.Errorf("Read record mismatch, expected: %+v, got: %+v", expected, record)
	}

	// non-existing
	_, err = log.Read(32)
	if !errors.Is(err, ErrOffsetNotFound) {
		t.Errorf("Expected ErrOffsetNotFound, got: %v", err)
	}
}
