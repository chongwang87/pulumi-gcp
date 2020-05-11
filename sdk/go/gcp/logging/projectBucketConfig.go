// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ProjectBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project pulumi.StringOutput `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewProjectBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewProjectBucketConfig(ctx *pulumi.Context,
	name string, args *ProjectBucketConfigArgs, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	if args == nil || args.BucketId == nil {
		return nil, errors.New("missing required argument 'BucketId'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &ProjectBucketConfigArgs{}
	}
	var resource ProjectBucketConfig
	err := ctx.RegisterResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectBucketConfig gets an existing ProjectBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectBucketConfigState, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	var resource ProjectBucketConfig
	err := ctx.ReadResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectBucketConfig resources.
type projectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project *string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type ProjectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigState)(nil)).Elem()
}

type projectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Project string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a ProjectBucketConfig resource.
type ProjectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigArgs)(nil)).Elem()
}