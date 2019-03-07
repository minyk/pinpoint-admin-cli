package queries

import (
	"bytes"
	"github.com/google/go-querystring/query"
)
import "github.com/minyk/pinpoint-cli/client"

type Remove struct {
	PrefixCb func() string
	ResponseOK func() []byte
}

func NewRemove() *Remove {
	return &Remove{
		PrefixCb: func() string { return "admin/" },
		ResponseOK: func() []byte { return []byte("\"OK\"")},
	}
}

func (q *Remove)RemoveApplication(applicationName string, durationDays int, inactive bool, locally bool) error {
	if inactive {
		err := q.removeApplicationWithInactivity(applicationName, durationDays)
		if err != nil {
			return nil
		}
	} else {
		err := removeApplicationRemote(applicationName, inactive)
		if err != nil {
			return nil
		}
	}
	return nil
}

// Remove inactive agents of an application
// First, find inactive agents, then remove one by one.
func (q *Remove)removeApplicationWithInactivity(applicationName string, durationDays int) error {
	opts := QueryParams{applicationName, "", durationDays }
	v, err := query.Values(opts)
	if err != nil {
		return err
	}

	response, err := client.HTTPServiceGetQuery("admin/getInactiveAgents.pinpoint",v.Encode())
	if err != nil {
		return err
	}

	apps := getJSONToStruct(response)

	client.PrintMessage("Found %d inactive agents for %s.", len(apps), applicationName)
	for key, _ := range(apps) {
		q.RemoveAgent(applicationName, key)
	}

	return nil
}


func removeApplicationRemote(applicationName string, inactive bool) error {
	opts := QueryParams{applicationName, "", 30 }
	v, err := query.Values(opts)
	if err != nil {
		return err
	}
	response, err := client.HTTPServiceGetQuery("admin/removeAgentId.pinpoint",v.Encode())
	if err != nil {
		return err
	} else {
		client.PrintJSONBytes(response)
	}

	return nil
}

func removeApplicationLocally(applicationName string, inactive bool) error {
	opts := QueryParams{applicationName, "", 30 }
	v, err := query.Values(opts)
	if err != nil {
		return err
	}

	response, err := client.HTTPServiceGetQuery("admin/getInactiveAgents.pinpoint",v.Encode())
	if err != nil {
		return err
	}

	apps := getJSONToStruct(response)
	client.PrintMessage("Total %d agents are found.", len(apps))
	client.PrintMessage("Start deleting one by one.")

	for index, key := range(apps) {
		removeOpts := QueryParams{applicationName, "", 30}
		v, err = query.Values(removeOpts)
		response, err = client.HTTPServiceGetQuery("admin/removeAgentId.pinpoint",v.Encode())
		if err != nil {
			return err
		}
		client.PrintMessage("Delete %s(%d th)...%s", key, index, response)
	}

	return nil
}


// Remove a agent of an application with specific id.
// applicationName and agentId is needed.
func (q *Remove)RemoveAgent(applicationName string, agentId string) error {
	opts := QueryParams{applicationName, agentId, 0 }
	v, err := query.Values(opts)
	if err != nil {
		return err
	}
	response, err := client.HTTPServiceGetQuery("admin/removeAgentId.pinpoint",v.Encode())
	if err != nil {
		return err
	} else {
		if bytes.Equal(response, q.ResponseOK()) {
			client.PrintMessage("Agent %s is successfully removed.", agentId)
		} else {
			client.PrintMessage(string(response))
		}
	}

	return nil
}
