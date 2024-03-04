// -------------------------------------------------------------------
// Generated by 365admin-publish 
// Service wrapper  v1
// -------------------------------------------------------------------
/*
---
title: Compile Lookup values
---
*/
package cmds
import (
"context"
"encoding/json"
"os"
"path"
"github.com/nats-io/nats.go"
"github.com/nats-io/nats.go/micro"
"github.com/spf13/cobra"
"github.com/365admin/nexiintra-profile/schemas"
"github.com/365admin/nexiintra-profile/execution"
"github.com/365admin/nexiintra-profile/utils"
)
func CompileLookupvaluesPost(ctx context.Context, args  []string)  ( *schemas.Profiledata, error) {

_, pwsherr := execution.ExecutePowerShell("john","*","nexiintra-profile","30-compile","10-newschannels.ps1","" )
if (pwsherr != nil) {
	return nil,pwsherr
}

resultingFile := path.Join(utils.WorkDir("nexiintra-profile"), "profiledata.json")
data, err := os.ReadFile(resultingFile)
if (err != nil) {
	return nil,err
}
resultObject := schemas.Profiledata{}
err = json.Unmarshal(data, &resultObject)
return &resultObject,nil

	
// end result mapping

func init(){
	
}
}