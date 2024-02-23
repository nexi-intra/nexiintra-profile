<#---
title: Forward To Database
connection: sharepoint
api: post
tag: forward
---#>
param ($databasename = "prod2")
if ($null -eq $env:WORKDIR) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
}
if (-not (Test-Path $env:WORKDIR)) {
    $x = New-Item -Path $env:WORKDIR -ItemType Directory 
}
$password = ""
$username = ""


$vars = kubectl exec "pod/$databasename-mongos-0" -n percona -- "env" 
    
foreach ($var in $vars) {
    $s = $var -split "="
    if ($s[0] -eq "MONGODB_DATABASE_ADMIN_PASSWORD") {
        $password = $s[1]
    }

    if ($s[0] -eq "MONGODB_DATABASE_ADMIN_USER") {
        $username = $s[1]
    }
}
    
    

$forwarded = "mongodb://"+$username+":"+$password+ "@localhost:27017/?directConnection=true&authMechanism=DEFAULT&tls=false"

write-host "Connection string is copied to clipboard"

Set-Clipboard -Value $forwarded
$destinationDir = "$env:WORKDIR/log/kubectl-portforward-mongo"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}

Push-Location
set-location $destinationDir

kubectl port-forward "pod/$databasename-mongos-0" -n percona 27017:27017
