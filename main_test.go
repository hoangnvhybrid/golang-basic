package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"
	if got := sayHello(name); got != want {
		t.Errorf("hell() = %q, want %q", got, want)
	}
}
