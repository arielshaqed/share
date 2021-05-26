package test

import (
	"fmt"
	"net/http"
	"os"
)


func getData(req *http.Request) error {
	p := req.URL.Query().Get("path")

	o, _ := os.Create("/tmp/" + p)
	if _, err := fmt.Fprintln(o, "Duly clobbered."); err != nil {
		return err
	}
	return o.Close()
}
