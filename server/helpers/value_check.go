package helpers

import (
	"reflect"

	"github.com/RecepieApp/server/models"
)

func GetFieldTitle(e *models.Recepie, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
