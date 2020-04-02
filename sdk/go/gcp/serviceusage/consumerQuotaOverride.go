// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package serviceusage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A consumer override is applied to the consumer on its own authority to limit its own quota usage.
// Consumer overrides cannot be used to grant more quota than would be allowed by admin overrides,
// producer overrides, or the default limit of the service.
//
// To get more information about ConsumerQuotaOverride, see:
//
// * How-to Guides
//     * [Getting Started](https://cloud.google.com/service-usage/docs/getting-started)
//     * [REST API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1beta1/services.consumerQuotaMetrics.limits.consumerOverrides)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_usage_consumer_quota_override.html.markdown.
type ConsumerQuotaOverride struct {
	pulumi.CustomResourceState

	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapOutput `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'.
	Limit pulumi.StringOutput `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The server-generated name of the quota override.
	Name pulumi.StringOutput `pulumi:"name"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringOutput `pulumi:"overrideValue"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewConsumerQuotaOverride registers a new resource with the given unique name, arguments, and options.
func NewConsumerQuotaOverride(ctx *pulumi.Context,
	name string, args *ConsumerQuotaOverrideArgs, opts ...pulumi.ResourceOption) (*ConsumerQuotaOverride, error) {
	if args == nil || args.Limit == nil {
		return nil, errors.New("missing required argument 'Limit'")
	}
	if args == nil || args.Metric == nil {
		return nil, errors.New("missing required argument 'Metric'")
	}
	if args == nil || args.OverrideValue == nil {
		return nil, errors.New("missing required argument 'OverrideValue'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &ConsumerQuotaOverrideArgs{}
	}
	var resource ConsumerQuotaOverride
	err := ctx.RegisterResource("gcp:serviceusage/consumerQuotaOverride:ConsumerQuotaOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerQuotaOverride gets an existing ConsumerQuotaOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerQuotaOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerQuotaOverrideState, opts ...pulumi.ResourceOption) (*ConsumerQuotaOverride, error) {
	var resource ConsumerQuotaOverride
	err := ctx.ReadResource("gcp:serviceusage/consumerQuotaOverride:ConsumerQuotaOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerQuotaOverride resources.
type consumerQuotaOverrideState struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force *bool `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'.
	Limit *string `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric *string `pulumi:"metric"`
	// The server-generated name of the quota override.
	Name *string `pulumi:"name"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue *string `pulumi:"overrideValue"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service *string `pulumi:"service"`
}

type ConsumerQuotaOverrideState struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapInput
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrInput
	// The limit on the metric, e.g. '/project/region'.
	Limit pulumi.StringPtrInput
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringPtrInput
	// The server-generated name of the quota override.
	Name pulumi.StringPtrInput
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service pulumi.StringPtrInput
}

func (ConsumerQuotaOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerQuotaOverrideState)(nil)).Elem()
}

type consumerQuotaOverrideArgs struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions map[string]string `pulumi:"dimensions"`
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force *bool `pulumi:"force"`
	// The limit on the metric, e.g. '/project/region'.
	Limit string `pulumi:"limit"`
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric string `pulumi:"metric"`
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue string `pulumi:"overrideValue"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a ConsumerQuotaOverride resource.
type ConsumerQuotaOverrideArgs struct {
	// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
	Dimensions pulumi.StringMapInput
	// If the new quota would decrease the existing quota by more than 10%, the request is rejected. If 'force' is 'true', that
	// safety check is ignored.
	Force pulumi.BoolPtrInput
	// The limit on the metric, e.g. '/project/region'.
	Limit pulumi.StringInput
	// The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.
	Metric pulumi.StringInput
	// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
	OverrideValue pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service that the metrics belong to, e.g. 'compute.googleapis.com'.
	Service pulumi.StringInput
}

func (ConsumerQuotaOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerQuotaOverrideArgs)(nil)).Elem()
}
