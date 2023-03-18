# kubedge-operator-ecds

Operator to deploy KUBEDGE ECDS Simulator

## Notes

tbd

## Getting Started

### Installing the Operator

```bash
$ make install-amd64
```

```
helm install kubedge-ecds-operator chart --set images.tags.operator=kubedge1/kubedge-ecds-operator-amd64:v0.1.0,images.pull_policy=Always --namespace default
NAME: kubedge-ecds-operator
LAST DEPLOYED: Sat Mar 18 20:01:46 2023
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

```bash
kubedge@kubedgesdk:~/proj/branchalignments/kubedge/kubedge-operator-ecds$ kubectl get all -A
```

```
NAMESPACE          NAME                                           READY   STATUS                   RESTARTS       AGE
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-4ksj7          0/1     ContainerStatusUnknown   1 (19d ago)    20d
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-9px6v          1/1     Running                  15 (80m ago)   20d
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-h7nzq          1/1     Running                  12 (80m ago)   19d
calico-system      pod/calico-kube-controllers-6b7b9c649d-92zrb   1/1     Running                  14 (80m ago)   20d
calico-system      pod/calico-node-s5648                          1/1     Running                  14 (80m ago)   20d
calico-system      pod/calico-typha-7c747fc89b-4xnrf              1/1     Running                  26 (79m ago)   20d
calico-system      pod/csi-node-driver-zgdwc                      2/2     Running                  28 (80m ago)   20d
default            pod/kubedge-ecds-operator-76d955dcb6-942b9     1/1     Running                  0              5m40s
kube-system        pod/coredns-787d4945fb-cm7x9                   1/1     Running                  14 (80m ago)   20d
kube-system        pod/coredns-787d4945fb-mhs4b                   1/1     Running                  14 (80m ago)   20d
kube-system        pod/etcd-node                                  1/1     Running                  14 (80m ago)   20d
kube-system        pod/kube-apiserver-node                        1/1     Running                  14 (80m ago)   20d
kube-system        pod/kube-controller-manager-node               1/1     Running                  19 (80m ago)   20d
kube-system        pod/kube-proxy-hl6z5                           1/1     Running                  14 (80m ago)   20d
kube-system        pod/kube-scheduler-node                        1/1     Running                  16 (80m ago)   20d
tigera-operator    pod/tigera-operator-54b47459dd-m4gl4           1/1     Running                  30 (80m ago)   20d

NAMESPACE          NAME                                      TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
calico-apiserver   service/calico-api                        ClusterIP   10.108.208.171   <none>        443/TCP                  20d
calico-system      service/calico-kube-controllers-metrics   ClusterIP   None             <none>        9094/TCP                 20d
calico-system      service/calico-typha                      ClusterIP   10.108.82.253    <none>        5473/TCP                 20d
default            service/kubernetes                        ClusterIP   10.96.0.1        <none>        443/TCP                  20d
kube-system        service/kube-dns                          ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   20d

NAMESPACE       NAME                             DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
calico-system   daemonset.apps/calico-node       1         1         1       1            1           kubernetes.io/os=linux   20d
calico-system   daemonset.apps/csi-node-driver   1         1         1       1            1           kubernetes.io/os=linux   20d
kube-system     daemonset.apps/kube-proxy        1         1         1       1            1           kubernetes.io/os=linux   20d

NAMESPACE          NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
calico-apiserver   deployment.apps/calico-apiserver          2/2     2            2           20d
calico-system      deployment.apps/calico-kube-controllers   1/1     1            1           20d
calico-system      deployment.apps/calico-typha              1/1     1            1           20d
default            deployment.apps/kubedge-ecds-operator     1/1     1            1           5m40s
kube-system        deployment.apps/coredns                   2/2     2            2           20d
tigera-operator    deployment.apps/tigera-operator           1/1     1            1           20d

