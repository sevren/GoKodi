package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sahilm/fuzzy"
	"github.com/sevren/go-home-kodi/pkgs/kodirpc"
	"github.com/sevren/go-home-kodi/pkgs/middlewares"
	"github.com/sevren/go-home-kodi/pkgs/models"
	log "github.com/sirupsen/logrus"
)

type config struct {
	ServerCreds kodirpc.KodiUser
	ServerURL   string `env:"KODI_SERVER_URL"`
	Token       string `env:"SERVER_TOKEN"`
	Port        int    `env:"SERVER_PORT" envDefault:"8080"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Authenticate(cfg.Token))

	serverURLWithBasic := fmt.Sprintf("http://%s:%s@%s:%s/jsonrpc", cfg.ServerCreds.User, cfg.ServerCreds.pass, cfg.ServerURL, cfg.Port)
	c, err := kodirpc.New(serverURLWithBasic)
	if err != nil {
		log.Fatal(err)
	}

	r.Route("/scanvideolibrary", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, x *http.Request) {

			// extracts the username from the url
			textingrediant := x.URL.Query().Get("q")
			log.Info(x)
			log.Info(textingrediant)

			response, err := c.ScanVideoLibrary()
			if err != nil {
				log.Errorf("Rpc called failed with %s", err)
			}

			if response.Error != nil {
				// check response.Error.Code, response.Error.Message, response.Error.Data  here
				log.Error(response)
			}

			log.Info(response)

		})

	})

	r.Route("/gettvshows", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, x *http.Request) {

			tvShows := models.TVResponse{}
			response, err := c.GetTVShows()
			if err != nil {
				log.Errorf("Rpc called failed with %s", err)
			}

			if response.Error != nil {
				// check response.Error.Code, response.Error.Message, response.Error.Data  here
				log.Error(response)
			}
			log.Info(response)
			response.GetObject(&tvShows)
			for _, show := range tvShows.TvShows {
				log.Infof("%+v", show)
			}

			results := fuzzy.FindFrom("doctor who", tvShows.TvShows)
			for _, r := range results {
				fmt.Println(tvShows.TvShows[r.Index])
			}

		})

	})

	http.ListenAndServe(":9991", r)

}
