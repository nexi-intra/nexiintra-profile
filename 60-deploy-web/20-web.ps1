<#---
title: Create Web deployment file
input: profiledataurl.json
output: web.yaml
connection: sharepoint
tag: createdeploymentfile
api: post
---#>
param (
    $name = "nexiintra-profile",
    $image = "ghcr.io/koksmat-com/ui:v1.0.0.profile-3"
)

$SPAUTH_TENANTID=$env:SPAUTH_TENANTID
$SPAUTH_CLIENTID=$env:SPAUTH_CLIENTID
$SPAUTH_CLIENTSECRET=$env:SPAUTH_CLIENTSECRET

$deployment = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $name
spec:
  selector:
    matchLabels:
      app: $name
  replicas: 1
  template:
    metadata:
      labels:
        app: $name
    spec: 
      containers:
      - name: $name
        image: $image
        ports:
          - containerPort: 3001
        env:
        - name: SPAUTH_TENANTID
          value: $SPAUTH_TENANTID
        - name: SPAUTH_CLIENTID
          value: $SPAUTH_CLIENTID
        - name: SPAUTH_CLIENTSECRET
          value: $SPAUTH_CLIENTSECRET

      
---
apiVersion: v1
kind: Service
metadata:
  name: $name
  labels:
    app: $name
    service: $name
spec:
  ports:
  - name: http
    port: 5301
    targetPort: 3001
  selector:
    app: $name

"@


$deployment | Out-File -FilePath "$env:WORKDIR/web.yaml" -Encoding utf8