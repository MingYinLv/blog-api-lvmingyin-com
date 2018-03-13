package main

import (
	"blog-api-lvmingyin-com/schema"
	"github.com/graphql-go/graphql"
	"os"
	"encoding/json"
)

func main() {
	result := graphql.Do(graphql.Params{
		Schema:        schema.RootSchema,
		RequestString: "{login(username:\"lvmingyin\",password:\"lvmingyin123\"){username,password}}",
	})

	json.NewEncoder(os.Stdout).Encode(result)
}
