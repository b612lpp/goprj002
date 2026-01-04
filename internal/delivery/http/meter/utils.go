package meter

import (
	"encoding/json"
	"net/http"

	"github.com/b612lpp/goprj002/internal/delivery/http/exeptions"
)

type values interface {
	*gasValues | *enValues
}

func parseIncJ[T values](r *http.Request, v T) error {

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {

		return exeptions.ErrParseData
	} else {

		return nil
	}

}
