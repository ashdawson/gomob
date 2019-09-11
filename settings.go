package main

var settings Settings

type Settings struct {
	TimerOnly bool
	Debug     bool
	ConfigFile string
}

func createSettings() {
	settings = newSettings()
}

func newSettings() Settings {
	return Settings{
		false,
		false,
		"mobSettings.json",
	}
}
