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
Wck-diagnostics-parser -zipfile /Users/francismarasouza/Downloads/eck-diagnostic-2022-05-31T12-45-19.zip

Welcome to the ECK diagnostic parser tool

Diagnostic version is 1.1.0
2022/05/31 12:45:26 ECK version is 2.2.0
Diagnostic Collection date is 2022-05-31 12:45:19.841005894 +0200 CEST

NODE NAME                                              CPU CAPACITY          CPU ALLOCATED          MEM CAPACITY          MEM ALLOCATED          VERSION             NODE READY
shoot--phbnrn--test-n1-q01-group-0-6bc59-9v79c          48                    47920m                 97600484Ki            96449508Ki             v1.21.12            True
shoot--phbnrn--test-n1-q01-group-0-6bc59-m5jtc          48                    47920m                 97600484Ki            96449508Ki             v1.21.12            True
shoot--phbnrn--test-n1-q01-group-0-6bc59-rzk49          48                    47920m                 97604772Ki            96453796Ki             v1.21.12            True

PODS NAME                                                                NAMESPACE                    STATUS              MEM REQUEST          MEM LIMIT           CPU REQUEST          CPU LIMIT           NODE NAME
test-eck-cluster-nonprod-elasticsearch-exporter-es-576d9479nbmcz          elastic-test-nonprod          Running             128Mi                512Mi               100m                 500m                shoot--phbnrn--test-n1-q01-group-0-6bc59-rzk49
test-eck-cluster-nonprod-es-default-0                                     elastic-test-nonprod          Running             15G                  18650M              4                    6                   shoot--phbnrn--test-n1-q01-group-0-6bc59-rzk49
test-eck-cluster-nonprod-es-default-1                                     elastic-test-nonprod          Running             15G                  18650M              4                    6                   shoot--phbnrn--test-n1-q01-group-0-6bc59-9v79c
test-eck-cluster-nonprod-es-default-2                                     elastic-test-nonprod          Running             15G                  18650M              4                    6                   shoot--phbnrn--test-n1-q01-group-0-6bc59-m5jtc
test-nbg-nonprod-kb-765549b8bc-f86wz                                      elastic-test-nonprod          Running             2G                   4G                  1                    2                   shoot--phbnrn--test-n1-q01-group-0-6bc59-rzk49
test-nbg-nonprod-kb-765549b8bc-qh5lh                                      elastic-test-nonprod          Running             2G                   4G                  1                    2                   shoot--phbnrn--test-n1-q01-group-0-6bc59-9v79c
elastic-operator-0                                                       elastic-test-nonprod          Running             150Mi                1Gi                 100m                 1                   shoot--phbnrn--test-n1-q01-group-0-6bc59-rzk49
prometheus-elastic-test-nonprod-0                                         elastic-test-nonprod          Running             Not set              Not set             Not set              Not set             shoot--phbnrn--test-n1-q01-group-0-6bc59-m5jtc

ES NAME                                    STATUS              VERSION             PHASE               NODES               NAMESPACE
test-eck-cluster-nonprod                    green               7.17.4              Ready               3                   elastic-test-nonprod

KB NAME                            STATUS              VERSION             PHASE                NODES               NAMESPACE
test-nbg-nonprod                    green               7.17.4              Established          2                   elastic-test-nonprod

STATEFULSET NAME                                      REPLICAS            NAMESPACE                    HEAP SIZE
test-eck-cluster-nonprod-es-default                    3                   elastic-test-nonprod          -Xms12G -Xmx12G
elastic-operator                                      1                   elastic-test-nonprod
prometheus-elastic-test-nonprod                        1                   elastic-test-nonprod

DEPLOYMENT NAME                                                      REPLICAS            NAMESPACE
test-eck-cluster-nonprod-elasticsearch-exporter-es                    1                   elastic-test-nonprod
test-nbg-nonprod-kb                                                   2                   elastic-test-nonprod

PVC NAME                                                         STATUS              CAPACITY            VOLUME NAME                                       SC NAME                 ACCESS MODE
elasticsearch-data-test-eck-cluster-nonprod-es-default-0          Bound               466Gi               pvc-c1b6d012-e54b-4e6a-8883-4c96f5623c6d          partition-gold          [ReadWriteOnce]
elasticsearch-data-test-eck-cluster-nonprod-es-default-1          Bound               466Gi               pvc-e726acd4-85bb-4d65-88f1-275b4a192268          partition-gold          [ReadWriteOnce]
elasticsearch-data-test-eck-cluster-nonprod-es-default-2          Bound               466Gi               pvc-add43951-4847-4b2c-938b-9808acae2f15          partition-gold          [ReadWriteOnce]
prometheus-db-prometheus-elastic-test-nonprod-0                   Bound               5Gi                 pvc-1d7a2dfa-5b6a-4364-b83a-9c856c727240          csi-lvm                 [ReadWriteOnce]

SERVICE NAME                                               TYPE                  NAMESPACE                    ENDPOINTS
test-eck-cluster-nonprod-elasticsearch-exporter-es          ClusterIP             elastic-test-nonprod          10.244.16.86,
test-eck-cluster-nonprod-es-default                         ClusterIP             elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-eck-cluster-nonprod-es-elastic-443-mpls                LoadBalancer          elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-eck-cluster-nonprod-es-http                            ClusterIP             elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-eck-cluster-nonprod-es-internal-http                   ClusterIP             elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-eck-cluster-nonprod-es-transport                       ClusterIP             elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-eck-cluster-nonprod-es-transport-300-mpls              LoadBalancer          elastic-test-nonprod          10.244.12.78, 10.244.16.94, 10.244.20.129,
test-nbg-nonprod-kb-http                                    ClusterIP             elastic-test-nonprod          10.244.12.77, 10.244.16.93,
elastic-webhook-server                                     ClusterIP             elastic-test-nonprod          10.244.16.92,
prometheus-elastic-test-nonprod                             ClusterIP             elastic-test-nonprod          10.244.20.125,

STORAGE CLASS NAME          PROVISIONER                     ALLOW EXPANSION          VOLUME BIND MODE              RECLAIM POLICY
csi-lvm                     metal-stack.io/csi-lvm          false                    WaitForFirstConsumer          Delete
csi-lvm-sc-linear           lvm.csi.metal-stack.io          true                     WaitForFirstConsumer          Delete
csi-lvm-sc-mirror           lvm.csi.metal-stack.io          true                     WaitForFirstConsumer          Delete
csi-lvm-sc-striped          lvm.csi.metal-stack.io          true                     WaitForFirstConsumer          Delete
partition-gold              csi.lightbitslabs.com           true                     Immediate                     Delete
partition-silver            csi.lightbitslabs.com           true                     Immediate                     Delete

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