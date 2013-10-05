package xbmctoollib

import (
	"path"
	"regexp"
	"strings"
)

//this is a horrible thing.
//split and joined on |
var filenameCrapRegex = []string{
	`studio\.ghibli\.[0-9][0-9]\.`,  //my ghibli rips are like this...
	`harry\.potter\.0[1-7][ab]?`,    //my harry potters...
	`\([^\)]+\)`,                    //anything in brackets
	`\.`, `,`, `_`, `-`, `\[`, `\]`, //single character stuff
	`retail.*`, `dvd.*`, `[bh][rd]rip.*`, `webrip.*`, `r[56].*`, `bluray.*`, //some quality types
	`720p.*`, `1080p.*`, `480p.*`, //hd specs
	`divx.*`, `xvid.*`, `ac3.*`, `x264.*`, //codec
	`extended`, `limited`, `re(rip|pack)`, //other modifiers
	`cd[0-9].*`, //,parts
	`[0-9]{4}`,  //a year
}
var stripCrap *regexp.Regexp

func init() {
	stripCrap = regexp.MustCompile(`(` + strings.Join(filenameCrapRegex, "|") + `)`)
}

func getTitle(s string) string {
	s = s[:len(s)-len(path.Ext(s))]
	s = strings.ToLower(s)
	s = stripCrap.ReplaceAllLiteralString(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}
