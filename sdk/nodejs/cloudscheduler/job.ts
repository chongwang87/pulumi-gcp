// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A scheduled job that can publish a pubsub message or a http request
 * every X interval of time, using crontab format string
 * 
 * 
 * To get more information about Job, see:
 * 
 * * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/scheduler/)
 * 
 * ## Example Usage - Scheduler Job Pubsub
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const job = new gcp.cloudscheduler.Job("job", {
 *     description: "test job",
 *     pubsubTarget: {
 *         data: Buffer.from("test").toString("base64"),
 *         topicName: topic.id,
 *     },
 *     schedule: "*&#47;2 * * * *",
 * });
 * ```
 * ## Example Usage - Scheduler Job Http
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const job = new gcp.cloudscheduler.Job("job", {
 *     description: "test http job",
 *     httpTarget: {
 *         httpMethod: "POST",
 *         uri: "https://example.com/ping",
 *     },
 *     schedule: "*&#47;8 * * * *",
 *     timeZone: "America/New_York",
 * });
 * ```
 * ## Example Usage - Scheduler Job App Engine
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const job = new gcp.cloudscheduler.Job("job", {
 *     appEngineHttpTarget: {
 *         appEngineRouting: {
 *             instance: "my-instance-001",
 *             service: "web",
 *             version: "prod",
 *         },
 *         httpMethod: "POST",
 *         relativeUri: "/ping",
 *     },
 *     description: "test app engine job",
 *     schedule: "*&#47;4 * * * *",
 *     timeZone: "Europe/London",
 * });
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:cloudscheduler/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Job.__pulumiType;
    }

    public readonly appEngineHttpTarget!: pulumi.Output<{ appEngineRouting?: { instance?: string, service?: string, version?: string }, body?: string, headers?: {[key: string]: string}, httpMethod?: string, relativeUri: string } | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly httpTarget!: pulumi.Output<{ body?: string, headers?: {[key: string]: string}, httpMethod?: string, uri: string } | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly pubsubTarget!: pulumi.Output<{ attributes?: {[key: string]: string}, data?: string, topicName: string } | undefined>;
    public readonly region!: pulumi.Output<string>;
    public readonly retryConfig!: pulumi.Output<{ maxBackoffDuration?: string, maxDoublings?: number, maxRetryDuration?: string, minBackoffDuration?: string, retryCount?: number } | undefined>;
    public readonly schedule!: pulumi.Output<string | undefined>;
    public readonly timeZone!: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["appEngineHttpTarget"] = state ? state.appEngineHttpTarget : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["httpTarget"] = state ? state.httpTarget : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pubsubTarget"] = state ? state.pubsubTarget : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["retryConfig"] = state ? state.retryConfig : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["timeZone"] = state ? state.timeZone : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            inputs["appEngineHttpTarget"] = args ? args.appEngineHttpTarget : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["httpTarget"] = args ? args.httpTarget : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pubsubTarget"] = args ? args.pubsubTarget : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["retryConfig"] = args ? args.retryConfig : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    readonly appEngineHttpTarget?: pulumi.Input<{ appEngineRouting?: pulumi.Input<{ instance?: pulumi.Input<string>, service?: pulumi.Input<string>, version?: pulumi.Input<string> }>, body?: pulumi.Input<string>, headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, httpMethod?: pulumi.Input<string>, relativeUri: pulumi.Input<string> }>;
    readonly description?: pulumi.Input<string>;
    readonly httpTarget?: pulumi.Input<{ body?: pulumi.Input<string>, headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, httpMethod?: pulumi.Input<string>, uri: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pubsubTarget?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, data?: pulumi.Input<string>, topicName: pulumi.Input<string> }>;
    readonly region?: pulumi.Input<string>;
    readonly retryConfig?: pulumi.Input<{ maxBackoffDuration?: pulumi.Input<string>, maxDoublings?: pulumi.Input<number>, maxRetryDuration?: pulumi.Input<string>, minBackoffDuration?: pulumi.Input<string>, retryCount?: pulumi.Input<number> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly timeZone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    readonly appEngineHttpTarget?: pulumi.Input<{ appEngineRouting?: pulumi.Input<{ instance?: pulumi.Input<string>, service?: pulumi.Input<string>, version?: pulumi.Input<string> }>, body?: pulumi.Input<string>, headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, httpMethod?: pulumi.Input<string>, relativeUri: pulumi.Input<string> }>;
    readonly description?: pulumi.Input<string>;
    readonly httpTarget?: pulumi.Input<{ body?: pulumi.Input<string>, headers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, httpMethod?: pulumi.Input<string>, uri: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pubsubTarget?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, data?: pulumi.Input<string>, topicName: pulumi.Input<string> }>;
    readonly region?: pulumi.Input<string>;
    readonly retryConfig?: pulumi.Input<{ maxBackoffDuration?: pulumi.Input<string>, maxDoublings?: pulumi.Input<number>, maxRetryDuration?: pulumi.Input<string>, minBackoffDuration?: pulumi.Input<string>, retryCount?: pulumi.Input<number> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly timeZone?: pulumi.Input<string>;
}
