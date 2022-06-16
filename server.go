package main

import ("fmt"; "net/http"; "io/ioutil"; "strings"; "text/template")

type text struct {
  Name string
  Body []byte
  Pass string
}

// showPage: Show and edit a web page.
func showPage(wrt http.ResponseWriter, req *http.Request){
  // fname: <current directory> + <string following the domain>
  var fname string = "." + req.URL.Path
  
  // pass: Editing function passcode
  var pass string = "passcode"
  
  // Get all received http.Request parameters
  req.ParseForm()
  body, err := ioutil.ReadFile(fname)
  if err != nil {
    // Web page is not exist
    fmt.Fprintf(wrt, "Bad Request: Not Exist")
    return
  }
  for key, value := range req.Form {
    if key == "pass" {
      // strings.Join: ([]string, string) -> string
      if strings.Join(value, "") == pass {
        temp, _ := template.ParseFiles("edit.html")
        
        // temp.Execute: Apply (Body, Name, Pass) to edit.html.
        temp.Execute(wrt, &text{Body:body, Name:req.URL.Path, Pass:pass})
        return
      }
    }else if key == "save" {
      // strings.Join: ([]string, string) ->  string
      if strings.Join(value, "") == pass {
        saveText(wrt, req)
      }
    }
  }
  fmt.Fprintf(wrt, "%s", body)
}

// saveText: Save the sent text data.
func saveText(wrt http.ResponseWriter, req *http.Request){
  // Get the sent data under the name "txt_body"
  body := req.FormValue("txt_body")
  
  // ioutil.WriteFile: Save data in <current directory> + <string following the domain>
  ioutil.WriteFile("." + req.URL.Path, []byte(body), 0600)
  
  // http.Redirect: Redirect to the saved web page with a status code.
  http.Redirect(wrt, req, req.URL.Path, http.StatusFound)
} 

func main(){
  // http.HandleFunc: Triggered by access to "/", execute the showPage function.
  http.HandleFunc("/", showPage)
  println("Running...")
  
  // http.ListenAndServe: Monitor port 8080
  http.ListenAndServe(":8080", nil)
}
