// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Publish Blob Storage
---
*/
package endpoints

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/swaggest/usecase"

	"github.com/365admin/nexiintra-profile/execution"
	"github.com/365admin/nexiintra-profile/schemas"
	"github.com/365admin/nexiintra-profile/utils"
)

func PublishchannelsUploadblobPost() usecase.Interactor {
	type Request struct {
		Body schemas.Washedprofiledata `json:"body" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {
		body, inputErr := json.Marshal(input.Body)
		if inputErr != nil {
			return inputErr
		}

		inputErr = os.WriteFile(path.Join(utils.WorkDir("nexiintra-profile"), "washedprofiledata.json"), body, 0644)
		if inputErr != nil {
			return inputErr
		}

		_, err := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "40-publish-channels", "10-blobstorage.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Publish Blob Storage")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Publish Channels")
	return u
}
