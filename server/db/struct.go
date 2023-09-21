package db

type Recepie struct {
	Made   bool
	Rating int
	Title  string
	Url    string
}

type HTTPError struct {
	Err      error
	Code     int
	FlashMsg string
	URL      string
}
