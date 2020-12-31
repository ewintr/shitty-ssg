package site

import (
	"io/ioutil"
	"path/filepath"

	"git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"
)

type StaticPage struct {
	Name string
	Path string
}

type Site struct {
	resourcesPath string
	config        *SiteConfig
	posts         Posts
	staticPages   []*StaticPage
}

func New(config *SiteConfig, resourcesPath string) (*Site, error) {
	if err := config.ParseTemplates(filepath.Join(resourcesPath, "template")); err != nil {
		return &Site{}, err
	}

	return &Site{
		resourcesPath: resourcesPath,
		config:        config,
		posts:         Posts{},
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
	post := NewPost(s.config, adoc.New(string(content)))
	if post.Kind != KIND_INVALID {
		s.posts = append(s.posts, post)
	}

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

	for _, tplConf := range s.config.TemplateConfigs {
		if err := tplConf.Render(targetPath, tplConf.Template, posts, s.staticPages); err != nil {
			return err
		}
	}

	return nil
}