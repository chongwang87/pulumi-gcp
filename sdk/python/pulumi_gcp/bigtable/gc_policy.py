# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GCPolicy']


class GCPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 column_family: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 max_age: Optional[pulumi.Input[pulumi.InputType['GCPolicyMaxAgeArgs']]] = None,
                 max_versions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GCPolicyMaxVersionArgs']]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
        [the official documentation](https://cloud.google.com/bigtable/) and
        [API](https://cloud.google.com/bigtable/docs/go/reference).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.bigtable.Instance("instance", clusters=[gcp.bigtable.InstanceClusterArgs(
            cluster_id="tf-instance-cluster",
            zone="us-central1-b",
            num_nodes=3,
            storage_type="HDD",
        )])
        table = gcp.bigtable.Table("table",
            instance_name=instance.name,
            column_families=[gcp.bigtable.TableColumnFamilyArgs(
                family="name",
            )])
        policy = gcp.bigtable.GCPolicy("policy",
            instance_name=instance.name,
            table=table.name,
            column_family="name",
            max_age=gcp.bigtable.GCPolicyMaxAgeArgs(
                duration="168h",
            ))
        ```

        Multiple conditions is also supported. `UNION` when any of its sub-policies apply (OR). `INTERSECTION` when all its sub-policies apply (AND)

        ```python
        import pulumi
        import pulumi_gcp as gcp

        policy = gcp.bigtable.GCPolicy("policy",
            instance_name=google_bigtable_instance["instance"]["name"],
            table=google_bigtable_table["table"]["name"],
            column_family="name",
            mode="UNION",
            max_age=gcp.bigtable.GCPolicyMaxAgeArgs(
                duration="168h",
            ),
            max_versions=[gcp.bigtable.GCPolicyMaxVersionArgs(
                number=10,
            )])
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_family: The name of the column family.
        :param pulumi.Input[str] instance_name: The name of the Bigtable instance.
        :param pulumi.Input[pulumi.InputType['GCPolicyMaxAgeArgs']] max_age: GC policy that applies to all cells older than the given age.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GCPolicyMaxVersionArgs']]]] max_versions: GC policy that applies to all versions of a cell except for the most recent.
        :param pulumi.Input[str] mode: If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] table: The name of the table.
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

            if column_family is None and not opts.urn:
                raise TypeError("Missing required property 'column_family'")
            __props__['column_family'] = column_family
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__['instance_name'] = instance_name
            __props__['max_age'] = max_age
            __props__['max_versions'] = max_versions
            __props__['mode'] = mode
            __props__['project'] = project
            if table is None and not opts.urn:
                raise TypeError("Missing required property 'table'")
            __props__['table'] = table
        super(GCPolicy, __self__).__init__(
            'gcp:bigtable/gCPolicy:GCPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            column_family: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            max_age: Optional[pulumi.Input[pulumi.InputType['GCPolicyMaxAgeArgs']]] = None,
            max_versions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GCPolicyMaxVersionArgs']]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            table: Optional[pulumi.Input[str]] = None) -> 'GCPolicy':
        """
        Get an existing GCPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_family: The name of the column family.
        :param pulumi.Input[str] instance_name: The name of the Bigtable instance.
        :param pulumi.Input[pulumi.InputType['GCPolicyMaxAgeArgs']] max_age: GC policy that applies to all cells older than the given age.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GCPolicyMaxVersionArgs']]]] max_versions: GC policy that applies to all versions of a cell except for the most recent.
        :param pulumi.Input[str] mode: If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] table: The name of the table.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["column_family"] = column_family
        __props__["instance_name"] = instance_name
        __props__["max_age"] = max_age
        __props__["max_versions"] = max_versions
        __props__["mode"] = mode
        __props__["project"] = project
        __props__["table"] = table
        return GCPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="columnFamily")
    def column_family(self) -> pulumi.Output[str]:
        """
        The name of the column family.
        """
        return pulumi.get(self, "column_family")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        The name of the Bigtable instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> pulumi.Output[Optional['outputs.GCPolicyMaxAge']]:
        """
        GC policy that applies to all cells older than the given age.
        """
        return pulumi.get(self, "max_age")

    @property
    @pulumi.getter(name="maxVersions")
    def max_versions(self) -> pulumi.Output[Optional[Sequence['outputs.GCPolicyMaxVersion']]]:
        """
        GC policy that applies to all versions of a cell except for the most recent.
        """
        return pulumi.get(self, "max_versions")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def table(self) -> pulumi.Output[str]:
        """
        The name of the table.
        """
        return pulumi.get(self, "table")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

