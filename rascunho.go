package main

import (
	"fmt"
	"github.com/servicesys/gosqlutils/query"
	"github.com/servicesys/gosqlutils/test"
)

func main() {

	fmt.Println("TESTE")

	pageable := query.PageableData{Page: 0, Size: 5}

	fmt.Println(pageable)

	contatoXRepository := test.ContatoXRepository{DBX: test.NewDBX()}

	fmt.Println("------------------------------------")
	pageData := query.PageableData{Page: 7, Size: 5, Sort: map[string]string{"Nome": "DESC", "DateCreate": "DESC"}}
	pageContato, errorName := contatoXRepository.GetByNamePage("lobo", pageData)
	fmt.Println("------------------------------------")

	fmt.Println(errorName)
	fmt.Println(pageContato.GetContent())
	fmt.Println(pageContato.GetTotalElements())

	var contatos = pageContato.GetContent().([]test.Contato)

	fmt.Println(pageContato.GetPageNumber())
	fmt.Println(pageContato.GetPageSize())

	fmt.Println(contatos)

	//""""""""""" FILTER

	var filters = make([]query.FilterParameter, 1)

	filters[0] = query.FilterParameter{Value: "%ana%", Operator: query.LIKE, Field: "Nome"}

	//filters[1] = 		query.FilterParameter{Value: "2020-06-01", Operator: query.GT, Field: "DateCreate"}

	//fmt.Println(query.BuildFilter(filters))
	pageContato2, error2 := contatoXRepository.GetByFilterPage(filters, pageable)

	fmt.Println(pageContato2)
	fmt.Println(error2)

	//*** REPORT
	report, errorReport := contatoXRepository.GetReportContato()
	fmt.Println(report)
	fmt.Println(errorReport)
}
