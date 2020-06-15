package custom

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Email is name and domain
type Email struct {
	Name   string
	Domain string
}

// Person is first and email
type Person struct {
	First string `json:"name"`
	Email Email  `json:"email"`
}

func (em *Email) String() string {
	return fmt.Sprintf("%s@%s", em.Name, em.Domain)
}

// MarshalJSON for marshalling email
func (em *Email) MarshalJSON() ([]byte, error) {
	if em.Name == "" {
		return []byte("null"), nil
	}
	return json.Marshal(em.String())
}

// UnmarshalJSON for unmarshalling email
func (em *Email) UnmarshalJSON(b []byte) error {

	if string(b) == `null` {
		return nil
	}

	var fullEmailAddress string
	if err := json.Unmarshal(b, &fullEmailAddress); err != nil {
		return err
	}

	mid := strings.Index(fullEmailAddress, "@")
	username := fullEmailAddress[:mid]
	domain := fullEmailAddress[mid+1:]

	*em = Email{Name: username, Domain: domain}

	return nil
}
