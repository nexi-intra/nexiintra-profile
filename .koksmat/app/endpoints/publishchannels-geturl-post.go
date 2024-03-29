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

	"github.com/swaggest/usecase"

	"github.com/365admin/nexiintra-profile/execution"
)

func PublishchannelsGeturlPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "40-publish-channels", "20-getsasurl.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Publish Blob Storage")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("40-publish-channels")
	return u
}
