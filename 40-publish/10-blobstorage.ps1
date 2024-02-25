<#---
title: Publish Blob Storage
connection: sharepoint
api: post
tag: uploadblob
input: profiledata.json
---#>


az storage blob upload  `
--account-name $env:AZURE_STORAGE_ACCOUNT `
--account-key $env:AZURE_STORAGE_KEY `
--container-name $env:AZURE_STORAGE_CONTAINER `
--overwrite $true  `
--file "$env:WORKDIR/profiledata.json" `
--name profiledata.json
