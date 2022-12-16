package runtime

import (
	"cloud.google.com/go/firestore"
	"context"
	. "github.com/RecepieApp/server/db"
	. "github.com/RecepieApp/server/helpers"
	"log"
)

func Start(client *firestore.Client) error {
	ctx := context.Background()
	r := Recepies("Sweedish meatballs", "https://therecipecritic.com/the-best-swedish-meatballs/")
	r.Modiify(4)

	AddRecepie(client, Recepie(r))
	err := DeleteRecepie(ctx, client, Recepie(r))
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("Couldn't delete : %s", err)
	}
	//err = UpdateRecepie(ctx, client, Recepie(r))
	//if err != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred in updating: %s", err)
	//}
	//ReadCollection(ctx, client)
	return nil
}
