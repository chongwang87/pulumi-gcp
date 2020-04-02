# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKMSSecretResult:
    """
    A collection of values returned by getKMSSecret.
    """
    def __init__(__self__, additional_authenticated_data=None, ciphertext=None, crypto_key=None, id=None, plaintext=None):
        if additional_authenticated_data and not isinstance(additional_authenticated_data, str):
            raise TypeError("Expected argument 'additional_authenticated_data' to be a str")
        __self__.additional_authenticated_data = additional_authenticated_data
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        __self__.ciphertext = ciphertext
        if crypto_key and not isinstance(crypto_key, str):
            raise TypeError("Expected argument 'crypto_key' to be a str")
        __self__.crypto_key = crypto_key
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        __self__.plaintext = plaintext
        """
        Contains the result of decrypting the provided ciphertext.
        """
class AwaitableGetKMSSecretResult(GetKMSSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSSecretResult(
            additional_authenticated_data=self.additional_authenticated_data,
            ciphertext=self.ciphertext,
            crypto_key=self.crypto_key,
            id=self.id,
            plaintext=self.plaintext)

def get_kms_secret(additional_authenticated_data=None,ciphertext=None,crypto_key=None,opts=None):
    """
    This data source allows you to use data encrypted with Google Cloud KMS
    within your resource definitions.

    For more information see
    [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).

    > **NOTE**: Using this data provider will allow you to conceal secret data within your
    resource definitions, but it does not take care of protecting that data in the
    logging output, plan output, or state output.  Please take care to secure your secret
    data outside of resource definitions.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_secret.html.markdown.


    :param str additional_authenticated_data: The [additional authenticated data](https://cloud.google.com/kms/docs/additional-authenticated-data) used for integrity checks during encryption and decryption.
    :param str ciphertext: The ciphertext to be decrypted, encoded in base64
    :param str crypto_key: The id of the CryptoKey that will be used to
           decrypt the provided ciphertext. This is represented by the format
           `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
    """
    __args__ = dict()


    __args__['additionalAuthenticatedData'] = additional_authenticated_data
    __args__['ciphertext'] = ciphertext
    __args__['cryptoKey'] = crypto_key
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSSecret:getKMSSecret', __args__, opts=opts).value

    return AwaitableGetKMSSecretResult(
        additional_authenticated_data=__ret__.get('additionalAuthenticatedData'),
        ciphertext=__ret__.get('ciphertext'),
        crypto_key=__ret__.get('cryptoKey'),
        id=__ret__.get('id'),
        plaintext=__ret__.get('plaintext'))
