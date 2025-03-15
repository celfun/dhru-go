package dhru

import (
	"fmt"
	"os"
	"testing"
)

func TestMapImeiServiceList(t *testing.T) {
	jsonData, err := os.ReadFile("services.json")
	if err != nil {
		t.Fatalf("error reading JSON: %v", err)
	}
	services, err := mapImeiServiceList(jsonData)
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

func TestMapAccountInfo(t *testing.T) {
	jsonData, err := os.ReadFile("services.json")
	if err != nil {
		t.Fatalf("error reading JSON: %v", err)
	}
	parsedAccountInfo, err := mapAccountInfo(jsonData)
	if err != nil {
		t.Fatalf("error unmarshaling JSON: %v", err)
	}
	fmt.Printf("\nAccountInfo:\n%v\n", parsedAccountInfo)
}
