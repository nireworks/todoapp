package server_test

import (
	"bytes"
	"testing"
	"todoapp/cmd/server"

	"github.com/stretchr/testify/assert"
)

func TestFailResponse_SendJSON(t *testing.T) {
	tests := []struct {
		name    string
		errMsg  string
		want    string
		wantErr bool
	}{
		{
			"All right",
			"failed to do something",
			"{\"error\":\"failed to do something\"}",
			false,
		},
		{
			"Empty error message",
			"",
			"{\"error\":\"an unexpected error has occurred\"}",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			fr := server.FailResponse{
				Error: tt.errMsg,
			}
			if err := fr.SendJSON(&buf); (err != nil) != tt.wantErr {
				t.Errorf("FailResponse.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			assert.Equal(t, tt.want, buf.String())
		})
	}
}
