package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {

		req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodGet, "http://yahoo.com", nil)
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html", data)
	})
}

// https://www.youtube.com/watch?v=VkGQFFl66X4
