package service

import (
	"encoding/json"
	"fmt"
	"log"
	"stockbit-test/omdb/constants"
	response "stockbit-test/omdb/dto"
	"stockbit-test/omdb/utils/http"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	envs map[string]string
	err  error
)

type OmdbService struct {
}

func (o OmdbService) List(page int, search string, apikey string) response.Response {
	var res response.Response
	var resError response.ResponseError
	var result response.Result

	envs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := envs["OMDB_URL"]

	resp, err := SendGetHTTP(url + "?apikey=" + apikey + "&s=" + search + "&page=" + strconv.Itoa(page))
	if err != nil {
		res.Response = false
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		return res
	}

	err = json.Unmarshal(resp, &resError)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		res.Response = false
		return res
	}

	if resError.Response == "False" {
		res.Response = false
		res.ErrCode = constants.ERR_CODE_05
		res.ErrDesc = constants.ERR_CODE_05_MSG
		return res
	}

	json.Unmarshal(resp, &result)

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Search = result.Search
	res.TotalResults = result.TotalResults
	res.Response = true

	return res
}

func (o OmdbService) Get(imdbID string, apikey string) response.ResponseDetail {
	var res response.ResponseDetail
	var result response.MovieDetail

	envs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := envs["OMDB_URL"]

	fmt.Println("url: ", url)

	resp, err := SendGetHTTP(url + "?apikey=" + apikey + "&i=" + imdbID)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_06
		res.ErrDesc = constants.ERR_CODE_06_MSG
		res.Response = false
		return res
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		res.Response = false
		return res
	}

	if result.Response == "False" {
		res.Response = false
		res.ErrCode = constants.ERR_CODE_06
		res.ErrDesc = constants.ERR_CODE_06_MSG
		return res
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Response = true
	res.Movie = result
	return res
}

// sendGetHTTP ...
func SendGetHTTP(url string) ([]byte, error) {
	return http.HttpGetByParam(url, "60s")
}
