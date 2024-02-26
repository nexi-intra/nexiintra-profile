// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Get SAS URL
---
*/
package endpoints

import (
	"context"

	"github.com/swaggest/usecase"

	"github.com/365admin/nexiintra-profile/execution"
)

func PublishGeturlPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "41-publish-domains", "20-getsasurl.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Get SAS URL")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("41-publish-domains")
	return u
}
