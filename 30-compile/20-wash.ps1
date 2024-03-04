<#---
title: Wash Profile Data
connection: sharepoint
api: post
tag: wash

input: profiledata.json
output: washedprofiledata.json
---
Here we call the `nexiintra-profile` cli with the argumenet to wash the profile data. 


You find the source for this in .koksmat/app/execution/wash.go

#>

nexiintra-profile wash run "$env:WORKDIR/profiledata.json" "$env:WORKDIR/washedprofiledata.json"
