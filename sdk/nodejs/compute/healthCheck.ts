// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Health Checks determine whether instances are responsive and able to do work.
 * They are an important part of a comprehensive load balancing configuration,
 * as they enable monitoring instances behind load balancers.
 * 
 * Health Checks poll instances at a specified interval. Instances that
 * do not respond successfully to some number of probes in a row are marked
 * as unhealthy. No new connections are sent to unhealthy instances,
 * though existing connections will continue. The health check will
 * continue to poll unhealthy instances. If an instance later responds
 * successfully to some number of consecutive probes, it is marked
 * healthy again and can receive new connections.
 * 
 * 
 * To get more information about HealthCheck, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
 * 
 * ## Example Usage - Health Check Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const internal_health_check = new gcp.compute.HealthCheck("internal-health-check", {
 *     checkIntervalSec: 1,
 *     tcpHealthCheck: {
 *         port: 80,
 *     },
 *     timeoutSec: 1,
 * });
 * ```
 */
export class HealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HealthCheckState, opts?: pulumi.CustomResourceOptions): HealthCheck {
        return new HealthCheck(name, <any>state, { ...opts, id: id });
    }

    public readonly checkIntervalSec: pulumi.Output<number | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly healthyThreshold: pulumi.Output<number | undefined>;
    public readonly httpHealthCheck: pulumi.Output<{ host?: string, port?: number, proxyHeader?: string, requestPath?: string, response?: string } | undefined>;
    public readonly httpsHealthCheck: pulumi.Output<{ host?: string, port?: number, proxyHeader?: string, requestPath?: string, response?: string } | undefined>;
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly sslHealthCheck: pulumi.Output<{ port?: number, proxyHeader?: string, request?: string, response?: string } | undefined>;
    public readonly tcpHealthCheck: pulumi.Output<{ port?: number, proxyHeader?: string, request?: string, response?: string } | undefined>;
    public readonly timeoutSec: pulumi.Output<number | undefined>;
    public /*out*/ readonly type: pulumi.Output<string>;
    public readonly unhealthyThreshold: pulumi.Output<number | undefined>;

    /**
     * Create a HealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HealthCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HealthCheckArgs | HealthCheckState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: HealthCheckState = argsOrState as HealthCheckState | undefined;
            inputs["checkIntervalSec"] = state ? state.checkIntervalSec : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            inputs["httpHealthCheck"] = state ? state.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = state ? state.httpsHealthCheck : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslHealthCheck"] = state ? state.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = state ? state.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HealthCheckArgs | undefined;
            inputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            inputs["httpHealthCheck"] = args ? args.httpHealthCheck : undefined;
            inputs["httpsHealthCheck"] = args ? args.httpsHealthCheck : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sslHealthCheck"] = args ? args.sslHealthCheck : undefined;
            inputs["tcpHealthCheck"] = args ? args.tcpHealthCheck : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        super("gcp:compute/healthCheck:HealthCheck", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HealthCheck resources.
 */
export interface HealthCheckState {
    readonly checkIntervalSec?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly healthyThreshold?: pulumi.Input<number>;
    readonly httpHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly httpsHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly sslHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly tcpHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly timeoutSec?: pulumi.Input<number>;
    readonly type?: pulumi.Input<string>;
    readonly unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HealthCheck resource.
 */
export interface HealthCheckArgs {
    readonly checkIntervalSec?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly healthyThreshold?: pulumi.Input<number>;
    readonly httpHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly httpsHealthCheck?: pulumi.Input<{ host?: pulumi.Input<string>, port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, requestPath?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly sslHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly tcpHealthCheck?: pulumi.Input<{ port?: pulumi.Input<number>, proxyHeader?: pulumi.Input<string>, request?: pulumi.Input<string>, response?: pulumi.Input<string> }>;
    readonly timeoutSec?: pulumi.Input<number>;
    readonly unhealthyThreshold?: pulumi.Input<number>;
}
