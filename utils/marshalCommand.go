package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MarshalCommand(r *http.Request, c interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}
