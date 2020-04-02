// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
 * [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
 * and
 * [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).
 * 
 * A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_crypto_key_version.html.markdown.
 */
export function getKMSCryptoKeyVersion(args: GetKMSCryptoKeyVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetKMSCryptoKeyVersionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:kms/getKMSCryptoKeyVersion:getKMSCryptoKeyVersion", {
        "cryptoKey": args.cryptoKey,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getKMSCryptoKeyVersion.
 */
export interface GetKMSCryptoKeyVersionArgs {
    /**
     * The `selfLink` of the Google Cloud Platform CryptoKey to which the key version belongs.
     */
    readonly cryptoKey: string;
    /**
     * The version number for this CryptoKeyVersion. Defaults to `1`.
     */
    readonly version?: number;
}

/**
 * A collection of values returned by getKMSCryptoKeyVersion.
 */
export interface GetKMSCryptoKeyVersionResult {
    /**
     * The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
     */
    readonly algorithm: string;
    readonly cryptoKey: string;
    /**
     * The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protectionLevel reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.
     */
    readonly protectionLevel: string;
    /**
     * If the enclosing CryptoKey has purpose `ASYMMETRIC_SIGN` or `ASYMMETRIC_DECRYPT`, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below.
     */
    readonly publicKey: outputs.kms.GetKMSCryptoKeyVersionPublicKey;
    /**
     * The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.
     */
    readonly state: string;
    readonly version?: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
