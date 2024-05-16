package main

import "fmt"

func main() {
	car := Car{
		Wheels: [4]Wheel{
			{Type: ""},
			{Type: ""},
			{Type: ""},
			{Type: ""},
		},
		Doors: [2]Door{
			{Side: "Kanan"},
			{Side: "Kiri"},
		},
	}

	car.Wheels[0] = Wheel{Type: "Karet"}
	car.Wheels[1] = Wheel{Type: "Kayu"}
	car.Wheels[2] = Wheel{Type: "Besi"}
	car.Wheels[3] = Wheel{Type: "Plastik"}

	for i, wheel := range car.Wheels {
		fmt.Printf("Roda %d diganti dengan ban berjenis %s\n", i+1, wheel.Type)
	}

	for _, door := range car.Doors {
		fmt.Printf("Mengetuk pintu %s berbunyi ", door.Side)
		door.Tap()
		fmt.Printf("Membuka pintu %s berbunyi ", door.Side)
		door.Open()
	}
}
