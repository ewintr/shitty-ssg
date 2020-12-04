package site

type HTMLPost struct {
	Slug      string
	Title     string
	DateLong  string
	DateShort string
	Content   string
}

type HTMLSummary struct {
	Link      string
	Title     string
	Language  Language
	DateShort string
	DateLong  string
	Summary   string
}
