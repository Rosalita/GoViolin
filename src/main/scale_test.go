package main

import "testing"

func TestSettingDefaultScaleOptionsNotNil(t *testing.T){

  options1, options2, options3, options4 := setDefaultScaleOptions()

  if options1 == nil {
    t.Error("Test failed, no default options for Scale / Arpeggio")
  }
  if options2 == nil {
    t.Error("Test failed, no default options for Major / Minor ")
  }
  if options3 == nil {
    t.Error("Test failed, no default options for Key")
  }
  if options4 == nil {
    t.Error("Test failed, no default options for Octave")
  }
}
