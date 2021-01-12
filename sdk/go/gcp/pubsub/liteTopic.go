// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A named resource to which messages are sent by publishers.
//
// To get more information about Topic, see:
//
// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
// * How-to Guides
//     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
//
// ## Example Usage
// ### Pubsub Lite Topic Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewLiteTopic(ctx, "example", &pubsub.LiteTopicArgs{
// 			Project: pulumi.String(project.Number),
// 			PartitionConfig: &pubsub.LiteTopicPartitionConfigArgs{
// 				Count: pulumi.Int(1),
// 				Capacity: &pubsub.LiteTopicPartitionConfigCapacityArgs{
// 					PublishMibPerSec:   pulumi.Int(4),
// 					SubscribeMibPerSec: pulumi.Int(8),
// 				},
// 			},
// 			RetentionConfig: &pubsub.LiteTopicRetentionConfigArgs{
// 				PerPartitionBytes: pulumi.String("32212254720"),
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
// Topic can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:pubsub/liteTopic:LiteTopic default projects/{{project}}/locations/{{zone}}/topics/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{name}}
// ```
type LiteTopic struct {
	pulumi.CustomResourceState

	// Name of the topic.
	Name pulumi.StringOutput `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrOutput `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrOutput `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewLiteTopic registers a new resource with the given unique name, arguments, and options.
func NewLiteTopic(ctx *pulumi.Context,
	name string, args *LiteTopicArgs, opts ...pulumi.ResourceOption) (*LiteTopic, error) {
	if args == nil {
		args = &LiteTopicArgs{}
	}

	var resource LiteTopic
	err := ctx.RegisterResource("gcp:pubsub/liteTopic:LiteTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiteTopic gets an existing LiteTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiteTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiteTopicState, opts ...pulumi.ResourceOption) (*LiteTopic, error) {
	var resource LiteTopic
	err := ctx.ReadResource("gcp:pubsub/liteTopic:LiteTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiteTopic resources.
type liteTopicState struct {
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig *LiteTopicPartitionConfig `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig *LiteTopicRetentionConfig `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

type LiteTopicState struct {
	// Name of the topic.
	Name pulumi.StringPtrInput
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*liteTopicState)(nil)).Elem()
}

type liteTopicArgs struct {
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig *LiteTopicPartitionConfig `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig *LiteTopicRetentionConfig `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LiteTopic resource.
type LiteTopicArgs struct {
	// Name of the topic.
	Name pulumi.StringPtrInput
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liteTopicArgs)(nil)).Elem()
}

type LiteTopicInput interface {
	pulumi.Input

	ToLiteTopicOutput() LiteTopicOutput
	ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput
}

func (LiteTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteTopic)(nil)).Elem()
}

func (i LiteTopic) ToLiteTopicOutput() LiteTopicOutput {
	return i.ToLiteTopicOutputWithContext(context.Background())
}

func (i LiteTopic) ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteTopicOutput)
}

type LiteTopicOutput struct {
	*pulumi.OutputState
}

func (LiteTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteTopicOutput)(nil)).Elem()
}

func (o LiteTopicOutput) ToLiteTopicOutput() LiteTopicOutput {
	return o
}

func (o LiteTopicOutput) ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LiteTopicOutput{})
}
