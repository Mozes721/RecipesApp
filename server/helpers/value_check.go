package helpers

import (
	"encoding/json"
	"reflect"

	"github.com/RecepieApp/server/db"
)

func GetFieldTitle(e *db.Recepie, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
func SetToMap(e *db.Recepie) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(e)
	json.Unmarshal(inrec, &inInterface)

	return inInterface

}
