package helpers

func Recepies(title string, url string) Recepie {
	return Recepie{
		Made:   false,
		Rating: 0,
		Title:  title,
		Url:    url,
	}
}

func (r *Recepie) Modify(rating int) {
	*r = Recepie{
		Made:   r.Made,
		Rating: rating,
		Title:  r.Title,
		Url:    r.Url,
	}

}
