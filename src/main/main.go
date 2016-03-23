package main

import(
  "fmt"
   "net/http"
   "html/template"
   "log"
)

type PageVars struct {
  Title string
  Headline string
}

func main(){

  // this code below makes a file server and serves everything as a file
  // fs := http.FileServer(http.Dir(""))
  // http.Handle("/", fs)


  // serve everything in the css folder as a file
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

  // when navigating to /home it should serve the home page
  http.HandleFunc("/", Index)
  http.HandleFunc("/home", Home)
  http.HandleFunc("/home2", Home2)
  http.HandleFunc("/scale", Scale)
  http.ListenAndServe(":8080", nil)

  }

  func Index(w http.ResponseWriter, r *http.Request){
    // prints a message to the server console
    fmt.Println("this is index page")
  }

  //handler for /home renders the home.html page
  func Home(w http.ResponseWriter, req *http.Request){
    pageVars := PageVars{
      Title:"This is a string written in main.go",
      Headline:"Practice Scales",
    }
    render(w, "home.html", pageVars)
  }

  func Home2(w http.ResponseWriter, req *http.Request){
    pageVars := PageVars{Title:"This is the same as home but it has a different title"}
    render(w, "home.html", pageVars)
  }

// handler for /scale which parses the values submitted from /home
  func Scale(w http.ResponseWriter, r *http.Request){


     pageVars := PageVars{Title:"This is a scale page"}
     render(w, "home.html", pageVars)
     r.ParseForm() //r is url.Values which is a map[string][]string

     var svalues []string
     sstring:= "Scale "

     for _, values := range r.Form {   // range over map
        for _, value := range values {    // range over []string
          svalues = append(svalues, value)
          sstring += value
          sstring += " "
    //    fmt.Println(key, value) // print all the keys and values to server console
    //    fmt.Fprintf(w, " %s : %s\n", key, value) // and print them in the browser

        }
      }
 fmt.Println(svalues)
 scaleString:= "Scale "
 scaleString += svalues[0]
 scaleString += " "
 scaleString += svalues[1]

 fmt.Println(scaleString)
 fmt.Println(sstring)
  }

  func render(w http.ResponseWriter, tmpl string, pageVars PageVars){

    tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
    t, err := template.ParseFiles(tmpl) //parse the template file held in the templates folder

    if err != nil{ // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
  }
