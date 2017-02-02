package main

import (
	  "testing"
    "net/http"
		"net/http/httptest"
		"github.com/stretchr/testify/assert"
    "bytes"
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
    r, err := http.NewRequest("GET", "/scale", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new responserecorder (which satisfies http.ResponseWriter) to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Home)

    // Handlers satisfy http.Handler, so can call their ServeHTTP method directly to pass in the request and responserecorder.
    handler.ServeHTTP(rr, r)

		// assert status code of response recorder is OK
    assert.Equal(t, rr.Code, http.StatusOK, "HandlerFunc(cale) returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
}

func TestSetRadioButton(t *testing.T){
	// Generate a request for 1 octave C major scale to pass to handler

    data := url.Values{}
    data.Set("Key", "C")
    data.Add("Octave", "1")
		data.Add("ScaleArp", "Scale")
		data.Add("Pitch", "Major")
  	fmt.Println(data)

	// url.Values{"Key": {"C"}, "Octave": {"1"},"ScaleArp": {"Scale"}, "Pitch": {"Major"}}

	r, err := http.NewRequest("GET", "/scaleshow", data)
	if err != nil {
			t.Fatal(err)
	}

fmt.Println(r.Body)

//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(ScaleShow)

	// Handlers satisfy http.Handler, so can call their ServeHTTP method directly to pass in the request and responserecorder.
//	handler.ServeHTTP(rr, r)

//	fmt.Println(rr.Body)
}
