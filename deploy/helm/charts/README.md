
#  OpenEBS LocalPV Provisioner

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Chart Lint and Test](https://github.com/openebs/device-localpv/workflows/Chart%20Lint%20and%20Test/badge.svg)
![Release Charts](https://github.com/openebs/device-localpv/workflows/Release%20Charts/badge.svg?branch=develop)

A Helm chart for openebs device localpv provisioner. This chart bootstraps OpenEBS Device LocalPV provisioner deployment on a [Kubernetes](http://kubernetes.io) cluster using the  [Helm](https://helm.sh) package manager.


**Homepage:** <http://www.openebs.io/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| pawanpraka1 | pawan@mayadata.io |  |
| prateekpandey14 | prateek.pandey@mayadata.io |  |


## Get Repo Info

```console
helm repo add openebs-devicelocalpv https://openebs.github.io/device-localpv
helm repo update
```

_See [helm repo](https://helm.sh/docs/helm/helm_repo/) for command documentation._

## Install Chart

Please visit the [link](https://openebs.github.io/device-localpv/) for install instructions via helm3.

```console
# Helm
$ helm install [RELEASE_NAME] openebs-devicelocalpv/device-localpv
```

**Note:** If moving from the operator to helm
- Make sure the namespace provided in the helm install command is same as `DEVICE_DRIVER_NAMESPACE` (by default it is `openebs`) env in the controller statefulset.
- Before installing, clean up the stale statefulset and daemonset from `kube-system` namespace using the below commands
```sh
kubectl delete sts openebs-device-controller -n kube-system
kubectl delete ds openebs-device-node -n kube-system
```


_See [configuration](#configuration) below._

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Uninstall Chart

```console
# Helm
$ helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

```console
# Helm
$ helm upgrade [RELEASE_NAME] [CHART] --install
```

## Configuration

The following table lists the configurable parameters of the OpenEBS Device Localpv chart and their default values.

| Parameter| Description| Default|
| -| -| -|
| `imagePullSecrets`| Provides image pull secrect| `""`|
| `devicePlugin.image.registry`| Registry for openebs-device-plugin image| `""`|
| `devicePlugin.image.repository`| Image repository for openebs-device-plugin| `openebs/device-driver`|
| `devicePlugin.image.pullPolicy`| Image pull policy for openebs-device-plugin| `IfNotPresent`|
| `devicePlugin.image.tag`| Image tag for openebs-device-plugin| `0.1.0`|
| `deviceNode.driverRegistrar.image.registry`| Registry for csi-node-driver-registrar image| `k8s.gcr.io/`|
| `deviceNode.driverRegistrar.image.repository`| Image repository for csi-node-driver-registrar| `sig-storage/csi-node-driver-registrar`|
| `deviceNode.driverRegistrar.image.pullPolicy`| Image pull policy for csi-node-driver-registrar| `IfNotPresent`|
| `deviceNode.driverRegistrar.image.tag`| Image tag for csi-node-driver-registrar| `v1.2.0`|
| `deviceNode.updateStrategy.type`| Update strategy for devicenode daemonset | `RollingUpdate` |
| `deviceNode.kubeletDir`| Kubelet mount point for devicenode daemonset| `"/var/lib/kubelet/"` |
| `deviceNode.annotations` | Annotations for devicenode daemonset metadata| `""`|
| `deviceNode.podAnnotations`| Annotations for devicenode daemonset's pods metadata | `""`|
| `deviceNode.resources`| Resource and request and limit for devicenode daemonset containers | `""`|
| `deviceNode.labels`| Labels for devicenode daemonset metadata | `""`|
| `deviceNode.podLabels`| Appends labels to the devicenode daemonset pods| `""`|
| `deviceNode.nodeSelector`| Nodeselector for devicenode daemonset pods| `""`|
| `deviceNode.tolerations` | devicenode daemonset's pod toleration values | `""`|
| `deviceNode.securityContext` | Security context for devicenode daemonset container | `""`|
| `deviceController.resizer.image.registry`| Registry for csi-resizer image| `k8s.gcr.io/`|
| `deviceController.resizer.image.repository`| Image repository for csi-resizer| `sig-storage/csi-resizer`|
| `deviceController.resizer.image.pullPolicy`| Image pull policy for csi-resizer| `IfNotPresent`|
| `deviceController.resizer.image.tag`| Image tag for csi-resizer| `v1.1.0`|
| `deviceController.snapshotter.image.registry`| Registry for csi-snapshotter image| `k8s.gcr.io/`|
| `deviceController.snapshotter.image.repository`| Image repository for csi-snapshotter| `sig-storage/csi-snapshotter`|
| `deviceController.snapshotter.image.pullPolicy`| Image pull policy for csi-snapshotter| `IfNotPresent`|
| `deviceController.snapshotter.image.tag`| Image tag for csi-snapshotter| `v4.0.0`|
| `deviceController.snapshotController.image.registry`| Registry for snapshot-controller image| `k8s.gcr.io/`|
| `deviceController.snapshotController.image.repository`| Image repository for snapshot-controller| `sig-storage/snapshot-controller`|
| `deviceController.snapshotController.image.pullPolicy`| Image pull policy for snapshot-controller| `IfNotPresent`|
| `deviceController.snapshotController.image.tag`| Image tag for snapshot-controller| `v4.0.0`|
| `deviceController.provisioner.image.registry`| Registry for csi-provisioner image| `k8s.gcr.io/`|
| `deviceController.provisioner.image.repository`| Image repository for csi-provisioner| `sig-storage/csi-provisioner`|
| `deviceController.provisioner.image.pullPolicy`| Image pull policy for csi-provisioner| `IfNotPresent`|
| `deviceController.provisioner.image.tag`| Image tag for csi-provisioner| `v2.1.0`|
| `deviceController.updateStrategy.type`| Update strategy for device localpv controller statefulset | `RollingUpdate` |
| `deviceController.annotations` | Annotations for device localpv controller statefulset metadata| `""`|
| `deviceController.podAnnotations`| Annotations for device localpv controller statefulset's pods metadata | `""`|
| `deviceController.resources`| Resource and request and limit for device localpv controller statefulset containers | `""`|
| `deviceController.labels`| Labels for device localpv controller statefulset metadata | `""`|
| `deviceController.podLabels`| Appends labels to the device localpv controller statefulset pods| `""`|
| `deviceController.nodeSelector`| Nodeselector for device localpv controller statefulset pods| `""`|
| `deviceController.tolerations` | device localpv controller statefulset's pod toleration values | `""`|
| `deviceController.securityContext` | Seurity context for device localpv controller statefulset container | `""`|
| `rbac.pspEnabled` | Enable PodSecurityPolicy | `false` |
| `serviceAccount.deviceNode.create` | Create a service account for devicenode or not| `true`|
| `serviceAccount.deviceNode.name` | Name for the devicenode service account| `openebs-device-node-sa`|
| `serviceAccount.deviceController.create` | Create a service account for device localpv controller or not| `true`|
| `serviceAccount.deviceController.name` | Name for the device localpv controller service account| `openebs-device-controller-sa`|
| `analytics.enabled` | Enable or Disable google analytics for the controller| `true`|

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
helm install <release-name> -f values.yaml openebs/device-localpv
```

> **Tip**: You can use the default [values.yaml](values.yaml)
