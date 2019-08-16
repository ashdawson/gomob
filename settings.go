package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/notif"
	"io/ioutil"
	"os"
)

var mobSettingsFile = "mobSettings.json"
var settings Settings

type Settings struct {
	BaseBranchName string `json:"BaseBranchName"`
	BaseRemoteName string `json:"BaseRemoteName"`
	BranchName     string `json:"BranchName"`
	RemoteName     string `json:"RemoteName"`
	CommitMessage  string `json:"CommitMessage"`
	IDE            string `json:"IDE"`
	Mob            string `json:"Mob"`
	Debug          bool
	TimeLimit      int
}


func createSettings() {
	remoteName, branchName := getBranchDetails()
	var getIDE string
	var getMob string
	getIDE, _ = notif.List("Please select your IDE", []string{"phpstorm", "vscode"})

	settings = Settings{
		"master",
		"origin",
		branchName,
		remoteName,
		"WIP - [MOB] ",
		getIDE,
		getMob,
		false,
		10,
	}

	possibleTeam := getPossibleTeam()
	if len(possibleTeam) > 0 {
		settings.Mob, _ = notif.MultiList("Please select your team", possibleTeam)
	}

	_, err := os.Create(mobSettingsFile)
	check(err)

	saveSettings()
}

func saveSettings() {
	file, _ := json.MarshalIndent(settings, "", " ")
	err := ioutil.WriteFile(mobSettingsFile, file, 0644)
	check(err)
}

func checkSettings() {
	file, err := os.Open(mobSettingsFile)
	file.Close()
	if err != nil {
		createSettings()
	} else {
		readSettings()
	}
}

func readSettings() {
	file, _ := ioutil.ReadFile(mobSettingsFile)
	settings = Settings{}

	err := json.Unmarshal([]byte(file), &settings)
	check(err)
}

func (settings *Settings) updateSetting(setting string, value string) {
	saveSettings()
}
