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

__all__ = ['InstanceTemplate']


class InstanceTemplate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_ip_forward: Optional[pulumi.Input[bool]] = None,
                 confidential_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateConfidentialInstanceConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateDiskArgs']]]]] = None,
                 enable_display: Optional[pulumi.Input[bool]] = None,
                 guest_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateGuestAcceleratorArgs']]]]] = None,
                 instance_description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 machine_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_startup_script: Optional[pulumi.Input[str]] = None,
                 min_cpu_platform: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateNetworkInterfaceArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduling: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateSchedulingArgs']]] = None,
                 service_account: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateServiceAccountArgs']]] = None,
                 shielded_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateShieldedInstanceConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a VM instance template resource within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_account = gcp.service_account.Account("defaultAccount",
            account_id="service-account-id",
            display_name="Service Account")
        my_image = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        foobar = gcp.compute.Disk("foobar",
            image=my_image.self_link,
            size=10,
            type="pd-ssd",
            zone="us-central1-a")
        default_instance_template = gcp.compute.InstanceTemplate("defaultInstanceTemplate",
            description="This template is used to create app server instances.",
            tags=[
                "foo",
                "bar",
            ],
            labels={
                "environment": "dev",
            },
            instance_description="description assigned to instances",
            machine_type="e2-medium",
            can_ip_forward=False,
            scheduling=gcp.compute.InstanceTemplateSchedulingArgs(
                automatic_restart=True,
                on_host_maintenance="MIGRATE",
            ),
            disks=[
                gcp.compute.InstanceTemplateDiskArgs(
                    source_image="debian-cloud/debian-9",
                    auto_delete=True,
                    boot=True,
                ),
                gcp.compute.InstanceTemplateDiskArgs(
                    source=foobar.name,
                    auto_delete=False,
                    boot=False,
                ),
            ],
            network_interfaces=[gcp.compute.InstanceTemplateNetworkInterfaceArgs(
                network="default",
            )],
            metadata={
                "foo": "bar",
            },
            service_account=gcp.compute.InstanceTemplateServiceAccountArgs(
                email=default_account.email,
                scopes=["cloud-platform"],
            ))
        ```
        ## Deploying the Latest Image

        A common way to use instance templates and managed instance groups is to deploy the
        latest image in a family, usually the latest build of your application. There are two
        ways to do this in the provider, and they have their pros and cons. The difference ends
        up being in how "latest" is interpreted. You can either deploy the latest image available
        when the provider runs, or you can have each instance check what the latest image is when
        it's being created, either as part of a scaling event or being rebuilt by the instance
        group manager.

        If you're not sure, we recommend deploying the latest image available when the provider runs,
        because this means all the instances in your group will be based on the same image, always,
        and means that no upgrades or changes to your instances happen outside of a `pulumi up`.
        You can achieve this by using the `compute.Image`
        data source, which will retrieve the latest image on every `pulumi apply`, and will update
        the template to use that specific image:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_image = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        instance_template = gcp.compute.InstanceTemplate("instanceTemplate",
            name_prefix="instance-template-",
            machine_type="e2-medium",
            region="us-central1",
            disks=[gcp.compute.InstanceTemplateDiskArgs(
                source_image=google_compute_image["my_image"]["self_link"],
            )])
        ```

        To have instances update to the latest on every scaling event or instance re-creation,
        use the family as the image for the disk, and it will use GCP's default behavior, setting
        the image for the template to the family:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance_template = gcp.compute.InstanceTemplate("instanceTemplate",
            disks=[gcp.compute.InstanceTemplateDiskArgs(
                source_image="debian-cloud/debian-9",
            )],
            machine_type="e2-medium",
            name_prefix="instance-template-",
            region="us-central1")
        ```

        ## Import

        Instance templates can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default projects/{{project}}/global/instanceTemplates/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{name}}
        ```

         [custom-vm-types]https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]https://cloud.google.com/network-tiers/docs/overview

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_ip_forward: Whether to allow sending and receiving of
               packets with non-matching source or destination IPs. This defaults to false.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateConfidentialInstanceConfigArgs']] confidential_instance_config: Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateDiskArgs']]]] disks: Disks to attach to instances created from this template.
               This can be specified multiple times for multiple disks. Structure is
               documented below.
        :param pulumi.Input[bool] enable_display: Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateGuestAcceleratorArgs']]]] guest_accelerators: List of the type and count of accelerator cards attached to the instance. Structure documented below.
        :param pulumi.Input[str] instance_description: A brief description to use for instances
               created from this template.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to instances
               created from this template,
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to make available from
               within instances created from this template.
        :param pulumi.Input[str] metadata_startup_script: An alternative to using the
               startup-script metadata key, mostly to match the compute_instance resource.
               This replaces the startup-script metadata key on the created instance and
               thus the two mechanisms are not allowed to be used simultaneously.
        :param pulumi.Input[str] min_cpu_platform: Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
               `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        :param pulumi.Input[str] name: The name of the instance template. If you leave
               this blank, the provider will auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateNetworkInterfaceArgs']]]] network_interfaces: Networks to attach to instances created from
               this template. This can be specified multiple times for multiple networks.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: An instance template is a global resource that is not
               bound to a zone or a region. However, you can still specify some regional
               resources in an instance template, which restricts the template to the
               region where that resource resides. For example, a custom `subnetwork`
               resource is tied to a specific region. Defaults to the region of the
               Provider if no value is given.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateSchedulingArgs']] scheduling: The scheduling strategy to use. More details about
               this configuration option are detailed below.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateServiceAccountArgs']] service_account: Service account to attach to the instance. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateShieldedInstanceConfigArgs']] shielded_instance_config: Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
               **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to attach to the instance.
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

            __props__['can_ip_forward'] = can_ip_forward
            __props__['confidential_instance_config'] = confidential_instance_config
            __props__['description'] = description
            if disks is None and not opts.urn:
                raise TypeError("Missing required property 'disks'")
            __props__['disks'] = disks
            __props__['enable_display'] = enable_display
            __props__['guest_accelerators'] = guest_accelerators
            __props__['instance_description'] = instance_description
            __props__['labels'] = labels
            if machine_type is None and not opts.urn:
                raise TypeError("Missing required property 'machine_type'")
            __props__['machine_type'] = machine_type
            __props__['metadata'] = metadata
            __props__['metadata_startup_script'] = metadata_startup_script
            __props__['min_cpu_platform'] = min_cpu_platform
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['network_interfaces'] = network_interfaces
            __props__['project'] = project
            __props__['region'] = region
            __props__['scheduling'] = scheduling
            __props__['service_account'] = service_account
            __props__['shielded_instance_config'] = shielded_instance_config
            __props__['tags'] = tags
            __props__['metadata_fingerprint'] = None
            __props__['self_link'] = None
            __props__['tags_fingerprint'] = None
        super(InstanceTemplate, __self__).__init__(
            'gcp:compute/instanceTemplate:InstanceTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_ip_forward: Optional[pulumi.Input[bool]] = None,
            confidential_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateConfidentialInstanceConfigArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateDiskArgs']]]]] = None,
            enable_display: Optional[pulumi.Input[bool]] = None,
            guest_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateGuestAcceleratorArgs']]]]] = None,
            instance_description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            machine_type: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            metadata_fingerprint: Optional[pulumi.Input[str]] = None,
            metadata_startup_script: Optional[pulumi.Input[str]] = None,
            min_cpu_platform: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateNetworkInterfaceArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scheduling: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateSchedulingArgs']]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            service_account: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateServiceAccountArgs']]] = None,
            shielded_instance_config: Optional[pulumi.Input[pulumi.InputType['InstanceTemplateShieldedInstanceConfigArgs']]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags_fingerprint: Optional[pulumi.Input[str]] = None) -> 'InstanceTemplate':
        """
        Get an existing InstanceTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_ip_forward: Whether to allow sending and receiving of
               packets with non-matching source or destination IPs. This defaults to false.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateConfidentialInstanceConfigArgs']] confidential_instance_config: Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateDiskArgs']]]] disks: Disks to attach to instances created from this template.
               This can be specified multiple times for multiple disks. Structure is
               documented below.
        :param pulumi.Input[bool] enable_display: Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
               **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateGuestAcceleratorArgs']]]] guest_accelerators: List of the type and count of accelerator cards attached to the instance. Structure documented below.
        :param pulumi.Input[str] instance_description: A brief description to use for instances
               created from this template.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to instances
               created from this template,
        :param pulumi.Input[str] machine_type: The machine type to create.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to make available from
               within instances created from this template.
        :param pulumi.Input[str] metadata_fingerprint: The unique fingerprint of the metadata.
        :param pulumi.Input[str] metadata_startup_script: An alternative to using the
               startup-script metadata key, mostly to match the compute_instance resource.
               This replaces the startup-script metadata key on the created instance and
               thus the two mechanisms are not allowed to be used simultaneously.
        :param pulumi.Input[str] min_cpu_platform: Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
               `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        :param pulumi.Input[str] name: The name of the instance template. If you leave
               this blank, the provider will auto-generate a unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceTemplateNetworkInterfaceArgs']]]] network_interfaces: Networks to attach to instances created from
               this template. This can be specified multiple times for multiple networks.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: An instance template is a global resource that is not
               bound to a zone or a region. However, you can still specify some regional
               resources in an instance template, which restricts the template to the
               region where that resource resides. For example, a custom `subnetwork`
               resource is tied to a specific region. Defaults to the region of the
               Provider if no value is given.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateSchedulingArgs']] scheduling: The scheduling strategy to use. More details about
               this configuration option are detailed below.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateServiceAccountArgs']] service_account: Service account to attach to the instance. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['InstanceTemplateShieldedInstanceConfigArgs']] shielded_instance_config: Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
               **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to attach to the instance.
        :param pulumi.Input[str] tags_fingerprint: The unique fingerprint of the tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["can_ip_forward"] = can_ip_forward
        __props__["confidential_instance_config"] = confidential_instance_config
        __props__["description"] = description
        __props__["disks"] = disks
        __props__["enable_display"] = enable_display
        __props__["guest_accelerators"] = guest_accelerators
        __props__["instance_description"] = instance_description
        __props__["labels"] = labels
        __props__["machine_type"] = machine_type
        __props__["metadata"] = metadata
        __props__["metadata_fingerprint"] = metadata_fingerprint
        __props__["metadata_startup_script"] = metadata_startup_script
        __props__["min_cpu_platform"] = min_cpu_platform
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["network_interfaces"] = network_interfaces
        __props__["project"] = project
        __props__["region"] = region
        __props__["scheduling"] = scheduling
        __props__["self_link"] = self_link
        __props__["service_account"] = service_account
        __props__["shielded_instance_config"] = shielded_instance_config
        __props__["tags"] = tags
        __props__["tags_fingerprint"] = tags_fingerprint
        return InstanceTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canIpForward")
    def can_ip_forward(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to allow sending and receiving of
        packets with non-matching source or destination IPs. This defaults to false.
        """
        return pulumi.get(self, "can_ip_forward")

    @property
    @pulumi.getter(name="confidentialInstanceConfig")
    def confidential_instance_config(self) -> pulumi.Output['outputs.InstanceTemplateConfidentialInstanceConfig']:
        """
        Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM.
        """
        return pulumi.get(self, "confidential_instance_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A brief description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disks(self) -> pulumi.Output[Sequence['outputs.InstanceTemplateDisk']]:
        """
        Disks to attach to instances created from this template.
        This can be specified multiple times for multiple disks. Structure is
        documented below.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="enableDisplay")
    def enable_display(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
        **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
        """
        return pulumi.get(self, "enable_display")

    @property
    @pulumi.getter(name="guestAccelerators")
    def guest_accelerators(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceTemplateGuestAccelerator']]]:
        """
        List of the type and count of accelerator cards attached to the instance. Structure documented below.
        """
        return pulumi.get(self, "guest_accelerators")

    @property
    @pulumi.getter(name="instanceDescription")
    def instance_description(self) -> pulumi.Output[Optional[str]]:
        """
        A brief description to use for instances
        created from this template.
        """
        return pulumi.get(self, "instance_description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to instances
        created from this template,
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> pulumi.Output[str]:
        """
        The machine type to create.
        """
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Metadata key/value pairs to make available from
        within instances created from this template.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataFingerprint")
    def metadata_fingerprint(self) -> pulumi.Output[str]:
        """
        The unique fingerprint of the metadata.
        """
        return pulumi.get(self, "metadata_fingerprint")

    @property
    @pulumi.getter(name="metadataStartupScript")
    def metadata_startup_script(self) -> pulumi.Output[Optional[str]]:
        """
        An alternative to using the
        startup-script metadata key, mostly to match the compute_instance resource.
        This replaces the startup-script metadata key on the created instance and
        thus the two mechanisms are not allowed to be used simultaneously.
        """
        return pulumi.get(self, "metadata_startup_script")

    @property
    @pulumi.getter(name="minCpuPlatform")
    def min_cpu_platform(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
        `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
        """
        return pulumi.get(self, "min_cpu_platform")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the instance template. If you leave
        this blank, the provider will auto-generate a unique name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceTemplateNetworkInterface']]]:
        """
        Networks to attach to instances created from
        this template. This can be specified multiple times for multiple networks.
        Structure is documented below.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        An instance template is a global resource that is not
        bound to a zone or a region. However, you can still specify some regional
        resources in an instance template, which restricts the template to the
        region where that resource resides. For example, a custom `subnetwork`
        resource is tied to a specific region. Defaults to the region of the
        Provider if no value is given.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def scheduling(self) -> pulumi.Output['outputs.InstanceTemplateScheduling']:
        """
        The scheduling strategy to use. More details about
        this configuration option are detailed below.
        """
        return pulumi.get(self, "scheduling")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[Optional['outputs.InstanceTemplateServiceAccount']]:
        """
        Service account to attach to the instance. Structure is documented below.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="shieldedInstanceConfig")
    def shielded_instance_config(self) -> pulumi.Output['outputs.InstanceTemplateShieldedInstanceConfig']:
        """
        Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
        **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
        """
        return pulumi.get(self, "shielded_instance_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags to attach to the instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsFingerprint")
    def tags_fingerprint(self) -> pulumi.Output[str]:
        """
        The unique fingerprint of the tags.
        """
        return pulumi.get(self, "tags_fingerprint")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

