package models

// Site представляет сайт, который необходимо проверять.
type Site struct {
	ID            int
	Name          string
	URL           string
	CheckType     string
	CheckInterval int
}
