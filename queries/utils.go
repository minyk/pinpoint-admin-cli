package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
)

type QueryParams struct {
	ApplicationName string `url:"applicationName,omitempty"`
	AgentID string `url:"agentId,omitempty"`
	DurationDays int `url:"durationDays,omitempty"`
}

type Application struct {
	ApplicationName string `json:"applicationName"`
	ServiceType     string `json:"serviceType"`
	Code            int    `json:"code"`
}

type Applications []Application

type GotApplications map[string]Applications

func getJSONToStruct(body []byte) GotApplications {
	dec := json.NewDecoder(bytes.NewReader(body))
	var gotApps GotApplications
	for {
		if err := dec.Decode(&gotApps); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return gotApps
}

func (apps *GotApplications) AppCountAndFilterByName(name string) (map[string]int, map[string]int) {
	appType := make(map[string]int)
	appCount := make(map[string]int)
	for _, app := range(*apps) {
		for _, ap := range(app) {
			if name == ap.ApplicationName || name == "*" {
				appCount[ap.ApplicationName]++
				appCount["*"]++
				appType[ap.ServiceType]++
				appType["*"]++
			}
		}
	}
	return appCount, appType
}

func (apps *GotApplications)GenerateAgentCountMessage(applicationName string) string {
	appCount, appType := apps.AppCountAndFilterByName(applicationName)
	appKeys := getIntMapKeys(appCount)
	typeKeys := getIntMapKeys(appType)
	var str strings.Builder
	str.WriteString(fmt.Sprintf("Application Name: %s",applicationName))
	if len(appCount) > 0 {
		for _, app := range (appKeys) {
			if app.String() != "*" {
				str.WriteString(fmt.Sprintf("\n  Application %s has %d agents.", app.String(), appCount[app.String()]))
			}
		}
		for _, serviceType := range (typeKeys) {
			if serviceType.String() != "*" {
				str.WriteString(fmt.Sprintf("\n  Service type %s has %d agents.", serviceType.String(), appType[serviceType.String()]))
			}
		}
	} else {
		str.WriteString(fmt.Sprintf("\n  Application %s has 0 agents.", applicationName))
	}
	return str.String()
}

func getStringMapKeys( m map[string]string ) []reflect.Value {
	return reflect.ValueOf(m).MapKeys()
}

func getIntMapKeys( m map[string]int ) []reflect.Value {
	return reflect.ValueOf(m).MapKeys()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
