<#---
title: Create Web deployment file
input: profiledataurl.json
output: web.yaml
tag: createdeploymentfile
api: post
---
#>
function EnvValue($name) {

$value = "****MISSING****"

foreach ($item in (Get-Item -Path Env:)) {
  if ($item.Key -eq $name) {
    $value = $item.Value
  }
}
  $r = @"
  - name: $name
         value: $value
"@ 
 
 return $r
}

<#
We start by finding which version tag to use
#>

$inputFile = join-path  $env:KITCHENROOT "nexiintra-profile" ".koksmat","koksmat.json"
$port = "4323"
if (!(Test-Path -Path $inputFile) ) {
   Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

<#
The we build the deployment file
#>
$appname = "nexiintra-profile"

$image = "ghcr.io/koksmat-com/$($appname)-web:$($version)"

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
          - containerPort: $port
        env:
        - name: SPAUTH_TENANTID
          value: $($env:SPAUTH_DOMAIN)
        $(EnvValue "SPAUTH_CLIENTID")
        $(EnvValue "NEWSCHANNELSBLOB")
        $(EnvValue "VALIDDOMAINSBLOB")
        $(EnvValue "AZURE_AD_CLIENT_ID")
        $(EnvValue "AZURE_AD_CLIENT_SECRET")
        $(EnvValue "AZURE_AD_TENANT_ID")
        $(EnvValue "SPAUTH_CLIENTSECRET")
        
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
    targetPort: $port
  selector:
    app: $appname
---    
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: $appname
spec:
  rules:
  - host: me.nexi-intra.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: $appname
            port:
              number: 5301
    
"@


$deployment | Out-File -FilePath "$env:WORKDIR/web.yaml" -Encoding utf8