# Infrastructure

The infrastructure which is being created is based on the `infrastructure/main.bicep` code. 

## CI/CD Github Workflows

| Workflow Name         | Trigger        | Description                                        |
|-----------------------|----------------|----------------------------------------------------|
| Build and publish     | Push to main   | This builds the application.                       |
| infrastructure deploy | Push to main   | Runs the bicep script and deploys azure resources. |
| Migrate database      | Manual trigger | Runs migration scripts.                            |
| PR test               | Pull request   | Runs `make test`.                                  |

## Azure

| Property             | Value                             |
|----------------------|-----------------------------------|
| Application          | gh-summerstudents-backend-uat-app |
| Kubernetes Namespace | rizz-and-compliance               |
| Subscription         | Summer students test              |
| Resource group       | backend                           |
| Database server      | sql-rafqkaoj4kdlc                 |
| APIM (stv2)          |    summerstudents-backend API     |
