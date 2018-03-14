package router

import (
	"blog-api-lvmingyin-com/schema"
	"fmt"
	"github.com/graphql-go/handler"
	"net/http"
)

func Start(listenPort int64) {

	h := handler.New(&handler.Config{
		Schema:   &schema.RootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)

}
