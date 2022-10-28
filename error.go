package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Error(err error, w http.ResponseWriter) {
	fmt.Println(err.Error())
	var e ErrorReply
	w.WriteHeader(http.StatusForbidden)
	e.Code = 404
	e.Message = err.Error()
	json.NewEncoder(w).Encode(e)

}
