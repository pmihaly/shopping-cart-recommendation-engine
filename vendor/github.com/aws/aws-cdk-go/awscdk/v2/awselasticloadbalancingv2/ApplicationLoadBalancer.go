package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define an Application Load Balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var asg autoScalingGroup
//   var vpc vpc
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener and open up the load balancer's security group
//   // to the world.
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//
//   	// 'open: true' is the default, you can leave it out if you want. Set it
//   	// to 'false' and use `listener.connections` if you want to be selective
//   	// about who can access the load balancer.
//   	Open: jsii.Boolean(true),
//   })
//
//   // Create an AutoScaling group and add it as a load balancing
//   // target to the listener.
//   listener.AddTargets(jsii.String("ApplicationFleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(8080),
//   	Targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type ApplicationLoadBalancer interface {
	BaseLoadBalancer
	IApplicationLoadBalancer
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The IP Address Type for this load balancer.
	IpAddressType() IpAddressType
	// A list of listeners that have been added to the load balancer.
	//
	// This list is only valid for owned constructs.
	Listeners() *[]ApplicationListener
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	// All metrics available for this load balancer.
	Metrics() IApplicationLoadBalancerMetrics
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
	// Add a new listener to this load balancer.
	AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener
	// Add a redirection listener to this load balancer.
	AddRedirect(props *ApplicationLoadBalancerRedirectConfig) ApplicationListener
	// Add a security group to this load balancer.
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Return the given named metric for this Application Load Balancer.
	// Default: Average over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.custom`` instead
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of concurrent TCP connections active from clients to the load balancer and from the load balancer to targets.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.activeConnectionCount`` instead
	MetricActiveConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the client that did not establish a session with the load balancer.
	//
	// Possible causes include a
	// mismatch of ciphers or protocols.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.clientTlsNegotiationErrorCount`` instead
	MetricClientTlsNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of load balancer capacity units (LCU) used by your load balancer.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.consumedLCUs`` instead
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of user authentications that could not be completed.
	//
	// Because an authenticate action was misconfigured, the load balancer
	// couldn't establish a connection with the IdP, or the load balancer
	// couldn't complete the authentication flow due to an internal error.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.elbAuthError`` instead
	MetricElbAuthError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of user authentications that could not be completed because the IdP denied access to the user or an authorization code was used more than once.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.elbAuthFailure`` instead
	MetricElbAuthFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in milliseconds, to query the IdP for the ID token and user info.
	//
	// If one or more of these operations fail, this is the time to failure.
	// Default: Average over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.elbAuthLatency`` instead
	MetricElbAuthLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of authenticate actions that were successful.
	//
	// This metric is incremented at the end of the authentication workflow,
	// after the load balancer has retrieved the user claims from the IdP.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.elbAuthSuccess`` instead
	MetricElbAuthSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 3xx/4xx/5xx codes that originate from the load balancer.
	//
	// This does not include any response codes generated by the targets.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.httpCodeElb`` instead
	MetricHttpCodeElb(code HttpCodeElb, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in the load balancer.
	//
	// This does not include any response codes generated by the load balancer.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.httpCodeTarget`` instead
	MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of fixed-response actions that were successful.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.httpFixedResponseCount`` instead
	MetricHttpFixedResponseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of redirect actions that were successful.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.httpRedirectCount`` instead
	MetricHttpRedirectCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of redirect actions that couldn't be completed because the URL in the response location header is larger than 8K.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.httpRedirectUrlLimitExceededCount`` instead
	MetricHttpRedirectUrlLimitExceededCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer over IPv6.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.ipv6ProcessedBytes`` instead
	MetricIpv6ProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of IPv6 requests received by the load balancer.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.ipv6RequestCount`` instead
	MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of new TCP connections established from clients to the load balancer and from the load balancer to targets.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.newConnectionCount`` instead
	MetricNewConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer over IPv4 and IPv6.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.processedBytes`` instead
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were rejected because the load balancer had reached its maximum number of connections.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.rejectedConnectionCount`` instead
	MetricRejectedConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of requests processed over IPv4 and IPv6.
	//
	// This count includes only the requests with a response generated by a target of the load balancer.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.requestCount`` instead
	MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of rules processed by the load balancer given a request rate averaged over an hour.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.ruleEvaluations`` instead
	MetricRuleEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were not successfully established between the load balancer and target.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.targetConnectionErrorCount`` instead
	MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
	// Default: Average over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.targetResponseTime`` instead
	MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
	//
	// Possible causes include a mismatch of ciphers or protocols.
	// Default: Sum over 5 minutes.
	//
	// Deprecated: Use ``ApplicationLoadBalancer.metrics.targetTLSNegotiationErrorCount`` instead
	MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	ResourcePolicyPrincipal() awsiam.IPrincipal
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
}

