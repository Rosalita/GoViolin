package main

import (
	"net/http"
)

func Duets(w http.ResponseWriter, r *http.Request) {

	// define default duet options
	dOptions := []ScaleOptions{
		ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		ScaleOptions{"Duet", "D Major", false, false, "D Major"},
		ScaleOptions{"Duet", "A Major", false, false, "A Major"},
	}

	pageVars := PageVars{
		Title:         "Practice Duets",
		Key:           "G Major",
		DuetImgPath:   "img/duet/gmajor.png",
		DuetAudioBoth: "mp3/duet/gmajorduetboth.mp3",
		DuetAudio1:    "mp3/duet/gmajorduetpt1.mp3",
		DuetAudio2:    "mp3/duet/gmajorduetpt2.mp3",
		DuetOptions:   dOptions,
	}
	render(w, "duets.html", pageVars)
}

func DuetShow(w http.ResponseWriter, r *http.Request) {

	// define default duet options
	dOptions := []ScaleOptions{
		ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		ScaleOptions{"Duet", "D Major", false, false, "D Major"},
		ScaleOptions{"Duet", "A Major", false, false, "A Major"},
	}

	// Set a placeholder image path, this will be changed later.
	DuetImgPath := "img/duet/gmajor.png"
	DuetAudioBoth := "mp3/duet/gmajorduetboth.mp3"
	DuetAudio1 := "mp3/duet/gmajorduetpt1"
	DuetAudio2 := "mp3/duet/gmajorduetpt2"

	r.ParseForm() //r is url.Values which is a map[string][]string
	var dvalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			dvalues = append(dvalues, value) // stick each value in a slice I know the name of
		}
	}

	switch dvalues[0] {
	case "D Major":
		dOptions = []ScaleOptions{
			ScaleOptions{"Duet", "G Major", false, false, "G Major"},
			ScaleOptions{"Duet", "D Major", false, true, "D Major"},
			ScaleOptions{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		dOptions = []ScaleOptions{
			ScaleOptions{"Duet", "G Major", false, true, "G Major"},
			ScaleOptions{"Duet", "D Major", false, false, "D Major"},
			ScaleOptions{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		dOptions = []ScaleOptions{
			ScaleOptions{"Duet", "G Major", false, false, "G Major"},
			ScaleOptions{"Duet", "D Major", false, false, "D Major"},
			ScaleOptions{"Duet", "A Major", false, true, "A Major"},
		}
		DuetImgPath = "img/duet/amajor.png"
		DuetAudioBoth = "mp3/duet/amajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/amajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/amajorduetpt2.mp3"
	}

	//	imgPath := "img/"

	// set default page variables
	pageVars := PageVars{
		Title:         "Practice Duets",
		Key:           "G Major",
		DuetImgPath:   DuetImgPath,
		DuetAudioBoth: DuetAudioBoth,
		DuetAudio1:    DuetAudio1,
		DuetAudio2:    DuetAudio2,
		DuetOptions:   dOptions,
	}
	render(w, "duets.html", pageVars)
}
