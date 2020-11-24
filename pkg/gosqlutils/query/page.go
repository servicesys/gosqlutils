package query

import "math"

type Page interface {

	// total number of pages
	GetTotalPages() int32

	// total number of items
	GetTotalElements() int64

	// current page number
	GetPageNumber() int32

	// page size
	GetPageSize() int32

	// number of items on the current page
	GetNumberOfElements() int32

	// list of items on this page
	GetContent() interface{}
}

type PageData struct {
	TotalPage            int32
	TotalElements        int64
	PageNumber           int32
	PageSize             int32
	PageNumberOfElements int32
	Content              interface{}
}

func (p PageData) GetTotalPages() int32 {
	return p.TotalPage
}

func (p PageData) GetTotalElements() int64 {
	return p.TotalElements
}

func (p PageData) GetPageNumber() int32 {
	return p.PageNumber
}

func (p PageData) GetPageSize() int32 {
	return p.PageSize
}

func (p PageData) GetNumberOfElements() int32 {
	return p.PageNumberOfElements
}

func (p PageData) GetContent() interface{} {
	return p.Content
}

func (p PageData) GetTotalPage(pageSize int32, totalElements int64) int32 {

	if pageSize == 0 {
		return 0
	} else {
		totalPage := int32(math.Ceil(float64(totalElements) / float64(pageSize)))
		return totalPage
	}
}

func TotalPage(pageSize int32, totalElements int64) int32 {

	if pageSize == 0 {
		return 0
	} else {
		totalPage := int32(math.Ceil(float64(totalElements) / float64(pageSize)))
		return totalPage
	}
}
