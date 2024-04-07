package controller

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader, v any) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(v)
	if err != nil {
        return err
    }
	return nil
}