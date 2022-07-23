package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Search string      `json:"search"`
	Limit  int         `json:"limit"`
	Page   int         `json:"page"`
	Sort   string      `json:"sort"`
	Rows   interface{} `json:"rows"`
}

func PaginationRequest(ctx *gin.Context) Pagination {
	search := ""
	limit := 10
	page := 1
	sort := "Id desc"

	query := ctx.Request.URL.Query()

	for k, v := range query {
		queryValue := v[len(v)-1]

		switch k {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break

		case "page":
			page, _ = strconv.Atoi(queryValue)
			break

		case "sort":
			sort = queryValue
			break

		case "search":
			search = queryValue
			break

		}
	}

	return Pagination{
		Search: search,
		Limit:  limit,
		Page:   page,
		Sort:   sort,
	}
}
