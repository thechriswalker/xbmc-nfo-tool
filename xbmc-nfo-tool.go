package main

import (
	lib "./lib"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	DIRECTORY = flag.String("dir", ".", "directory to scan")
	CONFIG    = flag.String("config", "config.toml", "config file")
	OVERWRITE = flag.Bool("nfo:overwrite", false, "if true, we ignore existing nfo's! OVERWRITES DATA!")
	MODE      = flag.String("mode", "list", "action to perform")
)

//This is the XBMC api key for themoviedb.
//freely available in their own source code
//as we are using TMDB on behalf of XBMC anyway, I don't see the harm in using this.
//@see https://github.com/xbmc/xbmc/blob/master/addons/metadata.common.themoviedb.org/tmdb.xml
const TMDB_API_KEY = "57983e31fb435df4df77afb854740ea9"

func main() {
	flag.Parse()

	dir, err := filepath.Abs(*DIRECTORY)
	if err != nil {
		log.Fatalln("Dir Error:", err)
	}

	var runner lib.MovieFunc

	switch *MODE {
	case "list":
		log.Println("Listing Absent NFOs:")
		runner = lib.NewAbsentNFOList()
	case "cli":
		log.Println("Starting CLI Mode")
		runner = lib.NewCliInterface(lib.NewTmdbMovieSearch(TMDB_API_KEY), *OVERWRITE)
	default:
		fmt.Println("Unkown Mode:", *MODE)
		fmt.Println("Command Line Flags:\n")
		flag.PrintDefaults()
		fmt.Println(help)
		os.Exit(1)
	}
	lib.ScanForMovies(dir, runner)
}

const help = `
Available Modes:
  list:    Lists all absent NFO files, i.e. Movies without NFOs
  cli:     Update NFOs on the console
`
