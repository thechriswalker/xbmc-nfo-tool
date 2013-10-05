package xbmctoollib

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//this is the type our scanner expects.
//allows you to do "something" with the result
//the path, will be the path to the "nfo",
//whether it exists or not.
type MovieFunc func(path string, hasNfo bool)

//scans a dir for movies and checks for the nfo.
func ScanForMovies(dir string, fn MovieFunc) error {
	var scanWalker filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		//is this a movie, and if part of a set, is it the first one!
		//if so is the nfo file present.
		if nfopath, ok := getNfoPath(path); ok {
			fn(nfopath, fileExists(nfopath))
		}
		return nil
	}
	return filepath.Walk(dir, scanWalker)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

//checks whether this is a valid movie file and works
//out the path to the nfo is it is. if it's not a valid
//movie file, "ok" will be false
//
//Check is fairly simple. look for extension in list of
//extensions.
//if present check for "cdX" suffix, if present only good
//if "cd1"
//after removing both extension and maybe "cd1" add ".nfo"
func getNfoPath(path string) (nfopath string, ok bool) {
	if m := movieRegexp.FindStringSubmatch(path); m != nil {
		//check the "cd" bit is either empty or "cd1"
		if m[1] == "" || strings.HasSuffix(m[1], "cd1") {
			//all good, the NFO path will be the match removed and .nfo added
			return path[:len(path)-len(m[0])] + ".nfo", true
		}
	}
	// fmt.Println("Miss:", path)
	return "", false
}

var movieRegexp = regexp.MustCompile(`([\.-]cd[0-9]+)?\.(avi|mkv|mp4|mov|m4v)$`)
