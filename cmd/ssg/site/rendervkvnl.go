package site

import (
	"os"
	"path/filepath"
	"text/template"
	"time"
)

func renderVKVNLPosts(targetPath string, tpl *template.Template, posts Posts, _ []*StaticPage) error {
	last, first := 0, len(posts)-1 // posts are sorted in reverse order
	for i, post := range posts {
		pData := post.HTMLPost()
		if pData.Slug == "" {
			return ErrInvalidPost
		}

		data := struct {
			Slug         string
			Title        string
			DateLong     string
			DateShort    string
			Content      string
			PreviousLink string
			NextLink     string
		}{
			Slug:      pData.Slug,
			Title:     pData.Title,
			DateLong:  pData.DateLong,
			DateShort: pData.DateShort,
			Content:   pData.Content,
		}

		path := targetPath
		if i != first {
			data.PreviousLink = posts[i+1].Link()
		}
		if i != last {
			data.NextLink = posts[i-1].Link()
			if i == last+1 {
				data.NextLink = "/"
			}
			path = filepath.Join(targetPath, post.Year(), data.Slug)
		}
		if i == last-1 {
		}

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

func renderVKVNLRSS(targetPath string, tpl *template.Template, posts Posts, _ []*StaticPage) error {
	rssPath := filepath.Join(targetPath, "index.xml")
	rssFile, err := os.Create(rssPath)
	if err != nil {
		return err
	}
	defer rssFile.Close()

	var xmlPosts []*XMLPost
	for _, p := range posts.RemoveKind(KIND_NOTE).Limit(10) {
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
