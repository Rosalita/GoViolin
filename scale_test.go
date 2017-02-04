package main

import (
	  "testing"
    "net/http"
		"net/http/httptest"
		"github.com/stretchr/testify/assert"
		"net/url"
		"fmt"

	)

type TD struct {
	Data     string
	Expected string
}

func TestSettingDefaultScaleOptionsNotNil(t *testing.T) {
	assert := assert.New(t)
	options1, options2, options3, options4 := setDefaultScaleOptions()
	assert.NotNil(options1)
	assert.NotNil(options2)
	assert.NotNil(options3)
	assert.NotNil(options4)
}

func TestSharpToS(t *testing.T) {
	assert := assert.New(t)
	testData := []TD{
		TD{"G#", "Gs"},
		TD{"D#", "Ds"},
		TD{"Bb", "Bb"},
	}
	for _, test := range testData {
		assert.Equal(changeSharpToS(test.Data), test.Expected)
	}
}

func TestScalePage(t *testing.T) {
    // Create a request to pass to handler
    r, err := http.NewRequest("GET", "http://localhost:8080/scale", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new responserecorder (which satisfies http.ResponseWriter) to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Scale)

    // Handlers satisfy http.Handler, so can call their ServeHTTP method directly to pass in the request and responserecorder.
    handler.ServeHTTP(rr, r)

		// assert status code of response recorder is OK
    assert.Equal(t, rr.Code, http.StatusOK, "HandlerFunc(Scale) returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
}

func TestSetRadioButtons(t *testing.T){
	// Generate the values for 1 octave C major scale so can pass these values to pass to handler
    form := url.Values{}
    form.Set("Key", "C")
    form.Add("Octave", "1")
		form.Add("ScaleArp", "Scale")
		form.Add("Pitch", "Major")

	// url.Values is a map[string][]string
	// url.Values{"Key": {"C"}, "Octave": {"1"},"ScaleArp": {"Scale"}, "Pitch": {"Major"}}

    // Create a new POST request with no body
	  r, err := http.NewRequest("POST", "http://localhost:8080/scaleshow", nil)
    if err != nil {
      	t.Fatal(err)
    }


	  // r.PostForm is currently an empty map
	   fmt.Println(r.PostForm)
		// Set r.PostForm to the url.Values that will be passed to the form
    r.PostForm = form

    // create a new responserecorder
		rr := httptest.NewRecorder()

    // create a handler which is set to the handler function that is being tested
    handler := http.HandlerFunc(ScaleShow)

    // Get the handler to serve the request storing result in the responserecorder
		handler.ServeHTTP(rr, r)


    fmt.Println(rr.Body)
    fmt.Println(rr.Code)
		fmt.Println(rr.HeaderMap)
		fmt.Println(rr.Flushed)

		assert.Equal(t, rr.Code, http.StatusOK, "HandlerFunc(Scaleshow) returned wrong status code: got %v want %v", rr.Code, http.StatusOK)

}
