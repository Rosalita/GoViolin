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
	AudioPath    string
	AudioPath2   string
	PitchOptions []ScaleOptions
	KeyOptions   []ScaleOptions
	OctaveOptions []ScaleOptions
}

type ScaleOptions struct {
	Name      string
	Value     string
	IsChecked bool
	Text      string
}

func main() {

	// this code below makes a file server and serves everything as a file
	// fs := http.FileServer(http.Dir(""))
	// http.Handle("/", fs)

	// serve everything in the css folder, the img folder and wav folder as a file
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/wav/", http.StripPrefix("/wav/", http.FileServer(http.Dir("wav"))))

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
	//populate the default PitchOptions for scales
	pOptions := []ScaleOptions{
		ScaleOptions{"Pitch", "Major", true, "Major"},
		ScaleOptions{"Pitch", "Minor", false, "Minor"},
	}

	// populate the default KeyOptions for scales
	kOptions := []ScaleOptions{
		ScaleOptions{"Key", "A", true, "A"},
		ScaleOptions{"Key", "Bb", false, "Bb"},
		ScaleOptions{"Key", "B", false, "B"},
		ScaleOptions{"Key", "C", false, "C"},
		ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
		ScaleOptions{"Key", "D", false, "D"},
		ScaleOptions{"Key", "Eb", false, "Eb"},
		ScaleOptions{"Key", "E", false, "E"},
		ScaleOptions{"Key", "F", false, "F"},
		ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
		ScaleOptions{"Key", "G", false, "G"},
		ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
	}

  // populate the default OctaveOptions for scales
  oOptions := []ScaleOptions{
		ScaleOptions{"Octave", "1", true, "1"},
		ScaleOptions{"Octave", "2", false, "2"},
	}

	// set page variables
	pageVars := PageVars{
		Title:        "Scale of A Major", // default scale initially displayed is A Major
		ScaleImgPath: "img/major/a1.png",
		AudioPath:    "wav/major/a1.wav",
		Pitch:        "Major",
		Key:          "A",
		PitchOptions: pOptions,
		KeyOptions:   kOptions,
		OctaveOptions: oOptions,
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

	//populate the default PitchOptions for scales
	pOptions := []ScaleOptions{
		ScaleOptions{"Pitch", "Major", true, "Major"},
		ScaleOptions{"Pitch", "Minor", false, "Minor"},
	}

	// populate the default KeyOptions for scales
	kOptions := []ScaleOptions{
		ScaleOptions{"Key", "A", true, "A"},
		ScaleOptions{"Key", "Bb", false, "Bb"},
		ScaleOptions{"Key", "B", false, "B"},
		ScaleOptions{"Key", "C", false, "C"},
		ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
		ScaleOptions{"Key", "D", false, "D"},
		ScaleOptions{"Key", "Eb", false, "Eb"},
		ScaleOptions{"Key", "E", false, "E"},
		ScaleOptions{"Key", "F", false, "F"},
		ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
		ScaleOptions{"Key", "G", false, "G"},
		ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
	}

  // populate the default OctaveOptions for scales
	oOptions := []ScaleOptions{
		ScaleOptions{"Octave", "1", true, "1"},
		ScaleOptions{"Octave", "2", false, "2"},
	}

	r.ParseForm() //r is url.Values which is a map[string][]string

	var svalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			svalues = append(svalues, value) // stick each value in a slice I know the name of
		}
	}
	sstring := "Scale of "
	key := ""
	pitch := ""
	octave := ""

	// identify selected key / pitch / octave and generate a string with the name of the scale from the form.

for i:= 0; i < 3; i++{
	switch svalues[i] {
	case "Major":
		pitch = svalues[i]
  case "Minor":
		pitch = svalues[i]
  case "1":
		octave = svalues[i]
	case "2":
		octave = svalues[i]
  default:
		key = svalues[i]
	}
}

fmt.Println(octave)
//	if svalues[0] == "Major" || svalues[0] == "Minor" {
//		pitch = svalues[0]
//		if svalues[1] == "1" || svalues[1] == "2"  {
//			octave = svalues[1]
//			key = svalues[2]
//		}
//	}

//	if svalues[1] == "Major" || svalues[1] == "Minor" {
//		pitch = svalues[1]
//		if svalues[0] == "1" || svalues[0] == "2"  {
//			octave = svalues[0]
//			key = svalues[2]
//		}
//	}


	// set the selected key's IsChecked value to true so can persist radio button selection
	switch key {
	case "A":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", true, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "Bb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", true, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "B":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", true, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "C":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", true, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "C#/Db":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", true, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "D":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", true, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "Eb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", true, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "E":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", true, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "F":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", true, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "F#/Gb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", true, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}
	case "G":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", true, "G"},
			ScaleOptions{"Key", "G#/Ab", false, "G#/Ab"},
		}

	case "G#/Ab":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, "A"},
			ScaleOptions{"Key", "Bb", false, "Bb"},
			ScaleOptions{"Key", "B", false, "B"},
			ScaleOptions{"Key", "C", false, "C"},
			ScaleOptions{"Key", "C#/Db", false, "C#/Db"},
			ScaleOptions{"Key", "D", false, "D"},
			ScaleOptions{"Key", "Eb", false, "Eb"},
			ScaleOptions{"Key", "E", false, "E"},
			ScaleOptions{"Key", "F", false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, "G"},
			ScaleOptions{"Key", "G#/Ab", true, "G#/Ab"},
		}
	}

	//generate a path to the associated img
	path := "img/"
	if pitch == "Major" {
		path += "major/"

		pOptions = []ScaleOptions{
			ScaleOptions{"Pitch", "Major", true, "Major"},
			ScaleOptions{"Pitch", "Minor", false, "Minor"},
		}

		if len(key) > 2 {
			key = key[3:]
		}
		sstring += key + " " + pitch

	}

	if pitch == "Minor" {
		path += "minor/"

		pOptions = []ScaleOptions{
			ScaleOptions{"Pitch", "Major", false, "Major"},
			ScaleOptions{"Pitch", "Minor", true, "Minor"},
		}

		if len(key) > 2 {
			key = key[:2]
		}
		sstring += key + " " + pitch
	}

	path += strings.ToLower(key)

	// if path string contains # replace the # with an s
	if strings.Contains(path, "#") {
		path = path[:len(path)-1] // remove the #
		path += "s"               // replace it with s
	}
	audioPath := path                                       //audioPath is a new copy of the img path
	audioPath2 := ""                                        //declare audiopath2 as a blank string, used for melodic minor scales
	audioPath = strings.Replace(audioPath, "img", "wav", 1) // but replace "img" with "wav"
	path += "1.png"

	if pitch == "Major" {
		audioPath += "1.wav"
	} else {
		audioPath2 = audioPath
		audioPath += "1h.wav"
		audioPath2 += "1m.wav"
	}

	pageVars := PageVars{
		Title:        sstring,
		Key:          key,
		Pitch:        pitch,
		ScaleImgPath: path,
		AudioPath:    audioPath,
		AudioPath2:   audioPath2,
		PitchOptions: pOptions,
		KeyOptions:   kOptions,
		OctaveOptions: oOptions,
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
