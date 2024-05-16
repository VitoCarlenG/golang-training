package main

import "fmt"

type Door struct {
    KnockSound, OpenSound string
}

func (d Door) Knock() {
    fmt.Println(d.KnockSound)
}

func (d Door) Open() {
    fmt.Println(d.OpenSound)
}
