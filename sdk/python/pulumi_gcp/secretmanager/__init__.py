# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_secret_version import *
from .secret import *
from .secret_iam_binding import *
from .secret_iam_member import *
from .secret_iam_policy import *
from .secret_version import *
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
            if typ == "gcp:secretmanager/secret:Secret":
                return Secret(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:secretmanager/secretIamBinding:SecretIamBinding":
                return SecretIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:secretmanager/secretIamMember:SecretIamMember":
                return SecretIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:secretmanager/secretIamPolicy:SecretIamPolicy":
                return SecretIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:secretmanager/secretVersion:SecretVersion":
                return SecretVersion(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "secretmanager/secret", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "secretmanager/secretIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "secretmanager/secretIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "secretmanager/secretIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "secretmanager/secretVersion", _module_instance)

_register_module()
