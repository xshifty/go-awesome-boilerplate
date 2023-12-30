package components

import "fmt"

type link struct {
	label      string
	pageUrl    string
	contentUrl string
}

func e(a, b int) string {
	return fmt.Sprint(a / b)
}

func GetLinks() []link {
	return []link{
		{"Home", "/", "/pages/home"},
		{"About Us", "/about-us", "/pages/about-us"},
	}
}
