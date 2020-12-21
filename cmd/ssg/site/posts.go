package site

import "sort"

type Posts []*Post

func (p Posts) Len() int           { return len(p) }
func (p Posts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Posts) Less(i, j int) bool { return p[i].Date.After(p[j].Date) }

// Sort sorts on reverse chronological order
func (p Posts) Sort() Posts {
	sort.Sort(p)

	return p
}

func (p Posts) FilterByKind(wanted Kind) Posts {
	pList := Posts{}
	for _, post := range p {
		if post.Kind == wanted {
			pList = append(pList, post)
		}
	}

	return pList
}

func (p Posts) FilterByYear(year string) Posts {
	pList := Posts{}
	for _, post := range p {
		if post.Year() == year {
			pList = append(pList, post)
		}
	}

	return pList
}

func (p Posts) FilterByTag(wanted Tag) Posts {
	pList := Posts{}

POSTS:
	for _, post := range p {
		for _, tag := range post.Tags {
			if tag == wanted {
				pList = append(pList, post)

				continue POSTS
			}
		}
	}

	return pList
}

func (p Posts) Limit(limit int) Posts {
	if len(p) <= limit {
		return p
	}

	return p[:limit]
}

func (p Posts) YearList() []string {
	fullList := []string{}
	for _, post := range p {
		fullList = append(fullList, post.Year())
	}
	list := removeDuplicates(fullList)
	sort.Strings(list)

	return list
}

func (p Posts) TagList() []string {
	fullList := []string{}
	for _, post := range p {
		for _, tag := range post.Tags {
			fullList = append(fullList, string(tag))
		}
	}

	list := removeDuplicates(fullList)
	sort.Strings(list)

	return list
}

func removeDuplicates(fullList []string) []string {
	list := []string{}
	for _, item := range fullList {
		isNew := true
		for _, li := range list {
			if item == li {
				isNew = false
			}
		}
		if isNew {
			list = append(list, item)
		}
	}

	return list
}
