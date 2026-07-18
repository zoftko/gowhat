package message

import "regexp"

// bsuidPattern matches regular BSUIDs (e.g. US.13491208655302741918) and
// parent BSUIDs (e.g. US.ENT.11815799212886844830):
// 2-letter ISO country code + "." + optional "ENT." + 1–128 alphanumeric chars.
var bsuidPattern = regexp.MustCompile(`^[A-Za-z]{2}\.(ENT\.)?[A-Za-z0-9]{1,128}$`)

func isBSUID(s string) bool {
	return bsuidPattern.MatchString(s)
}

func baseEnvelope(to string) Envelope {
	e := Envelope{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
	}
	if isBSUID(to) {
		e.Recipient = to
	} else {
		e.To = to
	}
	return e
}

type Envelope struct {
	MessagingProduct string           `json:"messaging_product"`
	RecipientType    string           `json:"recipient_type,omitempty"`
	To               string           `json:"to,omitempty"`
	Recipient        string           `json:"recipient,omitempty"`
	Type             string           `json:"type,omitempty"`
	Text             *Text            `json:"text,omitempty"`
	Image            *Image           `json:"image,omitempty"`
	Interactive      *Interactive     `json:"interactive,omitempty"`
	Document         *Document        `json:"document,omitempty"`
	Sticker          *Sticker         `json:"sticker,omitempty"`
	Status           string           `json:"status,omitempty"`
	TypingIndicator  *TypingIndicator `json:"typing_indicator,omitempty"`
	MessageID        string           `json:"message_id,omitempty"`
}

type TypingIndicator struct {
	Type string `json:"type"`
}

type Sticker struct {
	Link string `json:"link"`
}

type Document struct {
	Link     string `json:"link"`
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
}

type Text struct {
	Body       string `json:"body"`
	PreviewURL *bool  `json:"preview_url,omitempty"`
}

type TextObject struct {
	Text string `json:"text"`
}

type Image struct {
	ID      string `json:"id,omitempty"`
	Link    string `json:"link,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type Interactive struct {
	Type   string            `json:"type"`
	Header InteractiveHeader `json:"header"`
	Body   TextObject        `json:"body"`
	Footer *Text             `json:"footer,omitempty"`
	Action Action            `json:"action"`
}

type InteractiveHeader struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type Action struct {
	Name       *string     `json:"name"`
	Parameters *Parameters `json:"parameters,omitempty"`
	Buttons    *[]Button   `json:"buttons,omitempty"`
}

type Button struct {
	Type  string `json:"type"`
	Reply Reply  `json:"reply"`
}

type Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Parameters struct {
	Mode               string     `json:"mode,omitempty"`
	FlowMessageVersion string     `json:"flow_message_version"`
	FlowToken          string     `json:"flow_token"`
	FlowID             string     `json:"flow_id"`
	FlowCTA            string     `json:"flow_cta"`
	FlowAction         string     `json:"flow_action"`
	FlowActionPayload  FlowAction `json:"flow_action_payload,omitempty"`
}

type FlowAction struct {
	Screen string `json:"screen"`
}

type NewDocumentOpts struct {
	Link     string
	Caption  string
	Filename string
}

type NewTextOpts struct {
	Text       string
	PreviewURL *bool
}

type NewImageLinkOpts struct {
	Link    string
	Caption string
}

type NewFlowOpts struct {
	Header      InteractiveHeader
	Body        TextObject
	Footer      *Text
	FlowMode    string
	FlowId      string
	FlowToken   string
	FlowCTA     string
	FirstScreen string
}

func NewText(to string, opts NewTextOpts) Envelope {
	e := baseEnvelope(to)
	e.Type = "text"
	e.Text = &Text{
		Body:       opts.Text,
		PreviewURL: opts.PreviewURL,
	}
	return e
}

func NewImageLink(to string, opts NewImageLinkOpts) Envelope {
	e := baseEnvelope(to)
	e.Type = "image"
	e.Image = &Image{
		Link:    opts.Link,
		Caption: opts.Caption,
	}
	return e
}

func NewInteractiveFlow(to string, opts NewFlowOpts) Envelope {
	flow := "flow"
	e := baseEnvelope(to)
	e.Type = "interactive"
	e.Interactive = &Interactive{
		Type:   "flow",
		Header: opts.Header,
		Body:   opts.Body,
		Footer: opts.Footer,
		Action: Action{
			Name: &flow,
			Parameters: &Parameters{
				Mode:               opts.FlowMode,
				FlowMessageVersion: "3",
				FlowToken:          opts.FlowToken,
				FlowID:             opts.FlowId,
				FlowCTA:            opts.FlowCTA,
				FlowAction:         "navigate",
				FlowActionPayload:  FlowAction{Screen: opts.FirstScreen},
			},
		},
	}
	return e
}

func NewMessageRead(messageId string, typingIndicator bool) Envelope {
	var indicator *TypingIndicator
	if typingIndicator {
		indicator = &TypingIndicator{
			Type: "text",
		}
	}

	return Envelope{
		MessagingProduct: "whatsapp",
		Status:           "read",
		MessageID:        messageId,
		TypingIndicator:  indicator,
	}
}

func NewDocument(to string, opts NewDocumentOpts) Envelope {
	e := baseEnvelope(to)
	e.Type = "document"
	e.Document = &Document{
		Link:     opts.Link,
		Caption:  opts.Caption,
		Filename: opts.Filename,
	}
	return e
}

func NewSticker(to, link string) Envelope {
	e := baseEnvelope(to)
	e.Type = "sticker"
	e.Sticker = &Sticker{Link: link}
	return e
}
