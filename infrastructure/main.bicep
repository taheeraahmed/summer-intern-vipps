// this is an example bicep file. You will need to create your own bicep file to deploy your own resources.
// For more information on how to create a bicep file, see https://vippstech.no/docs/default/Component/azure-templates
// if you do not require any Azure resources, please delete.

param dbServerName string = 'rizz-and-comp-db-server'
param date string = utcNow('yyyy-MM-dd')
param location string = resourceGroup().location

var tenantId = '805bc25d-8e64-4ed6-8d24-3883c9068c5a'

var kvReadOnly = [
  'Get'
  'List'
]
var kvReadWrite = [
  'Get'
  'List'
  'Set'
  'Delete'
  'Recover'
  'Backup'
  'Restore'
]

var tags = {
  owner: 'team-summerstudents'
  'last-review': date
  environment: 'test'
  'personal-data': 'no'
  repo: 'https://github.com/vippsas'
}

module storage 'br:acrvippsakscmn.azurecr.io/vipps/vce/modules/vipps-mssql:2023-06-27' = {
  name: dbServerName // storage account names must be lowercase
  params: {
    databaseName: 'rizz-and-comp-db'
    dbAdLoginName: 'team-summerstudents'
    dbAdId: 'd86fa257-25c0-4f1f-a413-e9a3c5725e88'
    env: 'uat'
    tags: tags
    connectToVCE: false
  }
}

module keyvault 'br:acrvippsakscmn.azurecr.io/vipps/vce/modules/vipps-keyvault:2023-02-07' = {
  name: 'rizz-keyvault-deployment'
  params: {
    keyVaultName: 'rizz-kv'
    tags: tags
    allowVippsOfficeIp: true
    accessPolicies: [
      {
        tenantId: tenantId
        objectId: 'd86fa257-25c0-4f1f-a413-e9a3c5725e88' // team-summerstudents
        permissions: {
          secrets: kvReadWrite
        }
      }
      {
        tenantId: tenantId
        objectId: 'c208fad5-ec5e-4eb9-a7b0-fcb9de0bead5' // summerstudents-backend-app-test
        permissions: {
          secrets: kvReadOnly
        }
      }
    ]
    env: 'uat'
  }
}

module storageAccount 'br:acrvippsakscmn.azurecr.io/vipps/vce/modules/vipps-storage:2023-06-27' = {
  name: 'rizz-sa-deployment'
  params: {
    storageAccountName: 'rizzstorageaccount'
    location: location
    tags: tags
    env: 'uat'
    hasSharedAccessKeyEnabled: true
  }
}

