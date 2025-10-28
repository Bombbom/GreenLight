package main 
import(
	// "fmt"
	"net/http"
	// "encoding/json"
)

// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

// 	js := `{"status": "available", "environment": %q, "version": %q}`
// 	js = fmt.Sprintf(js, app.config.env, version)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(js))
// }

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){
	data := map[string]string{
		"status": "available",
		"environment": app.config.env, 
		"version": version, 
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)

	}
}