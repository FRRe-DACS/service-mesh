# Golang - gRPC - Service Mesh

## Deploy with no Serive Mesh

```console
kubectl apply -f deploy/no-service-mesh.yaml
```

## Send request to the app

```console
grpcurl -d '{"foo":{"id": 2,"bar":{"id": 2,"name":"Hello","description": "Hello"}}}' \
  -plaintext localhost:9091 ar.edu.utn.frre.dacs.foo.FooService/CreateFoo
```

## Clean up no Serive Mesh

```console
kubectl delete -f deploy/no-service-mesh.yaml
```

## Deploy with Service Mesh

```console
istioctl install --set profile=demo
kubectl apply -f ${ISTIO_HOME}/samples/addons -n istio-system
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/extras/zipkin.yaml
kubectl label namespace default istio-injection=enabled
kubectl apply -f deploy/service-mesh.yaml
```

## Send request to the Service Mesh

```console
grpcurl -d '{"foo":{"id": 2,"bar":{"id": 2,"name":"Hello","description": "Hello"}}}' \
  -plaintext localhost:80 ar.edu.utn.frre.dacs.foo.FooService/CreateFoo
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
istioctl dashboard zipkin
```

## Clean up - Service Mesh

```console
kubectl delete -f deploy/service-mesh.yaml
kubectl label namespace default istio-injection-
kubectl delete -f ${ISTIO_HOME}/samples/addons -n istio-system --ignore-not-found
kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.7/samples/addons/extras/zipkin.yaml
istioctl --set profile=demo manifest generate | kubectl delete --ignore-not-found -f -
```
