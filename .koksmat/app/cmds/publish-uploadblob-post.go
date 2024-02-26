// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Publish Blob Storage
---
*/
package cmds

import (
	"context"
	"os"
	"path"

	"github.com/365admin/nexiintra-profile/execution"
	"github.com/365admin/nexiintra-profile/utils"
)

func PublishUploadblobPost(ctx context.Context, body []byte, args []string) (*string, error) {
	inputErr := os.WriteFile(path.Join(utils.WorkDir("nexiintra-profile"), "whilelisteddomains.json"), body, 0644)
	if inputErr != nil {
		return nil, inputErr
	}

	_, pwsherr := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "41-publish-domains", "10-blobstorage.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
