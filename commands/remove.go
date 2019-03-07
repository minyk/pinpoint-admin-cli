package commands

import (
	"github.com/minyk/pinpoint-cli/queries"
)
import "gopkg.in/alecthomas/kingpin.v3-unstable"

type removeHandler struct {
	q         *queries.Remove
	applicationName string
	inactive        bool
	agentID         string
	durationDays    int
	processLocally  bool
}

func (cmd *removeHandler)handleRemoveApplications(a *kingpin.Application, e *kingpin.ParseElement, c *kingpin.ParseContext) error {
	if cmd.agentID == "" {
		return cmd.q.RemoveApplication(cmd.applicationName,cmd.durationDays,cmd.inactive,cmd.processLocally)
	} else {
		return cmd.q.RemoveAgent(cmd.applicationName,cmd.agentID)
	}
}

// HandleScheduleSection
func HandleRemoveSection(app *kingpin.Application, q *queries.Remove) {
	HandleRemoveAppCommands(app.Command("remove", "Remove application entries").Alias("removeapps"), q)
}

func HandleRemoveAppCommands(removeApp *kingpin.CmdClause, q *queries.Remove) {
	cmd := &removeHandler{q: q}
	remove := removeApp.Action(cmd.handleRemoveApplications)
	remove.Arg("applicationName", "Name of application on Pinpoint").Required().StringVar(&cmd.applicationName)
	remove.Flag("agent-id", "Agent ID to remove. If not supplied, all entries are removed from Pinpoint").StringVar(&cmd.agentID)
	remove.Flag("inactive-only", "Remove inactive agent only. Default is true").Default("true").BoolVar(&cmd.inactive)
	remove.Flag("duration-days", "Duration days of inactivity. Default is 30. Should be longer than 30").Default("30").PreAction(CheckDurationDays).IntVar(&cmd.durationDays)
	remove.Flag("process-locally", "Remove agents one by one to prevent server OOM").Default("false").BoolVar(&cmd.processLocally)
}
