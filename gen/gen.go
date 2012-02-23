package main

import (
	"encoding/json"
	"fmt"
	"github.com/DisposaBoy/JsonConfigReader"
	"log"
	"os"
	"strings"
	"text/template"
)

var fatal func(v ...interface{})

func Gen(filename string, newname string, templates []string, data interface{}) error {
	t := make([]string, 1+len(templates))
	t[0] = filename
	copy(t[1:], templates)
	temps, err := template.ParseFiles(t...)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(newname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	return temps.Execute(out, data)
}

func main() {
	var files, templates []string
	jf := ""
	e := ".go"
	jnext := false
	enext := false
	data := make(map[string]interface{})
	verbosity := 2
	which := 0
	for _, opt := range os.Args[1:] {
		switch {
		case jnext:
			jf = opt
			jnext = false
		case enext:
			e = opt
			enext = false
		case opt == "-q":
			verbosity--
		case opt == "-j":
			jnext = true
		case opt == "-e":
			enext = true
		case opt == "-t":
			which = 1
		case opt == "-f":
			which = 2
		case which == 1:
			templates = append(templates, opt)
		case which == 2:
			files = append(files, opt)
		default:
			log.Fatal("[-q [-q]] [-e] [-j <json>] [-t <templates>] -f <files>")
		}
	}
	if verbosity > 0 {
		fatal = log.Fatal
	} else {
		fatal = func(v ...interface{}) { os.Exit(1) }
	}
	if jnext {
		fatal("json file not given")
	}
	if enext {
		fatal("ending not given")
	}
	if len(files) == 0 {
		fatal("no files given")
	}
	if jf != "" {
		j, err := os.Open(jf)
		if err != nil {
			fatal("unable to open ", jf, ": ", err)
		}
		err = json.NewDecoder(JsonConfigReader.New(j)).Decode(&data)
		if err != nil {
			fatal("unable to load ", jf, ": ", err)
		}
	}
	for _, file := range files {
		var new string
		if i := strings.LastIndex(file, "."); i > 0 {
			new = file[:i] + e
		} else {
			new = file + e
		}
		if verbosity >= 2 {
			fmt.Printf("%s -> %s\n", file, new)
		}
		if err := Gen(file, new, templates, data); err != nil {
			fatal("error generating ", file, ": ", err)
		}
	}
}
