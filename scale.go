package main

import (
	"net/http"
	"strings"
)

func Scale(w http.ResponseWriter, req *http.Request) {

	//populate the default ScaleOptions, PitchOptions, KeyOptions, OctaveOptions for scales and arpeggios
	sOptions, pOptions, kOptions, oOptions := setDefaultScaleOptions()

	// set default page variables
	pageVars := PageVars{
		Title:         "Practice Scales and Arpeggios", // default scale initially displayed is A Major
		Scalearp:      "Scale",
		Pitch:         "Major",
		Key:           "A",
		ScaleImgPath:  "img/scale/major/a1.png",
		GifPath:       "",
		AudioPath:     "mp3/scale/major/a1.mp3",
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

	// Update options based on the user's selection

	// Set key options - set isChecked true for selected key and false for all other keys
	kOptions = setKeyOptions(key)

	// Set scale options
	if scalearp == "Scale" {
		// if scale is selected set scale isChecked to true and arpeggio isChecked to false
		sOptions = []ScaleOptions{
			ScaleOptions{"Scalearp", "Scale", false, true, "Scales"},
			ScaleOptions{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
		}
	} else {
		// if arpeggio is selected set arpeggio isChecked to true and scale isChecked to false
		sOptions = []ScaleOptions{
			ScaleOptions{"Scalearp", "Scale", false, false, "Scales"},
			ScaleOptions{"Scalearp", "Arpeggio", false, true, "Arpeggios"},
		}
	}

	// Set pitch options
	if pitch == "Major" {
		pOptions = []ScaleOptions{ // if major was selected, set major isChecked to true and minor isChecked to false
			ScaleOptions{"Pitch", "Major", false, true, "Major"},
			ScaleOptions{"Pitch", "Minor", false, false, "Minor"},
		}
	} else {
		pOptions = []ScaleOptions{ // if minor was selected, set minor isChecked to true and major isChecked to false
			ScaleOptions{"Pitch", "Major", false, false, "Major"},
			ScaleOptions{"Pitch", "Minor", false, true, "Minor"},
		}
	}

	// Set octave options
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

	// work out what the actual key is and set its value
	if pitch == "Major" {
		// for major scales if the key is longer than 2 characters, we only care about the last 2 characters
		if len(key) > 2 { // only select last two characters for keys which contain two possible names e.g. C#/Db
			key = key[3:]
		}
	} else { // pitch is minor
		// for minor scales if the key is longer than 2 characters, we only care about the first 2 characters
		if len(key) > 2 { // only select first two characters for keys which contain two possible names e.g. C#/Db
			key = key[:2]
		}
	}

	// Set the labels, Major have a scale and a drone, while minor have melodic and harmonic minor scales
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

	// Intialise paths to the associated images and mp3s
	imgPath, audioPath, audioPath2 := "img/", "mp3/", "mp3/"

	// Build paths to img and mp3 files that correspond to user selection
	if scalearp == "Scale" {
		imgPath += "scale/"
		audioPath += "scale/"

	} else {
		// if arpeggio is selected, add "arps/" to the img and mp3 paths
		imgPath += "arps/"
		audioPath += "arps/"
	}

	if pitch == "Major" {
		imgPath += "major/"
		audioPath += "major/"
	} else {
		imgPath += "minor/"
		audioPath += "minor/"
	}

	audioPath += strings.ToLower(key)
	imgPath += strings.ToLower(key)
	// if the img or audio path contain #, delete last character and replace it with s
	imgPath = changeSharpToS(imgPath)
	audioPath = changeSharpToS(audioPath)

	switch octave {
	case "1":
		imgPath += "1"
		audioPath += "1"
	case "2":
		imgPath += "2"
		audioPath += "2"
	}

	audioPath += ".mp3"
	imgPath += ".png"

	//generate audioPath2
	// audio path2 can either be a melodic minor scale or a drone note.
	// Set to melodic minor scale - if the first 16 characters of audio path are:
	if audioPath[:16] == "mp3/scale/minor/" {
		audioPath2 = audioPath                      // set audioPath2 to the original audioPath
		audioPath2 = audioPath2[:len(audioPath2)-4] // chop off the last 4 characters, this removes .mp3
		audioPath2 += "m.mp3"                       // then add m for melodic and the .mp3 suffix
	} else { // audioPath2 needs to be a drone note.
		audioPath2 += "drone/"
		audioPath2 += strings.ToLower(key)
		// may have just added a # to the path, so use the function to change # to s
		audioPath2 = changeSharpToS(audioPath2)
		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
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
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
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
