package test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gosqlutils/query"
	"strings"
)

type ContatoXRepository struct {
	DBX *sqlx.DB
}

func NewContatoXRepository(dbx *sqlx.DB) *ContatoXRepository {

	return &ContatoXRepository{DBX: dbx}

}

func (repository *ContatoXRepository) Save(contato Contato) (int64, error) {

	sqlStatement := `INSERT INTO public.contato (cont_nome, cont_celular, cont_email, cont_dt_criacao) VALUES($1 , $2 , $3 , now())
	`
	result, err := repository.DBX.Exec(sqlStatement, contato.Nome, contato.Celular, contato.Email)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repository *ContatoXRepository) GetByID(id int64) (Contato, error) {

	sqlStatement := `SELECT cont_id as ID , cont_nome as Nome , 
                                             cont_celular as Celular, 
                                             cont_email as Email, 
                                             cont_dt_criacao as DateCreate 
                                             FROM contato
                                             WHERE 1=1  AND cont_id= $1`

	contato := Contato{}
	err := repository.DBX.Get(&contato, sqlStatement, id)
	if err != nil {
		return contato, err
	}
	return contato, nil
}

func (repository *ContatoXRepository) GetByNamePage(strQuery string, pageable query.Pageable) (query.Page, error) {

	sqlStatementCount := `SELECT   count(*)
                                             FROM contato
                                             WHERE 1=1  AND lower (cont_nome)  like $1`

	sqlStatement := `SELECT cont_id as ID , cont_nome as Nome , 
                                             cont_celular as Celular, 
                                             cont_email  as Email,
                                             cont_dt_criacao as DateCreate
                                             FROM contato
                                             WHERE 1=1  AND lower (cont_nome)  like $3 
                                             %s 
                                             OFFSET $1 LIMIT $2` // as Email

	var contatos = []Contato{}

	strSqlSortField := query.BuildQuerySort(pageable)
	var like = "%" + strings.ToLower(strQuery) + "%"

	var sqlSort = fmt.Sprintf(sqlStatement, strSqlSortField)
	//fmt.Println(sqlSort)
	err := repository.DBX.Select(&contatos, sqlSort, pageable.GetOffset(), pageable.GetPageSize(), like)
	if err != nil {
		return nil, err
	}
	var total int64
	errcount := repository.DBX.Get(&total, sqlStatementCount, like)
	if errcount != nil {
		return nil, errcount
	}
	page := query.BuildPage(contatos, total, int32(len(contatos)), pageable)
	return page, nil
}

func (repository *ContatoXRepository) GetByNamePageNovo(strQuery string, pageable query.Pageable) (query.Page, error) {

	var contatos = []Contato{}

	mapFields := map[string]string{"cont_id": "ID", "cont_nome": "Nome", "cont_celular": "Celular", "cont_email": "Email",
		"cont_dt_criacao": "DateCreate"}
	strQueryTableFrom := "contato"
	strQueryPage := " OFFSET $2 LIMIT $3 "
	strFilter := "AND lower (cont_nome)  like $1"

	likeValue := "%" + strings.ToLower(strQuery) + "%"

	sqlStatementQuery, sqlStatementQueryCount := query.BuildSelectQuery(
		mapFields,
		strQueryTableFrom,
		strQueryPage,
		strFilter,
		pageable)

	fmt.Println(sqlStatementQuery)
	fmt.Println(sqlStatementQueryCount)

	//	//TODO ALTERAR PARA UTILIZAR NAME QUERY

	err := repository.DBX.Select(&contatos, string(sqlStatementQuery), "%"+strings.ToLower(strQuery)+"%", pageable.GetOffset(), pageable.GetPageSize())
	if err != nil {
		return nil, err
	}
	var total int64
	errcount := repository.DBX.Get(&total, string(sqlStatementQueryCount), likeValue)
	if errcount != nil {
		return nil, errcount
	}
	//page builder
	return query.BuildPage(contatos, total, int32(len(contatos)), pageable), nil
}

func (repository *ContatoXRepository) GetByFilterPage(filters []query.FilterParameter, pageable query.Pageable) (query.Page, error) {

	mapFields := map[string]string{"Nome": "cont_nome", "DateCreate": "cont_dt_criacao"}

	sqlStatementCount := `SELECT   count(*) FROM contato`

	sqlStatement := `SELECT cont_id as ID , cont_nome as Nome , 
                                             cont_celular as Celular, 
                                             cont_email  as Email,
                                             cont_dt_criacao as DateCreate
                                             FROM contato` // as Email

	var contatos = []Contato{}

	valuesFilter, values, sqlQuery, sqlQueryCount := query.BuildQueryFilterPage(filters, pageable, mapFields, sqlStatement, sqlStatementCount)

	err := repository.DBX.Select(&contatos, sqlQuery, values...)
	if err != nil {
		return nil, err
	}
	var total int64
	errcount := repository.DBX.Get(&total, sqlQueryCount, valuesFilter...)
	if errcount != nil {
		return nil, errcount
	}
	page := query.BuildPage(contatos, total, int32(len(contatos)), pageable)
	return page, nil
}

func (repository *ContatoXRepository) GetReportContato() (query.Report, error) {

	sqlStatement := `select cont_nome , cont_celular , count(*) as qtd  from contato c 
                     group by cont_nome , cont_celular
                     order by qtd desc`

	return query.GetReport(repository.DBX, sqlStatement)

}
