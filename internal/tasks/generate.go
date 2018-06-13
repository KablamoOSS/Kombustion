package tasks

import (
	"log"

	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/urfave/cli"
)

// GenerateFlags for kombustion upsert
var GenerateFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. ( --param Env=dev --param BucketName=test )",
	},
	cli.BoolFlag{
		Name:  "no-base-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
}

func init() {
	GenerateFlags = append(CloudFormationStackFlags)
}

// Generate a template and save it to disk, without upserting it
func Generate(c *cli.Context) {
	paramMap := cloudformation.GetParamMap(c)

	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		log.Fatal(err)
	}

	manifestFile := manifest.FindAndLoadManifest()
	if err != nil {
		log.Fatal(err)
	}

	// load all plugins
	loadedPlugins := plugins.LoadPlugins(manifestFile, lockFile)

	// if in devMode optionally load a devMode plugin

	tasks.GenerateTemplate(cloudformation.GenerateParams{
		Filename:           c.Args().Get(0),
		EnvFile:            c.String("env-file"),
		Env:                c.String("env"),
		DisableBaseOutputs: c.Bool("no-base-outputs"),
		ParamMap:           paramMap,
		Plugins:            loadedPlugins,
	})
}
