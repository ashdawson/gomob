package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/notif"
	"io/ioutil"
	"os"
)

var config Config

type Config struct {
	BaseBranchName string `json:"BaseBranchName"`
	BaseRemoteName string `json:"BaseRemoteName"`
	BranchName     string `json:"BranchName"`
	RemoteName     string `json:"RemoteName"`
	CommitMessage  string `json:"CommitMessage"`
	IDE            string `json:"IDE"`
	Mob            string `json:"Mob"`
	TimeLimit      int
}

func NewConfig() Config {
	return Config{
		"master",
		"origin",
		"",
		"",
		"WIP - [MOB] ",
		"",
		"",
		10,
	}
}

func createConfig() {
	config = NewConfig()
	config.RemoteName, config.BranchName = getBranchDetails()
	config.IDE, _ = notif.List("Please select your IDE", []string{"phpstorm", "vscode"})

	possibleTeam := getPossibleTeam()
	if len(possibleTeam) > 0 {
		config.Mob, _ = notif.MultiList("Please select your team", possibleTeam)
	}

	_, err := os.Create(settings.ConfigFile)
	check(err)

	config.save()
}

func checkConfig() {
	err := ReadConfig(settings.ConfigFile)
	if err != nil {
		createConfig()
	}
}

func ReadConfig(filename string) error {
	file, _ := ioutil.ReadFile(filename)

	err := json.Unmarshal([]byte(file), &config)
	return err
}

func (config Config) save() {
	file, _ := json.MarshalIndent(config, "", " ")
	err := ioutil.WriteFile(settings.ConfigFile, file, 0644)
	check(err)
}

