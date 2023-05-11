package docs

//go:generate swagger generate spec -o ../public/swagger.json --scan-models

// swagger:route POST /vacancies/search vacancy vacancySearchRequest
// Парсинг вакансий с habr-а.
//
// parameters:
// + name: query
// in: query
// description: Query string
// required: true
//
// responses:
//  200: description: Ok.
//  400: description: Bad request
//	500: description: Internal server error

// swagger:route POST /vacancies/delete vacancy vacancyDeleteRequest
// Удаление вакансии.
//
// parameters:
// + name: id
// in: query
// description: Vacancy ID
// required: true
//
// responses:
//  200:
//  400: description: Bad request
//	500: description: Internal server error

// swagger:route POST /vacancies/get vacancy vacancyGetRequest
// Получение вакансии по id.
//
// parameters:
// + name: id
// in: query
// description: Vacancy ID
// required: true
//
// responses:
//  200: vacancySearchResponse
//  400: description: Bad request
//	500: description: Internal server error

// swagger:response vacancySearchResponse
type vacancyGetResponse struct {
	// in:body
	Body model.Vacancy
}

// swagger:route Get /vacancies/list vacancy vacancyListRequest
// Все вакансии.
// responses:
//  200: vacancyListResponse
//  400: description: Bad request
//	500: description: Internal server error

// swagger:response vacancyListResponse
type vacancyListResponse struct {
	// in:body
	Body []model.Vacancy
}
