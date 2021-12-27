package response

type Response struct {
	ErrCode      string  `json:"ErrCode"`
	ErrDesc      string  `json:"ErrDesc"`
	Response     bool    `json:"Response"`
	TotalResults string  `json:"totalResults"`
	Search       []Movie `json:"Search"`
}

type Result struct {
	TotalResults string  `json:"totalResults"`
	Search       []Movie `json:"Search"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type ResponseDetail struct {
	ErrCode  string      `json:"ErrCode"`
	ErrDesc  string      `json:"ErrDesc"`
	Response bool        `json:"Response"`
	Movie    MovieDetail `json:"Movie"`
}

type MovieDetail struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

type ResponseError struct {
	Response string `json:"Response"`
	Error    string `json:"Error"`
}
