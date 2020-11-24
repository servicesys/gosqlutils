package query

import "github.com/jmoiron/sqlx"

type Report struct {
	Info    map[string]interface{}
	Header  []string
	Records []map[string]interface{}
	Total   int64
}

func GetReport(DBX *sqlx.DB, SQL string) (Report, error) {

	rows, errorRows := DBX.Query(SQL) // Note: Ignoring errors for brevity
	if errorRows != nil {
		return Report{}, errorRows
	}
	cols, _ := rows.Columns()
	var records []map[string]interface{}
	var reportResult = Report{}
	var header = make([]string, len(cols))

	for i, colName := range cols {
		header[i] = colName
	}
	reportResult.Header = header

	var count int64 = 0
	for rows.Next() {

		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]

		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return Report{}, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		//map[columnName:value columnName2:value2 columnName3:value3 ...]
		records = append(records, m)
		count++
	}
	reportResult.Total = count
	reportResult.Records = records
	return reportResult, nil
}
