package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/notif"
	"io/ioutil"
	"os"
)

var argsMap = map[string]string{}
var mobSettingsFile = "mobSettings.json"
var settings Settings

type Settings struct {
	BaseBranchName string `json:"BaseBranchName"`
	BaseRemoteName string `json:"BaseRemoteName"`
	BranchName     string `json:"BranchName"`
	RemoteName     string `json:"RemoteName"`
	CommitMessage  string `json:"CommitMessage"`
	IDE            string `json:"IDE"`
	TimeLimit      int
}

func setup() {
	checkSettings()
	readCommandLineArguments()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createSettings() {
	_, err := os.Create(mobSettingsFile)
	check(err)
	remoteName, branchName := getBranchDetails()
	getIDE, _ := notif.List("Please select your IDE", []string{"phpstorm", "vscode"})
	settings = Settings{
		"master",
		"origin",
		branchName,
		remoteName,
		"WIP - [MOB] ",
		getIDE,
		15,
	}
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
