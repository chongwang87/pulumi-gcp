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

__all__ = ['NodePool']


class NodePool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling: Optional[pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 initial_node_count: Optional[pulumi.Input[float]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management: Optional[pulumi.Input[pulumi.InputType['NodePoolManagementArgs']]] = None,
                 max_pods_per_node: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['NodePoolNodeConfigArgs']]] = None,
                 node_count: Optional[pulumi.Input[float]] = None,
                 node_locations: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 upgrade_settings: Optional[pulumi.Input[pulumi.InputType['NodePoolUpgradeSettingsArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
        the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
        and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools).

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']] autoscaling: Configuration required by cluster autoscaler to adjust
               the size of the node pool to the current cluster usage. Structure is documented below.
        :param pulumi.Input[str] cluster: The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        :param pulumi.Input[float] initial_node_count: The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone.
               Changing this will force recreation of the resource.
        :param pulumi.Input[str] location: The location (region or zone) of the cluster.
        :param pulumi.Input[pulumi.InputType['NodePoolManagementArgs']] management: Node management configuration, wherein auto-repair and
               auto-upgrade is configured. Structure is documented below.
        :param pulumi.Input[float] max_pods_per_node: The maximum number of pods per node in this node pool.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] name: The name of the node pool. If left blank, the provider will
               auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name for the node pool beginning
               with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[pulumi.InputType['NodePoolNodeConfigArgs']] node_config: The node configuration of the pool. See
               container.Cluster for schema.
        :param pulumi.Input[float] node_count: The number of nodes per instance group. This field can be used to
               update the number of nodes per instance group but should not be used alongside `autoscaling`.
        :param pulumi.Input[List[pulumi.Input[str]]] node_locations: The list of zones in which the node pool's nodes should be located. Nodes must
               be in the region of their regional cluster or in the same region as their
               cluster's zone for zonal clusters. If unspecified, the cluster-level
               `node_locations` will be used.
        :param pulumi.Input[str] project: The ID of the project in which to create the node pool. If blank,
               the provider-configured project will be used.
        :param pulumi.Input[pulumi.InputType['NodePoolUpgradeSettingsArgs']] upgrade_settings: Specify node upgrade settings to change how many nodes GKE attempts to
               upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
               The maximum number of nodes upgraded simultaneously is limited to 20.
        :param pulumi.Input[str] version: The Kubernetes version for the nodes in this pool. Note that if this field
               and `auto_upgrade` are both specified, they will fight each other for what the node version should
               be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
               recommended that you specify explicit versions as the provider will see spurious diffs
               when fuzzy versions are used. See the `container.getEngineVersions` data source's
               `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
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

            __props__['autoscaling'] = autoscaling
            if cluster is None:
                raise TypeError("Missing required property 'cluster'")
            __props__['cluster'] = cluster
            __props__['initial_node_count'] = initial_node_count
            __props__['location'] = location
            __props__['management'] = management
            __props__['max_pods_per_node'] = max_pods_per_node
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['node_config'] = node_config
            __props__['node_count'] = node_count
            __props__['node_locations'] = node_locations
            __props__['project'] = project
            __props__['upgrade_settings'] = upgrade_settings
            __props__['version'] = version
            __props__['instance_group_urls'] = None
        super(NodePool, __self__).__init__(
            'gcp:container/nodePool:NodePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling: Optional[pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']]] = None,
            cluster: Optional[pulumi.Input[str]] = None,
            initial_node_count: Optional[pulumi.Input[float]] = None,
            instance_group_urls: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            management: Optional[pulumi.Input[pulumi.InputType['NodePoolManagementArgs']]] = None,
            max_pods_per_node: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            node_config: Optional[pulumi.Input[pulumi.InputType['NodePoolNodeConfigArgs']]] = None,
            node_count: Optional[pulumi.Input[float]] = None,
            node_locations: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            upgrade_settings: Optional[pulumi.Input[pulumi.InputType['NodePoolUpgradeSettingsArgs']]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'NodePool':
        """
        Get an existing NodePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodePoolAutoscalingArgs']] autoscaling: Configuration required by cluster autoscaler to adjust
               the size of the node pool to the current cluster usage. Structure is documented below.
        :param pulumi.Input[str] cluster: The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        :param pulumi.Input[float] initial_node_count: The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone.
               Changing this will force recreation of the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] instance_group_urls: The resource URLs of the managed instance groups associated with this node pool.
        :param pulumi.Input[str] location: The location (region or zone) of the cluster.
        :param pulumi.Input[pulumi.InputType['NodePoolManagementArgs']] management: Node management configuration, wherein auto-repair and
               auto-upgrade is configured. Structure is documented below.
        :param pulumi.Input[float] max_pods_per_node: The maximum number of pods per node in this node pool.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] name: The name of the node pool. If left blank, the provider will
               auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name for the node pool beginning
               with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[pulumi.InputType['NodePoolNodeConfigArgs']] node_config: The node configuration of the pool. See
               container.Cluster for schema.
        :param pulumi.Input[float] node_count: The number of nodes per instance group. This field can be used to
               update the number of nodes per instance group but should not be used alongside `autoscaling`.
        :param pulumi.Input[List[pulumi.Input[str]]] node_locations: The list of zones in which the node pool's nodes should be located. Nodes must
               be in the region of their regional cluster or in the same region as their
               cluster's zone for zonal clusters. If unspecified, the cluster-level
               `node_locations` will be used.
        :param pulumi.Input[str] project: The ID of the project in which to create the node pool. If blank,
               the provider-configured project will be used.
        :param pulumi.Input[pulumi.InputType['NodePoolUpgradeSettingsArgs']] upgrade_settings: Specify node upgrade settings to change how many nodes GKE attempts to
               upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
               The maximum number of nodes upgraded simultaneously is limited to 20.
        :param pulumi.Input[str] version: The Kubernetes version for the nodes in this pool. Note that if this field
               and `auto_upgrade` are both specified, they will fight each other for what the node version should
               be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
               recommended that you specify explicit versions as the provider will see spurious diffs
               when fuzzy versions are used. See the `container.getEngineVersions` data source's
               `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling"] = autoscaling
        __props__["cluster"] = cluster
        __props__["initial_node_count"] = initial_node_count
        __props__["instance_group_urls"] = instance_group_urls
        __props__["location"] = location
        __props__["management"] = management
        __props__["max_pods_per_node"] = max_pods_per_node
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["node_config"] = node_config
        __props__["node_count"] = node_count
        __props__["node_locations"] = node_locations
        __props__["project"] = project
        __props__["upgrade_settings"] = upgrade_settings
        __props__["version"] = version
        return NodePool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def autoscaling(self) -> Optional['outputs.NodePoolAutoscaling']:
        """
        Configuration required by cluster autoscaler to adjust
        the size of the node pool to the current cluster usage. Structure is documented below.
        """
        return pulumi.get(self, "autoscaling")

    @property
    @pulumi.getter
    def cluster(self) -> str:
        """
        The cluster to create the node pool for. Cluster must be present in `location` provided for zonal clusters.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="initialNodeCount")
    def initial_node_count(self) -> float:
        """
        The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone.
        Changing this will force recreation of the resource.
        """
        return pulumi.get(self, "initial_node_count")

    @property
    @pulumi.getter(name="instanceGroupUrls")
    def instance_group_urls(self) -> List[str]:
        """
        The resource URLs of the managed instance groups associated with this node pool.
        """
        return pulumi.get(self, "instance_group_urls")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location (region or zone) of the cluster.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def management(self) -> 'outputs.NodePoolManagement':
        """
        Node management configuration, wherein auto-repair and
        auto-upgrade is configured. Structure is documented below.
        """
        return pulumi.get(self, "management")

    @property
    @pulumi.getter(name="maxPodsPerNode")
    def max_pods_per_node(self) -> float:
        """
        The maximum number of pods per node in this node pool.
        Note that this does not work on node pools which are "route-based" - that is, node
        pools belonging to clusters that do not have IP Aliasing enabled.
        See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
        for more information.
        """
        return pulumi.get(self, "max_pods_per_node")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the node pool. If left blank, the provider will
        auto-generate a unique name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> str:
        """
        Creates a unique name for the node pool beginning
        with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> 'outputs.NodePoolNodeConfig':
        """
        The node configuration of the pool. See
        container.Cluster for schema.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> float:
        """
        The number of nodes per instance group. This field can be used to
        update the number of nodes per instance group but should not be used alongside `autoscaling`.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="nodeLocations")
    def node_locations(self) -> List[str]:
        """
        The list of zones in which the node pool's nodes should be located. Nodes must
        be in the region of their regional cluster or in the same region as their
        cluster's zone for zonal clusters. If unspecified, the cluster-level
        `node_locations` will be used.
        """
        return pulumi.get(self, "node_locations")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which to create the node pool. If blank,
        the provider-configured project will be used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="upgradeSettings")
    def upgrade_settings(self) -> 'outputs.NodePoolUpgradeSettings':
        """
        Specify node upgrade settings to change how many nodes GKE attempts to
        upgrade at once. The number of nodes upgraded simultaneously is the sum of `max_surge` and `max_unavailable`.
        The maximum number of nodes upgraded simultaneously is limited to 20.
        """
        return pulumi.get(self, "upgrade_settings")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Kubernetes version for the nodes in this pool. Note that if this field
        and `auto_upgrade` are both specified, they will fight each other for what the node version should
        be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
        recommended that you specify explicit versions as the provider will see spurious diffs
        when fuzzy versions are used. See the `container.getEngineVersions` data source's
        `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

