package query

type Operator string

const (
	EQ   Operator = "="
	GT   Operator = ">"
	GTE  Operator = ">="
	LT   Operator = "<"
	LTE  Operator = "<="
	NE   Operator = "<>"
	LIKE Operator = " LIKE "
)

func (operator Operator) ToString() string {
	return string(operator)
}

type FilterParameter struct {
	Field    string
	Operator Operator
	Value    interface{}
}

/*
https://postgres.rest/query-statements/
	$eq 	Matches values that are equal to a specified value.
	$gt 	Matches values that are greater than a specified value.
	$gte 	Matches values that are greater than or equal to a specified value.
	$lt 	Matches values that are less than a specified value.
	$lte 	Matches values that are less than or equal to a specified value.
	$ne 	Matches all values that are not equal to a specified value.
	$in 	Matches any of the values specified in an array.
	$nin 	Matches none of the values specified in an array.
	$null 	Matches if field is null.
	$notnull 	Matches if field is not null.
	$true 	Matches if field is true.
	$nottrue 	Matches if field is not true.
	$false 	Matches if field is false.
	$notfalse 	Matches if field is not false.
	$like 	Matches always cover the entire string.
	$ilike 	Matches case-insensitive always cover the entire string.
*/
