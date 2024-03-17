package main

import (
	"Redeemer/Core/Client"
	"Redeemer/Core/Helpers"
	"time"
)

var configuration, _ = Helpers.LoadSettings()

func main() {
	Helpers.UpdateTitle("Initializing...")
	go func() {
		for {
			time.Sleep(time.Minute * 30)
		}
	}()

	Helpers.ClearScreen()
	Client.Redeemer()
}
