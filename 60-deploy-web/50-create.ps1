<#---
title: Create Kubernetes Deployment 
tag: create
api: post
input: web.yaml
---#>
param (
    $namespace = "magicbox-christianiabpos"
   
)

$deploymentFile = "$env:WORKDIR/web.yaml"

kubectl create -f $deploymentFile --namespace $namespace
