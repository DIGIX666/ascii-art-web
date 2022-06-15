package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const port = ":8080"

func getTable() []string {
	file, err := os.Open("standard.txt")
	content, _ := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	table := strings.Split(string(content), "\n")

	return table
}
func asciiGenerator(text string) {

	line := strings.Split(text, "\\n")
	content := getTable()
	for i := 0; i < len(line); i++ {
		if len(line[i]) > 0 {
			chars := []rune(line[i])
			for n := 0; n < 8; n++ {
				for v := 0; v < len(chars); v++ {
					group := int(chars[v]) - 32
					adress := group * 9
					charLine := content[adress+1+n]
					fmt.Print(charLine)
				}
				fmt.Print(string(rune('\n')))
			}
		} else {
			fmt.Print(string(rune('\n')))
		}
	}
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println("(http://localhost:8080) Server started on port ", port)
	http.ListenAndServe(port, nil)
}
