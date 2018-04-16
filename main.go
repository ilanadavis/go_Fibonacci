package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/julienschmidt/httprouter"
        "strconv"
  )

type Response struct {
  //input number for n
  Number int
  Results []string
}

/*
type JsonResponse struct {
  Meta interface{} `json:"meta"`
  Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
  Error *ApiError `json:"error"`
}

type ApiError struct {
  Status int `json:"status"`
  Title string `json:"title"`
}
*/

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
  		fmt.Println(f())
	  }
	}
}


func main() {
  router := httprouter.New()
//  router.GET("/", Results)
  router.GET("/fibonacci/:number", Results)

  log.Fatal(http.ListenAndServe(":8080", router))
}