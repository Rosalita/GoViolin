package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageVars struct {
	Title    string
	Headline string
	Key      string
	Pitch    string
}

func main() {

	// this code below makes a file server and serves everything as a file
	// fs := http.FileServer(http.Dir(""))
	// http.Handle("/", fs)

	// serve everything in the css folder as a file
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// when navigating to /home it should serve the home page
	http.HandleFunc("/", Home)
	//  http.HandleFunc("/home", Home)
	http.HandleFunc("/scale", Scale)
	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	// prints a message to the server console
	fmt.Println("this is index page")
}

//handler for /home renders the home.html page
func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title:    "Practice Scale",
		Headline: "Practice Scales",
	}
	render(w, "home.html", pageVars)
}

// handler for /scale which parses the values submitted from /home
func Scale(w http.ResponseWriter, r *http.Request) {

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
	if len(svalues[0]) == 1 {
		sstring += svalues[0] + " " + svalues[1]
		key = svalues[0]
		pitch = svalues[1]
	} else {
		sstring += svalues[1] + " " + svalues[0]
		key = svalues[1]
		pitch = svalues[0]
	}
	fmt.Println(sstring) // generate a string with the name of the scale from the form.

	pageVars := PageVars{
		Title:    sstring,
		Headline: sstring,
		Key:      key,
		Pitch:    pitch,
	}
	render(w, "home.html", pageVars)

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
