// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} * {{project}}/{{appId}}/{{service}} * {{appId}}/{{service}} * {{service}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class AppEngineServiceIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AppEngineServiceIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppEngineServiceIamPolicyState, opts?: pulumi.CustomResourceOptions): AppEngineServiceIamPolicy {
        return new AppEngineServiceIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy';

    /**
     * Returns true if the given object is an instance of AppEngineServiceIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppEngineServiceIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppEngineServiceIamPolicy.__pulumiType;
    }

    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a AppEngineServiceIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppEngineServiceIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppEngineServiceIamPolicyArgs | AppEngineServiceIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AppEngineServiceIamPolicyState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as AppEngineServiceIamPolicyArgs | undefined;
            if ((!args || args.appId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.policyData === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.service === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'service'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AppEngineServiceIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppEngineServiceIamPolicy resources.
 */
export interface AppEngineServiceIamPolicyState {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppEngineServiceIamPolicy resource.
 */
export interface AppEngineServiceIamPolicyArgs {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     */
    readonly appId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     */
    readonly service: pulumi.Input<string>;
}
