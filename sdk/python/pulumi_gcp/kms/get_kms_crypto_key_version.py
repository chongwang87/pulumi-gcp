# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKMSCryptoKeyVersionResult:
    """
    A collection of values returned by getKMSCryptoKeyVersion.
    """
    def __init__(__self__, algorithm=None, crypto_key=None, id=None, protection_level=None, public_key=None, state=None, version=None):
        if algorithm and not isinstance(algorithm, str):
            raise TypeError("Expected argument 'algorithm' to be a str")
        __self__.algorithm = algorithm
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        if crypto_key and not isinstance(crypto_key, str):
            raise TypeError("Expected argument 'crypto_key' to be a str")
        __self__.crypto_key = crypto_key
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if protection_level and not isinstance(protection_level, str):
            raise TypeError("Expected argument 'protection_level' to be a str")
        __self__.protection_level = protection_level
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.
        """
        if public_key and not isinstance(public_key, dict):
            raise TypeError("Expected argument 'public_key' to be a dict")
        __self__.public_key = public_key
        """
        If the enclosing CryptoKey has purpose `ASYMMETRIC_SIGN` or `ASYMMETRIC_DECRYPT`, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.
        """
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version
class AwaitableGetKMSCryptoKeyVersionResult(GetKMSCryptoKeyVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSCryptoKeyVersionResult(
            algorithm=self.algorithm,
            crypto_key=self.crypto_key,
            id=self.id,
            protection_level=self.protection_level,
            public_key=self.public_key,
            state=self.state,
            version=self.version)

def get_kms_crypto_key_version(crypto_key=None,version=None,opts=None):
    """
    Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).

    A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_crypto_key_version.html.markdown.


    :param str crypto_key: The `self_link` of the Google Cloud Platform CryptoKey to which the key version belongs.
    :param float version: The version number for this CryptoKeyVersion. Defaults to `1`.
    """
    __args__ = dict()


    __args__['cryptoKey'] = crypto_key
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSCryptoKeyVersion:getKMSCryptoKeyVersion', __args__, opts=opts).value

    return AwaitableGetKMSCryptoKeyVersionResult(
        algorithm=__ret__.get('algorithm'),
        crypto_key=__ret__.get('cryptoKey'),
        id=__ret__.get('id'),
        protection_level=__ret__.get('protectionLevel'),
        public_key=__ret__.get('publicKey'),
        state=__ret__.get('state'),
        version=__ret__.get('version'))
