// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Wash Profile Data
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

func CompileWashPost() usecase.Interactor {
	type Request struct {
		Body schemas.Profiledata `json:"body" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *schemas.Washedprofiledata) error {
		body, inputErr := json.Marshal(input.Body)
		if inputErr != nil {
			return inputErr
		}

		inputErr = os.WriteFile(path.Join(utils.WorkDir("nexiintra-profile"), "profiledata.json"), body, 0644)
		if inputErr != nil {
			return inputErr
		}

		_, err := execution.ExecutePowerShell("john", "*", "nexiintra-profile", "30-compile", "20-wash.ps1", "")
		if err != nil {
			return err
		}

		resultingFile := path.Join(utils.WorkDir("nexiintra-profile"), "washedprofiledata.json")
		data, err := os.ReadFile(resultingFile)
		if err != nil {
			return err
		}
		resultObject := schemas.Washedprofiledata{}
		err = json.Unmarshal(data, &resultObject)
		*output = resultObject
		return err

	})
	u.SetTitle("Wash Profile Data")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Compile")
	return u
}
