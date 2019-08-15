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
	Mob            string `json:"Mob"`
	Debug          bool
	TimeLimit      int
}

func setup() {
	//runChecks()
	checkSettings()
	readCommandLineArguments()
}

func createSettings() {
	_, err := os.Create(mobSettingsFile)
	check(err)
	remoteName, branchName := getBranchDetails()
	var getIDE string
	var getMob string
	if notif.CanUse() {
		getIDE, _ = notif.List("Please select your IDE", []string{"phpstorm", "vscode"})
		getMob, _ = notif.MultiList("Please select your team", getPossibleTeam())
	} else {
		getIDE = AskInput("Pleases enter your IDE:", []string{"phpstorm", "vscode"})
		getMob = AskInput("Please select your team:", getPossibleTeam())
	}

	settings = Settings{
		"master",
		"origin",
		branchName,
		remoteName,
		"WIP - [MOB] ",
		getIDE,
		getMob,
		false,
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

func runChecks() {
	file, err := os.Open(".git/info/exclude")
	file.Close()
	checkSay(err, "git has not been added to this directory")
}
