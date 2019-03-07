package commands

import (
	"github.com/minyk/pinpoint-cli/queries"
)
import "gopkg.in/alecthomas/kingpin.v3-unstable"

type getHandler struct {
	q               *queries.Get
	applicationName string
	inactive        bool
	agentID         string
	durationDays    int
	jsonOnly        bool
}

func (cmd *getHandler)handleGetApplications(a *kingpin.Application, e *kingpin.ParseElement, c *kingpin.ParseContext) error {
	if cmd.inactive {
		return cmd.q.GetInactiveAgents(cmd.applicationName, cmd.durationDays, cmd.jsonOnly)
	} else {
		return cmd.q.GetAgentsIDMap(cmd.applicationName, cmd.jsonOnly)
	}
}

func HandleGetSection(app *kingpin.Application, q *queries.Get) {
	HandleGetAppCommands(app.Command("get", "Get application entries").Alias("getapps"), q)
}

func HandleGetAppCommands(applications *kingpin.CmdClause, q *queries.Get) {
	cmd := &getHandler{q: q}
	get := applications.Action(cmd.handleGetApplications)
	get.Arg("applicationName", "Name of application on Pinpoint. If not specified, return all entries").Default("*").StringVar(&cmd.applicationName)
	get.Flag("inactive-only", "Return inactive agents only").Default("false").BoolVar(&cmd.inactive)
	get.Flag("duration-days", "Inactive duration when inactive is true. Should be greater than 30").Default("30").PreAction(CheckDurationDays).IntVar(&cmd.durationDays)
	get.Flag("raw-json", "Just print raw JSON result").Default("false").BoolVar(&cmd.jsonOnly)
}