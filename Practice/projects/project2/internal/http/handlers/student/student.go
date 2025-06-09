package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/kartik-pantelwar/new-api.git/internal/types"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF)
		//EOF means body is empty, means there is no input in the request
		slog.Info("Creating new Student")
		w.Write([]byte("hello"))
	}
}