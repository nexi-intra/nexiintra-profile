<#---
title: Get SAS URL
connection: sharepoint
api: post
tag: geturl
out: whilelisteddomainsaurl.json
---#>

$result = "$env:WORKDIR/whilelisteddomainsaurl.json"
$expiry = (Get-Date).AddMonths(30).ToString("yyyy-MM-ddTHH:mm:ssZ")

$sasurl = az storage blob generate-sas  `
--account-name $env:AZURE_STORAGE_ACCOUNT `
--account-key $env:AZURE_STORAGE_KEY `
--container-name $env:AZURE_STORAGE_CONTAINER `
--name whilelisteddomains.json `
--full-uri `
--permissions r `
--expiry  $expiry

$sasurl | Out-File -FilePath $result -Encoding utf8NoBOM