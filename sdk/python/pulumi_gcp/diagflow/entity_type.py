# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['EntityType']


class EntityType(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
                 entities: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an entity type. Entity types serve as a tool for extracting parameter values from natural language queries.

        To get more information about EntityType, see:

        * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.entityTypes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dialogflow/docs/)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enable_fuzzy_extraction'] = enable_fuzzy_extraction
            __props__['entities'] = entities
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['project'] = project
            __props__['name'] = None
        super(EntityType, __self__).__init__(
            'gcp:diagflow/entityType:EntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
            entities: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'EntityType':
        """
        Get an existing EntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[str] name: The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["enable_fuzzy_extraction"] = enable_fuzzy_extraction
        __props__["entities"] = entities
        __props__["kind"] = kind
        __props__["name"] = name
        __props__["project"] = project
        return EntityType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableFuzzyExtraction")
    def enable_fuzzy_extraction(self) -> Optional[bool]:
        """
        Enables fuzzy entity extraction during classification.
        """
        return pulumi.get(self, "enable_fuzzy_extraction")

    @property
    @pulumi.getter
    def entities(self) -> Optional[List['outputs.EntityTypeEntity']]:
        """
        The collection of entity entries associated with the entity type.
        Structure is documented below.
        """
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Indicates the kind of entity type.
        * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
        * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
        types can contain references to other entity types (with or without aliases).
        * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
        Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

