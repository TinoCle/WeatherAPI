package myip_test

import (
	"testing"

	"github.com/polds/MyIP"
)

func testMyIP(t *testing.T) {
	ip, err := myip.GetMyIP()

	if err != nil {
		t.Errorf("There was an error: %s\n", err.Error())
		t.Error("No error is expected.")
	}

	if ip == "" {
		t.Error("IP is not set, expecting value.")
	}
}
