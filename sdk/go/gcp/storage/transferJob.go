// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Transfer Job in Google Cloud Storage Transfer.
//
// To get more information about Google Cloud Storage Transfer, see:
//
// * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
// * How-to Guides
//     * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_transfer_job.html.markdown.
type TransferJob struct {
	pulumi.CustomResourceState

	// When the Transfer Job was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringOutput `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description pulumi.StringOutput `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringOutput `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule TransferJobScheduleOutput `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecOutput `pulumi:"transferSpec"`
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Schedule == nil {
		return nil, errors.New("missing required argument 'Schedule'")
	}
	if args == nil || args.TransferSpec == nil {
		return nil, errors.New("missing required argument 'TransferSpec'")
	}
	if args == nil {
		args = &TransferJobArgs{}
	}
	var resource TransferJob
	err := ctx.RegisterResource("gcp:storage/transferJob:TransferJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferJobState, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	var resource TransferJob
	err := ctx.ReadResource("gcp:storage/transferJob:TransferJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferJob resources.
type transferJobState struct {
	// When the Transfer Job was created.
	CreationTime *string `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime *string `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description *string `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime *string `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule *TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec *TransferJobTransferSpec `pulumi:"transferSpec"`
}

type TransferJobState struct {
	// When the Transfer Job was created.
	CreationTime pulumi.StringPtrInput
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringPtrInput
	// Unique description to identify the Transfer Job.
	Description pulumi.StringPtrInput
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringPtrInput
	// The name of the Transfer Job.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule TransferJobSchedulePtrInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecPtrInput
}

func (TransferJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobState)(nil)).Elem()
}

type transferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description string `pulumi:"description"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpec `pulumi:"transferSpec"`
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
	Schedule TransferJobScheduleInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecInput
}

func (TransferJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobArgs)(nil)).Elem()
}
