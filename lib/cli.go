package xbmctoollib

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func NewCliInterface(s Searcher, overwrite bool) (fn MovieFunc) {
	fn = func(nfoPath string, hasNfo bool) {
		//just output for now!
		if !hasNfo || overwrite {
			if overwrite {
				fmt.Println("Overwriting NFO:", nfoPath)
			} else {
				fmt.Println("Need to Create NFO for:", nfoPath)
			}
			handleCreateNfo(nfoPath, s)
		}
	}
	return
}

func handleCreateNfo(nfo string, s Searcher) {
	nfofile := path.Base(nfo)
	//search for movie, first strip crap...
	fmt.Println("Unstripped:", nfofile)
	title := getTitle(nfofile)
	fmt.Printf("Stripped Title: `%s`\n", title)
	r, err := s.Search(title)
	if err != nil {
		fmt.Println("SearchError:", err)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(r)
			fmt.Println(err)
		}
	}()
	if r.Hits() == 0 {
		fmt.Println("No Results... :(")
		return
	}
	var n int
	if r.Hits() > 1 {
		fmt.Println("Found", r.Hits(), "Results")
		for i := 0; i < r.Hits(); i++ {
			if url, name, _, ok := r.GetResult(i); ok {
				fmt.Println("\t", i, "] => ", name, "("+url+")")
			}
		}
		fmt.Print("\n")
		n = readInputInt(">>>Pick a Number:", 0)
	} else {
		//Only one result. pick it
		n = 0
	}
	urlToWrite, title, _, ok := r.GetResult(n)
	if !ok {
		fmt.Println("Bad Choice... :(")
		return
	}
	fmt.Println("Writing NFO File with url:", urlToWrite)
	//Write to Nfo file (no sets...)
	f, err := os.Create(nfo)
	if err != nil {
		fmt.Println("Could not create file:", nfofile)
		return
	}
	defer f.Close()
	fmt.Fprintln(f, urlToWrite)
}

func readInputInt(prompt string, def int) int {
	fmt.Printf("%s: [%d] ", prompt, def)
	b, err1 := in.ReadBytes('\n')
	n, err2 := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 32)
	if err1 != nil || err2 != nil {
		return def
	}
	return int(n)
}
