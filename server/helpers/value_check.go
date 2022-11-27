package helpers

import (
	"reflect"

	"github.com/RecepieApp/server/db"
)

func GetFieldTitle(e *db.Recepie, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
