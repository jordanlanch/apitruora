package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"sort"
	"os"
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
	urlAPIServer, err := url.Parse(os.Getenv("URL_APISERVER"))
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
	responseAPI := parseData(&apiServerModel, domain)
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
	Items := dbmodels.Items{}
	Items.Domain = domain
	responses := []dbmodels.Response{}
	responses = append(responses,*responseAPI)
	Items.Response  = responses

	_, err = persistence.CreateItems(db, &Items)
	if err != nil {
		return nil, err  
	}

	return responseAPI, nil
}
// GetItems get all Data
func GetItems(db *gorm.DB) (*[]dbmodels.Items, error) {
	var Items = []dbmodels.Items{}
	var dbResult = db.Set("gorm:auto_preload", true).Find(&Items)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &Items, nil
}

func parseData(apiServerModel *models.ApiServerResponse,domain string) *dbmodels.Response{
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
	resp.Logo, resp.Title, _ = getLogoAndTitle(domain)
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

func getLogoAndTitle(domain string) (string,string,error){
	resp, err := soup.Get("https://www."+domain)
	if err != nil {
		return "Error Get Logo","Error Get Title",err
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("link")
	img,title := "",""
	for _, link := range links {
		if link.Attrs()["type"] == "image/x-icon" {
			img =link.Attrs()["href"]
		}
	}
	titleFind := doc.FindAll("title")
	for i, t := range titleFind {
		if i == 0{
			title=t.Text()
		}
	}
	return img,title,nil
}

