# Create an endpoint

## Swagger

Go-swagger will generate a lot of boilerplate for us, but first 
we need to define an endpoint in `manifests/k8s/base/openApiSwagger/kustomization.yaml`. 

Look at the format of other enpoints or check out [how to define an operation](https://swagger.io/docs/specification/2-0/paths-and-operations/).

When the swagger definition is done run the following:

```bash
# generate boilerplate
make swagger

# clean up go dependencies 
go mod tidy
```

## Create a handler

