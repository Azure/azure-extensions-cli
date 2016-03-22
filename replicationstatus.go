package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func replicationStatus(c *cli.Context) {
	cl := mkClient(getFlag(c, flMgtURL.Name), checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	ns, name, version := checkFlag(c, flNamespace.Name), checkFlag(c, flName.Name), checkFlag(c, flVersion.Name)
	log.Debug("Requesting replication status.")
	rs, err := cl.GetReplicationStatus(ns, name, version)
	if err != nil {
		log.Fatal("Cannot fetch replication status: %v", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Location", "Status"})
	data := [][]string{}
	for _, s := range rs.Statuses {
		data = append(data, []string{s.Location, s.Status})
	}
	table.AppendBulk(data)
	table.Render()
}
