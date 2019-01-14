# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NodePool(pulumi.CustomResource):
    """
    Manages a Node Pool resource within GKE. For more information see
    [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
    and
    [API](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.nodePools).
    """
    def __init__(__self__, __name__, __opts__=None, autoscaling=None, cluster=None, initial_node_count=None, management=None, max_pods_per_node=None, name=None, name_prefix=None, node_config=None, node_count=None, project=None, region=None, version=None, zone=None):
        """Create a NodePool resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autoscaling'] = autoscaling

        if not cluster:
            raise TypeError('Missing required property cluster')
        __props__['cluster'] = cluster

        __props__['initial_node_count'] = initial_node_count

        __props__['management'] = management

        __props__['max_pods_per_node'] = max_pods_per_node

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['node_config'] = node_config

        __props__['node_count'] = node_count

        __props__['project'] = project

        __props__['region'] = region

        __props__['version'] = version

        __props__['zone'] = zone

        __props__['instance_group_urls'] = None

        super(NodePool, __self__).__init__(
            'gcp:container/nodePool:NodePool',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

