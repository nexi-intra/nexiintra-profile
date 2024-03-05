<#---
title: Delete Kubernetes Deployment 
tag: delete
api: post
input: web.yaml
---#>
param (
    $namespace = "magicbox-demo"
   
)

$deploymentFile = "$env:WORKDIR/web.yaml"

kubectl delete -f $deploymentFile --namespace $namespace
