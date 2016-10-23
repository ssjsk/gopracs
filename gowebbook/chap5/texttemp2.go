package main

import(
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}

const tmpl = `Notes are:
{{range .}}
	Title: {{.Title}}, Description: {{.Description}}
{{end}}
`

func main(){
	//Create slice of Note objects
	notes := []Note {
		{"text/template", "text/template generates textual output"},
		{"html/template", "text/template generates html output"},
	}

	//create new template with name
	t := template.New("note")

	//parse some content and generate a template

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//apply parsed template to data
	if err := t.Execute(os.Stdout, notes); err != nil{
		log.Fatal("Executing template:", err)
		return
	}
}