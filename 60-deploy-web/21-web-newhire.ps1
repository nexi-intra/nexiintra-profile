<#---
title: Web (new hire) deploy to production
input: profiledataurl.json
tag: webdeployproduction
api: post
---
We start by finding which version tag to use
#>

$appname = "nexiintra-profile"
$imagename = "nexiintra-profile"
$dnsname = "newhire.home.nexi-intra.com"
$inputFile = join-path  $env:KITCHENROOT $appname ".koksmat", "koksmat.json"
$port = "4323"
if (!(Test-Path -Path $inputFile) ) {
  Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

<#
The we build the deployment file
#>

$image = "ghcr.io/nexi-intra/$($imagename)-web:$($version)"

$config = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname-newhire
spec:
  selector:
    matchLabels:
      app: $appname-newhire
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname-newhire
    spec: 
      containers:
      - name: $appname-newhire
        image: $image
        ports:
          - containerPort: $port
        env:
        - name: SPAUTH_TENANTID
          value: $($env:SPAUTH_TENANTID)
        - name: SPAUTH_CLIENTID
          value: $($env:SPAUTH_CLIENTID)
        - name: NEWSCHANNELSBLOB
          value: $($env:NEWSCHANNELSBLOB)
        - name: VALIDDOMAINSBLOB
          value: $($env:VALIDDOMAINSBLOB)
        - name: AZURE_AD_CLIENT_ID
          value: $($env:AZURE_AD_CLIENT_ID)
        - name: AZURE_AD_CLIENT_SECRET
          value: $($env:AZURE_AD_CLIENT_SECRET)
        - name: AZURE_AD_TENANT_ID
          value: $($env:AZURE_AD_TENANT_ID)
        - name: SPAUTH_CLIENTSECRET
          value: $($env:SPAUTH_CLIENTSECRET)
        - name: PROFILETYPE
          value: newhire
        
---
apiVersion: v1
kind: Service
metadata:
  name: $appname-newhire
  labels:
    app: $appname-newhire
    service: $appname-newhire
spec:
  ports:
  - name: http
    port: 5301
    targetPort: $port
  selector:
    app: $appname-newhire
---    
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: $appname-newhire
spec:
  rules:
  - host: $dnsname
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: $appname-newhire
            port:
              number: 5301
    

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -