# mixr net/http
### mixr golang net/http

### using

``` go get github.com/b35t/mixr ```

`` Get, Post, Put, Patch, Delete ``

```

func main() {
	mux := http.NewServeMux()

	mux.Handle("/api/user", mixr.Get(func(w http.ResponseWriter, r *http.Request) {
		//code
	}))

	http.ListenAndServe(":1234", mux)
}

```
