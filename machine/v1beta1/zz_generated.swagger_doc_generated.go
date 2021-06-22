package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AWSMachineProviderCondition = map[string]string{
	"":                   "AWSMachineProviderCondition is a condition in a AWSMachineProviderStatus.",
	"type":               "Type is the type of the condition.",
	"status":             "Status is the status of the condition.",
	"lastProbeTime":      "LastProbeTime is the last time we probed the condition.",
	"lastTransitionTime": "LastTransitionTime is the last time the condition transitioned from one status to another.",
	"reason":             "Reason is a unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "Message is a human-readable message indicating details about last transition.",
}

func (AWSMachineProviderCondition) SwaggerDoc() map[string]string {
	return map_AWSMachineProviderCondition
}

var map_AWSMachineProviderConfig = map[string]string{
	"":                   "AWSMachineProviderConfig is the Schema for the awsmachineproviderconfigs API",
	"ami":                "AMI is the reference to the AMI from which to create the machine instance.",
	"instanceType":       "InstanceType is the type of instance to create. Example: m4.xlarge",
	"tags":               "Tags is the set of tags to add to apply to an instance, in addition to the ones added by default by the actuator. These tags are additive. The actuator will ensure these tags are present, but will not remove any other tags that may exist on the instance.",
	"iamInstanceProfile": "IAMInstanceProfile is a reference to an IAM role to assign to the instance",
	"userDataSecret":     "UserDataSecret contains a local reference to a secret that contains the UserData to apply to the instance",
	"credentialsSecret":  "CredentialsSecret is a reference to the secret with AWS credentials. Otherwise, defaults to permissions provided by attached IAM role where the actuator is running.",
	"keyName":            "KeyName is the name of the KeyPair to use for SSH",
	"deviceIndex":        "DeviceIndex is the index of the device on the instance for the network interface attachment. Defaults to 0.",
	"publicIp":           "PublicIP specifies whether the instance should get a public IP. If not present, it should use the default of its subnet.",
	"securityGroups":     "SecurityGroups is an array of references to security groups that should be applied to the instance.",
	"subnet":             "Subnet is a reference to the subnet to use for this instance",
	"placement":          "Placement specifies where to create the instance in AWS",
	"loadBalancers":      "LoadBalancers is the set of load balancers to which the new instance should be added once it is created.",
	"blockDevices":       "BlockDevices is the set of block device mapping associated to this instance, block device without a name will be used as a root device and only one device without a name is allowed https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html",
	"spotMarketOptions":  "SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.",
}

func (AWSMachineProviderConfig) SwaggerDoc() map[string]string {
	return map_AWSMachineProviderConfig
}

var map_AWSMachineProviderConfigList = map[string]string{
	"": "AWSMachineProviderConfigList contains a list of AWSMachineProviderConfig",
}

func (AWSMachineProviderConfigList) SwaggerDoc() map[string]string {
	return map_AWSMachineProviderConfigList
}

var map_AWSMachineProviderStatus = map[string]string{
	"":              "AWSMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field. It contains AWS-specific status information.",
	"instanceId":    "InstanceID is the instance ID of the machine created in AWS",
	"instanceState": "InstanceState is the state of the AWS instance for this machine",
	"conditions":    "Conditions is a set of conditions associated with the Machine to indicate errors or other status",
}

func (AWSMachineProviderStatus) SwaggerDoc() map[string]string {
	return map_AWSMachineProviderStatus
}

var map_AWSResourceReference = map[string]string{
	"":        "AWSResourceReference is a reference to a specific AWS resource by ID, ARN, or filters. Only one of ID, ARN or Filters may be specified. Specifying more than one will result in a validation error.",
	"id":      "ID of resource",
	"arn":     "ARN of resource",
	"filters": "Filters is a set of filters used to identify a resource",
}

func (AWSResourceReference) SwaggerDoc() map[string]string {
	return map_AWSResourceReference
}

var map_BlockDeviceMappingSpec = map[string]string{
	"":            "BlockDeviceMappingSpec describes a block device mapping",
	"deviceName":  "The device name exposed to the machine (for example, /dev/sdh or xvdh).",
	"ebs":         "Parameters used to automatically set up EBS volumes when the machine is launched.",
	"noDevice":    "Suppresses the specified device included in the block device mapping of the AMI.",
	"virtualName": "The virtual device name (ephemeralN). Machine store volumes are numbered starting from 0. An machine type with 2 available machine store volumes can specify mappings for ephemeral0 and ephemeral1.The number of available machine store volumes depends on the machine type. After you connect to the machine, you must mount the volume.\n\nConstraints: For M3 machines, you must specify machine store volumes in the block device mapping for the machine. When you launch an M3 machine, we ignore any machine store volumes specified in the block device mapping for the AMI.",
}

func (BlockDeviceMappingSpec) SwaggerDoc() map[string]string {
	return map_BlockDeviceMappingSpec
}

