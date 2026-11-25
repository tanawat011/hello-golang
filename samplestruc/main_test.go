package main

import "testing"

func TestCustomersJSON(t *testing.T) {
	got, err := customersJSON([]customer{{Firstname: "A", Lastname: "B", Code: 1, Phone: "0"}})
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"Firstname":"A","Lastname":"B","Code":1,"Phone":"0"}]`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEmptyListIsAnEmptyArray(t *testing.T) {
	got, err := customersJSON([]customer{})
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "[]" {
		t.Errorf("got %s, want []", got)
	}
}
