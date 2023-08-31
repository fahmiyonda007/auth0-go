package controllers

import (
	"math"
	"strconv"
	"strings"
)

type Filters struct {
	Page     int
	PageSize int
	Sort     string
	SortList []string
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) SortOrder() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func Limit(pageSize int) int {
	return pageSize
}

func Offset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

func CalculateMetadata(totalRecords int, page int, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

func ValidateFilter(m Metadata, page int, pageSize int) string {
	if page < 1 {
		return "page is not lower than 1"
	}

	if pageSize < 1 {
		return "length is not lower than 1"
	}

	if page > m.LastPage {
		return "page is not greater than " + strconv.Itoa(m.LastPage)
	}

	if pageSize > 50 {
		return "length is not greater than 50"
	}

	return ""
}
