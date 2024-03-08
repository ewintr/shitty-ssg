package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"code.ewintr.nl/shitty-ssg/cmd/notes/note"
)

func main() {
	notesPath := os.Getenv("NOTES_PATH")
	if notesPath == "" {
		log.Fatal("no notes directory to parse")
	}

	if len(os.Args) != 2 {
		log.Fatal("exactly one search term is required as parameter")
	}
	searchTerm := os.Args[1]

	var notes note.Notes
	if err := filepath.Walk(notesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".adoc" {
			if err := notes.AddFileNote(path); err != nil {
				return nil
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	notes = notes.FilterByTerm(searchTerm)
	if len(notes) == 0 {
		fmt.Println("Found nothing.")

		return
	}

	for i, n := range notes {
		fmt.Printf("%d) %s\n", i, n.Title)
	}
	reader := bufio.NewReader(os.Stdin)
	r, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	c := string(r)
	switch c {
	case "q":
		return
	default:
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		if i < 0 || i >= len(notes) {
			fmt.Println("number out of range")
			return
		}
		fmt.Printf("\n\n%s\n\n%s\n\n", notes[i].Title, notes[i].Content)
	}
}
