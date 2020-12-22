# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Document']


class Document(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 document_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
        that contains fields, which map to values. Each document is identified by a name.

        To get more information about Document, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)

        > **Warning:** This resource creates a Firestore Document on a project that already has
        Firestore enabled. If you haven't already enabled it, you can create a
        `appengine.Application` resource with `database_type` set to
        `"CLOUD_FIRESTORE"` to do so. Your Firestore location will be the same as
        the App Engine location specified.

        ## Example Usage

        ## Import

        Document can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firestore/document:Document default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
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

            if collection is None and not opts.urn:
                raise TypeError("Missing required property 'collection'")
            __props__['collection'] = collection
            __props__['database'] = database
            if document_id is None and not opts.urn:
                raise TypeError("Missing required property 'document_id'")
            __props__['document_id'] = document_id
            if fields is None and not opts.urn:
                raise TypeError("Missing required property 'fields'")
            __props__['fields'] = fields
            __props__['project'] = project
            __props__['create_time'] = None
            __props__['name'] = None
            __props__['path'] = None
            __props__['update_time'] = None
        super(Document, __self__).__init__(
            'gcp:firestore/document:Document',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collection: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            document_id: Optional[pulumi.Input[str]] = None,
            fields: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Document':
        """
        Get an existing Document resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 format.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        :param pulumi.Input[str] name: A server defined name for this index. Format:
               'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
        :param pulumi.Input[str] path: A relative path to the collection this document exists within
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: Last update timestamp in RFC3339 format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collection"] = collection
        __props__["create_time"] = create_time
        __props__["database"] = database
        __props__["document_id"] = document_id
        __props__["fields"] = fields
        __props__["name"] = name
        __props__["path"] = path
        __props__["project"] = project
        __props__["update_time"] = update_time
        return Document(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def collection(self) -> pulumi.Output[str]:
        """
        The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        """
        return pulumi.get(self, "collection")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="documentId")
    def document_id(self) -> pulumi.Output[str]:
        """
        The client-assigned document ID to use for this document during creation.
        """
        return pulumi.get(self, "document_id")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[str]:
        """
        The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. Format:
        'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        A relative path to the collection this document exists within
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update timestamp in RFC3339 format.
        """
        return pulumi.get(self, "update_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

