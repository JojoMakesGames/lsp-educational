package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 16 {
		t.Fatalf("Expected 16, got %d", contentLength)
	}
}
