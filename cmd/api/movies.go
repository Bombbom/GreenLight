package main 

import (
	"fmt"
	"net/http"
	"time"
	"greenlight.clone/internal/data"
	"encoding/json"
	// "errors"
	// "strconv"
	// "github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request){
	var input struct {
		Title string `json:"title"` 
		Year int32 `json:"year"`
		Runtime int32 `json:"runtime"` 
		Genres []string `json:"genres"` 
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return 
	}

	fmt.Fprintf(w, "%+v", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request){
	id , err := app.readIDParam(r)
	if err != nil {
		// http.NotFound(w,r)
		app.notFoundResponse(w, r)
		return 
	}
	
	movie := data.Movie{
		ID: id, 
		CreatedAt: time.Now(), 
		Title: "Casablance", 
		Runtime: 102, 
		Genres: []string{"drama", "romance", "war"},
		Version: 1, 
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil{
		// app.logger.Error(err.Error())
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
		
	}
}


