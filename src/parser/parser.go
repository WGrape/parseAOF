package parser

import (
	"parseAOF/src/global"
	"regexp"
)

func MatchLine(content string) int {
	var matched = false
	if matched, _ = regexp.MatchString(global.PatternOfLineCmdStart, content); matched {
		return global.MatchTypeCmdStart
	} else if matched, _ = regexp.MatchString(global.PatternOfLineArgLen, content); matched {
		return global.MatchTypeArgLen
	} else {
		return global.MatchTypeArgRaw
	}
}
