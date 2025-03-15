package test

import (
	"fmt"
	"github.com/celfun/dhru-go/dhru"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var apiURL string
var apiKey string
var username string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey = os.Getenv("DHRU_APIKEY")
	username = os.Getenv("DHRU_USERNAME")
	apiURL = os.Getenv("DHRU_URL")
}

func TestDhru_GetAccountInfoSuccess(t *testing.T) {
	dhruApi, err := dhru.FindApi(apiURL)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("api found in: %+v\n", dhruApi)
	dhruSession, err := dhru.NewDhruSession(dhruApi, username, apiKey)
	if err != nil {
		t.Error(err)
	}
	accountInfo, err := dhruSession.GetAccountInfo()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("accountInfo: %+v\n", accountInfo)
}

func TestDhru_GetAccountInfoFail(t *testing.T) {
	dhruApi, err := dhru.FindApi(apiURL)
	if err != nil {
		t.Error(err)
		return
	}
	println(dhruApi)
	dhruSession, err := dhru.NewDhruSession(dhruApi, username, apiKey)
	if err != nil {
		t.Error(err)
	}
	_, err = dhruSession.GetAccountInfo()
	if err == nil {
		t.Error(err)
		return
	}
	fmt.Printf("%v\n", err)
}

func TestDhru_GetImeiList(t *testing.T) {
	dhruApi, err := dhru.FindApi(apiURL)
	if err != nil {
		t.Error(err)
		return
	}
	println(dhruApi)
	dhruSession, err := dhru.NewDhruSession(dhruApi, username, apiKey)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = dhruSession.GetImeiServiceList()
	if err != nil {
		t.Error(err)
		return
	}
}
