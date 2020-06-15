package custom

import (
	"encoding/json"
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

	expected := "{\"name\":\"John\",\"email\":\"john.doe@example.com\"}"
	result := string(MarshaljsonBytes)
	if expected != result {
		t.Errorf("Marshalling failed, expected %s, got %s\n", expected, result)
	}
}
