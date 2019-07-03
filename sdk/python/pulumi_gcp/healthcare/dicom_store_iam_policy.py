# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DicomStoreIamPolicy(pulumi.CustomResource):
    dicom_store_id: pulumi.Output[str]
    """
    The DICOM store ID, in the form
    `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
    `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
    project setting will be used as a fallback.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the DICOM store's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `google_iam_policy` data source.
    """
    def __init__(__self__, resource_name, opts=None, dicom_store_id=None, policy_data=None, __name__=None, __opts__=None):
        """
        > **Warning:** These resources are in beta, and should be used with the terraform-provider-google-beta provider.
        See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
        
        Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:
        
        * `google_healthcare_dicom_store_iam_policy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
        * `google_healthcare_dicom_store_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
        * `google_healthcare_dicom_store_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.
        
        > **Note:** `google_healthcare_dicom_store_iam_policy` **cannot** be used in conjunction with `google_healthcare_dicom_store_iam_binding` and `google_healthcare_dicom_store_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_healthcare_dicom_store_iam_binding` resources **can be** used in conjunction with `google_healthcare_dicom_store_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `google_iam_policy` data source.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if dicom_store_id is None:
            raise TypeError("Missing required property 'dicom_store_id'")
        __props__['dicom_store_id'] = dicom_store_id

        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data

        __props__['etag'] = None

        super(DicomStoreIamPolicy, __self__).__init__(
            'gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

