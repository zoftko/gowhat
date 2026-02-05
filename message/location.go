package message

// Location represents a WhatsApp location message.
// See: https://developers.facebook.com/documentation/business-messaging/whatsapp/messages/location-messages/
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      *string `json:"name,omitempty"`
	Address   *string `json:"address,omitempty"`
}

// NewLocationOpts for creating a new location message.
type NewLocationOpts struct {
	Latitude  float64
	Longitude float64
	Name      *string
	Address   *string
}

// NewLocation creates a new WhatsApp location message.
func NewLocation(to string, opts NewLocationOpts) Envelope {
	return Envelope{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "location",
		Location: &Location{
			Latitude:  opts.Latitude,
			Longitude: opts.Longitude,
			Name:      opts.Name,
			Address:   opts.Address,
		},
	}
}
