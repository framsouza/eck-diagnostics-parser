# eck-diagnostics-parser

This tool is read, extracts and generate a friendly output of the ECK diagnostics tool.

You can install it in two ways:


### Installling with brew

```
$ brew tap framsouza/framsouza 
$ brew install framsouza/framsouza/eck-diagnostics-parser 
$ eck-diagnostics-parser -zipfile <pathtotheeckdiagnostic.zip>
```


### Cloning the repo

There are two options you could use if you choose to clone the repo:

#### Building the binary
```
$ git clone git@github.com:framsouza/eck-diagnostics-parser.git && cd eck-diagnostics-parser
$ go build -o eck-diagnostics-parser ./cmd
```

It will generate a `./eck-diagnostics-parser` command, you should it in the `PATH` or even create an alias 

#### Without building
```
$ git clone git@github.com:framsouza/eck-diagnostics-parser.git && cd eck-diagnostics-parser
$ go run cmd/main.go -zipfile  <pathtotheeckdiagnostic.zip>
```

## NOTES
It will create a temporary folder called `./tmp` in the current directory, make sure to not run it on `/` otherwise you will get the following error:
```
eck-diagnostics-parser -zipfile /Users/francismarasouza/Downloads/eck-diagnostic-2022-05-31T12-45-19.zip
2022/07/27 17:40:28 mkdir tmp: read-only file system
```

This directory will be automatically deleted once the program finishes.

### Expected output

