# go-example-app

## build code

```bash
argo submit -n argo .argoworkflow.yaml
```

## build deploy yaml

```bash
kustomize build deploy
```

## deploy

```bash
argocd app create go-example-app --repo https://github.com/kuops/go-example-app.git --path deploy --dest-server https://kubernetes.default.svc --dest-namespace default
```
