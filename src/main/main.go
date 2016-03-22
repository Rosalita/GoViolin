package main

import(
  "fmt"
   "net/http"
   "html/template"
   "log"
)
func main(){

  // this code below makes a file server and serves everything as a file
  // fs := http.FileServer(http.Dir(""))
  // http.Handle("/", fs)


  // serve everything in the css folder as a file
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

  // when navigating to /home it should serve the home page
  http.HandleFunc("/", Index)
  http.HandleFunc("/home", Home)
  http.HandleFunc("/scale", Scale)
  http.ListenAndServe(":8080", nil)

  }

  func Index(w http.ResponseWriter, r *http.Request){
    // prints a message to the server console
    fmt.Println("this is index page")
  }

  //handler for /home renders the home.html page
  func Home(w http.ResponseWriter, req *http.Request){
    render(w, "home.html")
  }

// handler for /scale which parses the values submitted from /home
  func Scale(w http.ResponseWriter, r *http.Request){
    r.ParseForm() //r is url.Values which is a map[string][]string
    for key, values := range r.Form {   // range over map
       for _, value := range values {    // range over []string
       fmt.Println(key, value) // print all the keys and values to server console
       fmt.Fprintf(w, " %s : %s\n", key, value) // and print them in the browser
       }
     }
  }

  func render(w http.ResponseWriter, tmpl string){

    tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
    t, err := template.ParseFiles(tmpl) //parse the template

    if err != nil{ // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, "") //execute the template

    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
  }
