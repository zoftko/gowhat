package message

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

//go:embed testdata/message_read_with_typing.json
var expectedMessageReadWithTypingJSON string

//go:embed testdata/message_read_without_typing.json
var expectedMessageReadWithoutTypingJSON string

func TestBaseEnvelopeRouting(t *testing.T) {
	tests := []struct {
		name          string
		to            string
		wantTo        string
		wantRecipient string
	}{
		{name: "phoneNumber", to: "16505551234", wantTo: "16505551234", wantRecipient: ""},
		{name: "bsuid", to: "US.123", wantTo: "", wantRecipient: "US.123"},
		{name: "parentBsuid", to: "US.ENT.123", wantTo: "", wantRecipient: "US.ENT.123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := baseEnvelope(tt.to)
			if e.To != tt.wantTo {
				t.Errorf("To: got %q, want %q", e.To, tt.wantTo)
			}
			if e.Recipient != tt.wantRecipient {
				t.Errorf("Recipient: got %q, want %q", e.Recipient, tt.wantRecipient)
			}
		})
	}
}

func TestNewMessageRead(t *testing.T) {
	tests := []struct {
		messageID       string
		typingIndicator bool
		expectedJSON    string
	}{
		{
			messageID:       "<WHATSAPP_MESSAGE_ID>",
			typingIndicator: true,
			expectedJSON:    expectedMessageReadWithTypingJSON,
		},
		{
			messageID:       "<WHATSAPP_MESSAGE_ID>",
			typingIndicator: false,
			expectedJSON:    expectedMessageReadWithoutTypingJSON,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vTypingIndicator", tt.typingIndicator), func(t *testing.T) {
			msg := NewMessageRead(tt.messageID, tt.typingIndicator)
			actualJSON, err := json.Marshal(msg)
			if err != nil {
				t.Fatalf("failed to marshal message: %v", err)
			}

			var expected, actual map[string]interface{}

			if err := json.Unmarshal([]byte(tt.expectedJSON), &expected); err != nil {
				t.Fatalf("failed to unmarshal expected JSON: %v", err)
			}

			if err := json.Unmarshal(actualJSON, &actual); err != nil {
				t.Fatalf("failed to unmarshal actual JSON: %v", err)
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got %v, want %v", actual, expected)
			}
		})
	}
}
