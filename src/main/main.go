package main

import(
//  "fmt"
  "net/http"

//  "html/template"
)
func main(){

  // need a serve mux
  fs := http.FileServer(http.Dir(""))
  http.Handle("/", fs)

  // when navigating to /home it should serve the home page

  http.ListenAndServe(":8080", nil)


  // page per major scale a - g
  // generate these pages from a templatePaths

}
