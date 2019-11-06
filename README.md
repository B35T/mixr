# mixr net/http
### mixr golang net/http

### using

```

func main() {
	mux := http.NewServeMux()
	var mixr NewMixr

	mux.Handle("/api/user", mixr.Get(func(w http.ResponseWriter, r *http.Request) {
		//code
	}))

	http.ListenAndServe(":1234", mux)
}

```
