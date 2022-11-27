package helpers

func CreateRecepie(title string, url string) RecepieHelper {
	return RecepieHelper{
		Made:   false,
		Rating: 0,
		Title:  title,
		Url:    url,
	}
}
