// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Read Channels
---
*/
package cmds

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/365admin/nexiintra-profile/execution"
	"github.com/365admin/nexiintra-profile/schemas"
	"github.com/365admin/nexiintra-profile/utils"
)

func ReadChannelsPost(ctx context.Context, args []string) (*schemas.Channels, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "20-read", "11-readchannels.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}

	resultingFile := path.Join(utils.WorkDir("nexiintra-profile"), "channels.json")
	data, err := os.ReadFile(resultingFile)
	if err != nil {
		return nil, err
	}
	resultObject := schemas.Channels{}
	err = json.Unmarshal(data, &resultObject)
	return &resultObject, nil

}
