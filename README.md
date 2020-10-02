# Golang - gRPC - Service Mesh

## Deploy with no Serive Mesh

```console
kubectl apply -f deploy/no-service-mesh.yaml
```

## Clean up no Serive Mesh

```console
kubectl delete -f deploy/no-service-mesh.yaml
```

## Deploy with Service Mesh

```console
istioctl install --set profile=demo
kubectl apply -f ${ISTIO_HOME}/samples/addons -n istio-system
kubectl label namespace default istio-injection=enabled
kubectl apply -f deploy/service-mesh.yaml
```

## Clean up - Service Mesh

```console
kubectl delete -f deploy/service-mesh.yaml
kubectl label namespace default istio-injection-
kubectl delete -f ${ISTIO_HOME}/samples/addons -n istio-system --ignore-not-found
istioctl --set profile=demo manifest generate | kubectl delete -f --ignore-not-found -
```

## Istioclt Dashboard

### Kiali

```console
istioctl dashboard kiali
```

### Jaeger

```console
istioctl dashboard jaeger
```

### Grafana

```console
istioctl dashboard grafana
```

### Zipkin

```console
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/extras/zipkin.yaml
istioctl dashboard zipkin
```
