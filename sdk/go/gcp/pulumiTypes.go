// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ProviderBatching struct {
	EnableBatching *bool   `pulumi:"enableBatching"`
	SendAfter      *string `pulumi:"sendAfter"`
}

type ProviderBatchingInput interface {
	pulumi.Input

	ToProviderBatchingOutput() ProviderBatchingOutput
	ToProviderBatchingOutputWithContext(context.Context) ProviderBatchingOutput
}

type ProviderBatchingArgs struct {
	EnableBatching pulumi.BoolPtrInput   `pulumi:"enableBatching"`
	SendAfter      pulumi.StringPtrInput `pulumi:"sendAfter"`
}

func (ProviderBatchingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderBatching)(nil)).Elem()
}

func (i ProviderBatchingArgs) ToProviderBatchingOutput() ProviderBatchingOutput {
	return i.ToProviderBatchingOutputWithContext(context.Background())
}

func (i ProviderBatchingArgs) ToProviderBatchingOutputWithContext(ctx context.Context) ProviderBatchingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderBatchingOutput)
}

func (i ProviderBatchingArgs) ToProviderBatchingPtrOutput() ProviderBatchingPtrOutput {
	return i.ToProviderBatchingPtrOutputWithContext(context.Background())
}

func (i ProviderBatchingArgs) ToProviderBatchingPtrOutputWithContext(ctx context.Context) ProviderBatchingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderBatchingOutput).ToProviderBatchingPtrOutputWithContext(ctx)
}

type ProviderBatchingPtrInput interface {
	pulumi.Input

	ToProviderBatchingPtrOutput() ProviderBatchingPtrOutput
	ToProviderBatchingPtrOutputWithContext(context.Context) ProviderBatchingPtrOutput
}

type providerBatchingPtrType ProviderBatchingArgs

func ProviderBatchingPtr(v *ProviderBatchingArgs) ProviderBatchingPtrInput {
	return (*providerBatchingPtrType)(v)
}

func (*providerBatchingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderBatching)(nil)).Elem()
}

func (i *providerBatchingPtrType) ToProviderBatchingPtrOutput() ProviderBatchingPtrOutput {
	return i.ToProviderBatchingPtrOutputWithContext(context.Background())
}

func (i *providerBatchingPtrType) ToProviderBatchingPtrOutputWithContext(ctx context.Context) ProviderBatchingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderBatchingPtrOutput)
}

type ProviderBatchingOutput struct{ *pulumi.OutputState }

func (ProviderBatchingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderBatching)(nil)).Elem()
}

func (o ProviderBatchingOutput) ToProviderBatchingOutput() ProviderBatchingOutput {
	return o
}

func (o ProviderBatchingOutput) ToProviderBatchingOutputWithContext(ctx context.Context) ProviderBatchingOutput {
	return o
}

func (o ProviderBatchingOutput) ToProviderBatchingPtrOutput() ProviderBatchingPtrOutput {
	return o.ToProviderBatchingPtrOutputWithContext(context.Background())
}

func (o ProviderBatchingOutput) ToProviderBatchingPtrOutputWithContext(ctx context.Context) ProviderBatchingPtrOutput {
	return o.ApplyT(func(v ProviderBatching) *ProviderBatching {
		return &v
	}).(ProviderBatchingPtrOutput)
}
func (o ProviderBatchingOutput) EnableBatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderBatching) *bool { return v.EnableBatching }).(pulumi.BoolPtrOutput)
}

func (o ProviderBatchingOutput) SendAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderBatching) *string { return v.SendAfter }).(pulumi.StringPtrOutput)
}

type ProviderBatchingPtrOutput struct{ *pulumi.OutputState }

func (ProviderBatchingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderBatching)(nil)).Elem()
}

func (o ProviderBatchingPtrOutput) ToProviderBatchingPtrOutput() ProviderBatchingPtrOutput {
	return o
}

func (o ProviderBatchingPtrOutput) ToProviderBatchingPtrOutputWithContext(ctx context.Context) ProviderBatchingPtrOutput {
	return o
}

func (o ProviderBatchingPtrOutput) Elem() ProviderBatchingOutput {
	return o.ApplyT(func(v *ProviderBatching) ProviderBatching { return *v }).(ProviderBatchingOutput)
}

func (o ProviderBatchingPtrOutput) EnableBatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderBatching) *bool { return v.EnableBatching }).(pulumi.BoolPtrOutput)
}

func (o ProviderBatchingPtrOutput) SendAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderBatching) *string { return v.SendAfter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderBatchingOutput{})
	pulumi.RegisterOutputType(ProviderBatchingPtrOutput{})
}
