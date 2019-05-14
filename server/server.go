package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"sort"
    "regexp"
    "strings"
	
    "github.com/likexian/whois-go"
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
	now := time.Now()
	then := now.Add(time.Duration(-1) * time.Hour)
	apiResponseList := []dbmodels.Response{}
	db.Preload("Servers").Where("created_at > ?", then).Order("created_at desc").Find(&apiResponseList)
	for i := 0; i < len(apiResponseList); i++ {
		if i ==0{
			responseAPI.PreviousSslGrade = apiResponseList[i].SslGrade
		}
		if responseAPI.SslGrade != apiResponseList[i].SslGrade && responseAPI.IsDown != apiResponseList[i].IsDown {
			 responseAPI.ServersChanged = true
			 break
		}
		responseAPI.ServersChanged = false
	
		

	}

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
		server.Country,server.Owner = getWhoIs(enpoints[i].IPAddress)
		servers = append(servers,server)
		grade = append(grade,enpoints[i].Grade)

	}
	resp := dbmodels.Response{}
	resp.Servers =servers
	if len(grade) > 0 {
		sort.Strings(grade)
		resp.SslGrade = grade[len(grade)-1]
	}
	resp.Logo, _ = getLogo()
	resp.IsDown = apiServerModel.Status != "READY"
	resp.CreatedAt = time.Now().UTC()
	return &resp
}
func getWhoIs(ip string)(string,string){
	result, err := whois.Whois(ip)
    var re = regexp.MustCompile(`.*Organization.*`)
    var re2 = regexp.MustCompile(`.*Country.*`)
    if err == nil {
        matches := re.FindStringSubmatch(result)
        organization := strings.Split(matches[0], ":")
        
        matches2 := re2.FindStringSubmatch(result)
        country := strings.Split(matches2[0], ":")
		
		return strings.TrimSpace(country[1]), strings.TrimSpace(organization[1])
    }
	return "", ""
}

func getLogo() (string,error){
	resp, err := soup.Get("https://www.truora.com")
	if err != nil {
		return "Error Get",err
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("link")
	img := ""
	for _, link := range links {
		if link.Attrs()["type"] == "image/x-icon" {
			img =link.Attrs()["href"]
			return img,nil
		}
	}
	return img,nil
}

