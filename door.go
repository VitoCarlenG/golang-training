package main

import "fmt"

type Door struct {
	Side string
}

func (door *Door) Tap() {
	if door.Side == "Kanan" {
		fmt.Println("Knock")
	} else {
		fmt.Println("Beep")
	}
}

func (door *Door) Open() {
	if door.Side == "Kanan" {
		fmt.Println("Beep")
	} else {
		fmt.Println("Knock")
	}
}
