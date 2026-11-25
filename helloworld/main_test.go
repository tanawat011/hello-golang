package main

import "testing"

func TestGetAddressAndYear(t *testing.T) {
	address, year := getAddressAndYear()
	if address != "นครศรีธรรมราช" || year != 2016 {
		t.Errorf("got (%q, %d), want (%q, 2016)", address, year, "นครศรีธรรมราช")
	}
}