```
Welcome to the ECK diagnostic parser tool

Diagnostic version is 1.0.2
2022/03/10 07:44:28 ECK version is 1.9.0
Diagnostic Collection date is 2022-03-10 07:44:25.579895118 +0000 UTC

NODE NAME                                             CPU CAPACITY          CPU ALLOCATED          MEM CAPACITY          MEM ALLOCATED          VERSION                      NODE READY
ip-10-52-1-46.eu-central-1.compute.internal           4                     3920m                  32499784Ki            31482952Ki             v1.21.5-eks-bc4871b          True
ip-10-52-2-170.eu-central-1.compute.internal          4                     3920m                  32499784Ki            31482952Ki             v1.21.5-eks-bc4871b          True
ip-10-52-3-15.eu-central-1.compute.internal           4                     3920m                  32499784Ki            31482952Ki             v1.21.5-eks-bc4871b          True

PODS NAME                           NAMESPACE           STATUS              MEM REQUEST          MEM LIMIT           CPU REQUEST          CPU LIMIT           NODE NAME                                             
test-es-data-hot-0               default             Running             16Gi                 16Gi                1                    1                   ip-10-52-1-46.eu-central-1.compute.internal 
test-es-data-hot-1               default             Running             16Gi                 16Gi                1                    1                   ip-10-52-2-170.eu-central-1.compute.internal
test-es-data-hot-2               default             Running             16Gi                 16Gi                1                    1                   ip-10-52-3-15.eu-central-1.compute.internal 
test-es-data-warm-0              default             Running             6Gi                  6Gi                 1                    1                   ip-10-52-3-15.eu-central-1.compute.internal 
test-es-data-warm-1              default             Running             6Gi                  6Gi                 1                    1                   ip-10-52-1-46.eu-central-1.compute.internal 
test-es-data-warm-2              default             Running             6Gi                  6Gi                 1                    1                   ip-10-52-2-170.eu-central-1.compute.internal
test-es-master-node-0            default             Running             2Gi                  2Gi                 Not set              Not set             ip-10-52-2-170.eu-central-1.compute.internal
test-es-master-node-1            default             Running             2Gi                  2Gi                 Not set              Not set             ip-10-52-3-15.eu-central-1.compute.internal 
test-es-master-node-2            default             Running             2Gi                  2Gi                 Not set              Not set             ip-10-52-1-46.eu-central-1.compute.internal 
test-kb-556cdcbdf-hrxh4          default             Running             1Gi                  1Gi                 Not set              Not set             ip-10-52-1-46.eu-central-1.compute.internal 

PODS NAME                   NAMESPACE               STATUS              MEM REQUEST          MEM LIMIT           CPU REQUEST          CPU LIMIT           NODE NAME                                             
elastic-operator-0          elastic-system          Running             150Mi                512Mi               100m                 1                   ip-10-52-2-170.eu-central-1.compute.internal

PODS NAME                         NAMESPACE            STATUS              MEM REQUEST          MEM LIMIT           CPU REQUEST          CPU LIMIT           NODE NAME                                             
aws-node-xmsvw                    kube-system          Running             Not set              Not set             10m                  Not set             ip-10-52-3-15.eu-central-1.compute.internal 
aws-node-xpfvp                    kube-system          Running             Not set              Not set             10m                  Not set             ip-10-52-1-46.eu-central-1.compute.internal 
aws-node-zjsdm                    kube-system          Running             Not set              Not set             10m                  Not set             ip-10-52-2-170.eu-central-1.compute.internal
coredns-745979c988-jrf48          kube-system          Running             70Mi                 170Mi               100m                 Not set             ip-10-52-1-46.eu-central-1.compute.internal 
coredns-745979c988-l9qk4          kube-system          Running             70Mi                 170Mi               100m                 Not set             ip-10-52-3-15.eu-central-1.compute.internal 
kube-proxy-ljb5l                  kube-system          Running             Not set              Not set             100m                 Not set             ip-10-52-1-46.eu-central-1.compute.internal 
kube-proxy-lqlhs                  kube-system          Running             Not set              Not set             100m                 Not set             ip-10-52-2-170.eu-central-1.compute.internal
kube-proxy-n5xr5                  kube-system          Running             Not set              Not set             100m                 Not set             ip-10-52-3-15.eu-central-1.compute.internal

ES NAME                       STATUS              VERSION             PHASE               NODES               NAMESPACE           
test                       green               7.16.3              Ready               9                   default

KB NAME                       STATUS              VERSION             PHASE                NODES               NAMESPACE           
test                       green               7.16.3              Established          1                   default

STATEFULSET NAME                          REPLICAS            NAMESPACE           HEAP SIZE              
test-es-data-hot                       3                   default             -Xms8g -Xmx8g          
test-es-data-warm                      3                   default             -Xms4g -Xmx4g          
test-es-master-node                    3                   default             

STATEFULSET NAME                    REPLICAS            NAMESPACE               HEAP SIZE           
elastic-operator                    1                   elastic-system

DEPLOYMENT NAME                    REPLICAS            NAMESPACE           
test-kb                         1                   default

PVC NAME                                             STATUS              CAPACITY            VOLUME NAME                                       SC NAME             ACCESS MODE              
elasticsearch-data-test-es-data-hot-0             Bound               613Gi               pvc-455f7a2d-fcc4-44ec-8750-11d09cf95e7b          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-data-hot-1             Bound               613Gi               pvc-2b3e1470-9ea0-4110-b2b9-bab8d31ed61f          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-data-hot-2             Bound               613Gi               pvc-355786de-07a4-4d28-9988-54698460933c          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-data-warm-0            Bound               1613Gi              pvc-496dc5dc-3584-456b-885c-960b45236193          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-data-warm-1            Bound               1613Gi              pvc-6147efb9-e5b6-4f47-9f7c-cfec69679c4c          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-data-warm-2            Bound               1613Gi              pvc-3a5fc7c6-d522-4aae-b2f2-217884294507          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-master-node-0          Bound               13Gi                pvc-237493f9-252e-4ad1-b3e2-4d78ac930338          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-master-node-1          Bound               13Gi                pvc-b2fc54a3-4dcf-4b18-a9b0-7e825533e8d3          gp2                 [ReadWriteOnce]
elasticsearch-data-test-es-master-node-2          Bound               13Gi                pvc-74179549-8367-412b-bb32-3b0f632d60ab          gp2                 [ReadWriteOnce]

SERVICE NAME                    TYPE                  NAMESPACE           ENDPOINTS           
test-es-data-hot             ClusterIP             default             10.52.1.19, 10.52.2.147, 10.52.3.245, 
test-es-data-warm            ClusterIP             default             10.52.1.81, 10.52.2.231, 10.52.3.240, 
test-es-http                 LoadBalancer          default             10.52.1.155, 10.52.1.19, 10.52.1.81, 10.52.2.147, 10.52.2.218, 10.52.2.231, 10.52.3.240, 10.52.3.245, 10.52.3.95, 
test-es-master-node          ClusterIP             default             10.52.1.155, 10.52.2.218, 10.52.3.95, 
test-es-transport            ClusterIP             default             10.52.1.155, 10.52.1.19, 10.52.1.81, 10.52.2.147, 10.52.2.218, 10.52.2.231, 10.52.3.240, 10.52.3.245, 10.52.3.95, 
test-kb-http                 LoadBalancer          default             10.52.1.120, 
kubernetes                      ClusterIP             default             10.52.1.83, 10.52.2.72, 

Storage Class file not found, make sure to use the latests version of ECK Diagnostics

BASED ON THE OUTPUT ABOVE, MAKE SURE THAT:
- The Elasticsearch services has an endpoint attached to it, if there's no endpoint the services won't be able to connect to the pods
- All the PVC has a Bound status
- All the Elasticsearch Resources are Green and the Phase is "READY", if the status is "ApplyingChanges" check the Operator logs
- All the Kibana Resources are Green and the Phase is "READY" or "ESTABLISHED", if the status is "ApplyingChanges" check the Operator logs
- All the Elasticsearch and Kibana pods has the RUNNING status, if not, check the pod logs
- All the Elasticsearch has the same MEM REQUEST & MEM LIMIT, it's very important to ensure quality of Serivce
- Starting with Elasticsearch 7.11, unless manually overridden, heap size is automatically calculated based on the available memory, if the HEAP SIZE column is empty for the Elasticsearch statefulsets, make sure Elasticsearch is >7.11
- If you desire to increase the disk size, make sure you are ran the latest ECK diagnostics and check the if ALLOW EXPANSION is true, if that is the case you can easily change the PVC volume size in the Elasticsearch manifest 
```