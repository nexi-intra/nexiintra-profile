<#---
magicbutton: Update News Channels metadata
title: Update News Channels metadata
output: magicbutton-newschannelsupdate.json
tag: channels
---

We start by reading data from SharePoint
#>


write-host "Reading data from SharePoint" -ForegroundColor Green -NoNewline

write-host " Countries" -NoNewline
nexiintra-profile read countries

write-host " Channels" -NoNewline
nexiintra-profile read channels

write-host " Categories" -NoNewline
nexiintra-profile read categories

write-host " Units" -NoNewline
nexiintra-profile read units

write-host ""
<#
Data are compiled and washed
#>
write-host "Processing" -ForegroundColor Green -NoNewline
write-host " Bundling" -NoNewline
nexiintra-profile compile lookupvalues
write-host " Washing" -NoNewline
nexiintra-profile wash run "$env:WORKDIR/profiledata.json" "$env:WORKDIR/washedprofiledata.json"
write-host ""

<#
We upload the data to the blob storage
#>
write-host "Publishing" -ForegroundColor Green -NoNewline
write-host " Blob" -NoNewline
nexiintra-profile publishchannels uploadblob "$env:WORKDIR/washedprofiledata.json"
write-host ""
<#
We get the Sas url to be used for getting the data from the blob storage
#>


nexiintra-profile publishchannels sasurl

<#
And write out the URL which can be used to read the data
#>
$SasUrl = Get-Content "$env:WORKDIR/profiledataurl.json" | ConvertFrom-Json

write-host "Done" -ForegroundColor Green

write-host "**************************************************" -ForegroundColor Green
write-host "Sasurl",$SasUrl -ForegroundColor Green
write-host "**************************************************" -ForegroundColor Green

