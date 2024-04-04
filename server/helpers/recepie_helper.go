package helpers

func Recepies(title, url, imageUrl string) Recepie {
	return Recepie{
		Made:     false,
		Rating:   0,
		Title:    title,
		Url:      url,
		ImageUrl: imageUrl,
	}
}

func (r *Recepie) Modify(rating int) {
	*r = Recepie{
		Made:     r.Made,
		Rating:   rating,
		Title:    r.Title,
		Url:      r.Url,
		ImageUrl: r.ImageUrl,
	}

}
