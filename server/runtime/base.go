package runtime

import (
	"cloud.google.com/go/firestore"
	//"context"
	. "github.com/RecepieApp/server/db"
	. "github.com/RecepieApp/server/helpers"
)

func Start(client *firestore.Client) error {
	r := CreateRecepie("Sweedish meatballs", "https://therecipecritic.com/the-best-swedish-meatballs/")
	AddRecepie(client, Recepie(r))
	//ReadCollection(context.Background(), client)

	return nil
}
