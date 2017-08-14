package controllers
import(
	"net/http"
	"text/template"
	"./../library"
)

func HomeController(w http.ResponseWriter, r *http.Request) bool{
	templates := template.Must(template.ParseFiles("app/templates/index.html"))

	Config := library.Config{
		Name:"Exodo",
		Oper :"command",
	}
	
	 if err := templates.ExecuteTemplate(w, "index.html", struct{Config interface{}}{Config}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return true
}
