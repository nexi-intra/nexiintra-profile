<#---
title: Read Channels
api: post
connection: sharepoint
tag: channels
output: channels.json
---
#>
param(
    
    [string]$siteUrl = "https://christianiabpos.sharepoint.com/sites/nexiintra-home"
)


$result = "$env:WORKDIR/channels.json"
Connect-PnPOnline -Url $siteUrl  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

$listItems = Get-PnpListItem -List "News Channels"   

write-host "Items in list: $($listItems.Count)"
$items = @()
foreach ($item in $listItems) {
    $item = @{
        ID = $item.FieldValues.ID
        Title = $item.FieldValues.Title
        RelevantUnits = $item.FieldValues.RelevantUnits
        Mandatory = $item.FieldValues.Mandatory
        RelevantCountires = $item.FieldValues.RelevantCountires
        NewsCategory = $item.FieldValues.NewsCategory

       
       
    }
    $items += $item
   
}

$items | ConvertTo-Json -Depth 10 | Out-File -FilePath $result -Encoding utf8NoBOM
