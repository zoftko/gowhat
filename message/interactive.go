package message

type InteractiveReplyButtonsOpts struct {
	Header  InteractiveHeader
	Body    string
	Buttons []Reply
}

func NewInteractiveReplyButtons(to string, opts InteractiveReplyButtonsOpts) Envelope {
	buttons := make([]Button, len(opts.Buttons))
	for i, b := range opts.Buttons {
		buttons[i] = Button{
			Type:  "reply",
			Reply: b,
		}
	}

	e := baseEnvelope(to)
	e.Type = "interactive"
	e.Interactive = &Interactive{
		Header: opts.Header,
		Body:   TextObject{Text: opts.Body},
		Type:   "button",
		Action: Action{
			Buttons: &buttons,
		},
	}
	return e
}
