package site

import "git.sr.ht/~ewintr/shitty-ssg/pkg/adoc"

var (
	SITE_CONFIG_EWNL = &SiteConfig{
		ID:      SITE_EWNL,
		BaseURL: "https://erikwinter.nl",
		TemplateConfigs: []*TemplateConfig{
			{
				Name:          "home",
				TemplateNames: []string{"list", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderHome,
			},
			{
				Name:          "listings",
				TemplateNames: []string{"list", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderListings,
			},
			{
				Name:          "archive",
				TemplateNames: []string{"archive", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderArchive,
			},
			{
				Name:          "static",
				TemplateNames: []string{"static", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderStaticPages,
			},
			{
				Name:          "posts",
				TemplateNames: []string{"post", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderPosts,
			},
			{
				Name:          "rss",
				TemplateNames: []string{"rss"},
				TemplateExt:   "goxml",
				Render:        renderRSS,
			},
		},
		KindMap: map[adoc.Kind]Kind{
			adoc.KIND_NOTE:     KIND_NOTE,
			adoc.KIND_VKV:      KIND_STORY,
			adoc.KIND_ESSAY:    KIND_ARTICLE,
			adoc.KIND_TUTORIAL: KIND_ARTICLE,
			adoc.KIND_ARTICLE:  KIND_ARTICLE,
		},
	}
	SITE_CONFIG_VKVNL = &SiteConfig{}
)
