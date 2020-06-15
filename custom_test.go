package custom

import (
	"encoding/json"
	"strconv"
	"testing"
)

func Test_should_marshall_customized(t *testing.T) {
	var p = Person{
		First: "John",
		Email: &Email{
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

func Test_should_not_marshall_when_name_nil(t *testing.T) {
	var p = Person{
		First: "John",
		Email: &Email{
			Name:   "",
			Domain: "example.com",
		},
	}

	MarshaljsonBytes, err := json.Marshal(&p)
	if err != nil {
		t.Error("Marshaller failed")
	}

	expected := "{\"name\":\"John\",\"email\":null}"
	result := string(MarshaljsonBytes)
	if expected != result {
		t.Errorf("Marshalling failed, expected %s, got %s\n", expected, result)
	}
}

func Test_should_unmarshall(t *testing.T) {
	personString := "{\"name\":\"John\",\"email\":\"john.doe@example.com\"}"

	person := Person{}
	json.Unmarshal([]byte(personString), &person)

	if person.First != "John" {
		t.Errorf("Firstname does not match, expected %s, got %s\n", "John", person.First)
	}
	if person.Email.Name != "john.doe" {
		t.Errorf("Email.Name does not match, expected %s, got %s\n", "john.doe", person.Email.Name)
	}
	if person.Email.Domain != "example.com" {
		t.Errorf("Email.Domain does not match, expected %s, got %s\n", "example.com", person.Email.Domain)
	}
}

func Test_should_marshall_nil_email(t *testing.T) {
	personString := "{\"name\":\"John\",\"email\":}"

	person := Person{}
	json.Unmarshal([]byte(personString), &person)

	if person.Email != nil {
		t.Errorf("Email marshall does not match, expected %s, got %s\n", "john.doe", person.Email.Name)
	}
}

var MarshaljsonBytes []byte

func benchmarkMarshall(b *testing.B, num int) {
	var p = Person{
		First: "John" + strconv.Itoa(num),
		Email: &Email{
			Name:   "",
			Domain: "example.com",
		},
	}
	for i := 0; i < num; i++ {
		MarshaljsonBytes, _ = json.Marshal(&p)
	}
}

func BenchmarkMarshall(b *testing.B) {
	benchmarkMarshall(b, 1000000)
}
