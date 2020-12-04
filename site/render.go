package site

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const dirMode = os.ModeDir | 0755

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

func renderStaticPages(targetPath string, tpl *template.Template, statics []*StaticPage) error {
	for _, static := range statics {
		destPath := filepath.Join(targetPath, static.Name)
		if err := os.MkdirAll(destPath, dirMode); err != nil {
			return err
		}
		pageFile, err := os.Create(filepath.Join(destPath, "index.html"))
		if err != nil {
			return err
		}
		defer pageFile.Close()

		mainHTML, err := ioutil.ReadFile(filepath.Join(static.Path, "main.html"))
		if err != nil {
			return err
		}

		data := struct {
			Title string
			Main  string
		}{
			Title: strings.Title(static.Name),
			Main:  string(mainHTML),
		}
		if err := tpl.Execute(pageFile, data); err != nil {
			return err
		}

		if err := copyFiles(filepath.Join(static.Path, "resources", "*"), destPath); err != nil {
			return err
		}
	}
	return nil
}

func renderHome(targetPath string, tpl *template.Template, posts Posts) error {
	data := struct {
		Title     string
		Summaries []*HTMLSummary
	}{
		Title:     "Recent",
		Summaries: posts.HTMLSummaries(),
	}

	hPath := filepath.Join(targetPath, "index.html")
	homeFile, err := os.Create(hPath)
	if err != nil {
		return err
	}
	defer homeFile.Close()

	return tpl.Execute(homeFile, data)
}

func renderArchive(targetPath string, tpl *template.Template, title string, posts Posts) error {
	archPath := filepath.Join(targetPath, "archive")
	if err := os.MkdirAll(archPath, dirMode); err != nil {
		return err
	}
	archFile, err := os.Create(filepath.Join(archPath, "index.html"))
	if err != nil {
		return err
	}

	type link struct {
		Name string
		Link string
	}

	tags := []link{}
	for _, tag := range posts.TagList() {
		tags = append(tags, link{
			Name: tag,
			Link: fmt.Sprintf("%s/", path.Join("/tags", tag)),
		})
	}

	yearLinks := map[Kind][]link{
		KIND_ARTICLE: {},
		KIND_NOTE:    {},
		KIND_STORY:   {},
	}
	for kind := range yearLinks {
		for _, year := range posts.FilterByKind(kind).YearList() {
			yearLinks[kind] = append(yearLinks[kind], link{
				Name: year,
				Link: fmt.Sprintf("%s/", path.Join("/", pluralKind[kind], year)),
			})
		}
	}

	data := struct {
		Title        string
		Tags         []link
		ArticleYears []link
		NoteYears    []link
		StoryYears   []link
	}{
		Title:        title,
		Tags:         tags,
		ArticleYears: yearLinks[KIND_ARTICLE],
		NoteYears:    yearLinks[KIND_NOTE],
		StoryYears:   yearLinks[KIND_STORY],
	}

	return tpl.Execute(archFile, data)
}

func renderListings(targetPath string, tpl *template.Template, posts Posts) error {
	for _, kind := range []Kind{KIND_NOTE, KIND_STORY, KIND_ARTICLE} {
		for _, year := range posts.FilterByKind(kind).YearList() {
			title := fmt.Sprintf("%s in %s", strings.Title(pluralKind[kind]), year)
			summaries := posts.FilterByKind(kind).FilterByYear(year).HTMLSummaries()
			path := filepath.Join(targetPath, pluralKind[kind], year)
			if err := renderListing(path, tpl, title, summaries); err != nil {
				return err
			}
		}
	}

	for _, tag := range posts.TagList() {
		title := fmt.Sprintf("Posts Tagged with \"%s\"", tag)
		summaries := posts.FilterByTag(Tag(tag)).HTMLSummaries()
		path := filepath.Join(targetPath, "tags", tag)
		if err := renderListing(path, tpl, title, summaries); err != nil {
			return err
		}
	}

	return nil
}

func renderListing(path string, tpl *template.Template, title string, summaries []*HTMLSummary) error {
	data := struct {
		Title     string
		Summaries []*HTMLSummary
	}{
		Title:     title,
		Summaries: summaries,
	}
	if err := os.MkdirAll(path, dirMode); err != nil {
		return err
	}
	lPath := filepath.Join(path, "index.html")
	f, err := os.Create(lPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, data)
}

func renderPosts(targetPath string, tpl *template.Template, posts Posts) error {
	for _, post := range posts {
		data := post.HTMLPost()
		if data.Slug == "" {
			return ErrInvalidPost
		}

		path := filepath.Join(targetPath, pluralKind[post.Kind], post.Year(), data.Slug)
		if err := os.MkdirAll(path, dirMode); err != nil {
			return err
		}

		nPath := filepath.Join(path, "index.html")
		f, err := os.Create(nPath)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := tpl.Execute(f, data); err != nil {
			return err
		}
	}

	return nil
}

func renderRSS(targetPath string, tpl *template.Template, posts Posts) error {
	rssPath := filepath.Join(targetPath, "index.xml")
	rssFile, err := os.Create(rssPath)
	if err != nil {
		return err
	}
	defer rssFile.Close()

	var xmlPosts []*XMLPost
	for _, p := range posts.Limit(10) {
		xmlPosts = append(xmlPosts, p.XMLPost())
	}

	data := struct {
		DateFormal string
		Posts      []*XMLPost
	}{
		DateFormal: time.Now().Format(time.RFC1123Z),
		Posts:      xmlPosts,
	}
	return tpl.Execute(rssFile, data)
}

func copyFiles(srcPattern, destPath string) error {
	filePaths, err := filepath.Glob(srcPattern)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(destPath, dirMode); err != nil {
		return err
	}

	for _, fPath := range filePaths {
		destFPath := filepath.Join(destPath, filepath.Base(fPath))
		content, err := ioutil.ReadFile(fPath)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(destFPath, content, 0644); err != nil {
			return err
		}
	}

	return nil
}
