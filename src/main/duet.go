package main

import (
	"net/http"
)

func Duets(w http.ResponseWriter, r *http.Request) {

  // define default duet options
	dOptions := []ScaleOptions{
		ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		ScaleOptions{"Duet", "D Major", false, false, "D Major"},
	}

	pageVars := PageVars{
		  Title:        "Practice Duets",
			Key:          "G Major",
			DuetImgPath:  "img/duet/gmajor.png",
			DuetAudio1:   "mp3/duet/gmajorduetboth.mp3",
	    DuetOptions:  dOptions,
	}
	render(w, "duets.html", pageVars)
}

func DuetShow(w http.ResponseWriter, r *http.Request) {

  // define default duet options
	dOptions := []ScaleOptions{
		ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		ScaleOptions{"Duet", "D Major", false, false, "D Major"},
	}

// Set a placeholder image path, this will be changed later.
DuetAudio1:= "mp3/duet/gmajorduetboth.mp3"

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
		}
		DuetAudio1 = "?"
	case "G Major":
			dOptions = []ScaleOptions{
				ScaleOptions{"Duet", "G Major", false, true, "G Major"},
				ScaleOptions{"Duet", "D Major", false, false, "D Major"},
			}
	    DuetAudio1 = "mp3/duet/gmajorduetboth.mp3"
		}


//	imgPath := "img/"

	// set default page variables
	pageVars := PageVars{
		Title:       "Practice Duets",
		Key:         "G Major",
		DuetImgPath: "img/duet/gmajor.png",
		DuetAudio1: DuetAudio1,
		DuetOptions: dOptions,
	}
	render(w, "duets.html", pageVars)
}
