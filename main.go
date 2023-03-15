package main

import (
	"log"
	"net/http"
	"os"
)

func indexHandler(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			talks := []string{}
			entries, err := os.ReadDir("./dist")
			if err != nil {
				log.Fatal(err)
				return
			}

			for _, e := range entries {
				if e.IsDir() {
					talks = append(talks, e.Name())
				}
			}
			if err := template(talks).Render(r.Context(), w); err != nil {
				log.Fatal(err)
				return
			}
			return
		}
		//if strings.HasSuffix(r.URL.Path, "/") {
		//r2 := new(http.Request)
		//*r2 = *r
		//r2.URL = new(url.URL)
		//*r2.URL = *r.URL
		//r2.URL.Path = r.URL.Path + "index.html"
		//r2.URL.RawPath = r.URL.RawPath + "index.html"
		//fs.ServeHTTP(w, r2)
		//return
		//}
		fs.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", indexHandler(http.FileServer(http.Dir("./dist"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