var map_EBSBlockDeviceSpec = map[string]string{
	"":                    "EBSBlockDeviceSpec describes a block device for an EBS volume. https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EbsBlockDevice",
	"deleteOnTermination": "Indicates whether the EBS volume is deleted on machine termination.",
	"encrypted":           "Indicates whether the EBS volume is encrypted. Encrypted Amazon EBS volumes may only be attached to machines that support Amazon EBS encryption.",
	"kmsKey":              "Indicates the KMS key that should be used to encrypt the Amazon EBS volume.",
	"iops":                "The number of I/O operations per second (IOPS) that the volume supports. For io1, this represents the number of IOPS that are provisioned for the volume. For gp2, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting. For more information about General Purpose SSD baseline performance, I/O credits, and bursting, see Amazon EBS Volume Types (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the Amazon Elastic Compute Cloud User Guide.\n\nMinimal and maximal IOPS for io1 and gp2 are constrained. Please, check https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html for precise boundaries for individual volumes.\n\nCondition: This parameter is required for requests to create io1 volumes; it is not used in requests to create gp2, st1, sc1, or standard volumes.",
	"volumeSize":          "The size of the volume, in GiB.\n\nConstraints: 1-16384 for General Purpose SSD (gp2), 4-16384 for Provisioned IOPS SSD (io1), 500-16384 for Throughput Optimized HDD (st1), 500-16384 for Cold HDD (sc1), and 1-1024 for Magnetic (standard) volumes. If you specify a snapshot, the volume size must be equal to or larger than the snapshot size.\n\nDefault: If you're creating the volume from a snapshot and don't specify a volume size, the default is the snapshot size.",
	"volumeType":          "The volume type: gp2, io1, st1, sc1, or standard. Default: standard",
}

func (EBSBlockDeviceSpec) SwaggerDoc() map[string]string {
	return map_EBSBlockDeviceSpec
}

var map_Filter = map[string]string{
	"":       "Filter is a filter used to identify an AWS resource",
	"name":   "Name of the filter. Filter names are case-sensitive.",
	"values": "Values includes one or more filter values. Filter values are case-sensitive.",
}

func (Filter) SwaggerDoc() map[string]string {
	return map_Filter
}

var map_LoadBalancerReference = map[string]string{
	"": "LoadBalancerReference is a reference to a load balancer on AWS.",
}

func (LoadBalancerReference) SwaggerDoc() map[string]string {
	return map_LoadBalancerReference
}

var map_Placement = map[string]string{
	"":                 "Placement indicates where to create the instance in AWS",
	"region":           "Region is the region to use to create the instance",
	"availabilityZone": "AvailabilityZone is the availability zone of the instance",
	"tenancy":          "Tenancy indicates if instance should run on shared or single-tenant hardware. There are supported 3 options: default, dedicated and host.",
}

func (Placement) SwaggerDoc() map[string]string {
	return map_Placement
}

var map_SpotMarketOptions = map[string]string{
	"":         "SpotMarketOptions defines the options available to a user when configuring Machines to run on Spot instances. Most users should provide an empty struct.",
	"maxPrice": "The maximum price the user is willing to pay for their instances Default: On-Demand price",
}

func (SpotMarketOptions) SwaggerDoc() map[string]string {
	return map_SpotMarketOptions
}

var map_TagSpecification = map[string]string{
	"":      "TagSpecification is the name/value pair for a tag",
	"name":  "Name of the tag",
	"value": "Value of the tag",
}

func (TagSpecification) SwaggerDoc() map[string]string {
	return map_TagSpecification
}

var map_LastOperation = map[string]string{
	"":            "LastOperation represents the detail of the last performed operation on the MachineObject.",
	"description": "Description is the human-readable description of the last operation.",
	"lastUpdated": "LastUpdated is the timestamp at which LastOperation API was last-updated.",
	"state":       "State is the current status of the last performed operation. E.g. Processing, Failed, Successful etc",
	"type":        "Type is the type of operation which was last performed. E.g. Create, Delete, Update etc",
}

func (LastOperation) SwaggerDoc() map[string]string {
	return map_LastOperation
}

var map_Machine = map[string]string{
	"": "Machine is the Schema for the machines API",
}

func (Machine) SwaggerDoc() map[string]string {
	return map_Machine
}

var map_MachineList = map[string]string{
	"": "MachineList contains a list of Machine",
}

func (MachineList) SwaggerDoc() map[string]string {
	return map_MachineList
}

var map_MachineSpec = map[string]string{
	"":             "MachineSpec defines the desired state of Machine",
	"metadata":     "ObjectMeta will autopopulate the Node created. Use this to indicate what labels, annotations, name prefix, etc., should be used when creating the Node.",
	"taints":       "The list of the taints to be applied to the corresponding Node in additive manner. This list will not overwrite any other taints added to the Node on an ongoing basis by other entities. These taints should be actively reconciled e.g. if you ask the machine controller to apply a taint and then manually remove the taint the machine controller will put it back) but not have the machine controller remove any taints",
	"providerSpec": "ProviderSpec details Provider-specific configuration to use during node creation.",
	"providerID":   "ProviderID is the identification ID of the machine provided by the provider. This field must match the provider ID as seen on the node object corresponding to this machine. This field is required by higher level consumers of cluster-api. Example use case is cluster autoscaler with cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find out machines at provider which could not get registered as Kubernetes nodes. With cluster-api as a generic out-of-tree provider for autoscaler, this field is required by autoscaler to be able to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserver and then a comparison is done to find out unregistered machines and are marked for delete. This field will be set by the actuators and consumed by higher level entities like autoscaler that will be interfacing with cluster-api as generic provider.",
}

