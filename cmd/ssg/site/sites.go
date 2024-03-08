package site

import "code.ewintr.nl/shitty-ssg/pkg/adoc"

var (
	SITE_CONFIG_EWNL = &SiteConfig{
		ID:            SITE_EWNL,
		BaseURL:       "https://erikwinter.nl",
		PathsWithKind: true,
		TemplateConfigs: []*TemplateConfig{
			{
				Name:          "home",
				TemplateNames: []string{"list", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderEWNLHome,
			},
			{
				Name:          "listings",
				TemplateNames: []string{"list", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderEWNLListings,
			},
			{
				Name:          "archive",
				TemplateNames: []string{"archive", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderEWNLArchive,
			},
			{
				Name:          "static",
				TemplateNames: []string{"static", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderEWNLStaticPages,
			},
			{
				Name:          "posts",
				TemplateNames: []string{"post", "head", "menu"},
				TemplateExt:   "gohtml",
				Render:        renderEWNLPosts,
			},
			{
				Name:          "rss",
				TemplateNames: []string{"rss"},
				TemplateExt:   "goxml",
				Render:        renderEWNLRSS,
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

	SITE_CONFIG_VKVNL = &SiteConfig{
		ID:            SITE_VKVNL,
		BaseURL:       "https://vrijkorteverhalen.nl",
		PathsWithKind: false,
		TemplateConfigs: []*TemplateConfig{
			{
				Name:          "post",
				TemplateNames: []string{"post"},
				TemplateExt:   "gohtml",
				Render:        renderVKVNLPosts,
			},
			{
				Name:          "rss",
				TemplateNames: []string{"rss"},
				TemplateExt:   "goxml",
				Render:        renderVKVNLRSS,
			},
		},
		KindMap: map[adoc.Kind]Kind{
			adoc.KIND_VKV: KIND_STORY,
		},
	}
)
