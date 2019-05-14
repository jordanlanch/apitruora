package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"sort"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/jinzhu/gorm"

	"../models"
	"../dbmodels"
	"../persistence"
)




//GetDataAPIServer return ListProviders from Bogota Address
func GetDataAPIServer(db  *gorm.DB, domain string) (*dbmodels.Response, error) {
	urlAPIServer, err := url.Parse("https://api.ssllabs.com/api/v3/analyze")
	if err != nil {
		return nil, err
	}
	var client = &http.Client{Timeout: 10 * time.Second}
	parameters := url.Values{}
	parameters.Add("host", domain)
	urlAPIServer.RawQuery = parameters.Encode()

	urlRead := urlAPIServer.String()
	resp, err := client.Get(urlRead)
	if err != nil {
		return nil, err
	}
	var apiServerModel models.ApiServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiServerModel); err != nil {
		return nil, err
	}
	responseAPI := parseData(&apiServerModel)
	
	apiResponseList := dbmodels.Response{}
	db.Preload("Servers").Order("created_at desc").Find(&apiResponseList)
	if responseAPI.SslGrade == apiResponseList.SslGrade && responseAPI.IsDown == apiResponseList.IsDown {
		responseAPI.ServersChanged = false
		}else{
			responseAPI.ServersChanged = true

	}
	fmt.Println("| oldApiResponse :\n", responseAPI)
	fmt.Println("| apiResponseList :\n", apiResponseList)
	//save dataBase

	persistence.CreateResponse(db, responseAPI)

	return responseAPI, nil
}

func parseData(apiServerModel *models.ApiServerResponse) *dbmodels.Response{
	enpoints := apiServerModel.Endpoints
	servers := []dbmodels.Servers{}
	server := dbmodels.Servers{}
	grade := []string{}
	for i := 0; i < len(enpoints); i++ {
		server.Address=enpoints[i].IPAddress
		server.SslGrade=enpoints[i].Grade
		server.CreatedAt = time.Now().UTC()
		servers = append(servers,server)
		grade = append(grade,enpoints[i].Grade)
		//TODO find whois Country Owner

	}
	resp := dbmodels.Response{}
	resp.Servers =servers
	resp.ServersChanged = isChange()
	if len(grade) > 0 {
		sort.Strings(grade)
		resp.SslGrade = grade[len(grade)-1]
	}
	resp.PreviousSslGrade = getPreviousSslGrade(resp.SslGrade)
	resp.Logo = getLogo()
	resp.IsDown = apiServerModel.Status != "READY"
	resp.CreatedAt = time.Now().UTC()
	return &resp
}

func getLogo() string{
	resp, err := soup.Get("https://www.truora.com")
	if err != nil {
		return ""
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("link")
	img := ""
	for _, link := range links {
		if link.Attrs()["type"] == "image/x-icon" {
			img =link.Attrs()["href"]
			return img
		}
	}
	return img
}

func getPreviousSslGrade(sslGrade string) string{
	return "E"
}

func isChange()bool{
	return true
}

