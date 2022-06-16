package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ResultAscii struct {
	TextAscii string
}

func Home(w http.ResponseWriter, r *http.Request) {
	namePolice := r.FormValue("police")
	//fmt.Println(namePolice)
	file, err := os.Open("assets/" + namePolice + ".txt")

	content, _ := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	table := strings.Split(string(content), "\n")

	var result []string

	line := strings.Split(r.FormValue("asciitext"), "\\n")
	for i := 0; i < len(line); i++ {
		if len(line[i]) > 0 {
			chars := []rune(line[i])
			for n := 0; n < 8; n++ {
				for v := 0; v < len(chars); v++ {
					group := int(chars[v]) - 32
					adress := group * 9
					charLine := table[adress+1+n]
					result = append(result, charLine)
				}
				result = append(result, string(rune('\n')))
			}
		} else {
			result = append(result, string(rune('\n')))
		}
	}
	sresult := ""
	for i := range result {
		sresult += result[i]
	}
	t, err := template.ParseFiles("./templates/home.html")

	if err != nil {
		fmt.Println(err)
	}
	sresult = strings.Replace(sresult, "\r", "\n", -1)
	resFinal := ResultAscii{sresult}
	t.Execute(w, resFinal)
}
