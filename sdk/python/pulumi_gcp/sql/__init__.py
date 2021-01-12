# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .database import *
from .database_instance import *
from .get_backup_run import *
from .get_ca_certs import *
from .get_database_instance import *
from .source_representation_instance import *
from .ssl_cert import *
from .user import *
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
            if typ == "gcp:sql/database:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:sql/databaseInstance:DatabaseInstance":
                return DatabaseInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance":
                return SourceRepresentationInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:sql/sslCert:SslCert":
                return SslCert(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:sql/user:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "sql/database", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "sql/databaseInstance", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "sql/sourceRepresentationInstance", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "sql/sslCert", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "sql/user", _module_instance)

_register_module()
