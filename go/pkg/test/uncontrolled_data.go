package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


func getData(w http.ResponseWriter, req *http.Request) error {
	p := string(req.URL.Query().Get("path")[0])

	o, _ := os.Create(filepath.Join("/tmp/", p))
	if _, err := fmt.Fprintln(o, "Duly clobbered."); err != nil {
		return err
	}
	_ = o.Close()

	c, _ := os.ReadFile(filepath.Join("/var/tmp/", p))
	w.Write(c)

	return nil
}

func sanitize1(s string) (string, error) {
	s = filepath.Clean(s)
	if !strings.HasPrefix(s, "/tmp/") {
		return s, fmt.Errorf("BAD PATH")
	}
	return s, nil
}

func getData1(w http.ResponseWriter, req *http.Request) error {
	p := string(req.URL.Query().Get("path")[0])

	var err error

	if p, err = sanitize1(filepath.Join("/tmp/", p)); err != nil {
		return err
	}

	c, _ := os.ReadFile(p)
	w.Write(c)

	o, _ := os.Create(p)
	if _, err := fmt.Fprintln(o, "Duly clobbered."); err != nil {
		return err
	}
	_ = o.Close()

	return nil
}

func sanitize2(s string) error {
	s = filepath.Clean(s)
	if !strings.HasPrefix(s, "/tmp/") {
		return fmt.Errorf("BAD PATH")
	}
	return nil
}

func getData2(w http.ResponseWriter, req *http.Request) error {
	p := string(req.URL.Query().Get("path")[0])
	p = filepath.Join("/tmp/", p)

	var err error

	if err = sanitize2(p); err != nil {
		return err
	}

	// p is now safe.

	c, _ := os.ReadFile(p)
	w.Write(c)

	o, _ := os.Create(p)
	if _, err := fmt.Fprintln(o, "Duly clobbered."); err != nil {
		return err
	}
	_ = o.Close()

	return nil
}
