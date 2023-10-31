package global

import (
	"fmt"
	"os"
	"regexp"
)

const PatternOfAOFFile = "^aof\\.split_[a-z0-9]+$"
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
var DataDir = ""
var Debug = false

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
	if Debug {
		fmt.Printf(text)
	}
}

func LogInfo(text string) {
	fmt.Printf(text)
}

func LogError(text string) {
	fmt.Printf(text)
}
