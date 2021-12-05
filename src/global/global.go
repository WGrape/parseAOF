package global

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

const PatternOfAOFFile = "^aof\\.split_[a-z]+$"
const PatternOfLineCmdStart = "^\\*\\d$"
const PatternOfLineArgLen = "^\\$\\d$"

const MatchTypeCmdStart = 1
const MatchTypeArgLen = 2
const MatchTypeArgRaw = 3

const EmptyString = ""
const WhitespaceString = " "

const ExitCodeFatal = 1
const DefaultFileMode = 0644
const Delta = 1

const ConfigDefaultDebug = false
const ConfigDefaultMaxRoutines = 1024

var RootDir, _ = os.Getwd()
var DataDir = path.Clean(fmt.Sprintf("%s/%s", RootDir, "../data"))
var ConfigFile = path.Clean(fmt.Sprintf("%s/%s", RootDir, "../config/config.yml"))

var Config AppConfig

func init() {
	bytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		LogError(fmt.Sprintf("Failed to read config: %s\n", err.Error()))
		Config.SetDebug(ConfigDefaultDebug)
		Config.SetMaxRoutines(ConfigDefaultMaxRoutines)
	} else {
		err = yaml.Unmarshal(bytes, &Config)
		if err != nil {
			LogError(fmt.Sprintf("Failed to create config: %s", err.Error()))
			os.Exit(ExitCodeFatal)
		}
	}
}

func IsAOFFile(fileName string) bool {
	matched, _ := regexp.MatchString(PatternOfAOFFile, fileName)
	return matched
}

func GetSplitFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", DataDir, fileName)
}

func GetParsedFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s.parsed", DataDir, fileName)
}

func LogDebug(text string) {
	if Config.GetDebug() {
		fmt.Printf(text)
	}
}

func LogInfo(text string) {
	fmt.Printf(text)
}

func LogError(text string) {
	fmt.Printf(text)
}
