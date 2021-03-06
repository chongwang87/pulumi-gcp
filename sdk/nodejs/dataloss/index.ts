// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./preventionDeidentifyTemplate";
export * from "./preventionInspectTemplate";
export * from "./preventionJobTrigger";
export * from "./preventionStoredInfoType";

// Import resources to register:
import { PreventionDeidentifyTemplate } from "./preventionDeidentifyTemplate";
import { PreventionInspectTemplate } from "./preventionInspectTemplate";
import { PreventionJobTrigger } from "./preventionJobTrigger";
import { PreventionStoredInfoType } from "./preventionStoredInfoType";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate":
                return new PreventionDeidentifyTemplate(name, <any>undefined, { urn })
            case "gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate":
                return new PreventionInspectTemplate(name, <any>undefined, { urn })
            case "gcp:dataloss/preventionJobTrigger:PreventionJobTrigger":
                return new PreventionJobTrigger(name, <any>undefined, { urn })
            case "gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType":
                return new PreventionStoredInfoType(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "dataloss/preventionDeidentifyTemplate", _module)
pulumi.runtime.registerResourceModule("gcp", "dataloss/preventionInspectTemplate", _module)
pulumi.runtime.registerResourceModule("gcp", "dataloss/preventionJobTrigger", _module)
pulumi.runtime.registerResourceModule("gcp", "dataloss/preventionStoredInfoType", _module)
