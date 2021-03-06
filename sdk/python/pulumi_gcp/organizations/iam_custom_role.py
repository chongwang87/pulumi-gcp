# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['IAMCustomRole']


class IAMCustomRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows management of a customized Cloud IAM organization role. For more information see
        [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
        and
        [API](https://cloud.google.com/iam/reference/rest/v1/organizations.roles).

        > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
         from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
         same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
         after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
         made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
         by the provider, and new roles cannot share that name.

        ## Example Usage

        This snippet creates a customized IAM organization role.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_custom_role = gcp.organizations.IAMCustomRole("my-custom-role",
            description="A description",
            org_id="123456789",
            permissions=[
                "iam.roles.list",
                "iam.roles.create",
                "iam.roles.delete",
            ],
            role_id="myCustomRole",
            title="My Custom Role")
        ```

        ## Import

        Customized IAM organization role can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMCustomRole:IAMCustomRole my-custom-role organizations/123456789/roles/myCustomRole
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description for the role.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        :param pulumi.Input[str] role_id: The role id to use for this role.
        :param pulumi.Input[str] stage: The current launch stage of the role.
               Defaults to `GA`.
               List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        :param pulumi.Input[str] title: A human-readable title for the role.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__['org_id'] = org_id
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__['role_id'] = role_id
            __props__['stage'] = stage
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['deleted'] = None
            __props__['name'] = None
        super(IAMCustomRole, __self__).__init__(
            'gcp:organizations/iAMCustomRole:IAMCustomRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deleted: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            stage: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'IAMCustomRole':
        """
        Get an existing IAMCustomRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deleted: (Optional) The current deleted state of the role.
        :param pulumi.Input[str] description: A human-readable description for the role.
        :param pulumi.Input[str] name: The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        :param pulumi.Input[str] role_id: The role id to use for this role.
        :param pulumi.Input[str] stage: The current launch stage of the role.
               Defaults to `GA`.
               List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        :param pulumi.Input[str] title: A human-readable title for the role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["deleted"] = deleted
        __props__["description"] = description
        __props__["name"] = name
        __props__["org_id"] = org_id
        __props__["permissions"] = permissions
        __props__["role_id"] = role_id
        __props__["stage"] = stage
        __props__["title"] = title
        return IAMCustomRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def deleted(self) -> pulumi.Output[bool]:
        """
        (Optional) The current deleted state of the role.
        """
        return pulumi.get(self, "deleted")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable description for the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The numeric ID of the organization in which you want to create a custom role.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The role id to use for this role.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[Optional[str]]:
        """
        The current launch stage of the role.
        Defaults to `GA`.
        List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        """
        return pulumi.get(self, "stage")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        A human-readable title for the role.
        """
        return pulumi.get(self, "title")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

