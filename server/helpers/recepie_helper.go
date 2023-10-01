package helpers

func Recepies(title string, url string) RecepieHelper {
	return RecepieHelper{
		Made:   false,
		Rating: 0,
		Title:  title,
		Url:    url,
	}
}

func (r *RecepieHelper) Modify(rating int) {
	*r = RecepieHelper{
		Made:   r.Made,
		Rating: rating,
		Title:  r.Title,
		Url:    r.Url,
	}

}
