package site

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

type StaticPage struct {
	Name string
	Path string
}

type Site struct {
	resourcesPath string
	templates     map[string]*template.Template
	posts         Posts
	staticPages   []*StaticPage
}

func New(resourcesPath string) (*Site, error) {
	templates, err := parseTemplates(resourcesPath)
	if err != nil {
		return &Site{}, err
	}

	return &Site{
		resourcesPath: resourcesPath,
		templates:     templates,
		posts:         []Post{},
		staticPages:   []*StaticPage{},
	}, nil
}

func (s *Site) AddStaticPage(staticPath string) {
	s.staticPages = append(s.staticPages, &StaticPage{
		Name: filepath.Base(staticPath),
		Path: staticPath,
	})
}

func (s *Site) AddFilePost(fPath string) error {
	content, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	s.posts = append(s.posts, NewPost(string(content)))

	return nil
}

func (s *Site) AddFolderPost(kind Kind, fPath string) error {
	// TODO implement
	return nil
}

func (s *Site) RenderHTML(targetPath string) error {
	posts := s.posts.Sort()

	if err := resetTarget(targetPath); err != nil {
		return err
	}
	if err := moveResources(targetPath, s.resourcesPath); err != nil {
		return err
	}
	if err := renderStaticPages(targetPath, s.templates["static"], s.staticPages); err != nil {
		return err
	}

	if err := renderArchive(targetPath, s.templates["archive"], "Archive", posts); err != nil {
		return err
	}

	if err := renderHome(targetPath, s.templates["list"], posts.Limit(10)); err != nil {
		return err
	}

	if err := renderListings(targetPath, s.templates["list"], posts); err != nil {
		return err
	}

	if err := renderPosts(targetPath, s.templates["post"], posts); err != nil {
		return err
	}

	if err := renderRSS(targetPath, s.templates["rss"], posts); err != nil {
		return err
	}
	return nil
}

func parseTemplates(resourcesPath string) (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}
	tPath := filepath.Join(resourcesPath, "template")
	for _, tName := range []string{"post", "list", "archive", "static"} {
		var tFiles []string
		for _, tf := range []string{tName, "head", "menu"} {
			tFiles = append(tFiles, filepath.Join(tPath, fmt.Sprintf("%s.gohtml", tf)))
		}
		tpl, err := template.ParseFiles(tFiles...)
		if err != nil {
			return map[string]*template.Template{}, err
		}
		templates[tName] = tpl
	}

	rss, err := template.ParseFiles(filepath.Join(tPath, "rss.goxml"))
	if err != nil {
		return map[string]*template.Template{}, err
	}
	templates["rss"] = rss

	return templates, nil
}
