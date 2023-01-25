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
kubectl config set-context fullcycle
```

### Get Nodes Running
```sh
kubectl get nodes
```

### Apply changes to your cluster using a spec file
```sh
kubectl apply -f k8s/FILE_NAME.yaml
```

### Get Pods running
```sh
kubectl get pods
```

### Get Deployment status
```sh
kubectl get deployment.apps goserver
```

### Describe Deployment
```sh
kubectl describe deployment.apps goserver
```

### Get Statefulset Status
```sh
kubectl get statefulset.apps goserver
```

### Describe StatefulSet
```sh
kubectl describe statefulset.apps goserver
```

### Describe Deployment
```sh
kubectl describe deployment.apps goserver
```

### Get Resource Revision
```sh
kkubectl describe deployment.apps goserver
```

### Get Rollout History
```sh
kubectl rollout history deployments/goserver
```

### Rollout Undo Change
```sh
kubectl rollout undo deployments/goserver --to_revision=REVISON_NUMBER
```

