package render

import (
	"bytes"
	"fmt"
	"github.com/KabirGupta07/bookings/pkg/config"
	"github.com/KabirGupta07/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

var app *config.AppConfig
//NewTemplates sets the config for render package
func NewTemplates(a *config.AppConfig){
	app = a
}

//func AddDefaultData(td *models.TemplateData)*models.TemplateData{
//
//
//	//return td
//}


func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache{
		tc= app.TemplateCache
	}else{
		tc,_= CreateTemplateCache()
	}

	//get the template cache from the app config

	t,ok:= tc[tmpl]
	if !ok{
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	//td = AddDefaultData(td)

	_ = t.Execute(buf,td)
	_,err := buf.WriteTo(w)
	if err!= nil{
		fmt.Println("error writing template to browser",err)
	}
}


//CreateTemplateCache creates a template cache
func CreateTemplateCache()(map[string]*template.Template, error){
	//myCache stores the cache data
	myCache := map[string]*template.Template{}

	//pages store the filenames ending with .page.layout
	pages,err:=filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return myCache,err
	}

	var ts *template.Template
	for _,page :=range pages{
		name:=filepath.Base(page)
		ts, err = template.New(name).Funcs(functions).ParseFiles(page)
		if err!=nil{
			return myCache, err
		}

		matches,err := filepath.Glob("./templates/*.layout.tmpl")
		if err!=nil{
			return myCache, err
		}

		if len(matches)>0{
			ts,err= ts.ParseGlob("./templates/*.layout.tmpl")
			if err!=nil{
				return myCache, err
			}
		}

		myCache[name] = ts
	}
return myCache,nil
}
