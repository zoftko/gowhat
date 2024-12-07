package webhook

import (
	"strconv"
	"time"
)

// Time parses the message's UNIX timestamp into a time.Time instance.
func (m Message) Time() (time.Time, error) {
	i, err := strconv.ParseInt(m.Timestamp, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(i, 0), nil
}
