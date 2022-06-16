package main

import ("fmt"; "net/http"; "io/ioutil"; "strings"; "text/template")

type text struct {
  Name string
  Body []byte
  Pass string
}

func showPage(wrt http.ResponseWriter, req *http.Request){
  var fname string = "." + req.URL.Path // <current directory> + <string following the domain>
  var pass string = "passcode"
  req.ParseForm()
  body, err := ioutil.ReadFile(fname)
  if err != nil {
    fmt.Fprintf(wrt, "Bad Request: Not Exist")
    return
  }
  for key, value := range req.Form {
    if key == "pass" {
      if strings.Join(value, "") == pass {
        temp, _ := template.ParseFiles("edit.html");
        temp.Execute(wrt, &text{Body:body, Name:req.URL.Path, Pass:pass})
        return
      }
    }else if key == "save" {
      if strings.Join(value, "") == pass {
        saveText(wrt, req)
      }
    }
  }
  fmt.Fprintf(wrt, "%s", body)
}

func saveText(wrt http.ResponseWriter, req *http.Request){
  body := req.FormValue("txt_body")
  ioutil.WriteFile("." + req.URL.Path, []byte(body), 0600)
  http.Redirect(wrt, req, req.URL.Path, http.StatusFound)
} 

func main(){
  http.HandleFunc("/", showPage)
  println("Running...")
  http.ListenAndServe(":8080", nil)
}
