# api-management-poc

## Declarative Kong YAML

- As we've not deployed a database in the istio based cluster _(ie. Kong is deployed in DB-less mode)_ in order to
  update the kong config using the declarative syntax, the `kong.yaml` file has to be part of the helm `values.yaml`
  file. To apply it you then need to run
```sh
helm upgrade -n kong-istio kong-istio kong/kong -f kong/values.yaml
```
