package components

type link struct {
	label      string
	pageUrl    string
	contentUrl string
}

func GetLinks() []link {
	return []link{
		{"Home", "/", "/pages/home"},
		{"About Us", "/about-us", "/pages/about-us"},
	}
}