func (MachineSpec) SwaggerDoc() map[string]string {
	return map_MachineSpec
}

var map_MachineStatus = map[string]string{
	"":               "MachineStatus defines the observed state of Machine",
	"nodeRef":        "NodeRef will point to the corresponding Node if it exists.",
	"lastUpdated":    "LastUpdated identifies when this status was last observed.",
	"errorReason":    "ErrorReason will be set in the event that there is a terminal problem reconciling the Machine and will contain a succinct value suitable for machine interpretation.\n\nThis field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output.",
	"errorMessage":   "ErrorMessage will be set in the event that there is a terminal problem reconciling the Machine and will contain a more verbose string suitable for logging and human consumption.\n\nThis field should not be set for transitive errors that a controller faces that are expected to be fixed automatically over time (like service outages), but instead indicate that something is fundamentally wrong with the Machine's spec or the configuration of the controller, and that manual intervention is required. Examples of terminal errors would be invalid combinations of settings in the spec, values that are unsupported by the controller, or the responsible controller itself being critically misconfigured.\n\nAny transient errors that occur during the reconciliation of Machines can be added as events to the Machine object and/or logged in the controller's output.",
	"providerStatus": "ProviderStatus details a Provider-specific status. It is recommended that providers maintain their own versioned API types that should be serialized/deserialized from this field.",
	"addresses":      "Addresses is a list of addresses assigned to the machine. Queried from cloud provider, if available.",
	"lastOperation":  "LastOperation describes the last-operation performed by the machine-controller. This API should be useful as a history in terms of the latest operation performed on the specific machine. It should also convey the state of the latest-operation for example if it is still on-going, failed or completed successfully.",
	"phase":          "Phase represents the current phase of machine actuation. One of: Failed, Provisioning, Provisioned, Running, Deleting",
	"conditions":     "Conditions defines the current state of the Machine",
}

func (MachineStatus) SwaggerDoc() map[string]string {
	return map_MachineStatus
}

var map_Condition = map[string]string{
	"":                   "Condition defines an observation of a Machine API resource operational state.",
	"type":               "Type of condition in CamelCase or in foo.example.com/CamelCase. Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"severity":           "Severity provides an explicit classification of Reason code, so the users or machines can immediately understand the current situation and act accordingly. The Severity field MUST be set only when Status=False.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another. This should be when the underlying condition changed. If that is not known, then using the time when the API field changed is acceptable.",
	"reason":             "The reason for the condition's last transition in CamelCase. The specific API may choose whether or not this field is considered a guaranteed API. This field may not be empty.",
	"message":            "A human readable message indicating details about the transition. This field may be empty.",
}

func (Condition) SwaggerDoc() map[string]string {
	return map_Condition
}

var map_ObjectMeta = map[string]string{
	"":                "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. This is a copy of customizable fields from metav1.ObjectMeta.\n\nObjectMeta is embedded in `Machine.Spec`, `MachineDeployment.Template` and `MachineSet.Template`, which are not top-level Kubernetes objects. Given that metav1.ObjectMeta has lots of special cases and read-only fields which end up in the generated CRD validation, having it as a subset simplifies the API and some issues that can impact user experience.\n\nDuring the [upgrade to controller-tools@v2](https://github.com/kubernetes-sigs/cluster-api/pull/1054) for v1alpha2, we noticed a failure would occur running Cluster API test suite against the new CRDs, specifically `spec.metadata.creationTimestamp in body must be of type string: \"null\"`. The investigation showed that `controller-tools@v2` behaves differently than its previous version when handling types from [metav1](k8s.io/apimachinery/pkg/apis/meta/v1) package.\n\nIn more details, we found that embedded (non-top level) types that embedded `metav1.ObjectMeta` had validation properties, including for `creationTimestamp` (metav1.Time). The `metav1.Time` type specifies a custom json marshaller that, when IsZero() is true, returns `null` which breaks validation because the field isn't marked as nullable.\n\nIn future versions, controller-tools@v2 might allow overriding the type and validation for embedded types. When that happens, this hack should be revisited.",
	"name":            "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
	"generateName":    "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
	"namespace":       "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
	"labels":          "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
	"annotations":     "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
	"ownerReferences": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
}

func (ObjectMeta) SwaggerDoc() map[string]string {
	return map_ObjectMeta
}

var map_ProviderSpec = map[string]string{
	"":      "ProviderSpec defines the configuration to use during node creation.",
	"value": "Value is an inlined, serialized representation of the resource configuration. It is recommended that providers maintain their own versioned API types that should be serialized/deserialized from this field, akin to component config.",
}

func (ProviderSpec) SwaggerDoc() map[string]string {
	return map_ProviderSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
