// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Generate SAS URL
---
*/
package cmds

import (
	"context"

	"github.com/365admin/nexiintra-profile/execution"
)

func PublishchannelsSasurlPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "40-publish-channels", "20-getsasurl.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
