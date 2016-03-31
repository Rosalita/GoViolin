package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageVars struct {
	Title        string
	Headline     string
	Key          string
	Pitch        string
	ScaleImgPath string
}

func main() {

	// this code below makes a file server and serves everything as a file
	// fs := http.FileServer(http.Dir(""))
	// http.Handle("/", fs)

	// serve everything in the css folder and the img folder as a file
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	// when navigating to /home it should serve the home page
	http.HandleFunc("/", Home)
	http.HandleFunc("/scale", Scale)
		http.HandleFunc("/scaleshow", ScaleShow)
	http.HandleFunc("/arp", Arp)
	http.HandleFunc("/diary", Diary)
	http.ListenAndServe(":8080", nil)

}

//handler for / renders the home.html
func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "Welcome",
	}
	render(w, "home.html", pageVars)
}


func Scale(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "Practice Scales",
	}
	render(w, "scale.html", pageVars)
}

func Arp(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "Practice Arpeggios",
	}
	render(w, "arp.html", pageVars)
}

func Diary(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "Practice Diary",
	}
	render(w, "diary.html", pageVars)
}



// handler for /scaleshow which parses the values submitted from /scale
func ScaleShow(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //r is url.Values which is a map[string][]string

	var svalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			svalues = append(svalues, value) // stick each value in a slice I know the name of
		}
	}
	sstring := "This is a scale of "
	key := ""
	pitch := ""

	// identify selected key / pitch and generate a string with the name of the scale from the form.
	if len(svalues[0]) == 1 || len(svalues[0]) == 2 {
		sstring += svalues[0] + " " + svalues[1]
		key = svalues[0]
		pitch = svalues[1]
	} else {
		sstring += svalues[1] + " " + svalues[0]
		key = svalues[1]
		pitch = svalues[0]
	}

	//generate a path to the associated img

	path := "img/"
	if pitch == "Major" {
		path += "major/"
	}
	if pitch == "Minor" {
		path += "minor/"
	}

	path += strings.ToLower(key)
	path += "1.png"

	pageVars := PageVars{
		Title:        "Practice Scales",
		Headline:     sstring,
		Key:          key,
		Pitch:        pitch,
		ScaleImgPath: path,
	}
	render(w, "scale.html", pageVars)
}

func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {

	tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
	t, err := template.ParseFiles(tmpl)      //parse the template file held in the templates folder

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
