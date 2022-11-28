package runtime

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	. "github.com/RecepieApp/server/db"
	. "github.com/RecepieApp/server/helpers"
)

func Start(client *firestore.Client) error {
	ctx := context.Background()
	r := Recepies("Sweedish meatballs", "https://therecipecritic.com/the-best-swedish-meatballs/")
	fmt.Println(r)
	r.Modiify(6)
	fmt.Println(r)
	//AddRecepie(client, Recepie(r))
	//
	//ReadCollection(ctx, client)
	UpdateRecepie(ctx, client, Recepie(r))
	//ReadCollection(ctx, client)
	return nil
}
