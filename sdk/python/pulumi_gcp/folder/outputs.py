# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'IAMBindingCondition',
    'IAMMemberCondition',
    'IamAuditConfigAuditLogConfig',
    'OrganizationPolicyBooleanPolicy',
    'OrganizationPolicyListPolicy',
    'OrganizationPolicyListPolicyAllow',
    'OrganizationPolicyListPolicyDeny',
    'OrganizationPolicyRestorePolicy',
    'GetOrganizationPolicyBooleanPolicyResult',
    'GetOrganizationPolicyListPolicyResult',
    'GetOrganizationPolicyListPolicyAllowResult',
    'GetOrganizationPolicyListPolicyDenyResult',
    'GetOrganizationPolicyRestorePolicyResult',
]

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
                 exempted_members: Optional[List[str]] = None):
        pulumi.set(__self__, "log_type", log_type)
        if exempted_members is not None:
            pulumi.set(__self__, "exempted_members", exempted_members)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[List[str]]:
        return pulumi.get(self, "exempted_members")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyBooleanPolicy(dict):
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
class OrganizationPolicyListPolicy(dict):
    def __init__(__self__, *,
                 allow: Optional['outputs.OrganizationPolicyListPolicyAllow'] = None,
                 deny: Optional['outputs.OrganizationPolicyListPolicyDeny'] = None,
                 inherit_from_parent: Optional[bool] = None,
                 suggested_value: Optional[str] = None):
        """
        :param 'OrganizationPolicyListPolicyAllowArgs' allow: or `deny` - (Optional) One or the other must be set.
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
    def allow(self) -> Optional['outputs.OrganizationPolicyListPolicyAllow']:
        """
        or `deny` - (Optional) One or the other must be set.
        """
        return pulumi.get(self, "allow")

    @property
    @pulumi.getter
    def deny(self) -> Optional['outputs.OrganizationPolicyListPolicyDeny']:
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
class OrganizationPolicyListPolicyAllow(dict):
    def __init__(__self__, *,
                 all: Optional[bool] = None,
                 values: Optional[List[str]] = None):
        """
        :param bool all: The policy allows or denies all values.
        :param List[str] values: The policy can define specific values that are allowed or denied.
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
    def values(self) -> Optional[List[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyListPolicyDeny(dict):
    def __init__(__self__, *,
                 all: Optional[bool] = None,
                 values: Optional[List[str]] = None):
        """
        :param bool all: The policy allows or denies all values.
        :param List[str] values: The policy can define specific values that are allowed or denied.
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
    def values(self) -> Optional[List[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyRestorePolicy(dict):
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
class GetOrganizationPolicyBooleanPolicyResult(dict):
    def __init__(__self__, *,
                 enforced: bool):
        pulumi.set(__self__, "enforced", enforced)

    @property
    @pulumi.getter
    def enforced(self) -> bool:
        return pulumi.get(self, "enforced")


@pulumi.output_type
class GetOrganizationPolicyListPolicyResult(dict):
    def __init__(__self__, *,
                 allows: List['outputs.GetOrganizationPolicyListPolicyAllowResult'],
                 denies: List['outputs.GetOrganizationPolicyListPolicyDenyResult'],
                 inherit_from_parent: bool,
                 suggested_value: str):
        pulumi.set(__self__, "allows", allows)
        pulumi.set(__self__, "denies", denies)
        pulumi.set(__self__, "inherit_from_parent", inherit_from_parent)
        pulumi.set(__self__, "suggested_value", suggested_value)

    @property
    @pulumi.getter
    def allows(self) -> List['outputs.GetOrganizationPolicyListPolicyAllowResult']:
        return pulumi.get(self, "allows")

    @property
    @pulumi.getter
    def denies(self) -> List['outputs.GetOrganizationPolicyListPolicyDenyResult']:
        return pulumi.get(self, "denies")

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> bool:
        return pulumi.get(self, "inherit_from_parent")

    @property
    @pulumi.getter(name="suggestedValue")
    def suggested_value(self) -> str:
        return pulumi.get(self, "suggested_value")


@pulumi.output_type
class GetOrganizationPolicyListPolicyAllowResult(dict):
    def __init__(__self__, *,
                 all: bool,
                 values: List[str]):
        pulumi.set(__self__, "all", all)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> bool:
        return pulumi.get(self, "all")

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        return pulumi.get(self, "values")


@pulumi.output_type
class GetOrganizationPolicyListPolicyDenyResult(dict):
    def __init__(__self__, *,
                 all: bool,
                 values: List[str]):
        pulumi.set(__self__, "all", all)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> bool:
        return pulumi.get(self, "all")

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        return pulumi.get(self, "values")


@pulumi.output_type
class GetOrganizationPolicyRestorePolicyResult(dict):
    def __init__(__self__, *,
                 default: bool):
        pulumi.set(__self__, "default", default)

    @property
    @pulumi.getter
    def default(self) -> bool:
        return pulumi.get(self, "default")


