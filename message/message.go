package message

type Envelope struct {
	MessagingProduct string       `json:"messaging_product"`
	RecipientType    string       `json:"recipient_type,omitempty"`
	To               string       `json:"to,omitempty"`
	Type             string       `json:"type,omitempty"`
	Text             *Text        `json:"text,omitempty"`
	Image            *Image       `json:"image,omitempty"`
	Interactive      *Interactive `json:"interactive,omitempty"`
	Status           string       `json:"status,omitempty"`
	MessageID        string       `json:"message_id,omitempty"`
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
	Name       string     `json:"name"`
	Parameters Parameters `json:"parameters"`
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
	return Envelope{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text: &Text{
			Body:       opts.Text,
			PreviewURL: opts.PreviewURL,
		},
	}
}

func NewImageLink(to string, opts NewImageLinkOpts) Envelope {
	return Envelope{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "image",
		Image: &Image{
			Link:    opts.Link,
			Caption: opts.Caption,
		},
	}
}

func NewInteractiveFlow(to string, opts NewFlowOpts) Envelope {
	return Envelope{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "interactive",
		Interactive: &Interactive{
			Type:   "flow",
			Header: opts.Header,
			Body:   opts.Body,
			Footer: opts.Footer,
			Action: Action{
				Name: "flow",
				Parameters: Parameters{
					Mode:               opts.FlowMode,
					FlowMessageVersion: "3",
					FlowToken:          opts.FlowToken,
					FlowID:             opts.FlowId,
					FlowCTA:            opts.FlowCTA,
					FlowAction:         "navigate",
					FlowActionPayload:  FlowAction{Screen: opts.FirstScreen},
				},
			},
		},
	}
}

func NewMessageRead(messageId string) Envelope {
	return Envelope{
		MessagingProduct: "whatsapp",
		Status:           "read",
		MessageID:        messageId,
	}
}
