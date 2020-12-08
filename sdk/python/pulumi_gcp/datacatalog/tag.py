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

__all__ = ['Tag']


class Tag(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 column: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagFieldArgs']]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Tags are used to attach custom metadata to Data Catalog resources. Tags conform to the specifications within their tag template.

        See [Data Catalog IAM](https://cloud.google.com/data-catalog/docs/concepts/iam) for information on the permissions needed to create or view tags.

        To get more information about Tag, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.tags)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Entry Tag Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_entry_group")
        entry = gcp.datacatalog.Entry("entry",
            entry_group=entry_group.id,
            entry_id="my_entry",
            user_specified_type="my_custom_type",
            user_specified_system="SomethingExternal")
        tag_template = gcp.datacatalog.TagTemplate("tagTemplate",
            tag_template_id="my_template",
            region="us-central1",
            display_name="Demo Tag Template",
            fields=[
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="source",
                    display_name="Source of data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="STRING",
                    ),
                    is_required=True,
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="num_rows",
                    display_name="Number of rows in the data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="DOUBLE",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="pii_type",
                    display_name="PII type",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        enum_type=gcp.datacatalog.TagTemplateFieldTypeEnumTypeArgs(
                            allowed_values=[
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="EMAIL",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="SOCIAL SECURITY NUMBER",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="NONE",
                                ),
                            ],
                        ),
                    ),
                ),
            ],
            force_delete=False)
        basic_tag = gcp.datacatalog.Tag("basicTag",
            parent=entry.id,
            template=tag_template.id,
            fields=[gcp.datacatalog.TagFieldArgs(
                field_name="source",
                string_value="my-string",
            )])
        ```
        ### Data Catalog Entry Group Tag

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_entry_group")
        first_entry = gcp.datacatalog.Entry("firstEntry",
            entry_group=entry_group.id,
            entry_id="first_entry",
            user_specified_type="my_custom_type",
            user_specified_system="SomethingExternal")
        second_entry = gcp.datacatalog.Entry("secondEntry",
            entry_group=entry_group.id,
            entry_id="second_entry",
            user_specified_type="another_custom_type",
            user_specified_system="SomethingElseExternal")
        tag_template = gcp.datacatalog.TagTemplate("tagTemplate",
            tag_template_id="my_template",
            region="us-central1",
            display_name="Demo Tag Template",
            fields=[
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="source",
                    display_name="Source of data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="STRING",
                    ),
                    is_required=True,
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="num_rows",
                    display_name="Number of rows in the data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="DOUBLE",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="pii_type",
                    display_name="PII type",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        enum_type=gcp.datacatalog.TagTemplateFieldTypeEnumTypeArgs(
                            allowed_values=[
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="EMAIL",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="SOCIAL SECURITY NUMBER",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="NONE",
                                ),
                            ],
                        ),
                    ),
                ),
            ],
            force_delete=False)
        entry_group_tag = gcp.datacatalog.Tag("entryGroupTag",
            parent=entry_group.id,
            template=tag_template.id,
            fields=[gcp.datacatalog.TagFieldArgs(
                field_name="source",
                string_value="my-string",
            )])
        ```
        ### Data Catalog Entry Tag Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_entry_group")
        entry = gcp.datacatalog.Entry("entry",
            entry_group=entry_group.id,
            entry_id="my_entry",
            user_specified_type="my_custom_type",
            user_specified_system="SomethingExternal",
            schema=\"\"\"{
          "columns": [
            {
              "column": "first_name",
              "description": "First name",
              "mode": "REQUIRED",
              "type": "STRING"
            },
            {
              "column": "last_name",
              "description": "Last name",
              "mode": "REQUIRED",
              "type": "STRING"
            },
            {
              "column": "address",
              "description": "Address",
              "mode": "REPEATED",
              "subcolumns": [
                {
                  "column": "city",
                  "description": "City",
                  "mode": "NULLABLE",
                  "type": "STRING"
                },
                {
                  "column": "state",
                  "description": "State",
                  "mode": "NULLABLE",
                  "type": "STRING"
                }
              ],
              "type": "RECORD"
            }
          ]
        }
        \"\"\")
        tag_template = gcp.datacatalog.TagTemplate("tagTemplate",
            tag_template_id="my_template",
            region="us-central1",
            display_name="Demo Tag Template",
            fields=[
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="source",
                    display_name="Source of data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="STRING",
                    ),
                    is_required=True,
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="num_rows",
                    display_name="Number of rows in the data asset",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        primitive_type="DOUBLE",
                    ),
                ),
                gcp.datacatalog.TagTemplateFieldArgs(
                    field_id="pii_type",
                    display_name="PII type",
                    type=gcp.datacatalog.TagTemplateFieldTypeArgs(
                        enum_type=gcp.datacatalog.TagTemplateFieldTypeEnumTypeArgs(
                            allowed_values=[
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="EMAIL",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="SOCIAL SECURITY NUMBER",
                                ),
                                gcp.datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs(
                                    display_name="NONE",
                                ),
                            ],
                        ),
                    ),
                ),
            ],
            force_delete=False)
        basic_tag = gcp.datacatalog.Tag("basicTag",
            parent=entry.id,
            template=tag_template.id,
            fields=[
                gcp.datacatalog.TagFieldArgs(
                    field_name="source",
                    string_value="my-string",
                ),
                gcp.datacatalog.TagFieldArgs(
                    field_name="num_rows",
                    double_value=5,
                ),
                gcp.datacatalog.TagFieldArgs(
                    field_name="pii_type",
                    enum_value="EMAIL",
                ),
            ],
            column="address")
        second_tag = gcp.datacatalog.Tag("second-tag",
            parent=entry.id,
            template=tag_template.id,
            fields=[
                gcp.datacatalog.TagFieldArgs(
                    field_name="source",
                    string_value="my-string",
                ),
                gcp.datacatalog.TagFieldArgs(
                    field_name="pii_type",
                    enum_value="NONE",
                ),
            ],
            column="first_name")
        ```

        ## Import

        Tag can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datacatalog/tag:Tag default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column: Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
               individual column based on that schema.
               For attaching a tag to a nested column, use `.` to separate the column names. Example:
               `outer_column.inner_column`
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagFieldArgs']]]] fields: This maps the ID of a tag field to the value of and additional information about that field.
               Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[str] parent: The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
               all entries in that group.
        :param pulumi.Input[str] template: The resource name of the tag template that this tag uses. Example:
               projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
               This field cannot be modified after creation.
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

            __props__['column'] = column
            if fields is None and not opts.urn:
                raise TypeError("Missing required property 'fields'")
            __props__['fields'] = fields
            __props__['parent'] = parent
            if template is None and not opts.urn:
                raise TypeError("Missing required property 'template'")
            __props__['template'] = template
            __props__['name'] = None
            __props__['template_displayname'] = None
        super(Tag, __self__).__init__(
            'gcp:datacatalog/tag:Tag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            column: Optional[pulumi.Input[str]] = None,
            fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagFieldArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None,
            template_displayname: Optional[pulumi.Input[str]] = None) -> 'Tag':
        """
        Get an existing Tag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column: Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
               individual column based on that schema.
               For attaching a tag to a nested column, use `.` to separate the column names. Example:
               `outer_column.inner_column`
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagFieldArgs']]]] fields: This maps the ID of a tag field to the value of and additional information about that field.
               Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
               Structure is documented below.
        :param pulumi.Input[str] name: The resource name of the tag in URL format. Example:
               projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
               projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
               identifier. Note that this Tag may not actually be stored in the location in this name.
        :param pulumi.Input[str] parent: The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
               all entries in that group.
        :param pulumi.Input[str] template: The resource name of the tag template that this tag uses. Example:
               projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
               This field cannot be modified after creation.
        :param pulumi.Input[str] template_displayname: The display name of the tag template.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["column"] = column
        __props__["fields"] = fields
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["template"] = template
        __props__["template_displayname"] = template_displayname
        return Tag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def column(self) -> pulumi.Output[Optional[str]]:
        """
        Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
        individual column based on that schema.
        For attaching a tag to a nested column, use `.` to separate the column names. Example:
        `outer_column.inner_column`
        """
        return pulumi.get(self, "column")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence['outputs.TagField']]:
        """
        This maps the ID of a tag field to the value of and additional information about that field.
        Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
        Structure is documented below.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the tag in URL format. Example:
        projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
        projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated
        identifier. Note that this Tag may not actually be stored in the location in this name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
        all entries in that group.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[str]:
        """
        The resource name of the tag template that this tag uses. Example:
        projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
        This field cannot be modified after creation.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="templateDisplayname")
    def template_displayname(self) -> pulumi.Output[str]:
        """
        The display name of the tag template.
        """
        return pulumi.get(self, "template_displayname")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

