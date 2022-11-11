package helpers

import (
	"fmt"
	"reflect"

	"github.com/RecepieApp/server/db"
)

func GetFieldTitle(e *db.Recepie, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
func GetFieldDBTitles(recepie []interface{}, field string) string {

	for k, v := range recepie {
		fmt.Println(k, v)
	}
	// r := reflect.ValueOf(recepie)
	// f := reflect.Indirect(r).FieldByName(field)
	return "asds"

}
