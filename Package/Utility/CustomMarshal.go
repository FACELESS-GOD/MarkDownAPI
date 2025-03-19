package Utility

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(Req *http.Request, ReqInterface interface{}) {
	if body, err := ioutil.ReadAll(Req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), ReqInterface); err == nil {
			return
		}
	}
}
