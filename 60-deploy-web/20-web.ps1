<#---
title: Create Web deployment file
input: profiledataurl.json
output: web.yaml
tag: createdeploymentfile
api: post
---



NEWSCHANNELSBLOB=https://magicbox.blob.core.windows.net/cache/profiledata.json?se=2026-08-25T19%3A55%3A34Z&sp=r&sv=2022-11-02&sr=b&sig=%2B%2FkGv%2FTNHoK9cI%2BHdvVxCbCJXv4C8k8kL5nify3F0Cg%3D
VALIDDOMAINSBLOB=https://magicbox.blob.core.windows.net/cache/whilelisteddomains.json?se=2026-08-25T20%3A10%3A05Z&sp=r&sv=2022-11-02&sr=b&sig=13rv5iigS8u5jdoqvtCkl8j%2Bokc0iZLkl33kR%2BUekyI%3D


#>


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
        - name: NEWSCHANNELSBLOB
          value: $($env:NEWSCHANNELSBLOB)
        - name: VALIDDOMAINSBLOB
          value: $($env:VALIDDOMAINSBLOB)      
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