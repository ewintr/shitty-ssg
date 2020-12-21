package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~ewintr/shitty-ssg/cmd/ssg/site"
)

var (
	resources = flag.String("resources", "./resources", "folder with templates and other resources")
	content   = flag.String("content", "./content,/projectx", "comma separated list of folders search for content")
	statics   = flag.String("statics", "./statics", "folder with static content")
	public    = flag.String("public", "./public", "target folder for generated site")
)

func main() {
	flag.Parse()
	if *resources == "" || *content == "" || *public == "" || *statics == "" {
		log.Fatal("missing parameter")
	}

	// initialize site
	s, err := site.New(*resources)
	if err != nil {
		log.Fatal(err)
	}

	// add statics
	staticNames, err := ioutil.ReadDir(*statics)
	if err != nil {
		log.Fatal(err)
	}
	for _, sn := range staticNames {
		if sn.IsDir() {
			s.AddStaticPage(filepath.Join(*statics, sn.Name()))
		}
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
