package router

import (
	"blog-api-lvmingyin-com/schema"
	"blog-api-lvmingyin-com/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"net/http"
)

func Start(listenPort int64) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/graphql", func(c *gin.Context) {
		postQuery, postOK := c.GetPostForm("query")
		getQuery, getOK := c.GetQuery("query")
		if postOK || getOK {
			result := graphql.Do(graphql.Params{
				Schema:        schema.RootSchema,
				RequestString: util.If(postOK, postQuery, getQuery).(string),
			})

			c.JSON(http.StatusOK, result)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "缺少query参数",
		})
	})

	router.Run(fmt.Sprintf(":%d", listenPort))
}
