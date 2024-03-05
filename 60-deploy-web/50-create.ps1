<#---
title: Create Kubernetes Deployment 
tag: create
api: post
input: web.yaml
---#>
param (
    $namespace = "magicbox-demo"
   
)

$deploymentFile = "$env:WORKDIR/web.yaml"

kubectl create -f $deploymentFile --namespace $namespace
