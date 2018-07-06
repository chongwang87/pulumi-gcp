# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Subnetwork(pulumi.CustomResource):
    """
    Manages a subnetwork within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/vpc/#vpc_networks_and_subnets)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/subnetworks).
    """
    def __init__(__self__, __name__, __opts__=None, description=None, enable_flow_logs=None, ip_cidr_range=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None):
        """Create a Subnetwork resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        Description of this subnetwork.
        """
        __props__['description'] = description

        if enable_flow_logs and not isinstance(enable_flow_logs, bool):
            raise TypeError('Expected property enable_flow_logs to be a bool')
        __self__.enable_flow_logs = enable_flow_logs
        """
        )
        Set to `true` to enable [flow logs](https://cloud.google.com/vpc/docs/using-flow-logs)
        for this subnetwork.
        """
        __props__['enableFlowLogs'] = enable_flow_logs

        if not ip_cidr_range:
            raise TypeError('Missing required property ip_cidr_range')
        elif not isinstance(ip_cidr_range, basestring):
            raise TypeError('Expected property ip_cidr_range to be a basestring')
        __self__.ip_cidr_range = ip_cidr_range
        """
        The range of IP addresses belonging to this subnetwork secondary range. Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network.
        """
        __props__['ipCidrRange'] = ip_cidr_range

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        A unique name for the resource, required by GCE.
        Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not network:
            raise TypeError('Missing required property network')
        elif not isinstance(network, basestring):
            raise TypeError('Expected property network to be a basestring')
        __self__.network = network
        """
        The network name or resource link to the parent
        network of this subnetwork. The parent network must have been created
        in custom subnet mode.
        """
        __props__['network'] = network

        if private_ip_google_access and not isinstance(private_ip_google_access, bool):
            raise TypeError('Expected property private_ip_google_access to be a bool')
        __self__.private_ip_google_access = private_ip_google_access
        """
        Whether the VMs in this subnet
        can access Google services without assigned external IP
        addresses.
        """
        __props__['privateIpGoogleAccess'] = private_ip_google_access

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region this subnetwork will be created in. If
        unspecified, this defaults to the region configured in the provider.
        """
        __props__['region'] = region

        if secondary_ip_ranges and not isinstance(secondary_ip_ranges, list):
            raise TypeError('Expected property secondary_ip_ranges to be a list')
        __self__.secondary_ip_ranges = secondary_ip_ranges
        """
        ) An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. Structure is documented below.
        """
        __props__['secondaryIpRanges'] = secondary_ip_ranges

        __self__.fingerprint = pulumi.runtime.UNKNOWN
        __self__.gateway_address = pulumi.runtime.UNKNOWN
        """
        The IP address of the gateway.
        """
        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """

        super(Subnetwork, __self__).__init__(
            'gcp:compute/subnetwork:Subnetwork',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'description' in outs:
            self.description = outs['description']
        if 'enableFlowLogs' in outs:
            self.enable_flow_logs = outs['enableFlowLogs']
        if 'fingerprint' in outs:
            self.fingerprint = outs['fingerprint']
        if 'gatewayAddress' in outs:
            self.gateway_address = outs['gatewayAddress']
        if 'ipCidrRange' in outs:
            self.ip_cidr_range = outs['ipCidrRange']
        if 'name' in outs:
            self.name = outs['name']
        if 'network' in outs:
            self.network = outs['network']
        if 'privateIpGoogleAccess' in outs:
            self.private_ip_google_access = outs['privateIpGoogleAccess']
        if 'project' in outs:
            self.project = outs['project']
        if 'region' in outs:
            self.region = outs['region']
        if 'secondaryIpRanges' in outs:
            self.secondary_ip_ranges = outs['secondaryIpRanges']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']