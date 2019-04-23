# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TargetHttpsProxy(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    proxy_id: pulumi.Output[float]
    quic_override: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    ssl_certificates: pulumi.Output[list]
    ssl_policy: pulumi.Output[str]
    url_map: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, name=None, project=None, quic_override=None, ssl_certificates=None, ssl_policy=None, url_map=None, __name__=None, __opts__=None):
        """
        Represents a TargetHttpsProxy resource, which is used by one or more
        global forwarding rule to route incoming HTTPS requests to a URL map.
        
        
        To get more information about TargetHttpsProxy, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

        __props__['description'] = description

        __props__['name'] = name

        __props__['project'] = project

        __props__['quic_override'] = quic_override

        if ssl_certificates is None:
            raise TypeError("Missing required property 'ssl_certificates'")
        __props__['ssl_certificates'] = ssl_certificates

        __props__['ssl_policy'] = ssl_policy

        if url_map is None:
            raise TypeError("Missing required property 'url_map'")
        __props__['url_map'] = url_map

        __props__['creation_timestamp'] = None
        __props__['proxy_id'] = None
        __props__['self_link'] = None

        super(TargetHttpsProxy, __self__).__init__(
            'gcp:compute/targetHttpsProxy:TargetHttpsProxy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

