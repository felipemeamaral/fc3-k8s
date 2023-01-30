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

### Get ReplicaSets
```sh
kubectl get replicasets.apps
```

### Get Deployments
```sh
kubectl get deployments.apps
```

### Get Deployment status
```sh
kubectl get deployments.apps goserver
```

### Describe Deployment
```sh
kubectl describe deployments.apps goserver
```

### Get Statefulset Status
```sh
kubectl get statefulsets.apps goserver
```

### Describe StatefulSet
```sh
kubectl describe statefulsets.apps goserver
```

### Describe Deployment
```sh
kubectl describe deployments.apps goserver
```

### Get Resource Revision
```sh
kkubectl describe deployments.apps goserver
```

### Get Rollout History
```sh
kubectl rollout history deployments/goserver
```

### Rollout Undo Change
```sh
kubectl rollout undo deployments/goserver --to_revision=REVISON_NUMBER
```

### Port forward to one of your Pods
```sh
kubectl port-forward pods/goserver 8080:8080
```

### Run stress test using fortio on your application
```sh
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service:8000/healthz"
```