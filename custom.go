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
	Email *Email `json:"email"`
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
	emailAddress, error := unMarshallString(b)
	if error != nil {
		return error
	}
	username, domain := splitEmail(emailAddress)
	*em = Email{Name: username, Domain: domain}
	return nil
}

func unMarshallString(b []byte) (string, error) {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return "", err
	}
	return s, nil
}

func splitEmail(e string) (string, string) {
	mid := strings.Index(e, "@")
	username := e[:mid]
	domain := e[mid+1:]
	return username, domain
}
