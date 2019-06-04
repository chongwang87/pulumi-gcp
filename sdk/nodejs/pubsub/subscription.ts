// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A named resource representing the stream of messages from a single,
 * specific topic, to be delivered to the subscribing application.
 * 
 * 
 * To get more information about Subscription, see:
 * 
 * * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
 * * How-to Guides
 *     * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)
 * 
 * ## Example Usage - Pubsub Subscription Push
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const exampleTopic = new gcp.pubsub.Topic("example", {});
 * const exampleSubscription = new gcp.pubsub.Subscription("example", {
 *     ackDeadlineSeconds: 20,
 *     labels: {
 *         foo: "bar",
 *     },
 *     pushConfig: {
 *         attributes: {
 *             "x-goog-version": "v1",
 *         },
 *         pushEndpoint: "https://example.com/push",
 *     },
 *     topic: exampleTopic.name,
 * });
 * ```
 * ## Example Usage - Pubsub Subscription Pull
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const exampleTopic = new gcp.pubsub.Topic("example", {});
 * const exampleSubscription = new gcp.pubsub.Subscription("example", {
 *     ackDeadlineSeconds: 20,
 *     expirationPolicy: {
 *         ttl: "300000.5s",
 *     },
 *     labels: {
 *         foo: "bar",
 *     },
 *     // 20 minutes
 *     messageRetentionDuration: "1200s",
 *     retainAckedMessages: true,
 *     topic: exampleTopic.name,
 * });
 * ```
 * ## Example Usage - Pubsub Subscription Different Project
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const exampleTopic = new gcp.pubsub.Topic("example", {
 *     project: "topic-project",
 * });
 * const exampleSubscription = new gcp.pubsub.Subscription("example", {
 *     project: "subscription-project",
 *     topic: exampleTopic.id,
 * });
 * ```
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:pubsub/subscription:Subscription';

    /**
     * Returns true if the given object is an instance of Subscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subscription {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Subscription.__pulumiType;
    }

    public readonly ackDeadlineSeconds!: pulumi.Output<number>;
    public readonly expirationPolicy!: pulumi.Output<{ ttl?: string }>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly messageRetentionDuration!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly pushConfig!: pulumi.Output<{ attributes?: {[key: string]: string}, pushEndpoint: string } | undefined>;
    public readonly retainAckedMessages!: pulumi.Output<boolean | undefined>;
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubscriptionState | undefined;
            inputs["ackDeadlineSeconds"] = state ? state.ackDeadlineSeconds : undefined;
            inputs["expirationPolicy"] = state ? state.expirationPolicy : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["messageRetentionDuration"] = state ? state.messageRetentionDuration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pushConfig"] = state ? state.pushConfig : undefined;
            inputs["retainAckedMessages"] = state ? state.retainAckedMessages : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if (!args || args.topic === undefined) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["ackDeadlineSeconds"] = args ? args.ackDeadlineSeconds : undefined;
            inputs["expirationPolicy"] = args ? args.expirationPolicy : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["messageRetentionDuration"] = args ? args.messageRetentionDuration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pushConfig"] = args ? args.pushConfig : undefined;
            inputs["retainAckedMessages"] = args ? args.retainAckedMessages : undefined;
            inputs["topic"] = args ? args.topic : undefined;
            inputs["path"] = undefined /*out*/;
        }
        super(Subscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    readonly ackDeadlineSeconds?: pulumi.Input<number>;
    readonly expirationPolicy?: pulumi.Input<{ ttl?: pulumi.Input<string> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly messageRetentionDuration?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pushConfig?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pushEndpoint: pulumi.Input<string> }>;
    readonly retainAckedMessages?: pulumi.Input<boolean>;
    readonly topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    readonly ackDeadlineSeconds?: pulumi.Input<number>;
    readonly expirationPolicy?: pulumi.Input<{ ttl?: pulumi.Input<string> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly messageRetentionDuration?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pushConfig?: pulumi.Input<{ attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pushEndpoint: pulumi.Input<string> }>;
    readonly retainAckedMessages?: pulumi.Input<boolean>;
    readonly topic: pulumi.Input<string>;
}
