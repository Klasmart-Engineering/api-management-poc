# api-management-poc

## Declarative Kong YAML

- As we've not deployed a database in the istio based cluster _(ie. Kong is deployed in DB-less mode)_ in order to
  update the kong config using the declarative syntax, the `kong.yaml` file has to be part of the helm `values.yaml`
  file. To apply it you then need to run
```sh
helm upgrade -n kong-istio kong-istio kong/kong -f kong/values.yaml
```

### K3D Local Cluster Setup
Install K3D to run a local cluster.    
Install Istio onto your machine.
```
k3d cluster create local-cluster    

istioctl install --set profile=demo -y    
kubectl label namespace default istio-injection=enabled     
k create ns edge    
k create ns opa-authorisation     
kubectl label namespace edge istio-injection=enabled     
kubectl label namespace opa-authorisation istio-injection=enabled     
```

Update the istio-ingressgateway.yaml to use your new namespaces. A find-replace of 'istio-kong' for 'edge' should do it if the file stays as it is now :)
I'll look like this:
```
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: kong-gateway
  namespace: edge
spec:
  selector:
    istio: ingressgateway # use Istio default gateway implementation
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: kong-gateway
  namespace: edge
spec:
  hosts:
  - "*"
  gateways:
  - kong-gateway
  exportTo:
    - "istio-system"
  http:
  - route:
    - destination:
        port:
          number: 80
        host: edge-kong-proxy
---
apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: "turn-off-mtls"
  namespace: "istio-system"
spec:
  mtls:
    mode: DISABLE
```

```
Nb. a good way of jumping between namespaces is to install kubens
```
kubens     
cd kong    
helm repo add kong https://charts.konghq.com      
helm install -n edge edge kong/kong -f values.yaml     
k get svc      
k apply -f istio-ingressgateway.yaml     
helm upgrade -n edge edge kong/kong -f values.yaml      
k get svc      
helm install -n opa-authorisation opa kong/kong -f opa-values.yaml     
helm upgrade -n edge edge kong/kong -f values.yaml     
helm upgrade -n opa-authorisation opa kong/kong -f opa-values.yaml     
```

Install Lens for a shortcut for port forwarding out of your cluster.

Click on the local-cluster -> Workloads -> Pods.
You should find a Pod called opa-kong-9999999-dkfgkd (id will be different).
Click on that Pod and then scroll down the side panel that emerges until you find the 'Ports' section. Click on the 'Forward' button for the proxy entry.
```
curl http://localhost:64891/starwars
```


