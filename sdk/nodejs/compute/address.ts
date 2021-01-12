// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents an Address resource.
 *
 * Each virtual machine instance has an ephemeral internal IP address and,
 * optionally, an external IP address. To communicate between instances on
 * the same network, you can use an instance's internal IP address. To
 * communicate with the Internet and instances outside of the same network,
 * you must specify the instance's external IP address.
 *
 * Internal IP addresses are ephemeral and only belong to an instance for
 * the lifetime of the instance; if the instance is deleted and recreated,
 * the instance is assigned a new internal IP address, either by Compute
 * Engine or by you. External IP addresses can be either ephemeral or
 * static.
 *
 * To get more information about Address, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
 *     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
 *
 * ## Example Usage
 * ### Address Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const ipAddress = new gcp.compute.Address("ip_address", {});
 * ```
 * ### Address With Subnetwork
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const internalWithSubnetAndAddress = new gcp.compute.Address("internalWithSubnetAndAddress", {
 *     subnetwork: defaultSubnetwork.id,
 *     addressType: "INTERNAL",
 *     address: "10.0.42.42",
 *     region: "us-central1",
 * });
 * ```
 * ### Address With Gce Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const internalWithGceEndpoint = new gcp.compute.Address("internal_with_gce_endpoint", {
 *     addressType: "INTERNAL",
 *     purpose: "GCE_ENDPOINT",
 * });
 * ```
 * ### Instance With Ip
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const static = new gcp.compute.Address("static", {});
 * const debianImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const instanceWithIp = new gcp.compute.Instance("instanceWithIp", {
 *     machineType: "f1-micro",
 *     zone: "us-central1-a",
 *     bootDisk: {
 *         initializeParams: {
 *             image: debianImage.then(debianImage => debianImage.selfLink),
 *         },
 *     },
 *     networkInterfaces: [{
 *         network: "default",
 *         accessConfigs: [{
 *             natIp: static.address,
 *         }],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Address can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default projects/{{project}}/regions/{{region}}/addresses/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{name}}
 * ```
 */
export class Address extends pulumi.CustomResource {
    /**
     * Get an existing Address resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressState, opts?: pulumi.CustomResourceOptions): Address {
        return new Address(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/address:Address';

    /**
     * Returns true if the given object is an instance of Address.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Address {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Address.__pulumiType;
    }

    /**
     * The static external IP address represented by this resource. Only
     * IPv4 is supported. An address may only be specified for INTERNAL
     * address types. The IP address must be inside the specified subnetwork,
     * if any.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The type of address to reserve.
     * Default value is `EXTERNAL`.
     * Possible values are `INTERNAL` and `EXTERNAL`.
     */
    public readonly addressType!: pulumi.Output<string | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     * Possible values are `PREMIUM` and `STANDARD`.
     */
    public readonly networkTier!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The purpose of this resource, which can be one of the following values:
     * * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
     * * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
     * * VPC_PEERING for addresses that are reserved for VPC peer networks.
     * This should only be set when using an Internal address.
     * Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The URL of the subnetwork in which to reserve the address. If an IP
     * address is specified, it must be within the subnetwork's IP range.
     * This field can only be used with INTERNAL type with
     * GCE_ENDPOINT/DNS_RESOLVER purposes.
     */
    public readonly subnetwork!: pulumi.Output<string>;
    /**
     * The URLs of the resources that are using this address.
     */
    public /*out*/ readonly users!: pulumi.Output<string[]>;

    /**
     * Create a Address resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressArgs | AddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AddressState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["addressType"] = state ? state.addressType : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkTier"] = state ? state.networkTier : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["purpose"] = state ? state.purpose : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as AddressArgs | undefined;
            inputs["address"] = args ? args.address : undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkTier"] = args ? args.networkTier : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Address.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Address resources.
 */
export interface AddressState {
    /**
     * The static external IP address represented by this resource. Only
     * IPv4 is supported. An address may only be specified for INTERNAL
     * address types. The IP address must be inside the specified subnetwork,
     * if any.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The type of address to reserve.
     * Default value is `EXTERNAL`.
     * Possible values are `INTERNAL` and `EXTERNAL`.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     * Possible values are `PREMIUM` and `STANDARD`.
     */
    readonly networkTier?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The purpose of this resource, which can be one of the following values:
     * * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
     * * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
     * * VPC_PEERING for addresses that are reserved for VPC peer networks.
     * This should only be set when using an Internal address.
     * Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
     */
    readonly purpose?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The URL of the subnetwork in which to reserve the address. If an IP
     * address is specified, it must be within the subnetwork's IP range.
     * This field can only be used with INTERNAL type with
     * GCE_ENDPOINT/DNS_RESOLVER purposes.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * The URLs of the resources that are using this address.
     */
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Address resource.
 */
export interface AddressArgs {
    /**
     * The static external IP address represented by this resource. Only
     * IPv4 is supported. An address may only be specified for INTERNAL
     * address types. The IP address must be inside the specified subnetwork,
     * if any.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The type of address to reserve.
     * Default value is `EXTERNAL`.
     * Possible values are `INTERNAL` and `EXTERNAL`.
     */
    readonly addressType?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     * Possible values are `PREMIUM` and `STANDARD`.
     */
    readonly networkTier?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The purpose of this resource, which can be one of the following values:
     * * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
     * * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
     * * VPC_PEERING for addresses that are reserved for VPC peer networks.
     * This should only be set when using an Internal address.
     * Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
     */
    readonly purpose?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URL of the subnetwork in which to reserve the address. If an IP
     * address is specified, it must be within the subnetwork's IP range.
     * This field can only be used with INTERNAL type with
     * GCE_ENDPOINT/DNS_RESOLVER purposes.
     */
    readonly subnetwork?: pulumi.Input<string>;
}
