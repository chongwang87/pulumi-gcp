# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'AccessApprovalSettingsEnrolledService',
    'IAMBindingCondition',
    'IAMMemberCondition',
    'IamAuditConfigAuditLogConfig',
    'PolicyBooleanPolicy',
    'PolicyListPolicy',
    'PolicyListPolicyAllow',
    'PolicyListPolicyDeny',
    'PolicyRestorePolicy',
    'GetIAMPolicyAuditConfigResult',
    'GetIAMPolicyAuditConfigAuditLogConfigResult',
    'GetIAMPolicyBindingResult',
    'GetIAMPolicyBindingConditionResult',
]

@pulumi.output_type
class AccessApprovalSettingsEnrolledService(dict):
    def __init__(__self__, *,
                 cloud_product: str,
                 enrollment_level: Optional[str] = None):
        pulumi.set(__self__, "cloud_product", cloud_product)
        if enrollment_level is not None:
            pulumi.set(__self__, "enrollment_level", enrollment_level)

    @property
    @pulumi.getter(name="cloudProduct")
    def cloud_product(self) -> str:
        return pulumi.get(self, "cloud_product")

    @property
    @pulumi.getter(name="enrollmentLevel")
    def enrollment_level(self) -> Optional[str]:
        return pulumi.get(self, "enrollment_level")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IAMBindingCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IAMMemberCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamAuditConfigAuditLogConfig(dict):
    def __init__(__self__, *,
                 log_type: str,
                 exempted_members: Optional[Sequence[str]] = None):
        """
        :param str log_type: Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        :param Sequence[str] exempted_members: Identities that do not cause logging for this type of permission.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        pulumi.set(__self__, "log_type", log_type)
        if exempted_members is not None:
            pulumi.set(__self__, "exempted_members", exempted_members)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        """
        Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        """
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[Sequence[str]]:
        """
        Identities that do not cause logging for this type of permission.
        Each entry can have one of the following values:
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "exempted_members")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyBooleanPolicy(dict):
    def __init__(__self__, *,
                 enforced: bool):
        """
        :param bool enforced: If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        pulumi.set(__self__, "enforced", enforced)

    @property
    @pulumi.getter
    def enforced(self) -> bool:
        """
        If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        return pulumi.get(self, "enforced")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicy(dict):
    def __init__(__self__, *,
                 allow: Optional['outputs.PolicyListPolicyAllow'] = None,
                 deny: Optional['outputs.PolicyListPolicyDeny'] = None,
                 inherit_from_parent: Optional[bool] = None,
                 suggested_value: Optional[str] = None):
        """
        :param 'PolicyListPolicyAllowArgs' allow: or `deny` - (Optional) One or the other must be set.
        :param bool inherit_from_parent: If set to true, the values from the effective Policy of the parent resource
               are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        :param str suggested_value: The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        if allow is not None:
            pulumi.set(__self__, "allow", allow)
        if deny is not None:
            pulumi.set(__self__, "deny", deny)
        if inherit_from_parent is not None:
            pulumi.set(__self__, "inherit_from_parent", inherit_from_parent)
        if suggested_value is not None:
            pulumi.set(__self__, "suggested_value", suggested_value)

    @property
    @pulumi.getter
    def allow(self) -> Optional['outputs.PolicyListPolicyAllow']:
        """
        or `deny` - (Optional) One or the other must be set.
        """
        return pulumi.get(self, "allow")

    @property
    @pulumi.getter
    def deny(self) -> Optional['outputs.PolicyListPolicyDeny']:
        return pulumi.get(self, "deny")

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> Optional[bool]:
        """
        If set to true, the values from the effective Policy of the parent resource
        are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        """
        return pulumi.get(self, "inherit_from_parent")

    @property
    @pulumi.getter(name="suggestedValue")
    def suggested_value(self) -> Optional[str]:
        """
        The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        return pulumi.get(self, "suggested_value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicyAllow(dict):
    def __init__(__self__, *,
                 all: Optional[bool] = None,
                 values: Optional[Sequence[str]] = None):
        """
        :param bool all: The policy allows or denies all values.
        :param Sequence[str] values: The policy can define specific values that are allowed or denied.
        """
        if all is not None:
            pulumi.set(__self__, "all", all)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> Optional[bool]:
        """
        The policy allows or denies all values.
        """
        return pulumi.get(self, "all")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicyDeny(dict):
    def __init__(__self__, *,
                 all: Optional[bool] = None,
                 values: Optional[Sequence[str]] = None):
        """
        :param bool all: The policy allows or denies all values.
        :param Sequence[str] values: The policy can define specific values that are allowed or denied.
        """
        if all is not None:
            pulumi.set(__self__, "all", all)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> Optional[bool]:
        """
        The policy allows or denies all values.
        """
        return pulumi.get(self, "all")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyRestorePolicy(dict):
    def __init__(__self__, *,
                 default: bool):
        """
        :param bool default: May only be set to true. If set, then the default Policy is restored.
        """
        pulumi.set(__self__, "default", default)

    @property
    @pulumi.getter
    def default(self) -> bool:
        """
        May only be set to true. If set, then the default Policy is restored.
        """
        return pulumi.get(self, "default")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetIAMPolicyAuditConfigResult(dict):
    def __init__(__self__, *,
                 audit_log_configs: Sequence['outputs.GetIAMPolicyAuditConfigAuditLogConfigResult'],
                 service: str):
        """
        :param Sequence['GetIAMPolicyAuditConfigAuditLogConfigArgs'] audit_log_configs: A nested block that defines the operations you'd like to log.
        :param str service: Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        pulumi.set(__self__, "audit_log_configs", audit_log_configs)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="auditLogConfigs")
    def audit_log_configs(self) -> Sequence['outputs.GetIAMPolicyAuditConfigAuditLogConfigResult']:
        """
        A nested block that defines the operations you'd like to log.
        """
        return pulumi.get(self, "audit_log_configs")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        return pulumi.get(self, "service")


@pulumi.output_type
class GetIAMPolicyAuditConfigAuditLogConfigResult(dict):
    def __init__(__self__, *,
                 log_type: str,
                 exempted_members: Optional[Sequence[str]] = None):
        """
        :param str log_type: Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
        :param Sequence[str] exempted_members: Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        """
        pulumi.set(__self__, "log_type", log_type)
        if exempted_members is not None:
            pulumi.set(__self__, "exempted_members", exempted_members)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        """
        Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
        """
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[Sequence[str]]:
        """
        Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        """
        return pulumi.get(self, "exempted_members")


@pulumi.output_type
class GetIAMPolicyBindingResult(dict):
    def __init__(__self__, *,
                 members: Sequence[str],
                 role: str,
                 condition: Optional['outputs.GetIAMPolicyBindingConditionResult'] = None):
        """
        :param Sequence[str] members: An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. Some resources **don't** support this identity.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. Some resources **don't** support this identity.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param str role: The role/permission that will be granted to the members.
               See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
               Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param 'GetIAMPolicyBindingConditionArgs' condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. Some resources **don't** support this identity.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. Some resources **don't** support this identity.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role/permission that will be granted to the members.
        See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
        Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def condition(self) -> Optional['outputs.GetIAMPolicyBindingConditionResult']:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
        """
        return pulumi.get(self, "condition")


@pulumi.output_type
class GetIAMPolicyBindingConditionResult(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        """
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str title: A title for the expression, i.e. a short string describing its purpose.
        :param str description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")


