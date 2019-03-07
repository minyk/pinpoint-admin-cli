package queries

import (
	"github.com/google/go-querystring/query"
)
import "github.com/minyk/pinpoint-cli/client"

type Get struct {
	PrefixCb func() string
}

func NewGet() *Get {
	return &Get{
		PrefixCb: func() string { return "admin/" },
	}
}

func (q *Get)GetAgentsIDMap(applicationName string, jsonOnly bool) error {

	response, err := client.HTTPServiceGet(q.PrefixCb() + "agentIdMap.pinpoint")
	if err != nil {
		return err
	}

	if jsonOnly {
		client.PrintJSONBytes(response)
	} else {
		apps := getJSONToStruct(response)
		client.PrintMessage(apps.GenerateAgentCountMessage(applicationName))
	}

	return nil
}

func (q *Get)GetInactiveAgents(applicationName string, durationDays int, jsonOnly bool) error {
	if applicationName == "*" {
		response, err := client.HTTPServiceGet(q.PrefixCb() + "agentIdMap.pinpoint")
		if err != nil {
			return err
		}

		apps := getJSONToStruct(response)
		appCount, _ := apps.AppCountAndFilterByName(applicationName)
		names := getIntMapKeys(appCount)
		for _,name := range(names) {
			if name.String() != "*" {
				client.PrintVerbose("Quering inactive agents of %s", name.String())
				opts := QueryParams{name.String(), "", durationDays}
				v, err := query.Values(opts)
				if err != nil {
					return err
				}

				response, err := client.HTTPServiceGetQuery(q.PrefixCb()+"getInactiveAgents.pinpoint", v.Encode())
				if err != nil {
					return err
				}

				if jsonOnly {
					client.PrintJSONBytes(response)
				} else {
					apps := getJSONToStruct(response)
					client.PrintMessage(apps.GenerateAgentCountMessage(name.String()))
				}
			}
		}
	} else {
		opts := QueryParams{applicationName, "", durationDays }
		v, err := query.Values(opts)
		if err != nil {
			return err
		}

		response, err := client.HTTPServiceGetQuery(q.PrefixCb() + "getInactiveAgents.pinpoint",v.Encode())
		if err != nil {
			return err
		}

		if jsonOnly {
			client.PrintJSONBytes(response)
		} else {
			apps := getJSONToStruct(response)
			client.PrintMessage(apps.GenerateAgentCountMessage(applicationName))
		}
	}

	return nil
}

