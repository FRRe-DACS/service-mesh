# Golang - gRPC - Service Mesh

## Deploy with no Serive Mesh

```console
kubectl apply -f deploy/no-service-mesh.yaml
```

## Deploy with Service Mesh

```console
istioctl install --set profile=demo
kubectl label namespace default istio-injection=enabled
kubectl apply -f deploy/service-mesh.yaml
```
