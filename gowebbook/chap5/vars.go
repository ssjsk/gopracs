package main

import (
	"log"
	"os"
	"text/template"
)

//const tml = `{{definte "T"}}{{ $a := 5, $b := 5}} {{ eq $a $b | if }} a and b are equal !{{end}}`

func main(){
	/*t := template.New("varb")
	t, err := t.Parse(tml)
	if err != nil{
		log.Fatal("Parse err: ", err)
		return
	}

	if err := t.Execute(os.Stdout); err != nil{
		log.Fatal("Execute error: ", err)
		return
	}*/
	t, err := template.New("text").Parse(`{{definte "T"}}{{ $a := 5, $b := 5}} {{ eq $a $b | if }} a and b are equal !{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", 5)

	if err != nil {
		log.Fatal("Execute: ", err)
	}
}