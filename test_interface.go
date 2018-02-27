package testing_interface

import "fmt"

type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() { fmt.Println("meow") }

func GetACat() Cat {
	var myTabby *Tabby = nil
	// Oops, we forgot to set myTabby to a real value
	return myTabby
}

func TestGetACat(t *testing.T) {
	if GetACat() == nil {
		t.Errorf("Forgot to return a real cat!")
	}
}
