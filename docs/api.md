# Api 

## General
The api is accessible from this url:

`https://ece46ec4-6f9c-489b-8fe5-146a89e11635.tech-02.net/summerstudents-backend/`

The API offers endpoints to mock some of the behaviour of the recurring-api to
support our prototype.


## Specs

The swagger definition is located at `manifests/k8s/base/openApiSwagger/summerstudents-backend-apim.yaml`

When updated, run `make swagger` to generate code.
### V0 Endpoints

The V0 enpoints are used for testing by the frontend team, and only sends dummy data.
It has no connection to the database.


### V1 Endpoints

These are the endpoints that are used by the IOS app and Merchant portal for our 
prototype.

[API Definition](https://vippstech.no/catalog/default/api/summerstudents-backend-api/definition)

