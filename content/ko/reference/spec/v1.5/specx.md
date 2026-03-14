---
title: "Experimental"
weight: 20
description: "Gateway API v1.5 specification reference for experimental channel resources"
---
- [gateway.networking.x-k8s.io/v1alpha1](#gatewaynetworkingx-k8siov1alpha1)


## gateway.networking.x-k8s.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the gateway.networking.k8s-x.io
API group.


### Resource Types
- [XBackendTrafficPolicy](#xbackendtrafficpolicy)
- [XMesh](#xmesh)



#### BackendTrafficPolicySpec



BackendTrafficPolicySpec define the desired state of BackendTrafficPolicy
Note: there is no Override or Default policy configuration.



_Appears in:_

- [XBackendTrafficPolicy](#xbackendtrafficpolicy)


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `targetRefs` _[LocalPolicyTargetReference](#localpolicytargetreference) array_| TargetRefs identifies API object(s) to apply this policy to.<br />Currently, Backends (A grouping of like endpoints such as Service,<br />ServiceImport, or any implementation-specific backendRef) are the only<br />valid API target references.<br />Currently, a TargetRef cannot be scoped to a specific port on a<br />Service. |  | MaxItems: 16 <br />MinItems: 1 <br /> |
| `retryConstraint` _[RetryConstraint](#retryconstraint)_<br /> :warning: **Experimental**| RetryConstraint defines the configuration for when to allow or prevent<br />further retries to a target backend, by dynamically calculating a 'retry<br />budget'. This budget is calculated based on the percentage of incoming<br />traffic composed of retries over a given time interval. Once the budget<br />is exceeded, additional retries will be rejected.<br />For example, if the retry budget interval is 10 seconds, there have been<br />1000 active requests in the past 10 seconds, and the allowed percentage<br />of requests that can be retried is 20% (the default), then 200 of those<br />requests may be composed of retries. Active requests will only be<br />considered for the duration of the interval when calculating the retry<br />budget. Retrying the same original request multiple times within the<br />retry budget interval will lead to each retry being counted towards<br />calculating the budget.<br />Configuring a RetryConstraint in BackendTrafficPolicy is compatible with<br />HTTPRoute Retry settings for each HTTPRouteRule that targets the same<br />backend. While the HTTPRouteRule Retry stanza can specify whether a<br />request will be retried, and the number of retry attempts each client<br />may perform, RetryConstraint helps prevent cascading failures such as<br />retry storms during periods of consistent failures.<br />After the retry budget has been exceeded, additional retries to the<br />backend MUST return a 503 response to the client.<br />Additional configurations for defining a constraint on retries MAY be<br />defined in the future.<br />Support: Extended<br /><gateway:experimental> |  |  |
| `sessionPersistence` _[SessionPersistence](#sessionpersistence)_| SessionPersistence defines and configures session persistence<br />for the backend.<br />Support: Extended |  |  |


#### BudgetDetails



BudgetDetails specifies the details of the budget configuration, like
the percentage of requests in the budget, and the interval between
checks.



_Appears in:_

- [RetryConstraint](#retryconstraint)


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `percent` _integer_| Percent defines the maximum percentage of active requests that may<br />be made up of retries.<br />Support: Extended | 20 | Maximum: 100 <br />Minimum: 0 <br /> |
| `interval` _[Duration](#duration)_| Interval defines the duration in which requests will be considered<br />for calculating the budget for retries.<br />Support: Extended | 10s |  |










#### MeshSpec



MeshSpec defines the desired state of an XMesh.



_Appears in:_

- [XMesh](#xmesh)


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `controllerName` _[GatewayController](#gatewaycontroller)_| ControllerName is the name of a controller that is managing Gateway API<br />resources for mesh traffic management. The value of this field MUST be a<br />domain prefixed path.<br />Example: "example.com/awesome-mesh".<br />This field is not mutable and cannot be empty.<br />Support: Core |  |  |
| `parametersRef` _[ParametersReference](#parametersreference)_| ParametersRef is an optional reference to a resource that contains<br />implementation-specific configuration for this Mesh. If no<br />implementation-specific parameters are needed, this field MUST be<br />omitted.<br />ParametersRef can reference a standard Kubernetes resource, i.e.<br />ConfigMap, or an implementation-specific custom resource. The resource<br />can be cluster-scoped or namespace-scoped.<br />If the referent cannot be found, refers to an unsupported kind, or when<br />the data within that resource is malformed, the Mesh MUST be rejected<br />with the "Accepted" status condition set to "False" and an<br />"InvalidParameters" reason.<br />Support: Implementation-specific |  |  |
| `description` _string_| Description optionally provides a human-readable description of a Mesh. |  | MaxLength: 64 <br /> |


#### MeshStatus



MeshStatus is the current status for the Mesh.



_Appears in:_

- [XMesh](#xmesh)


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `conditions` _[Condition](https://kubernetes.io/reference/generated/kubernetes-api/v1.32/#condition-v1-meta) array_| Conditions is the current status from the controller for<br />this Mesh.<br />Controllers should prefer to publish conditions using values<br />of MeshConditionType for the type of each Condition. | [map[lastTransitionTime:1970-01-01T00:00:00Z message:Waiting for controller reason:Pending status:Unknown type:Accepted] map[lastTransitionTime:1970-01-01T00:00:00Z message:Waiting for controller reason:Pending status:Unknown type:Programmed]] | MaxItems: 8 <br /> |
| `supportedFeatures` _SupportedFeature array_| SupportedFeatures is the set of features the Mesh support.<br />It MUST be sorted in ascending alphabetical order by the Name key. |  | MaxItems: 64 <br /> |




#### RequestRate



RequestRate expresses a rate of requests over a given period of time.



_Appears in:_

- [RetryConstraint](#retryconstraint)


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `count` _integer_| Count specifies the number of requests per time interval.<br />Support: Extended |  | Maximum: 1e+06 <br />Minimum: 1 <br /> |
| `interval` _[Duration](#duration)_| Interval specifies the divisor of the rate of requests, the amount of<br />time during which the given count of requests occur.<br />Support: Extended |  |  |


#### RetryConstraint



RetryConstraint defines the configuration for when to retry a request.



_Appears in:_

- [BackendTrafficPolicySpec](#backendtrafficpolicyspec) :warning: Experimental in `retryConstraint` field


| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `budget` _[BudgetDetails](#budgetdetails)_| Budget holds the details of the retry budget configuration. | \{ interval:10s percent:20 \} |  |
| `minRetryRate` _[RequestRate](#requestrate)_| MinRetryRate defines the minimum rate of retries that will be allowable<br />over a specified duration of time.<br />The effective overall minimum rate of retries targeting the backend<br />service may be much higher, as there can be any number of clients which<br />are applying this setting locally.<br />This ensures that requests can still be retried during periods of low<br />traffic, where the budget for retries may be calculated as a very low<br />value.<br />Support: Extended | \{ count:10 interval:1s \} |  |




#### XBackendTrafficPolicy



XBackendTrafficPolicy defines the configuration for how traffic to a
target backend should be handled.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gateway.networking.x-k8s.io/v1alpha1` | | |
| `kind` _string_ | `XBackendTrafficPolicy` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta)_| Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[BackendTrafficPolicySpec](#backendtrafficpolicyspec)_| Spec defines the desired state of BackendTrafficPolicy. |  |  |
| `status` _[PolicyStatus](#policystatus)_| Status defines the current state of BackendTrafficPolicy. |  |  |


#### XMesh



XMesh defines mesh-wide characteristics of a GAMMA-compliant service mesh.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gateway.networking.x-k8s.io/v1alpha1` | | |
| `kind` _string_ | `XMesh` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta)_| Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MeshSpec](#meshspec)_| Spec defines the desired state of XMesh. |  |  |
| `status` _[MeshStatus](#meshstatus)_| Status defines the current state of XMesh.<br /> | \{ conditions:[map[lastTransitionTime:1970-01-01T00:00:00Z message:Waiting for controller reason:Pending status:Unknown type:Accepted]] \} |  |

