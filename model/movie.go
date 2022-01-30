package model

type Movie struct {
	Id           int
	Name         string
	Year         string
	Director     string
	Screenwriter string
	Starring     string
	Type         string
	Country      string
	Language     string
	Length       string
	IMDb         string
	StarNum      int
	Score        string
	Star         string
	HaveWatched  int
	WantToWatch  int
	Synopsis     string
	URL          string
	PeopleURL    string
	NameInfo     string
	CoverInfo    string
}

type MovieBrief struct {
	Id   int
	Name string
	URL  string
}
