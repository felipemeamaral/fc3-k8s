# Setup a Kubernetes node using Kind

### Install kind on macOS
```sh
brew install kind
```

### Create the cluster using the kubernetes specification file we created
```sh
kind create cluster --config=k8s/kind.yaml --name=fullcycle
```

### Change between your active clusters
List your active clusters:
```sh
kubectl config get-clusters
```

Change to your desired cluster:
```sh
kubectl config set-context CLUSTER_NAME
```

