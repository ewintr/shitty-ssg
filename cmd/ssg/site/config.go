package site

import (
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"go-mod.ewintr.nl/shitty-ssg/pkg/adoc"
)

var (
	ErrUnknownSiteID = errors.New("unknown site id")
)

type SiteID string

const (
	SITE_EWNL  = SiteID("ewnl")
	SITE_VKVNL = SiteID("vkvnl")
)

type TemplateConfig struct {
	Name          string
	TemplateNames []string
	TemplateExt   string
	Template      *template.Template
	Render        func(targetPath string, tpl *template.Template, posts Posts, staticPages []*StaticPage) error
}

type SiteConfig struct {
	ID              SiteID
	BaseURL         string
	PathsWithKind   bool
	TemplateConfigs []*TemplateConfig
	StaticPages     []*StaticPage
	KindMap         map[adoc.Kind]Kind
}

func NewSiteConfig(id SiteID) (*SiteConfig, error) {
	var config *SiteConfig

	switch id {
	case SITE_EWNL:
		config = SITE_CONFIG_EWNL
	case SITE_VKVNL:
		config = SITE_CONFIG_VKVNL
	default:
		return &SiteConfig{}, ErrUnknownSiteID
	}

	return config, nil
}

func (sc *SiteConfig) ParseTemplates(tplPath string) error {
	for _, tplConf := range sc.TemplateConfigs {
		var tFiles []string
		for _, tName := range tplConf.TemplateNames {
			tFiles = append(tFiles, filepath.Join(tplPath, fmt.Sprintf("%s.%s", tName, tplConf.TemplateExt)))
		}
		tpl, err := template.ParseFiles(tFiles...)
		if err != nil {
			return err
		}
		tplConf.Template = tpl
	}

	return nil
}

func (sc *SiteConfig) MapKind(docKind adoc.Kind) Kind {
	siteKind, ok := sc.KindMap[docKind]
	if !ok {
		return KIND_INVALID
	}

	return siteKind
}
