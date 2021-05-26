package main

import (
	"fmt"
	"ioutil"
	"net/http"
	"os"
	"path/filepath"
)


func getData(w http.ResponseWriter, req *http.Request) error {
	p := req.URL.Query().Get("path")[0]

	o, _ := os.Create(filepath.Join("/tmp/", p))
	if _, err := fmt.Fprintln(o, "Duly clobbered."); err != nil {
		return err
	}

	w.Write(ioutil.ReadFile(filepath.Join("/var/tmp/", p)))
	
	return o.Close()
}
