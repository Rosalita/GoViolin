package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageVars struct {
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	LeftLabel     string
	RightLabel    string
	ScaleOptions  []ScaleOptions
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

	//populate the default ScaleOptions, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
	sOptions, pOptions, kOptions, oOptions := setDefaultScaleOptions()

	// set page variables
	pageVars := PageVars{
		Title:         "Practice Scales and Arpeggios", // default scale initially displayed is A Major
		Scalearp:      "Scale",
		Pitch:         "Major",
		Key:           "A",
		ScaleImgPath:  "img/major/a1.png",
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     "mp3/major/a1.mp3",
		AudioPath2:    "mp3/drone/a1.mp3",
		LeftLabel:     "Listen to Major scale",
		RightLabel:    "Listen to Drone",
		ScaleOptions:  sOptions,
		PitchOptions:  pOptions,
		KeyOptions:    kOptions,
		OctaveOptions: oOptions,
	}
	render(w, "scale.html", pageVars)
}

// handler for /scaleshow which parses the values submitted from /scale
func ScaleShow(w http.ResponseWriter, r *http.Request) {

	//populate the default ScaleOptions, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
	sOptions, pOptions, kOptions, oOptions := setDefaultScaleOptions()

	r.ParseForm() //r is url.Values which is a map[string][]string

	var svalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			svalues = append(svalues, value) // stick each value in a slice I know the name of
		}
	}

	scalearp, key, pitch, octave, leftlabel, rightlabel := "", "", "", "", "", ""

	// the slice of values return by the request can be arranged in any order
	// so identify selected scale / arpeggio, pitch, key and octave and store values in variables for later use.
	for i := 0; i < 4; i++ {
		switch svalues[i] {
		case "Major":
			pitch = svalues[i]
		case "Minor":
			pitch = svalues[i]
		case "1":
			octave = svalues[i]
		case "2":
			octave = svalues[i]
		case "Scale":
			scalearp = svalues[i]
		case "Arpeggio":
			scalearp = svalues[i]
		default:
			key = svalues[i]
		}
	}

	//set the labels

	if pitch == "Major" {
		leftlabel = "Listen to Major "
		rightlabel = "Listen to Drone"
		if scalearp == "Scale" {
			leftlabel += "Scale"
		} else {
			leftlabel += "Arpeggio"
		}
	} else {
		if scalearp == "Arpeggio" {
			leftlabel += "Listen to Minor Arpeggio"
			rightlabel = "Listen to Drone"
		} else {
			leftlabel += "Listen to Harmonic Minor Scale"
			rightlabel += "Listen to Melodic Minor Scale"
		}
	}

	// update the key options based on users selection to set isChecked true for selected key and false for all other keys
	kOptions = setKeyOptions(key)

	//intialise paths to the associated images and mp3s
	imgPath := "img/"
	audioPath := "mp3/"
	audioPath2 := "mp3/" // audio path2 is used for drone notes when major scale or arpeggio selected and melodic minors when minor scale is selected

	if scalearp == "Scale" {
		// if scale is selected set scale isChecked to true and arpeggio isChecked to false
		sOptions = []ScaleOptions{
			ScaleOptions{"Scalearp", "Scale", false, true, "Scales"},
			ScaleOptions{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
		}

		if pitch == "Minor"{
		 audioPath2 += "minor/"
		} else {
		audioPath2 += "drone/"
		}

	} else {
		// if arpeggio is selected set arpeggio isChecked to true and scale isChecked to false
		sOptions = []ScaleOptions{
			ScaleOptions{"Scalearp", "Scale", false, false, "Scales"},
			ScaleOptions{"Scalearp", "Arpeggio", false, true, "Arpeggios"},
		}

		// if arpeggio is selected, add "arps/" to the img and mp3 paths and "drone/" to the second audioPath
		imgPath += "arps/"
		audioPath += "arps/"
		audioPath2 += "drone/"
	}

	if pitch == "Major" {
		imgPath += "major/"        // add "major/" to the image path to find the image
		audioPath += "major/"      // add "major/" to the audio path to find the mp3
		pOptions = []ScaleOptions{ // if major was selected, set major isChecked to true and minor isChecked to false
			ScaleOptions{"Pitch", "Major", false, true, "Major"},
			ScaleOptions{"Pitch", "Minor", false, false, "Minor"},
		}
		// for major scales if the key is longer than 2 characters, we only care about the last 2 characters
		if len(key) > 2 {
			key = key[3:]
		}
		imgPath += strings.ToLower(key) // keys must be added to the path as lower case
		audioPath += strings.ToLower(key)
		audioPath2 += strings.ToLower(key)
		switch octave {
		case "1":
			imgPath += "1.png"
			audioPath += "1.mp3"
			audioPath2 += "1.mp3"

		case "2":
			imgPath += "2.png"
			audioPath += "2.mp3"
			audioPath2 += "2.mp3"
		}

	}

	if pitch == "Minor" {
		imgPath += "minor/"        // add "minor/" to the image path to find the image
		audioPath += "minor/"      // add "minor/" to the audio path to find the mp3
		pOptions = []ScaleOptions{ // if minor was selected, set minor isChecked to true and major isChecked to false
			ScaleOptions{"Pitch", "Major", false, false, "Major"},
			ScaleOptions{"Pitch", "Minor", false, true, "Minor"},
		}
		// for minor scales if the key is longer than 2 characters, we only care about the first 2 characters
		if len(key) > 2 {
			key = key[:2]
		}
		imgPath += strings.ToLower(key) // keys must be added to the path as lower case
		audioPath += strings.ToLower(key)
		audioPath2 += strings.ToLower(key)

		// minor scales can contain # change this to an s in the path
		imgPath = changeSharpToS(imgPath)
		audioPath = changeSharpToS(audioPath)
		audioPath2 = changeSharpToS(audioPath2)

		switch octave {
		case "1":
			imgPath += "1.png"
			if scalearp == "Scale" {
				audioPath += "1h.mp3"
				audioPath2 += "1m.mp3"
			} else { // this is an arpeggio
				audioPath += "1.mp3"
				audioPath2 += "1.mp3"
			}
		case "2":
			imgPath += "2.png"
			if scalearp == "Scale" {
				audioPath += "2h.mp3"
				audioPath2 += "2m.mp3"
			} else { // this is an arpeggio
				audioPath += "2.mp3"
				audioPath2 += "2.mp3"
			}
		}
	}

	// persist the selected octave options
	if octave == "1" {
		oOptions = []ScaleOptions{
			ScaleOptions{"Octave", "1", false, true, "1 Octave"},
			ScaleOptions{"Octave", "2", false, false, "2 Octave"},
		}
	} else {
		oOptions = []ScaleOptions{
			ScaleOptions{"Octave", "1", false, false, "1 Octave"},
			ScaleOptions{"Octave", "2", false, true, "2 Octave"},
		}
	}

	// if arpeggios is selected, disable the 1 octave radio button and default to 2 octaves
	if scalearp == "Arpeggio" {

		oOptions = []ScaleOptions{
			ScaleOptions{"Octave", "1", true, false, "1 Octave"},
			ScaleOptions{"Octave", "2", false, true, "2 Octave"},
		}

		// if any of the existing paths contain a 1, replace with 2
		if strings.Contains(imgPath, "1") {
			imgPath = imgPath[:len(imgPath)-5]
			imgPath += "2.png"
		}
		if strings.Contains(audioPath, "1") {
			audioPath = audioPath[:len(audioPath)-5]
			audioPath += "2.mp3"
		}

		// if any of the existing audioPath2's contain a 1, replace with a 2
		if strings.Contains(audioPath2, "1") {
			audioPath2 = audioPath2[:len(audioPath2)-5]
			audioPath2 += "2.mp3"
		}

	}

	pageVars := PageVars{
		Title:         "Practice Scales and Arpeggios",
		Scalearp:      scalearp,
		Key:           key,
		Pitch:         pitch,
		ScaleImgPath:  imgPath,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     audioPath,
		AudioPath2:    audioPath2,
		LeftLabel:     leftlabel,
		RightLabel:    rightlabel,
		ScaleOptions:  sOptions,
		PitchOptions:  pOptions,
		KeyOptions:    kOptions,
		OctaveOptions: oOptions,
	}
	render(w, "scale.html", pageVars)
}

