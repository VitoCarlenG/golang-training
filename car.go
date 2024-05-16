package main

import "fmt"

type Wheel interface {
    Type() string
}

type RubberWheel struct{}
type WoodWheel struct{}
type IronWheel struct{}

func (w RubberWheel) Type() string {
    return "Rubber Wheel"
}

func (w WoodWheel) Type() string {
    return "Wood Wheel"
}

func (w IronWheel) Type() string {
    return "Iron Wheel"
}

type Car struct {
    Wheels [4]Wheel
    LeftDoor, RightDoor Door
}

func NewCar() *Car {
    return &Car{
        Wheels: [4]Wheel{RubberWheel{}, RubberWheel{}, RubberWheel{}, RubberWheel{}},
        LeftDoor:  Door{"beep", "Knock"},
        RightDoor: Door{"Knock", "beep"},
    }
}

func (c *Car) ReplaceWheel(position int, wheel Wheel) {
    if position >= 0 && position < len(c.Wheels) {
        c.Wheels[position] = wheel
    } else {
        fmt.Println("Invalid wheel position")
    }
}
