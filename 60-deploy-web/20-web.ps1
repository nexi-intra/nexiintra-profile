<#---
title: Create Web deployment file
input: profiledataurl.json
output: web.yaml
tag: createdeploymentfile
api: post
---#>


$inputFile = join-path  $env:KITCHENROOT "nexiintra-profile" ".koksmat","koksmat.json"



if (!(Test-Path -Path $inputFile) ) {
   Throw "Cannot find file at expected path: $inputFile"
} 

$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

$appname = "nexiintra-profile"

$image = "ghcr.io/koksmat-com/$($appname):$($version)"

$deployment = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname
spec:
  selector:
    matchLabels:
      app: $appname
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname
    spec: 
      containers:
      - name: $appname
        image: $image
        ports:
          - containerPort: 3001
        env:
        - name: SPAUTH_TENANTID
          value: $($env:SPAUTH_DOMAIN)
        - name: SPAUTH_CLIENTID
          value: $($env:SPAUTH_CLIENTID)
        - name: SPAUTH_CLIENTSECRET
          value: $$($env:SPAUTH_CLIENTSECRET)

      
---
apiVersion: v1
kind: Service
metadata:
  name: $appname
  labels:
    app: $appname
    service: $appname
spec:
  ports:
  - name: http
    port: 5301
    targetPort: 3001
  selector:
    app: $appname

"@


$deployment | Out-File -FilePath "$env:WORKDIR/web.yaml" -Encoding utf8