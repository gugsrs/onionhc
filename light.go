package main

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

const LED_PIN = "17"
const RELE_PIN = "1"

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	led := gpioreg.ByName(LED_PIN)
	rele := gpioreg.ByName(RELE_PIN)
	defer led.Halt()
	defer rele.Halt()

	info := accessory.Info{
		Name:         "Light Bulb",
		Manufacturer: "Skr",
	}

	acc := accessory.NewLightbulb(info)

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			rele.Out(gpio.High)
			led.Out(gpio.Low)
			turnLightOn()
		} else {
			rele.Out(gpio.Low)
			led.Out(gpio.High)
			turnLightOff()
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
