// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * ConsentStore can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStore:ConsentStore default {{dataset}}/consentStores/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStore:ConsentStore default {{dataset}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStore:ConsentStore default {{name}}
 * ```
 */
export class ConsentStore extends pulumi.CustomResource {
    /**
     * Get an existing ConsentStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConsentStoreState, opts?: pulumi.CustomResourceOptions): ConsentStore {
        return new ConsentStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/consentStore:ConsentStore';

    /**
     * Returns true if the given object is an instance of ConsentStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConsentStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConsentStore.__pulumiType;
    }

    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
     */
    public readonly defaultConsentTtl!: pulumi.Output<string | undefined>;
    /**
     * If true, [consents.patch] [google.cloud.healthcare.v1beta1.consent.UpdateConsent] creates the consent if it does not already exist.
     */
    public readonly enableConsentCreateOnUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * User-supplied key-value pairs used to organize Consent stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of this ConsentStore, for example:
     * "consent1"
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ConsentStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConsentStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConsentStoreArgs | ConsentStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConsentStoreState | undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["defaultConsentTtl"] = state ? state.defaultConsentTtl : undefined;
            inputs["enableConsentCreateOnUpdate"] = state ? state.enableConsentCreateOnUpdate : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ConsentStoreArgs | undefined;
            if ((!args || args.dataset === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataset'");
            }
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["defaultConsentTtl"] = args ? args.defaultConsentTtl : undefined;
            inputs["enableConsentCreateOnUpdate"] = args ? args.enableConsentCreateOnUpdate : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConsentStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConsentStore resources.
 */
export interface ConsentStoreState {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset?: pulumi.Input<string>;
    /**
     * Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
     */
    readonly defaultConsentTtl?: pulumi.Input<string>;
    /**
     * If true, [consents.patch] [google.cloud.healthcare.v1beta1.consent.UpdateConsent] creates the consent if it does not already exist.
     */
    readonly enableConsentCreateOnUpdate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize Consent stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of this ConsentStore, for example:
     * "consent1"
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConsentStore resource.
 */
export interface ConsentStoreArgs {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset: pulumi.Input<string>;
    /**
     * Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
     * A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
     */
    readonly defaultConsentTtl?: pulumi.Input<string>;
    /**
     * If true, [consents.patch] [google.cloud.healthcare.v1beta1.consent.UpdateConsent] creates the consent if it does not already exist.
     */
    readonly enableConsentCreateOnUpdate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize Consent stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of this ConsentStore, for example:
     * "consent1"
     */
    readonly name?: pulumi.Input<string>;
}