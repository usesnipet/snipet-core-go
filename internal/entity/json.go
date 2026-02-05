package entity

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/datatypes"
)

type JSONMap = datatypes.JSONMap

type EncryptedJSON map[string]any

func (e EncryptedJSON) Value() (driver.Value, error) {
	// TODO: encrypt before saving
	return json.Marshal(e)
}

func (e *EncryptedJSON) Scan(value any) error {
	// TODO: decrypt after reading
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, e)
}
