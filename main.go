package main

import (
    "fmt"
    "golang-training/car"
)

func main() {
    // Buat mobil baru
    myCar := car.NewCar()

    // Ganti roda
    myCar.ReplaceWheel(0, car.WoodWheel{})
    myCar.ReplaceWheel(1, car.IronWheel{})
    myCar.ReplaceWheel(2, car.RubberWheel{})
    myCar.ReplaceWheel(3, car.WoodWheel{})

    // Tampilkan tipe roda
    for i, wheel := range myCar.Wheels {
        fmt.Printf("Wheel %d: %s\n", i+1, wheel.Type())
    }

    // Ketuk dan buka pintu
    fmt.Println("Knocking and opening right door:")
    myCar.RightDoor.Knock()
    myCar.RightDoor.Open()

    fmt.Println("Knocking and opening left door:")
    myCar.LeftDoor.Knock()
    myCar.LeftDoor.Open()
}
