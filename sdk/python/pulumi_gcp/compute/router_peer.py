# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RouterPeer(pulumi.CustomResource):
    advertise_mode: pulumi.Output[str]
    """
    User-specified flag to indicate which mode to use for advertisement.
    Valid values of this enum field are: `DEFAULT`, `CUSTOM`
    """
    advertised_groups: pulumi.Output[list]
    """
    User-specified list of prefix groups to advertise in custom
    mode, which can take one of the following options:
    * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
    * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
    * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
    """
    advertised_ip_ranges: pulumi.Output[list]
    """
    User-specified list of individual IP ranges to advertise in
    custom mode. This field can only be populated if advertiseMode
    is `CUSTOM` and is advertised to all peers of the router. These IP
    ranges will be advertised in addition to any specified groups.
    Leave this field blank to advertise no custom IP ranges.  Structure is documented below.

      * `description` (`str`) - User-specified description for the IP range.
      * `range` (`str`) - The IP range to advertise. The value must be a
        CIDR-formatted string.
    """
    advertised_route_priority: pulumi.Output[float]
    """
    The priority of routes advertised to this BGP peer.
    Where there is more than one matching route of maximum
    length, the routes with the lowest priority value win.
    """
    interface: pulumi.Output[str]
    """
    Name of the interface the BGP peer is associated with.
    """
    ip_address: pulumi.Output[str]
    """
    IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
    """
    management_type: pulumi.Output[str]
    """
    The resource that configures and manages this BGP peer. * 'MANAGED_BY_USER' is the default value and can be managed by
    you or other users * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and managed by Cloud Interconnect,
    specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type
    of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
    """
    name: pulumi.Output[str]
    """
    Name of this BGP peer. The name must be 1-63 characters long,
    and comply with RFC1035. Specifically, the name must be 1-63 characters
    long and match the regular expression `a-z?` which
    means the first character must be a lowercase letter, and all
    following characters must be a dash, lowercase letter, or digit,
    except the last character, which cannot be a dash.
    """
    peer_asn: pulumi.Output[float]
    """
    Peer BGP Autonomous System Number (ASN).
    Each BGP interface may use a different value.
    """
    peer_ip_address: pulumi.Output[str]
    """
    IP address of the BGP interface outside Google Cloud Platform.
    Only IPv4 is supported.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Region where the router and BgpPeer reside.
    If it is not provided, the provider region is used.
    """
    router: pulumi.Output[str]
    """
    The name of the Cloud Router in which this BgpPeer will be configured.
    """
    def __init__(__self__, resource_name, opts=None, advertise_mode=None, advertised_groups=None, advertised_ip_ranges=None, advertised_route_priority=None, interface=None, name=None, peer_asn=None, peer_ip_address=None, project=None, region=None, router=None, __props__=None, __name__=None, __opts__=None):
        """
        BGP information that must be configured into the routing stack to
        establish BGP peering. This information must specify the peer ASN
        and either the interface name, IP address, or peer IP address.
        Please refer to RFC4273.


        To get more information about RouterBgpPeer, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
        * How-to Guides
            * [Google Cloud Router](https://cloud.google.com/router/docs/)

        ## Example Usage

        ### Router Peer Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        peer = gcp.compute.RouterPeer("peer",
            advertised_route_priority=100,
            interface="interface-1",
            peer_asn=65513,
            peer_ip_address="169.254.1.2",
            region="us-central1",
            router="my-router")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advertise_mode: User-specified flag to indicate which mode to use for advertisement.
               Valid values of this enum field are: `DEFAULT`, `CUSTOM`
        :param pulumi.Input[list] advertised_groups: User-specified list of prefix groups to advertise in custom
               mode, which can take one of the following options:
               * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
               * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
               * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
        :param pulumi.Input[list] advertised_ip_ranges: User-specified list of individual IP ranges to advertise in
               custom mode. This field can only be populated if advertiseMode
               is `CUSTOM` and is advertised to all peers of the router. These IP
               ranges will be advertised in addition to any specified groups.
               Leave this field blank to advertise no custom IP ranges.  Structure is documented below.
        :param pulumi.Input[float] advertised_route_priority: The priority of routes advertised to this BGP peer.
               Where there is more than one matching route of maximum
               length, the routes with the lowest priority value win.
        :param pulumi.Input[str] interface: Name of the interface the BGP peer is associated with.
        :param pulumi.Input[str] name: Name of this BGP peer. The name must be 1-63 characters long,
               and comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[float] peer_asn: Peer BGP Autonomous System Number (ASN).
               Each BGP interface may use a different value.
        :param pulumi.Input[str] peer_ip_address: IP address of the BGP interface outside Google Cloud Platform.
               Only IPv4 is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the router and BgpPeer reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] router: The name of the Cloud Router in which this BgpPeer will be configured.

        The **advertised_ip_ranges** object supports the following:

          * `description` (`pulumi.Input[str]`) - User-specified description for the IP range.
          * `range` (`pulumi.Input[str]`) - The IP range to advertise. The value must be a
            CIDR-formatted string.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['advertise_mode'] = advertise_mode
            __props__['advertised_groups'] = advertised_groups
            __props__['advertised_ip_ranges'] = advertised_ip_ranges
            __props__['advertised_route_priority'] = advertised_route_priority
            if interface is None:
                raise TypeError("Missing required property 'interface'")
            __props__['interface'] = interface
            __props__['name'] = name
            if peer_asn is None:
                raise TypeError("Missing required property 'peer_asn'")
            __props__['peer_asn'] = peer_asn
            if peer_ip_address is None:
                raise TypeError("Missing required property 'peer_ip_address'")
            __props__['peer_ip_address'] = peer_ip_address
            __props__['project'] = project
            __props__['region'] = region
            if router is None:
                raise TypeError("Missing required property 'router'")
            __props__['router'] = router
            __props__['ip_address'] = None
            __props__['management_type'] = None
        super(RouterPeer, __self__).__init__(
            'gcp:compute/routerPeer:RouterPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, advertise_mode=None, advertised_groups=None, advertised_ip_ranges=None, advertised_route_priority=None, interface=None, ip_address=None, management_type=None, name=None, peer_asn=None, peer_ip_address=None, project=None, region=None, router=None):
        """
        Get an existing RouterPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advertise_mode: User-specified flag to indicate which mode to use for advertisement.
               Valid values of this enum field are: `DEFAULT`, `CUSTOM`
        :param pulumi.Input[list] advertised_groups: User-specified list of prefix groups to advertise in custom
               mode, which can take one of the following options:
               * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
               * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
               * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
        :param pulumi.Input[list] advertised_ip_ranges: User-specified list of individual IP ranges to advertise in
               custom mode. This field can only be populated if advertiseMode
               is `CUSTOM` and is advertised to all peers of the router. These IP
               ranges will be advertised in addition to any specified groups.
               Leave this field blank to advertise no custom IP ranges.  Structure is documented below.
        :param pulumi.Input[float] advertised_route_priority: The priority of routes advertised to this BGP peer.
               Where there is more than one matching route of maximum
               length, the routes with the lowest priority value win.
        :param pulumi.Input[str] interface: Name of the interface the BGP peer is associated with.
        :param pulumi.Input[str] ip_address: IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
        :param pulumi.Input[str] management_type: The resource that configures and manages this BGP peer. * 'MANAGED_BY_USER' is the default value and can be managed by
               you or other users * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and managed by Cloud Interconnect,
               specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type
               of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
        :param pulumi.Input[str] name: Name of this BGP peer. The name must be 1-63 characters long,
               and comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[float] peer_asn: Peer BGP Autonomous System Number (ASN).
               Each BGP interface may use a different value.
        :param pulumi.Input[str] peer_ip_address: IP address of the BGP interface outside Google Cloud Platform.
               Only IPv4 is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the router and BgpPeer reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] router: The name of the Cloud Router in which this BgpPeer will be configured.

        The **advertised_ip_ranges** object supports the following:

          * `description` (`pulumi.Input[str]`) - User-specified description for the IP range.
          * `range` (`pulumi.Input[str]`) - The IP range to advertise. The value must be a
            CIDR-formatted string.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["advertise_mode"] = advertise_mode
        __props__["advertised_groups"] = advertised_groups
        __props__["advertised_ip_ranges"] = advertised_ip_ranges
        __props__["advertised_route_priority"] = advertised_route_priority
        __props__["interface"] = interface
        __props__["ip_address"] = ip_address
        __props__["management_type"] = management_type
        __props__["name"] = name
        __props__["peer_asn"] = peer_asn
        __props__["peer_ip_address"] = peer_ip_address
        __props__["project"] = project
        __props__["region"] = region
        __props__["router"] = router
        return RouterPeer(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

