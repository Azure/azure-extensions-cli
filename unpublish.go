package main

import (
	"bytes"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func unpublishVersion(c *cli.Context) {
	p := struct {
		Namespace, Name, Version string
	}{
		Namespace: checkFlag(c, flNamespace.Name),
		Name:      checkFlag(c, flName.Name),
		Version:   checkFlag(c, flVersion.Name)}

	manifestXml := `<?xml version="1.0" encoding="utf-8" ?>
<ExtensionImage xmlns="http://schemas.microsoft.com/windowsazure"  xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
  <!-- WARNING: Ordering of fields matter in this file. -->
  <ProviderNameSpace>{{.Namespace}}</ProviderNameSpace>
  <Type>{{.Name}}</Type>
  <Version>{{.Version}}</Version>
  <IsInternalExtension>true</IsInternalExtension>
  <IsJsonExtension>true</IsJsonExtension>
</ExtensionImage>`
	tpl, err := template.New("unregisterManifest").Parse(manifestXml)
	if err != nil {
		log.Fatalf("template parse error: %v", err)
	}

	var b bytes.Buffer
	if err = tpl.Execute(&b, p); err != nil {
		log.Fatalf("template execute error: %v", err)
	}

	cl := mkClient(checkFlag(c, flSubsID.Name), checkFlag(c, flSubsCert.Name))
	op, err := cl.UpdateExtension(b.Bytes())
	if err != nil {
		log.Fatalf("UpdateExtension failed: %v", err)
	}
	lg := log.WithField("x-ms-operation-id", op)
	lg.Info("UpdateExtension operation started.")
	if err := cl.WaitForOperation(op); err != nil {
		lg.Fatalf("UpdateExtension failed: %v", err)
	}
	lg.Info("UpdateExtension operation finished.")
}
