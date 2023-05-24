package util

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetData(r *http.Request, dst interface{}) error {
	if g, ok := dst.(ParamsGetter); ok {
		if err := g.GetParams(r); err != nil {
			return err
		}
	}
	if u, ok := dst.(json.Unmarshaler); ok {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		err = u.UnmarshalJSON(body)
		if err != nil {
			return err
		}
	}
	if v, ok := dst.(Validator); ok {
		return v.Validate()
	}
	return nil
}

func GetIDFromPath(pattern string, path string) (int64, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(path, pattern))
	return int64(id), err
}

func SendData(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		SendError(w, err, http.StatusInternalServerError)
		return
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		SendError(w, err, http.StatusInternalServerError)
		return
	}

}

func SendError(w http.ResponseWriter, err error, statusCode int) {
	out, err := json.Marshal(APIResponse{
		Status: "error",
		Result: err,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

type APIResponse struct {
	Status string `json:"status"`
	Result any    `json:"result"`
}

func NewSuccessResponse(result any) APIResponse {
	return APIResponse{
		Status: "success",
		Result: result,
	}
}

func NewEmptySuccessResponse() APIResponse {
	return APIResponse{
		Status: "success",
	}
}

type ParamsGetter interface {
	GetParams(req *http.Request) error
}

type Validator interface {
	Validate() error
}
