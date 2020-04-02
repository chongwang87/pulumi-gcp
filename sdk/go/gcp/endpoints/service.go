// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package endpoints

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/endpoints_service.html.markdown.
type Service struct {
	pulumi.CustomResourceState

	Apis               ServiceApiArrayOutput      `pulumi:"apis"`
	ConfigId           pulumi.StringOutput        `pulumi:"configId"`
	DnsAddress         pulumi.StringOutput        `pulumi:"dnsAddress"`
	Endpoints          ServiceEndpointArrayOutput `pulumi:"endpoints"`
	GrpcConfig         pulumi.StringPtrOutput     `pulumi:"grpcConfig"`
	OpenapiConfig      pulumi.StringPtrOutput     `pulumi:"openapiConfig"`
	Project            pulumi.StringOutput        `pulumi:"project"`
	ProtocOutputBase64 pulumi.StringPtrOutput     `pulumi:"protocOutputBase64"`
	ServiceName        pulumi.StringOutput        `pulumi:"serviceName"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("gcp:endpoints/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:endpoints/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	Apis               []ServiceApi      `pulumi:"apis"`
	ConfigId           *string           `pulumi:"configId"`
	DnsAddress         *string           `pulumi:"dnsAddress"`
	Endpoints          []ServiceEndpoint `pulumi:"endpoints"`
	GrpcConfig         *string           `pulumi:"grpcConfig"`
	OpenapiConfig      *string           `pulumi:"openapiConfig"`
	Project            *string           `pulumi:"project"`
	ProtocOutputBase64 *string           `pulumi:"protocOutputBase64"`
	ServiceName        *string           `pulumi:"serviceName"`
}

type ServiceState struct {
	Apis               ServiceApiArrayInput
	ConfigId           pulumi.StringPtrInput
	DnsAddress         pulumi.StringPtrInput
	Endpoints          ServiceEndpointArrayInput
	GrpcConfig         pulumi.StringPtrInput
	OpenapiConfig      pulumi.StringPtrInput
	Project            pulumi.StringPtrInput
	ProtocOutputBase64 pulumi.StringPtrInput
	ServiceName        pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	GrpcConfig         *string `pulumi:"grpcConfig"`
	OpenapiConfig      *string `pulumi:"openapiConfig"`
	Project            *string `pulumi:"project"`
	ProtocOutputBase64 *string `pulumi:"protocOutputBase64"`
	ServiceName        string  `pulumi:"serviceName"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	GrpcConfig         pulumi.StringPtrInput
	OpenapiConfig      pulumi.StringPtrInput
	Project            pulumi.StringPtrInput
	ProtocOutputBase64 pulumi.StringPtrInput
	ServiceName        pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
