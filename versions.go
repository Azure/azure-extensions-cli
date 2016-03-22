package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func listVersions(c *cli.Context) {
	cl := mkClient(getFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	v, err := cl.ListVersions()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Namespace", "Type", "Version", "Replicated?", "Internal?", "Regions"})
	data := [][]string{}
	for _, e := range v.Extensions {
		data = append(data, []string{e.Ns, e.Name, e.Version, fmt.Sprintf("%v", e.ReplicationCompleted), fmt.Sprintf("%v", e.IsInternal), e.Regions})
	}
	table.AppendBulk(data)
	table.Render()
}
