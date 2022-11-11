package helpers

import (
	"encoding/json"

	"github.com/RecepieApp/server/db"
)

func (r *db.Data) ToJson(recepies interface{}) error {
	if err := json.Marshaler(recepies, &r.Recepies); err != nil {
		return err
	}

	return nil
}

func ToJson(recepies map[string]interface{}) (json, error) {
	json := json.Marshaler(recepies)
	return json, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
