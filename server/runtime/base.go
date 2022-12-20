package runtime

import (
	"github.com/RecepieApp/server/api"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8000")
	return nil
}

//r.GET("/test", func(ctx *gin.Context) {
//	ctx.JSON(200, gin.H{
//		"Message": "OK!",
//	})
//})

//ctx := context.Background()
//r := Recepies("Sweedish meatballs", "https://therecipecritic.com/the-best-swedish-meatballs/")
//r.Modiify(4)

//AddRecepie(client, Recepie(r))
//err := DeleteRecepie(ctx, client, Recepie(r))
//if err != nil {
//	// Handle any errors in an appropriate way, such as returning them.
//	log.Printf("Couldn't delete : %s", err)
//}
//err = UpdateRecepie(ctx, client, Recepie(r))
//if err != nil {
//	// Handle any errors in an appropriate way, such as returning them.
//	log.Printf("An error has occurred in updating: %s", err)
//}
//ReadCollection(ctx, client)
