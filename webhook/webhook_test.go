package webhook

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed testdata/button_reply.json
var buttonReply []byte

func TestParseButtonReply(t *testing.T) {
	var notification Notification
	err := json.Unmarshal(buttonReply, &notification)
	if err != nil {
		t.Fatalf("unable to unmarshall notification: %v", err)
	}

	message := notification.Entry[0].Changes[0].Value.Messages[0]
	if message.Type != "interactive" {
		t.Errorf("message type - want %s, got %s", "interactive", message.Type)
	}

	if message.Interactive.ButtonReply == nil {
		t.Fatalf("expected ButtonReply to be non nil")
	}
	buttonReply := message.Interactive.ButtonReply

	expectedID := "bug"
	expectedTitle := "bug in code"
	if buttonReply.ID != expectedID {
		t.Errorf("button id - got: %s, want: %s", buttonReply.ID, expectedID)
	}

	if buttonReply.Title != expectedTitle {
		t.Errorf("button title-  got: %s, want: %s", buttonReply.Title, expectedTitle)
	}
}

func TestParseContext(t *testing.T) {
	var notification Notification
	err := json.Unmarshal(buttonReply, &notification)
	if err != nil {
		t.Fatalf("unable to unmarshall notification: %v", err)
	}

	message := notification.Entry[0].Changes[0].Value.Messages[0]
	if message.Context == nil {
		t.Fatalf("expected message context to be non nil")
	}

	wantFrom := "15550783881"
	if message.Context.From != wantFrom {
		t.Errorf("context.From - got: %s, want: %s", message.Context.From, wantFrom)
	}

	wantId := "wamid.HBgLMTQxMjU1NTA4MjkVAgASGBQzQUNCNjk5RDUwNUZGMUZEM0VBRAA="
	if message.Context.ID != wantId {
		t.Errorf("context.ID - got: %s, want: %s", message.Context.ID, wantId)
	}
}
