package controllers
import(
	"net/http"
	"text/template"
	"../library"
	"log"
	"os/exec"
)

func HomeController(w http.ResponseWriter, r *http.Request) bool{
	templates := template.Must(template.ParseFiles("app/templates/index.html"))

	//testing commands
	command := "cat /proc/meminfo | grep 'MemTotal\\|MemFree\\|MemAvailable'| awk '{print \"<b>\"}{print $1}{print \"</b>\"}{print $2}' "
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() err %s\n", err)
	}
	
	Config := library.Config{
		Name:"Exodo",
		Oper : string(out),
	}
	
	 if err := templates.ExecuteTemplate(w, "index.html", struct{Config interface{}}{Config}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return true
}
