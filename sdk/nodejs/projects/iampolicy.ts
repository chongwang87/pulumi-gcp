// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
 *
 * * `gcp.projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
 * * `gcp.projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
 * * `gcp.projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
 * * `gcp.projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
 *
 * > **Note:** `gcp.projects.IAMPolicy` **cannot** be used in conjunction with `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.projects.IAMBinding` resources **can be** used in conjunction with `gcp.projects.IAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * > **Note:** It is not possible to grant the `roles/owner` role using any of these resources due to this being disallowed by the underlying `projects.setIamPolicy` API method. See the method [documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy) for full details. It is, however, possible to remove all owners from the project by passing in an empty `members = []` list to the `gcp.projects.IAMBinding` resource. This is useful for removing the owner role from a project upon creation, however, precautions should be taken to avoid inadvertently locking oneself out of a project such as by granting additional roles to alternate entities.
 *
 * ## google\_project\_iam\_policy
 *
 * > **Be careful!** You can accidentally lock yourself out of your project
 *    using this resource. Deleting a `gcp.projects.IAMPolicy` removes access
 *    from anyone without organization-level access to the project. Proceed with caution.
 *    It's not recommended to use `gcp.projects.IAMPolicy` with your provider project
 *    to avoid locking yourself out, and it should generally only be used with projects
 *    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
 *    applying the change.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const project = new gcp.projects.IAMPolicy("project", {
 *     project: "your-project-id",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         condition: {
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *             title: "expires_after_2019_12_31",
 *         },
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * }, { async: true }));
 * const project = new gcp.projects.IAMPolicy("project", {
 *     policyData: admin.policyData,
 *     project: "your-project-id",
 * });
 * ```
 *
 * ## google\_project\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.IAMBinding("project", {
 *     members: ["user:jane@example.com"],
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.IAMBinding("project", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     members: ["user:jane@example.com"],
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_project\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.IAMMember("project", {
 *     member: "user:jane@example.com",
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.IAMMember("project", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     member: "user:jane@example.com",
 *     project: "your-project-id",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_project\_iam\_audit\_config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.projects.IAMAuditConfig("project", {
 *     auditLogConfigs: [
 *         {
 *             logType: "ADMIN_READ",
 *         },
 *         {
 *             exemptedMembers: ["user:joebloggs@hashicorp.com"],
 *             logType: "DATA_READ",
 *         },
 *     ],
 *     project: "your-project-id",
 *     service: "allServices",
 * });
 * ```
 *
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `project_id`, role, and member e.g.
 *
 * ```sh
 *  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `project_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `project_id`.
 *
 * ```sh
 *  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project your-project-id
 * ```
 *
 *  IAM audit config imports use the identifier of the resource in question and the service, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id foo.googleapis.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/iAMPolicy:IAMPolicy';

    /**
     * Returns true if the given object is an instance of IAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the project's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
     * Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a IAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMPolicyArgs | IAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    /**
     * (Computed) The etag of the project's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
     * Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the project. The policy will be
     * merged with any existing policy applied to the project.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
     * Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
     * will not be inferred from the provider.
     */
    readonly project: pulumi.Input<string>;
}
