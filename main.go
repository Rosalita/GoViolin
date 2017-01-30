package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageVars struct {
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	DuetImgPath   string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	DuetAudioBoth string
	DuetAudio1    string
	DuetAudio2    string
	LeftLabel     string
	RightLabel    string
	ScaleOptions  []ScaleOptions
	DuetOptions   []ScaleOptions
	PitchOptions  []ScaleOptions
	KeyOptions    []ScaleOptions
	OctaveOptions []ScaleOptions
}

type ScaleOptions struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

func main() {

	// this code below makes a file server and serves everything as a file
	// fs := http.FileServer(http.Dir(""))
	// http.Handle("/", fs)

	// serve everything in the css folder, the img folder and mp3 folder as a file
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/mp3/", http.StripPrefix("/mp3/", http.FileServer(http.Dir("mp3"))))

	// when navigating to /home it should serve the home page
	http.HandleFunc("/", Home)
	http.HandleFunc("/scale", Scale)
	http.HandleFunc("/scaleshow", ScaleShow)
	http.HandleFunc("/duets", Duets)
	http.HandleFunc("/duetshow", DuetShow)
	http.ListenAndServe(getPort(), nil)

}

// Detect $PORT and if present uses it for listen and serve else defaults to :8080
// This is so that app can run on Heroku
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
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
