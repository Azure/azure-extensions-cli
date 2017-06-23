package main

import (
	"fmt"
	"os"

	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func listVersions(c *cli.Context) {
	cl := mkClient(checkFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	v, err := cl.ListVersions()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	json := c.Bool(flJSON.Name)
	var f func(_ ListVersionsResponse) error
	if json {
		f = printListVersionsAsJSON
	} else {
		f = printListVersionsAsTable
	}
	if err := f(v); err != nil {
		log.Fatal(err)
	}
}

func printListVersionsAsJSON(r ListVersionsResponse) error {
	b, err := json.MarshalIndent(r.Extensions, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format as json: %+v", err)
	}
	fmt.Fprintf(os.Stdout, "%s", string(b))
	return nil
}

func printListVersionsAsTable(v ListVersionsResponse) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(4000)
	table.SetHeader([]string{"Namespace", "Type", "Version", "Replicated?", "Internal?", "Regions"})
	data := [][]string{}
	for _, e := range v.Extensions {
		data = append(data, []string{e.Ns, e.Name, e.Version, fmt.Sprintf("%v", e.ReplicationCompleted), fmt.Sprintf("%v", e.IsInternal), e.Regions})
	}
	table.AppendBulk(data)
	table.Render()

	return nil
}
