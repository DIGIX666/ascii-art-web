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
	//renderTemplate(w, "home")
	file, err := os.Open("./assets/shadow.txt")
	content, _ := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	table := strings.Split(string(content), "\n")

	var result []string

	line := strings.Split(r.FormValue("asciitext"), "\\n")
	//content := table
	for i := 0; i < len(line); i++ {
		if len(line[i]) > 0 {
			chars := []rune(line[i])
			for n := 0; n < 8; n++ {
				for v := 0; v < len(chars); v++ {
					group := int(chars[v]) - 32
					adress := group * 9
					charLine := table[adress+1+n]
					result = append(result, charLine)
					//fmt.Fprintf(w, charLine)
				}
				//fmt.Fprint(w, string(rune('\n')))
				result = append(result, string(rune('\n')))
			}
		} else {
			//fmt.Fprint(w, string(rune('\n')))
			result = append(result, string(rune('\n')))
		}
	}
	sresult := ""
	for i := range result {
		sresult += result[i]
	}

	t, err := template.ParseFiles("./templates/" + "home.html")

	if err != nil {
		fmt.Println(err)
	}
	resFinal := ResultAscii{sresult}
	t.Execute(w, resFinal)
	fmt.Print(w, resFinal)
}
