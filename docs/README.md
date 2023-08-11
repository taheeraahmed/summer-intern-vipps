# Welcome to the summerstudents backend

This is the backend service for the 2023 summerstudent project where we
bring recurring donations to charities in vippsnummer.



## Vippsservice 

### Getting started

This repository is created from vippstech.no and contain what you need to start creating an app on the Vipps Tech Platform. Your team should have admin access for this repository, but if not reach out to us at `#help-github` in Slack.

### Workflows
Under `manifests/k8s` folder, you will find a starting point for deploying a VippsService (continuously running application) and/or a Vibe (CronJob). The `overlays` are there to help you deploy different versions and environment variables of your applications to the different environments. 

The repository is configured with GitHub variables to allow your workflow(s) to push images to the container registry and deploy them to our Kubernetes cluster. As a starting point two workflows are created for you, one for deploying a VippsService and one for infrastructure components using Bicep.

See the [general documentation for the Vipps Tech Platform](https://vippstech.no/docs/default/component/vipps-tech-platform-doc) for more information regarding these services. If you have any questions, reach out to us in `#platform-support`.


Your entity is already registered in Vipps' service catalog. To make changes, see `catalog-info.yaml`.
Here you can edit and enable plugins like *Azure*, *Azure DevOps*, *Kubernetes* via the `annotations` section.
By default, you application is added with `experimental` as its lifecycle. Once your application is production ready, change to `production`

[For general information see our documentation](https://vippstech.no/docs/default/component/vipps-backstage)