func Diary(w http.ResponseWriter, r *http.Request) {
	pageVars := PageVars{
		Title: "Practice Diary",
	}
	render(w, "diary.html", pageVars)
}

func Duets(w http.ResponseWriter, r *http.Request) {
	pageVars := PageVars{
		Title: "Practice Duets",
	}
	render(w, "duets.html", pageVars)
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

func setDefaultScaleOptions() ([]ScaleOptions, []ScaleOptions, []ScaleOptions, []ScaleOptions) {

	//set the default scaleOptions for scales and arpeggios
	sOptions := []ScaleOptions{
		ScaleOptions{"Scalearp", "Scale", false, true, "Scales"},
		ScaleOptions{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
	}

	//set the default PitchOptions for scales and arpeggios
	pOptions := []ScaleOptions{
		ScaleOptions{"Pitch", "Major", false, true, "Major"},
		ScaleOptions{"Pitch", "Minor", false, false, "Minor"},
	}

	// set the default KeyOptions for scales and arpeggios
	kOptions := []ScaleOptions{
		ScaleOptions{"Key", "A", false, true, "A"},
		ScaleOptions{"Key", "Bb", false, false, "Bb"},
		ScaleOptions{"Key", "B", false, false, "B"},
		ScaleOptions{"Key", "C", false, false, "C"},
		ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
		ScaleOptions{"Key", "D", false, false, "D"},
		ScaleOptions{"Key", "Eb", false, false, "Eb"},
		ScaleOptions{"Key", "E", false, false, "E"},
		ScaleOptions{"Key", "F", false, false, "F"},
		ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
		ScaleOptions{"Key", "G", false, false, "G"},
		ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	// set the default OctaveOptions for scales and arpeggios
	oOptions := []ScaleOptions{
		ScaleOptions{"Octave", "1", false, true, "1 Octave"},
		ScaleOptions{"Octave", "2", false, false, "2 Octave"},
	}
	return sOptions, pOptions, kOptions, oOptions
}

func setKeyOptions(key string) (kOptions []ScaleOptions) {
	switch key {
	case "A":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, true, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Bb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, true, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "B":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, true, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, true, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C#/Db":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, true, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "D":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, true, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Eb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, true, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "E":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, true, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, true, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F#/Gb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, true, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "G":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, true, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}

	case "G#/Ab":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, true, "G#/Ab"},
		}
	}

	return kOptions
}

func changeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1] // remove the last character
		path += "s"               // replace it with s
	}
	return path
}
