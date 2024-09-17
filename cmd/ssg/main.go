package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-mod.ewintr.nl/shitty-ssg/cmd/ssg/site"
)

var (
	siteName  = flag.String("site", "ewnl", "site id, either 'ewnl' or 'vkvnl'")
	resources = flag.String("resources", "./resources", "folder with templates and other resources")
	content   = flag.String("content", "./content,/projectx", "comma separated list of folders search for content")
	statics   = flag.String("statics", "./statics", "folder with static content")
	public    = flag.String("public", "./public", "target folder for generated site")
)

func main() {
	flag.Parse()
	if *siteName == "" || *resources == "" || *content == "" || *public == "" || *statics == "" {
		log.Fatal("missing parameter")
	}

	var siteId site.SiteID
	switch *siteName {
	case "ewnl":
		siteId = site.SITE_EWNL
	case "vkvnl":
		siteId = site.SITE_VKVNL
	default:
		log.Fatal(errors.New("unknown site"))
	}

	// initialize site
	config, err := site.NewSiteConfig(siteId)
	if err != nil {
		log.Fatal(err)
	}
	s, err := site.New(config, *resources)
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
