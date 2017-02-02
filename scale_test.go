package main

import "testing"
import "github.com/stretchr/testify/assert"

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
