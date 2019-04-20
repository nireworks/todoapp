package server

import (
	"encoding/json"
	"fmt"
	"io"
)

const emptyMessagePlaceholder = "an unexpected error has occurred"

type FailResponse struct {
	Error string `json:"error"`
}

func (fr FailResponse) SendJSON(w io.Writer) error {
	if fr.Error == "" {
		fr.Error = emptyMessagePlaceholder
	}

	msg, err := json.Marshal(fr)
	if err != nil {
		return fmt.Errorf("failresponse.Send json marshal: %v", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("failresponse.Send: %v", err)
	}

	return nil
}
