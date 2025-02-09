package main

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {
	// Initialize periph.io
	if _, err := host.Init(); err != nil {
		fmt.Println("Failed to initialize periph:", err)
		return
	}

	// Define GPIO pins
	trig := gpioreg.ByName("GPIO23")
	echo := gpioreg.ByName("GPIO24")

	if trig == nil || echo == nil {
		fmt.Println("Failed to find GPIO pins")
		return
	}

	fmt.Println("Distance Measurement In Progress")

	// Set up pins
	trig.Out(gpio.Low)
	time.Sleep(2 * time.Second)

	// Send trigger pulse
	trig.Out(gpio.High)
	time.Sleep(10 * time.Microsecond)
	trig.Out(gpio.Low)

	// Wait for echo to start
	start := time.Now()
	for echo.Read() == gpio.Low {
		start = time.Now()
	}

	// Wait for echo to end
	end := time.Now()
	for echo.Read() == gpio.High {
		end = time.Now()
	}

	// Calculate duration and distance
	duration := end.Sub(start).Seconds()
	distance := duration * 17150 // Speed of sound in air (343m/s)

	fmt.Printf("Distance: %.2f cm\n", distance)
}