// The jsii proxy struct for ApplicationLoadBalancer
type jsiiProxy_ApplicationLoadBalancer struct {
	jsiiProxy_BaseLoadBalancer
	jsiiProxy_IApplicationLoadBalancer
}

func (j *jsiiProxy_ApplicationLoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) IpAddressType() IpAddressType {
	var returns IpAddressType
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Listeners() *[]ApplicationListener {
	var returns *[]ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Metrics() IApplicationLoadBalancerMetrics {
	var returns IApplicationLoadBalancerMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewApplicationLoadBalancer(scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) ApplicationLoadBalancer {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationLoadBalancer_Override(a ApplicationLoadBalancer, scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing Application Load Balancer.
func ApplicationLoadBalancer_FromApplicationLoadBalancerAttributes(scope constructs.Construct, id *string, attrs *ApplicationLoadBalancerAttributes) IApplicationLoadBalancer {
	_init_.Initialize()

	if err := validateApplicationLoadBalancer_FromApplicationLoadBalancerAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IApplicationLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"fromApplicationLoadBalancerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Look up an application load balancer.
func ApplicationLoadBalancer_FromLookup(scope constructs.Construct, id *string, options *ApplicationLoadBalancerLookupOptions) IApplicationLoadBalancer {
	_init_.Initialize()

	if err := validateApplicationLoadBalancer_FromLookupParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns IApplicationLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApplicationLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ApplicationLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancer_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApplicationLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplicationLoadBalancer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener {
	if err := a.validateAddListenerParameters(id, props); err != nil {
		panic(err)
	}
	var returns ApplicationListener

	_jsii_.Invoke(
		a,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) AddRedirect(props *ApplicationLoadBalancerRedirectConfig) ApplicationListener {
	if err := a.validateAddRedirectParameters(props); err != nil {
		panic(err)
	}
	var returns ApplicationListener

	_jsii_.Invoke(
		a,
		"addRedirect",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	if err := a.validateAddSecurityGroupParameters(securityGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	if err := a.validateLogAccessLogsParameters(bucket); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricActiveConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricActiveConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricActiveConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricClientTlsNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricClientTlsNegotiationErrorCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricClientTlsNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricConsumedLCUsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricConsumedLCUs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricElbAuthErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricElbAuthFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricElbAuthLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricElbAuthSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpCodeElb(code HttpCodeElb, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHttpCodeElbParameters(code, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpCodeElb",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHttpCodeTargetParameters(code, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpCodeTarget",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpFixedResponseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHttpFixedResponseCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpFixedResponseCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpRedirectCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHttpRedirectCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpRedirectCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpRedirectUrlLimitExceededCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHttpRedirectUrlLimitExceededCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpRedirectUrlLimitExceededCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricIpv6ProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricIpv6ProcessedBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricIpv6ProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricIpv6RequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricIpv6RequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricNewConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricNewConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricNewConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricProcessedBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricRejectedConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricRejectedConnectionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRejectedConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricRuleEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricRuleEvaluationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRuleEvaluations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTargetConnectionErrorCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetConnectionErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTargetResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTargetTLSNegotiationErrorCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetTLSNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) RemoveAttribute(key *string) {
	if err := a.validateRemoveAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) ResourcePolicyPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal

	_jsii_.Invoke(
		a,
		"resourcePolicyPrincipal",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) SetAttribute(key *string, value *string) {
	if err := a.validateSetAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"setAttribute",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

