// -------------------------------------------------------------------
// Generated by 365admin-publish 
// Service wrapper  v1
// -------------------------------------------------------------------
/*
---
title: Wash Profile Data
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
func CompileWashPost(ctx context.Context, body []byte,args  []string)  ( *schemas.Washedprofiledata, error) {
inputErr := os.WriteFile(path.Join(utils.WorkDir("nexiintra-profile"), "profiledata.json"), body, 0644)
if inputErr != nil {
	return nil,inputErr
}

_, pwsherr := execution.ExecutePowerShell("john","*","nexiintra-profile","30-compile","20-wash.ps1","" )
if (pwsherr != nil) {
	return nil,pwsherr
}

resultingFile := path.Join(utils.WorkDir("nexiintra-profile"), "washedprofiledata.json")
data, err := os.ReadFile(resultingFile)
if (err != nil) {
	return nil,err
}
resultObject := schemas.Washedprofiledata{}
err = json.Unmarshal(data, &resultObject)
return &resultObject,nil

	
// end result mapping

func init(){
	
}
}
