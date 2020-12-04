package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~ewintr/shitty-ssg/site"
)

var (
	resources = flag.String("resources", "./resources", "folder with templates and other resources")
	content   = flag.String("content", "./content,/projectx", "comma separated list of folders search for content")
	public    = flag.String("public", "./public", "target folder for generated site")
)

func main() {
	flag.Parse()
	if *resources == "" || *content == "" || *public == "" {
		log.Fatal("missing parameter")
	}

	// initialize site
	s, err := site.New(*resources)
	if err != nil {
		log.Fatal(err)
	}

	// add content
	for _, cp := range strings.Split(*content, ",") {
		log.Printf("checking %s for content\n", cp)
		if err := filepath.Walk(cp, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".adoc" {
				s.AddFilePost(path)
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}

	// render site
	if err := s.RenderHTML(*public); err != nil {
		log.Fatal(err)
	}
}
