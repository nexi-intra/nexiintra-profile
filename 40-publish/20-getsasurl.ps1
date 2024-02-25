<#---
title: Publish Blob Storage
connection: sharepoint
api: post
tag: geturl
out: profiledataurl.json
---#>

$result = "$env:WORKDIR/profiledataurl.json"
$expiry = (Get-Date).AddMonths(30).ToString("yyyy-MM-ddTHH:mm:ssZ")

$sasurl = az storage blob generate-sas  `
--account-name $env:AZURE_STORAGE_ACCOUNT `
--account-key $env:AZURE_STORAGE_KEY `
--container-name $env:AZURE_STORAGE_CONTAINER `
--name profiledata.json `
--full-uri `
--permissions r `
--expiry  $expiry

$sasurl | Out-File -FilePath $result -Encoding utf8NoBOM