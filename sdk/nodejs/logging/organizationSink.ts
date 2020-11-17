// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a organization-level logging sink. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/) and
 * [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
 *
 * Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
 * granted to the credentials used with this provider.
 */
export class OrganizationSink extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSinkState, opts?: pulumi.CustomResourceOptions): OrganizationSink {
        return new OrganizationSink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/organizationSink:OrganizationSink';

    /**
     * Returns true if the given object is an instance of OrganizationSink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSink.__pulumiType;
    }

    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    public readonly bigqueryOptions!: pulumi.Output<outputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
     * one of exclusion_filters it will not be exported.
     */
    public readonly exclusions!: pulumi.Output<outputs.logging.OrganizationSinkExclusion[] | undefined>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    public readonly includeChildren!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the logging sink.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity!: pulumi.Output<string>;

    /**
     * Create a OrganizationSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSinkArgs | OrganizationSinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationSinkState | undefined;
            inputs["bigqueryOptions"] = state ? state.bigqueryOptions : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["exclusions"] = state ? state.exclusions : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["includeChildren"] = state ? state.includeChildren : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as OrganizationSinkArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            inputs["bigqueryOptions"] = args ? args.bigqueryOptions : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["exclusions"] = args ? args.exclusions : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["includeChildren"] = args ? args.includeChildren : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrganizationSink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSink resources.
 */
export interface OrganizationSinkState {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    readonly bigqueryOptions?: pulumi.Input<inputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
     * one of exclusion_filters it will not be exported.
     */
    readonly exclusions?: pulumi.Input<pulumi.Input<inputs.logging.OrganizationSinkExclusion>[]>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    readonly includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    readonly writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationSink resource.
 */
export interface OrganizationSinkArgs {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    readonly bigqueryOptions?: pulumi.Input<inputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and
     * one of exclusion_filters it will not be exported.
     */
    readonly exclusions?: pulumi.Input<pulumi.Input<inputs.logging.OrganizationSinkExclusion>[]>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    readonly includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    readonly orgId: pulumi.Input<string>;
}
