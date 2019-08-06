// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for GCE instance. Each of these resources serves a different use case:
 * 
 * * `google.compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `google.compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `google.compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * > **Note:** `google.compute.InstanceIAMPolicy` **cannot** be used in conjunction with `google.compute.InstanceIAMBinding` and `google.compute.InstanceIAMMember` or they will fight over what your policy should be.
 * 
 * > **Note:** `google.compute.InstanceIAMBinding` resources **can be** used in conjunction with `google.compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_compute\_instance\_iam\_policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["user:jane@example.com"],
 *         role: "roles/compute.osLogin",
 *     }],
 * }));
 * const instance = new gcp.compute.InstanceIAMPolicy("instance", {
 *     instanceName: "your-instance-name",
 *     policyData: admin.policyData,
 * });
 * ```
 * 
 * ## google\_compute\_instance\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.compute.InstanceIAMBinding("instance", {
 *     instanceName: "your-instance-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/compute.osLogin",
 * });
 * ```
 * 
 * ## google\_compute\_instance\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const instance = new gcp.compute.InstanceIAMMember("instance", {
 *     instanceName: "your-instance-name",
 *     member: "user:jane@example.com",
 *     role: "roles/compute.osLogin",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_iam_policy.html.markdown.
 */
export class InstanceIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIAMPolicyState, opts?: pulumi.CustomResourceOptions): InstanceIAMPolicy {
        return new InstanceIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceIAMPolicy:InstanceIAMPolicy';

    /**
     * Returns true if the given object is an instance of InstanceIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `googleIamPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIAMPolicyArgs | InstanceIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceIAMPolicyArgs | undefined;
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIAMPolicy resources.
 */
export interface InstanceIAMPolicyState {
    /**
     * (Computed) The etag of the instance's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `googleIamPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIAMPolicy resource.
 */
export interface InstanceIAMPolicyArgs {
    /**
     * The name of the instance.
     */
    readonly instanceName: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `googleIamPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone of the instance. If
     * unspecified, this defaults to the zone configured in the provider.
     */
    readonly zone?: pulumi.Input<string>;
}
