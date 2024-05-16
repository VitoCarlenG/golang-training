package main

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
