// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class OrganizationSecurityPolicyAssociation extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSecurityPolicyAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSecurityPolicyAssociationState, opts?: pulumi.CustomResourceOptions): OrganizationSecurityPolicyAssociation {
        return new OrganizationSecurityPolicyAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation';

    /**
     * Returns true if the given object is an instance of OrganizationSecurityPolicyAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSecurityPolicyAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSecurityPolicyAssociation.__pulumiType;
    }

    /**
     * The resource that the security policy is attached to.
     */
    public readonly attachmentId!: pulumi.Output<string>;
    /**
     * The display name of the security policy of the association.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * The name for an association.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The security policy ID of the association.
     */
    public readonly policyId!: pulumi.Output<string>;

    /**
     * Create a OrganizationSecurityPolicyAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSecurityPolicyAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSecurityPolicyAssociationArgs | OrganizationSecurityPolicyAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationSecurityPolicyAssociationState | undefined;
            inputs["attachmentId"] = state ? state.attachmentId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
        } else {
            const args = argsOrState as OrganizationSecurityPolicyAssociationArgs | undefined;
            if (!args || args.attachmentId === undefined) {
                throw new Error("Missing required property 'attachmentId'");
            }
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            inputs["attachmentId"] = args ? args.attachmentId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["displayName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrganizationSecurityPolicyAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSecurityPolicyAssociation resources.
 */
export interface OrganizationSecurityPolicyAssociationState {
    /**
     * The resource that the security policy is attached to.
     */
    readonly attachmentId?: pulumi.Input<string>;
    /**
     * The display name of the security policy of the association.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name for an association.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The security policy ID of the association.
     */
    readonly policyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationSecurityPolicyAssociation resource.
 */
export interface OrganizationSecurityPolicyAssociationArgs {
    /**
     * The resource that the security policy is attached to.
     */
    readonly attachmentId: pulumi.Input<string>;
    /**
     * The name for an association.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The security policy ID of the association.
     */
    readonly policyId: pulumi.Input<string>;
}