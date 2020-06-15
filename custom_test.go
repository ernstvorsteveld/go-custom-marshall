package custom

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_should_marshall_customized(t *testing.T) {
	var p = Person{
		First: "John",
		Email: Email{
			Name:   "john.doe",
			Domain: "example.com",
		},
	}

	MarshaljsonBytes, err := json.Marshal(&p)
	if err != nil {
		t.Error("Marshaller failed")
	}

	fmt.Printf("Got string %s", string(MarshaljsonBytes))
}
