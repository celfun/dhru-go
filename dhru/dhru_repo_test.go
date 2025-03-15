package dhru

import (
	"fmt"
	"os"
	"testing"
)

func TestList_UnmarshalJSON(t *testing.T) {
	jsonData, err := os.ReadFile("services.json")
	if err != nil {
		t.Fatalf("error reading JSON: %v", err)
	}
	services, err := flattenIMEIServiceList(jsonData)
	if err != nil {
		t.Fatalf("error unmarshaling JSON: %v", err)
	}
	for i, service := range services {
		fmt.Printf("\n%d:\n", i)
		println(service.ServiceID)
		println(service.ServiceName)
		println(service.GroupName)
		println(service.Credit)
	}
}

func TestDhru_GetAccountInfo(t *testing.T) {
	jsonData, err := os.ReadFile("services.json")
	if err != nil {
		t.Fatalf("error reading JSON: %v", err)
	}
	accountInfo, err := mapAccountInfo(jsonData)
	if err != nil {
		t.Fatalf("error unmarshaling JSON: %v", err)
	}
	fmt.Printf("\nAccountInfo:\n%v\n", accountInfo)
}
