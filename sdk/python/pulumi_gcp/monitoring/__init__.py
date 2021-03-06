# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .alert_policy import *
from .custom_service import *
from .dashboard import *
from .get_app_engine_service import *
from .get_cluster_istio_service import *
from .get_mesh_istio_service import *
from .get_notification_channel import *
from .get_secret_version import *
from .get_uptime_check_i_ps import *
from .group import *
from .metric_descriptor import *
from .notification_channel import *
from .slo import *
from .uptime_check_config import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:monitoring/alertPolicy:AlertPolicy":
                return AlertPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/customService:CustomService":
                return CustomService(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/dashboard:Dashboard":
                return Dashboard(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/group:Group":
                return Group(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/metricDescriptor:MetricDescriptor":
                return MetricDescriptor(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/notificationChannel:NotificationChannel":
                return NotificationChannel(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/slo:Slo":
                return Slo(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig":
                return UptimeCheckConfig(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "monitoring/alertPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/customService", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/dashboard", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/group", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/metricDescriptor", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/notificationChannel", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/slo", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "monitoring/uptimeCheckConfig", _module_instance)

_register_module()
