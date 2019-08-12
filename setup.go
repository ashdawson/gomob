package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/command"
	"io/ioutil"
	"os"
	"strings"
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
	command.Read()
	getArguments()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createSettings() {
	_, err := os.Create(mobSettingsFile)
	check(err)
	branchName, remoteName := getBranchDetails()
	settings = Settings{
		"master",
		"origin",
		branchName,
		remoteName,
		"WIP - [MOB] ",
		"PHPSTORM",
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
	_, err := os.Open(mobSettingsFile)
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

func (currentSettings *Settings) updateSetting(setting string, value string) {
	saveSettings()
}

func getArguments() {
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "--") {
			if _, ok := argsMap[os.Args[i]]; !ok {
				var hasValue = strings.Index(os.Args[i], "=")
				if hasValue > 0 {
					argsMap[os.Args[i][2:hasValue]] = os.Args[i][hasValue:]
				} else {
					argsMap[os.Args[i][2:]] = ""
				}
			}
		}
	}
}
