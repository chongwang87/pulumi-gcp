// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Binary Authorization Attestor. Each of these resources serves a different use case:
 *
 * * `gcp.binaryauthorization.AttestorIamPolicy`: Authoritative. Sets the IAM policy for the attestor and replaces any existing policy already attached.
 * * `gcp.binaryauthorization.AttestorIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the attestor are preserved.
 * * `gcp.binaryauthorization.AttestorIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the attestor are preserved.
 *
 * > **Note:** `gcp.binaryauthorization.AttestorIamPolicy` **cannot** be used in conjunction with `gcp.binaryauthorization.AttestorIamBinding` and `gcp.binaryauthorization.AttestorIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.binaryauthorization.AttestorIamBinding` resources **can be** used in conjunction with `gcp.binaryauthorization.AttestorIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_binary\_authorization\_attestor\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.binaryauthorization.AttestorIamPolicy("policy", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_binary\_authorization\_attestor\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.binaryauthorization.AttestorIamBinding("binding", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_binary\_authorization\_attestor\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.binaryauthorization.AttestorIamMember("member", {
 *     project: google_binary_authorization_attestor.attestor.project,
 *     attestor: google_binary_authorization_attestor.attestor.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/attestors/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Binary Authorization attestor IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding editor "projects/{{project}}/attestors/{{attestor}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding editor "projects/{{project}}/attestors/{{attestor}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding editor projects/{{project}}/attestors/{{attestor}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class AttestorIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing AttestorIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttestorIamBindingState, opts?: pulumi.CustomResourceOptions): AttestorIamBinding {
        return new AttestorIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding';

    /**
     * Returns true if the given object is an instance of AttestorIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttestorIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttestorIamBinding.__pulumiType;
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly attestor!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.binaryauthorization.AttestorIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a AttestorIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttestorIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttestorIamBindingArgs | AttestorIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttestorIamBindingState | undefined;
            inputs["attestor"] = state ? state.attestor : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as AttestorIamBindingArgs | undefined;
            if ((!args || args.attestor === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'attestor'");
            }
            if ((!args || args.members === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'role'");
            }
            inputs["attestor"] = args ? args.attestor : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AttestorIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttestorIamBinding resources.
 */
export interface AttestorIamBindingState {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly attestor?: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.binaryauthorization.AttestorIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttestorIamBinding resource.
 */
export interface AttestorIamBindingArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly attestor: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.binaryauthorization.AttestorIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
