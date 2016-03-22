package main

import(
  "fmt"
   "net/http"
   "html/template"
   "log"
)
func main(){

  // this code below makes a file server and serves everything as a file
  //  fs := http.FileServer(http.Dir(""))
  //  http.Handle("/", fs)


  // serve everything in the css folder as a file
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

    // when navigating to /home it should serve the home page

  http.HandleFunc("/", Index)
  http.HandleFunc("/home", Home)
  http.ListenAndServe(":8080", nil)
    // page per major scale a - g
    // generate these pages from a templatePaths



  }

  func Index(w http.ResponseWriter, r *http.Request){
    fmt.Println("this is index page")
  }
  func Home(w http.ResponseWriter, req *http.Request){
    render(w, "home.html")
  }

  func render(w http.ResponseWriter, tmpl string){
      fmt.Printf("%s", tmpl)
    tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
    fmt.Printf("%s", tmpl)
    t, err := template.ParseFiles(tmpl)
    if err != nil{
      log.Print("template parsing error: ", err)
    }

    err = t.Execute(w, "")
    if err != nil {
      log.Print("template executing error: ", err)
    }
  }
