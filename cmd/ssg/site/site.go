package site

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"code.ewintr.nl/shitty-ssg/pkg/adoc"
)

const dirMode = os.ModeDir | 0755

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
	doc := adoc.New(string(content))
	if !doc.Public {
		return nil
	}

	post := NewPost(s.config, doc)
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

func resetTarget(targetPath string) error {
	if err := os.RemoveAll(targetPath); err != nil {
		return err
	}

	return os.Mkdir(targetPath, dirMode)
}

func moveResources(targetPath, resourcesPath string) error {
	for _, dir := range []string{"css", "font"} {
		srcPath := filepath.Join(resourcesPath, dir)
		destPath := filepath.Join(targetPath, dir)
		if err := copyFiles(filepath.Join(srcPath, "*"), destPath); err != nil {
			return err
		}
	}

	return nil
}
