package query

import (
	"strconv"
)

type SQLCount string
type SQL string

func BuildSelectQuery(fields map[string]string, strQueryFrom string,
	strQueryPage string, strFilter string, page Pageable) (SQL, SQLCount) {

	var querySQL SQL
	var querySQLCount SQLCount

	//FIELDS
	strQueryFields := BuildQueryFields(fields)
	//SORT
	strQuerySort := BuildQuerySort(page)
	//FILTER?

	strQuerySQL := "SELECT " + strQueryFields + " FROM " + strQueryFrom + " WHERE 1=1 " + strFilter + strQuerySort + strQueryPage
	strQuerySQLCount := " SELECT count(*)  as total_rows  " + " FROM " + strQueryFrom + " WHERE 1=1 " + strFilter

	querySQLCount = SQLCount(strQuerySQLCount)
	querySQL = SQL(strQuerySQL)
	return querySQL, querySQLCount
}

func BuildQuerySort(page Pageable) string {
	strQuerySort := " "
	mapSort := page.GetSort()
	if len(mapSort) > 0 {
		strQuerySort = " ORDER BY "
	}
	var totalSort = len(mapSort)
	var idxSort = 1
	for field, direction := range mapSort {
		strQuerySort = strQuerySort + field + " " + direction
		if idxSort < totalSort {
			strQuerySort = strQuerySort + " , "
		}
		idxSort++
	}
	return strQuerySort
}

func BuildQueryFields(fields map[string]string) string {
	var totalFields = len(fields)
	var idxFields = 1
	strQueryFields := " "
	for field, alias := range fields {
		strQueryFields = strQueryFields + field + " as " + alias
		if idxFields < totalFields {
			strQueryFields = strQueryFields + " , "
		}
		idxFields++
	}
	return strQueryFields
}

func BuildQueryFilter(filters []FilterParameter, mapFields map[string]string) (query string, values []interface{}, qtd int) {
	var total = len(filters)
	var idx = 1
	strQueryFilter := " WHERE 1=1 AND "
	for i, filter := range filters {
		if filter.Operator == LIKE {
			strQueryFilter = strQueryFilter + "lower(" + mapFields[filter.Field] + ") " + (filter.Operator.ToString()) + " $" + strconv.Itoa(i+1)
		} else {
			strQueryFilter = strQueryFilter + mapFields[filter.Field] + " " + (filter.Operator.ToString()) + " $" + strconv.Itoa(i+1)
		}

		values = append(values, filter.Value)
		if idx < total {
			strQueryFilter = strQueryFilter + " AND "
		}
		idx++
	}
	return strQueryFilter, values, idx
}

func BuildQueryPage(qtd int) (query string) {

	query = " OFFSET $" + strconv.Itoa(qtd) + " LIMIT $" + strconv.Itoa(qtd+1)
	return
}

func BuildQueryFilterPage(filters []FilterParameter, pageable Pageable, mapFields map[string]string,
	sqlStatement string, sqlStatementCount string) ([]interface{}, []interface{}, string, string) {
	strSqlSortField := BuildQuerySort(pageable)

	strSqlFilterField, valuesFilter, qtd := BuildQueryFilter(filters, mapFields)
	strSqlPage := BuildQueryPage(qtd)

	values := append(valuesFilter, pageable.GetOffset(), pageable.GetPageSize())

	//var like = "%" + strings.ToLower(strQuery) + "%"

	var sqlQuery = sqlStatement + strSqlFilterField + strSqlSortField + strSqlPage

	var sqlQueryCount = sqlStatementCount + strSqlFilterField + strSqlSortField

	return valuesFilter, values, sqlQuery, sqlQueryCount
}

func BuildPage(content interface{}, totalElements int64, totalPageElements int32, pageable Pageable) Page {

	page := PageData{}
	page.Content = content
	page.PageSize = pageable.GetPageSize()
	page.PageNumber = pageable.GetPageNumber()
	page.TotalPage = TotalPage(pageable.GetPageSize(), totalElements)
	page.TotalElements = totalElements
	page.PageNumberOfElements = totalPageElements
	return page
}
