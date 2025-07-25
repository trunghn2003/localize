package localize

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

func (tf TranslatableField) Value() (driver.Value, error) {
	return json.Marshal(tf)
}

func (tf *TranslatableField) Scan(value interface{}) error {
	if value == nil {
		*tf = make(map[string]string)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to assert DB value to []byte")
	}
	return json.Unmarshal(bytes, tf)
}
