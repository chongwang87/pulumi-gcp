// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./api";
export * from "./apiConfig";
export * from "./apiConfigIamBinding";
export * from "./apiConfigIamMember";
export * from "./apiConfigIamPolicy";
export * from "./apiIamBinding";
export * from "./apiIamMember";
export * from "./apiIamPolicy";
export * from "./gateway";
export * from "./gatewayIamBinding";
export * from "./gatewayIamMember";
export * from "./gatewayIamPolicy";

// Import resources to register:
import { Api } from "./api";
import { ApiConfig } from "./apiConfig";
import { ApiConfigIamBinding } from "./apiConfigIamBinding";
import { ApiConfigIamMember } from "./apiConfigIamMember";
import { ApiConfigIamPolicy } from "./apiConfigIamPolicy";
import { ApiIamBinding } from "./apiIamBinding";
import { ApiIamMember } from "./apiIamMember";
import { ApiIamPolicy } from "./apiIamPolicy";
import { Gateway } from "./gateway";
import { GatewayIamBinding } from "./gatewayIamBinding";
import { GatewayIamMember } from "./gatewayIamMember";
import { GatewayIamPolicy } from "./gatewayIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:apigateway/api:Api":
                return new Api(name, <any>undefined, { urn })
            case "gcp:apigateway/apiConfig:ApiConfig":
                return new ApiConfig(name, <any>undefined, { urn })
            case "gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding":
                return new ApiConfigIamBinding(name, <any>undefined, { urn })
            case "gcp:apigateway/apiConfigIamMember:ApiConfigIamMember":
                return new ApiConfigIamMember(name, <any>undefined, { urn })
            case "gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy":
                return new ApiConfigIamPolicy(name, <any>undefined, { urn })
            case "gcp:apigateway/apiIamBinding:ApiIamBinding":
                return new ApiIamBinding(name, <any>undefined, { urn })
            case "gcp:apigateway/apiIamMember:ApiIamMember":
                return new ApiIamMember(name, <any>undefined, { urn })
            case "gcp:apigateway/apiIamPolicy:ApiIamPolicy":
                return new ApiIamPolicy(name, <any>undefined, { urn })
            case "gcp:apigateway/gateway:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "gcp:apigateway/gatewayIamBinding:GatewayIamBinding":
                return new GatewayIamBinding(name, <any>undefined, { urn })
            case "gcp:apigateway/gatewayIamMember:GatewayIamMember":
                return new GatewayIamMember(name, <any>undefined, { urn })
            case "gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy":
                return new GatewayIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "apigateway/api", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiConfigIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiConfigIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiConfigIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/apiIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/gateway", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/gatewayIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/gatewayIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "apigateway/gatewayIamPolicy", _module)
