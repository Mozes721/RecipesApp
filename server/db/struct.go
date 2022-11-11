package db

// type Recepie map[string]interface{}
type Data struct {
	Recepies *Recepie
}

type Recepie struct {
	Made   bool
	Rating int
	Title  string
	Url    string
}
