package main 
import (
	"fmt"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int 
	env string 
}

type application struct {
	config config 
	logger *slog.Logger 
}

func main(){
	var cfg config 

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse() 


	logger := slog.New(slog.NewtextHandler(os.Stdout, nil))

	app := &application{
		config: cfg, 
		logger: logger, 
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckhandler)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", cfg.port), 
		Handler: mux, 
		IdleTimeout: time.Minute, 
		ReadTimeout: 5* time.Second, 
		WriteTimeout: 10*time.Second, 
		ErrorLog: slog.NewLogger(logger.Handler(), slog.LevelError), 
	}
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe() 
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}