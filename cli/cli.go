/*
Package cli is the entrypoint to writing a custom CLI for a service built using the DC/OS SDK.
It provides a number of standard subcommands that allow users to interact with their service.
*/
package cli

import (
	"fmt"
	"github.com/minyk/pinpoint-cli/commands"
	"github.com/minyk/pinpoint-cli/config"
	"github.com/minyk/pinpoint-cli/queries"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
	"os"
	"strings"
)

// GetModuleName returns the module name, if it was passed in, or an error otherwise.
func GetModuleName() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf(
			"Must have at least one argument for the CLI module name: %s <modname>", os.Args[0])
	}
	return os.Args[1], nil
}

// GetArguments returns an array of the arguments passed into this CLI.
func GetArguments() []string {
	// Exercise validation of argument count:
	if len(os.Args) < 1 {
		return make([]string, 0)
	}
	return os.Args[1:]
}

// HandleDefaultSections is a utility method to allow applications built around this library to provide
// all of the standard subcommands of the CLI.
func HandleDefaultSections(app *kingpin.Application) {
	removeQueries := queries.NewRemove()
	getQueries := queries.NewGet()

	commands.HandleRemoveSection(app, removeQueries)
	commands.HandleGetSection(app, getQueries)
}

// New instantiates a new kingpin.Application and returns a reference to it.
// This contains basic flags that are universally applicable, e.g. --name.
func New() *kingpin.Application {
	// Before doing anything else, check for envvars relating to logging, and update config.Verbose to reflect them.
	// We do this outside of the normal flag handling for two reasons:
	// - We want the Verbose bit to be set as early as possible, even before arg handling starts.
	// - In kingpin, setting an envvar just changes a default value and doesn't trigger any actions.
	if strings.EqualFold(os.Getenv("DCOS_DEBUG"), "true") {
		config.Verbose = true
	} else {
		logLevel := os.Getenv("DCOS_LOG_LEVEL")
		// Treat either "info" or "debug" as verbose:
		if strings.EqualFold(logLevel, "info") || strings.EqualFold(logLevel, "debug") {
			config.Verbose = true
		}
	}

	//var err error
	//config.ModuleName, err = GetModuleName()
	//if err != nil {
	//	client.PrintMessage(err.Error())
	//	os.Exit(1)
	//}
	//app := kingpin.New(fmt.Sprintf("pinpoint-cli %s", config.ModuleName), "")
	app := kingpin.New("pinpoint-cli", "")

	app.GetFlag("help").Short('h') // in addition to default '--help'

	// Enable verbose logging with '-v', in addition to DCOS_DEBUG/DCOS_LOG_LEVEL which are handled above.
	app.Flag("verbose", "Enable extra logging of requests/responses").Short('v').BoolVar(&config.Verbose)

	// --info and --config-schema are required by the main DC/OS CLI:
	// Prints a description of the module.
	app.Flag("info", "Show short description.").Hidden().PreAction(func(*kingpin.Application, *kingpin.ParseElement, *kingpin.ParseContext) error {
		fmt.Fprintf(os.Stdout, "Pinpoint Admin API\n")
		os.Exit(0)
		return nil
	}).Bool()

	app.Flag("pinpoint", "URL of Target Pinpoint. e.g. http://pinpoint:8080").PreAction(commands.CheckURL).URLVar(&config.PinpointURL) //.StringVar(&config.PinpointURL)

	app.Flag("pinpoint-password", "Admin password of Target Pinpoint").StringVar(&config.PinpointPassword)

	return app
}