NAMESPACE          NAME                                                 DESIRED   CURRENT   READY   AGE
calico-apiserver   replicaset.apps/calico-apiserver-6d6df8c6b5          2         2         2       20d
calico-system      replicaset.apps/calico-kube-controllers-6b7b9c649d   1         1         1       20d
calico-system      replicaset.apps/calico-typha-7c747fc89b              1         1         1       20d
default            replicaset.apps/kubedge-ecds-operator-76d955dcb6     1         1         1       5m40s
kube-system        replicaset.apps/coredns-787d4945fb                   2         2         2       20d
tigera-operator    replicaset.apps/tigera-operator-54b47459dd           1         1         1       20d
```

### Deploy the cluster
Deploy the CR

```bash
$ kubectl apply -f examples/ecds-amd64.yaml
```

```code
ecdscluster.kubedgeoperators.kubedge.cloud/kubedge-ecds-cluster created
```

```bash
$ kubectl get all -A
```

```code
NAMESPACE          NAME                                           READY   STATUS                   RESTARTS       AGE
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-4ksj7          0/1     ContainerStatusUnknown   1 (19d ago)    20d
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-9px6v          1/1     Running                  15 (84m ago)   20d
calico-apiserver   pod/calico-apiserver-6d6df8c6b5-h7nzq          1/1     Running                  12 (84m ago)   19d
calico-system      pod/calico-kube-controllers-6b7b9c649d-92zrb   1/1     Running                  14 (84m ago)   20d
calico-system      pod/calico-node-s5648                          1/1     Running                  14 (84m ago)   20d
calico-system      pod/calico-typha-7c747fc89b-4xnrf              1/1     Running                  26 (83m ago)   20d
calico-system      pod/csi-node-driver-zgdwc                      2/2     Running                  28 (84m ago)   20d
default            pod/businesslogic-0                            1/1     Running                  0              62s
default            pod/enrichment-0                               1/1     Running                  0              62s
default            pod/frontend-0                                 1/1     Running                  0              62s
default            pod/kubedge-ecds-operator-76d955dcb6-942b9     1/1     Running                  0              9m18s
default            pod/loadbalancer-0                             1/1     Running                  0              61s
default            pod/platform-0                                 1/1     Running                  0              61s
kube-system        pod/coredns-787d4945fb-cm7x9                   1/1     Running                  14 (84m ago)   20d
kube-system        pod/coredns-787d4945fb-mhs4b                   1/1     Running                  14 (84m ago)   20d
kube-system        pod/etcd-node                                  1/1     Running                  14 (84m ago)   20d
kube-system        pod/kube-apiserver-node                        1/1     Running                  14 (84m ago)   20d
kube-system        pod/kube-controller-manager-node               1/1     Running                  19 (84m ago)   20d
kube-system        pod/kube-proxy-hl6z5                           1/1     Running                  14 (84m ago)   20d
kube-system        pod/kube-scheduler-node                        1/1     Running                  16 (84m ago)   20d
tigera-operator    pod/tigera-operator-54b47459dd-m4gl4           1/1     Running                  30 (84m ago)   20d

NAMESPACE          NAME                                      TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
calico-apiserver   service/calico-api                        ClusterIP   10.108.208.171   <none>        443/TCP                  20d
calico-system      service/calico-kube-controllers-metrics   ClusterIP   None             <none>        9094/TCP                 20d
calico-system      service/calico-typha                      ClusterIP   10.108.82.253    <none>        5473/TCP                 20d
default            service/businesslogic                     ClusterIP   None             <none>        <none>                   62s
default            service/enrichment                        ClusterIP   None             <none>        <none>                   62s
default            service/frontend                          ClusterIP   None             <none>        <none>                   62s
default            service/kubernetes                        ClusterIP   10.96.0.1        <none>        443/TCP                  20d
default            service/loadbalancer                      ClusterIP   None             <none>        <none>                   62s
default            service/platform                          ClusterIP   None             <none>        <none>                   61s
kube-system        service/kube-dns                          ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   20d

NAMESPACE       NAME                             DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
calico-system   daemonset.apps/calico-node       1         1         1       1            1           kubernetes.io/os=linux   20d
calico-system   daemonset.apps/csi-node-driver   1         1         1       1            1           kubernetes.io/os=linux   20d
kube-system     daemonset.apps/kube-proxy        1         1         1       1            1           kubernetes.io/os=linux   20d

NAMESPACE          NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
calico-apiserver   deployment.apps/calico-apiserver          2/2     2            2           20d
calico-system      deployment.apps/calico-kube-controllers   1/1     1            1           20d
calico-system      deployment.apps/calico-typha              1/1     1            1           20d
default            deployment.apps/kubedge-ecds-operator     1/1     1            1           9m18s
kube-system        deployment.apps/coredns                   2/2     2            2           20d
tigera-operator    deployment.apps/tigera-operator           1/1     1            1           20d

NAMESPACE          NAME                                                 DESIRED   CURRENT   READY   AGE
calico-apiserver   replicaset.apps/calico-apiserver-6d6df8c6b5          2         2         2       20d
calico-system      replicaset.apps/calico-kube-controllers-6b7b9c649d   1         1         1       20d
calico-system      replicaset.apps/calico-typha-7c747fc89b              1         1         1       20d
default            replicaset.apps/kubedge-ecds-operator-76d955dcb6     1         1         1       9m18s
kube-system        replicaset.apps/coredns-787d4945fb                   2         2         2       20d
tigera-operator    replicaset.apps/tigera-operator-54b47459dd           1         1         1       20d

NAMESPACE   NAME                             READY   AGE
default     statefulset.apps/businesslogic   1/1     62s
default     statefulset.apps/enrichment      1/1     62s
default     statefulset.apps/frontend        1/1     62s
default     statefulset.apps/loadbalancer    1/1     62s
default     statefulset.apps/platform        1/1     61s
```


