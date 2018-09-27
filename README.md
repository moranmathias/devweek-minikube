# Kubernetes demo

## Requirements

- Docker
- [Minikube](https://kubernetes.io/docs/setup/minikube/#installation)
- [Helm](https://docs.helm.sh/using_helm/#from-the-binary-releases)

## Setup

- Run `eval $(minikube docker-env)` to build the docker images in minikube vm
- Build images:
```
docker build --force-rm -t devweek/echosrv -f ./build/echosrv.dockerfile .
docker build --force-rm -t devweek/uppercasesrv -f ./build/uppercasesrv.dockerfile .
```

### Deploy echosrv
```
kubectl create -f deployment/deployments/echosrv.yml
kubectl create -f deployment/services/echosrv.yml
```


### Deploy uppercasesrv

```
kubectl create -f deployment/deployments/uppercasesrv.yml
kubectl create -f deployment/services/uppercasesrv.yml
```

### Helm usage
```
helm install --name pg-storage --set service.type=NodePort stable/postgresql
```