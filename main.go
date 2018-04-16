package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/julienschmidt/httprouter"
        "strconv"
        "github.com/rs/cors"
  )

func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}

func Results(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if number, err := strconv.Atoi(params.ByName("number")); err == nil {
  	f := fibonacci()
  	for i := 0; i < number; i++ {
  	  	fmt.Fprint(w, f())
//  		print by collecting all with comma in between
//      convert into json object
	  }
	}
}

func main() {
  router := httprouter.New()
  router.GET("/fibonacci/:number", Results)
  
  handler := cors.Default().Handler(router)
  log.Fatal(http.ListenAndServe(":8080", handler))
}