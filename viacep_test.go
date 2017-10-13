package viacep

import (
	"testing"
)

func TestValid(t *testing.T) {
	resp, err := Consulta("0 1 243-000")
	if err != nil {
		t.Error(err)
	}
	if resp.Logradouro != "Rua Sergipe" {
		t.Errorf("invalid response %v != %v", resp.Logradouro, "Rua Sergipe")
	}
}

func TestBadRequest(t *testing.T) {
	_, err := Consulta("000000000000000")
	if err == nil {
		t.Error("invalid CEP should fail")
	}
	if err.Error() != "http status code 400" {
		t.Errorf("invalid response %v != %v", err.Error(), "http status code 400")
	}
}

func TestNotFound(t *testing.T) {
	_, err := Consulta("00000000")
	if err == nil {
		t.Error("invalid CEP should fail")
	}
	if !NotFound(err) {
		t.Errorf("invalid response %v != %v", err.Error(), "not found")
	}
}
