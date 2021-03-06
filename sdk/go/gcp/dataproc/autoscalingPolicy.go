// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes an autoscaling policy for Dataproc cluster autoscaler.
//
// ## Example Usage
// ### Dataproc Autoscaling Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		asp, err := dataproc.NewAutoscalingPolicy(ctx, "asp", &dataproc.AutoscalingPolicyArgs{
// 			PolicyId: pulumi.String("dataproc-policy"),
// 			Location: pulumi.String("us-central1"),
// 			WorkerConfig: &dataproc.AutoscalingPolicyWorkerConfigArgs{
// 				MaxInstances: pulumi.Int(3),
// 			},
// 			BasicAlgorithm: &dataproc.AutoscalingPolicyBasicAlgorithmArgs{
// 				YarnConfig: &dataproc.AutoscalingPolicyBasicAlgorithmYarnConfigArgs{
// 					GracefulDecommissionTimeout: pulumi.String("30s"),
// 					ScaleUpFactor:               pulumi.Float64(0.5),
// 					ScaleDownFactor:             pulumi.Float64(0.5),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewCluster(ctx, "basic", &dataproc.ClusterArgs{
// 			Region: pulumi.String("us-central1"),
// 			ClusterConfig: &dataproc.ClusterClusterConfigArgs{
// 				AutoscalingConfig: &dataproc.ClusterClusterConfigAutoscalingConfigArgs{
// 					PolicyUri: asp.Name,
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// AutoscalingPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{project}}/{{location}}/{{policy_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{location}}/{{policy_id}}
// ```
type AutoscalingPolicy struct {
	pulumi.CustomResourceState

	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm AutoscalingPolicyBasicAlgorithmPtrOutput `pulumi:"basicAlgorithm"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig AutoscalingPolicySecondaryWorkerConfigPtrOutput `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig AutoscalingPolicyWorkerConfigPtrOutput `pulumi:"workerConfig"`
}

// NewAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *AutoscalingPolicyArgs, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	var resource AutoscalingPolicy
	err := ctx.RegisterResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalingPolicy gets an existing AutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalingPolicyState, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	var resource AutoscalingPolicy
	err := ctx.ReadResource("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscalingPolicy resources.
type autoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm *AutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	Location *string `pulumi:"location"`
	// The "resource name" of the autoscaling policy.
	Name *string `pulumi:"name"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	PolicyId *string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig *AutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig *AutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

type AutoscalingPolicyState struct {
	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm AutoscalingPolicyBasicAlgorithmPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	Location pulumi.StringPtrInput
	// The "resource name" of the autoscaling policy.
	Name pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	PolicyId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig AutoscalingPolicySecondaryWorkerConfigPtrInput
	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig AutoscalingPolicyWorkerConfigPtrInput
}

func (AutoscalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyState)(nil)).Elem()
}

type autoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm *AutoscalingPolicyBasicAlgorithm `pulumi:"basicAlgorithm"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	Location *string `pulumi:"location"`
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	PolicyId string `pulumi:"policyId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig *AutoscalingPolicySecondaryWorkerConfig `pulumi:"secondaryWorkerConfig"`
	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig *AutoscalingPolicyWorkerConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a AutoscalingPolicy resource.
type AutoscalingPolicyArgs struct {
	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm AutoscalingPolicyBasicAlgorithmPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	Location pulumi.StringPtrInput
	// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 50 characters.
	PolicyId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig AutoscalingPolicySecondaryWorkerConfigPtrInput
	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig AutoscalingPolicyWorkerConfigPtrInput
}

func (AutoscalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyArgs)(nil)).Elem()
}

type AutoscalingPolicyInput interface {
	pulumi.Input

	ToAutoscalingPolicyOutput() AutoscalingPolicyOutput
	ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput
}

func (AutoscalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicy)(nil)).Elem()
}

func (i AutoscalingPolicy) ToAutoscalingPolicyOutput() AutoscalingPolicyOutput {
	return i.ToAutoscalingPolicyOutputWithContext(context.Background())
}

func (i AutoscalingPolicy) ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyOutput)
}

type AutoscalingPolicyOutput struct {
	*pulumi.OutputState
}

func (AutoscalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicyOutput)(nil)).Elem()
}

func (o AutoscalingPolicyOutput) ToAutoscalingPolicyOutput() AutoscalingPolicyOutput {
	return o
}

func (o AutoscalingPolicyOutput) ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutoscalingPolicyOutput{})
}
