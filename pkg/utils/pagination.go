package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Sort       string `json:"sort"`
	Order      string `json:"order"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

const (
	defaultPage  = 1
	defaultLimit = 10
	maxLimit     = 100
)

func GetPaginationFromRequest(c *gin.Context) *Pagination {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	if page < 1 {
		page = defaultPage
	}

	if limit < 1 {
		limit = defaultLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return &Pagination{
		Page: page,
		Limit: limit,
		Sort: sort,
		Order: order,
	}
}

func (p *Pagination) GetOffset() int{
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) SetTotalRows(total int64) {
	p.TotalRows = total
	p.TotalPages = int(math.Ceil(float64(total) / float64(p.Limit)))
}