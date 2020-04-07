// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Flexible App Version resource to create a new version of flexible GAE Application. Based on Google Compute Engine,
 * the App Engine flexible environment automatically scales your app up and down while also balancing the load.
 * Learn about the differences between the standard environment and the flexible environment
 * at https://cloud.google.com/appengine/docs/the-appengine-environments.
 * 
 * 
 * To get more information about FlexibleAppVersion, see:
 * 
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/appengine/docs/flexible)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_flexible_app_version.html.markdown.
 */
export class FlexibleAppVersion extends pulumi.CustomResource {
    /**
     * Get an existing FlexibleAppVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlexibleAppVersionState, opts?: pulumi.CustomResourceOptions): FlexibleAppVersion {
        return new FlexibleAppVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/flexibleAppVersion:FlexibleAppVersion';

    /**
     * Returns true if the given object is an instance of FlexibleAppVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlexibleAppVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlexibleAppVersion.__pulumiType;
    }

    /**
     * Serving configuration for Google Cloud Endpoints.
     */
    public readonly apiConfig!: pulumi.Output<outputs.appengine.FlexibleAppVersionApiConfig | undefined>;
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     */
    public readonly automaticScaling!: pulumi.Output<outputs.appengine.FlexibleAppVersionAutomaticScaling | undefined>;
    /**
     * Metadata settings that are supplied to this version to enable beta runtime features.
     */
    public readonly betaSettings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding
     * StaticFilesHandler does not specify its own expiration time.
     */
    public readonly defaultExpiration!: pulumi.Output<string | undefined>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.    
     */
    public readonly deleteServiceOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Code and application artifacts that make up this version.
     */
    public readonly deployment!: pulumi.Output<outputs.appengine.FlexibleAppVersionDeployment | undefined>;
    /**
     * Code and application artifacts that make up this version.
     */
    public readonly endpointsApiService!: pulumi.Output<outputs.appengine.FlexibleAppVersionEndpointsApiService | undefined>;
    /**
     * The entrypoint for the application.
     */
    public readonly entrypoint!: pulumi.Output<outputs.appengine.FlexibleAppVersionEntrypoint | undefined>;
    /**
     * Environment variables available to the application. As these are not returned in the API request, Terraform will not
     * detect any changes made outside of the Terraform config.
     */
    public readonly envVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Before an application can receive email or XMPP messages, the application must be configured to enable the service.
     */
    public readonly inboundServices!: pulumi.Output<string[] | undefined>;
    /**
     * Instance class that is used to run this version. Valid values are AutomaticScaling: F1, F2, F4, F4_1G ManualScaling:
     * B1, B2, B4, B8, B4_1G Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
     */
    public readonly instanceClass!: pulumi.Output<string | undefined>;
    /**
     * Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
     */
    public readonly livenessCheck!: pulumi.Output<outputs.appengine.FlexibleAppVersionLivenessCheck>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the
     * state of its memory over time.
     */
    public readonly manualScaling!: pulumi.Output<outputs.appengine.FlexibleAppVersionManualScaling | undefined>;
    /**
     * The identifier for this object. Format specified above.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Extra network settings
     */
    public readonly network!: pulumi.Output<outputs.appengine.FlexibleAppVersionNetwork | undefined>;
    /**
     * Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
     */
    public readonly nobuildFilesRegex!: pulumi.Output<string | undefined>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    public readonly noopOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic
     * rotation.
     */
    public readonly readinessCheck!: pulumi.Output<outputs.appengine.FlexibleAppVersionReadinessCheck>;
    /**
     * Machine resources for a version.
     */
    public readonly resources!: pulumi.Output<outputs.appengine.FlexibleAppVersionResources | undefined>;
    /**
     * Desired runtime. Example python27.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at
     * https://cloud.google.com/appengine/docs/standard//config/appref
     */
    public readonly runtimeApiVersion!: pulumi.Output<string>;
    /**
     * The channel of the runtime to use. Only available for some runtimes.
     */
    public readonly runtimeChannel!: pulumi.Output<string | undefined>;
    /**
     * The path or name of the app's main executable.
     */
    public readonly runtimeMainExecutablePath!: pulumi.Output<string | undefined>;
    /**
     * AppEngine service resource
     */
    public readonly service!: pulumi.Output<string | undefined>;
    /**
     * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
     * Defaults to SERVING.
     */
    public readonly servingStatus!: pulumi.Output<string | undefined>;
    /**
     * Relative name of the version within the service. For example, 'v1'. Version names can contain only lowercase
     * letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    public readonly versionId!: pulumi.Output<string | undefined>;
    /**
     * Enables VPC connectivity for standard apps.
     */
    public readonly vpcAccessConnector!: pulumi.Output<outputs.appengine.FlexibleAppVersionVpcAccessConnector | undefined>;

    /**
     * Create a FlexibleAppVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlexibleAppVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlexibleAppVersionArgs | FlexibleAppVersionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FlexibleAppVersionState | undefined;
            inputs["apiConfig"] = state ? state.apiConfig : undefined;
            inputs["automaticScaling"] = state ? state.automaticScaling : undefined;
            inputs["betaSettings"] = state ? state.betaSettings : undefined;
            inputs["defaultExpiration"] = state ? state.defaultExpiration : undefined;
            inputs["deleteServiceOnDestroy"] = state ? state.deleteServiceOnDestroy : undefined;
            inputs["deployment"] = state ? state.deployment : undefined;
            inputs["endpointsApiService"] = state ? state.endpointsApiService : undefined;
            inputs["entrypoint"] = state ? state.entrypoint : undefined;
            inputs["envVariables"] = state ? state.envVariables : undefined;
            inputs["inboundServices"] = state ? state.inboundServices : undefined;
            inputs["instanceClass"] = state ? state.instanceClass : undefined;
            inputs["livenessCheck"] = state ? state.livenessCheck : undefined;
            inputs["manualScaling"] = state ? state.manualScaling : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["nobuildFilesRegex"] = state ? state.nobuildFilesRegex : undefined;
            inputs["noopOnDestroy"] = state ? state.noopOnDestroy : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["readinessCheck"] = state ? state.readinessCheck : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["runtimeApiVersion"] = state ? state.runtimeApiVersion : undefined;
            inputs["runtimeChannel"] = state ? state.runtimeChannel : undefined;
            inputs["runtimeMainExecutablePath"] = state ? state.runtimeMainExecutablePath : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["servingStatus"] = state ? state.servingStatus : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
            inputs["vpcAccessConnector"] = state ? state.vpcAccessConnector : undefined;
        } else {
            const args = argsOrState as FlexibleAppVersionArgs | undefined;
            if (!args || args.livenessCheck === undefined) {
                throw new Error("Missing required property 'livenessCheck'");
            }
            if (!args || args.readinessCheck === undefined) {
                throw new Error("Missing required property 'readinessCheck'");
            }
            if (!args || args.runtime === undefined) {
                throw new Error("Missing required property 'runtime'");
            }
            inputs["apiConfig"] = args ? args.apiConfig : undefined;
            inputs["automaticScaling"] = args ? args.automaticScaling : undefined;
            inputs["betaSettings"] = args ? args.betaSettings : undefined;
            inputs["defaultExpiration"] = args ? args.defaultExpiration : undefined;
            inputs["deleteServiceOnDestroy"] = args ? args.deleteServiceOnDestroy : undefined;
            inputs["deployment"] = args ? args.deployment : undefined;
            inputs["endpointsApiService"] = args ? args.endpointsApiService : undefined;
            inputs["entrypoint"] = args ? args.entrypoint : undefined;
            inputs["envVariables"] = args ? args.envVariables : undefined;
            inputs["inboundServices"] = args ? args.inboundServices : undefined;
            inputs["instanceClass"] = args ? args.instanceClass : undefined;
            inputs["livenessCheck"] = args ? args.livenessCheck : undefined;
            inputs["manualScaling"] = args ? args.manualScaling : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["nobuildFilesRegex"] = args ? args.nobuildFilesRegex : undefined;
            inputs["noopOnDestroy"] = args ? args.noopOnDestroy : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["readinessCheck"] = args ? args.readinessCheck : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["runtimeApiVersion"] = args ? args.runtimeApiVersion : undefined;
            inputs["runtimeChannel"] = args ? args.runtimeChannel : undefined;
            inputs["runtimeMainExecutablePath"] = args ? args.runtimeMainExecutablePath : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["servingStatus"] = args ? args.servingStatus : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
            inputs["vpcAccessConnector"] = args ? args.vpcAccessConnector : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FlexibleAppVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlexibleAppVersion resources.
 */
export interface FlexibleAppVersionState {
    /**
     * Serving configuration for Google Cloud Endpoints.
     */
    readonly apiConfig?: pulumi.Input<inputs.appengine.FlexibleAppVersionApiConfig>;
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     */
    readonly automaticScaling?: pulumi.Input<inputs.appengine.FlexibleAppVersionAutomaticScaling>;
    /**
     * Metadata settings that are supplied to this version to enable beta runtime features.
     */
    readonly betaSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding
     * StaticFilesHandler does not specify its own expiration time.
     */
    readonly defaultExpiration?: pulumi.Input<string>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.    
     */
    readonly deleteServiceOnDestroy?: pulumi.Input<boolean>;
    /**
     * Code and application artifacts that make up this version.
     */
    readonly deployment?: pulumi.Input<inputs.appengine.FlexibleAppVersionDeployment>;
    /**
     * Code and application artifacts that make up this version.
     */
    readonly endpointsApiService?: pulumi.Input<inputs.appengine.FlexibleAppVersionEndpointsApiService>;
    /**
     * The entrypoint for the application.
     */
    readonly entrypoint?: pulumi.Input<inputs.appengine.FlexibleAppVersionEntrypoint>;
    /**
     * Environment variables available to the application. As these are not returned in the API request, Terraform will not
     * detect any changes made outside of the Terraform config.
     */
    readonly envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Before an application can receive email or XMPP messages, the application must be configured to enable the service.
     */
    readonly inboundServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance class that is used to run this version. Valid values are AutomaticScaling: F1, F2, F4, F4_1G ManualScaling:
     * B1, B2, B4, B8, B4_1G Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
     */
    readonly instanceClass?: pulumi.Input<string>;
    /**
     * Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
     */
    readonly livenessCheck?: pulumi.Input<inputs.appengine.FlexibleAppVersionLivenessCheck>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the
     * state of its memory over time.
     */
    readonly manualScaling?: pulumi.Input<inputs.appengine.FlexibleAppVersionManualScaling>;
    /**
     * The identifier for this object. Format specified above.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Extra network settings
     */
    readonly network?: pulumi.Input<inputs.appengine.FlexibleAppVersionNetwork>;
    /**
     * Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
     */
    readonly nobuildFilesRegex?: pulumi.Input<string>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    readonly noopOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic
     * rotation.
     */
    readonly readinessCheck?: pulumi.Input<inputs.appengine.FlexibleAppVersionReadinessCheck>;
    /**
     * Machine resources for a version.
     */
    readonly resources?: pulumi.Input<inputs.appengine.FlexibleAppVersionResources>;
    /**
     * Desired runtime. Example python27.
     */
    readonly runtime?: pulumi.Input<string>;
    /**
     * The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at
     * https://cloud.google.com/appengine/docs/standard//config/appref
     */
    readonly runtimeApiVersion?: pulumi.Input<string>;
    /**
     * The channel of the runtime to use. Only available for some runtimes.
     */
    readonly runtimeChannel?: pulumi.Input<string>;
    /**
     * The path or name of the app's main executable.
     */
    readonly runtimeMainExecutablePath?: pulumi.Input<string>;
    /**
     * AppEngine service resource
     */
    readonly service?: pulumi.Input<string>;
    /**
     * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
     * Defaults to SERVING.
     */
    readonly servingStatus?: pulumi.Input<string>;
    /**
     * Relative name of the version within the service. For example, 'v1'. Version names can contain only lowercase
     * letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    readonly versionId?: pulumi.Input<string>;
    /**
     * Enables VPC connectivity for standard apps.
     */
    readonly vpcAccessConnector?: pulumi.Input<inputs.appengine.FlexibleAppVersionVpcAccessConnector>;
}

/**
 * The set of arguments for constructing a FlexibleAppVersion resource.
 */
export interface FlexibleAppVersionArgs {
    /**
     * Serving configuration for Google Cloud Endpoints.
     */
    readonly apiConfig?: pulumi.Input<inputs.appengine.FlexibleAppVersionApiConfig>;
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     */
    readonly automaticScaling?: pulumi.Input<inputs.appengine.FlexibleAppVersionAutomaticScaling>;
    /**
     * Metadata settings that are supplied to this version to enable beta runtime features.
     */
    readonly betaSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding
     * StaticFilesHandler does not specify its own expiration time.
     */
    readonly defaultExpiration?: pulumi.Input<string>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.    
     */
    readonly deleteServiceOnDestroy?: pulumi.Input<boolean>;
    /**
     * Code and application artifacts that make up this version.
     */
    readonly deployment?: pulumi.Input<inputs.appengine.FlexibleAppVersionDeployment>;
    /**
     * Code and application artifacts that make up this version.
     */
    readonly endpointsApiService?: pulumi.Input<inputs.appengine.FlexibleAppVersionEndpointsApiService>;
    /**
     * The entrypoint for the application.
     */
    readonly entrypoint?: pulumi.Input<inputs.appengine.FlexibleAppVersionEntrypoint>;
    /**
     * Environment variables available to the application. As these are not returned in the API request, Terraform will not
     * detect any changes made outside of the Terraform config.
     */
    readonly envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Before an application can receive email or XMPP messages, the application must be configured to enable the service.
     */
    readonly inboundServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance class that is used to run this version. Valid values are AutomaticScaling: F1, F2, F4, F4_1G ManualScaling:
     * B1, B2, B4, B8, B4_1G Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
     */
    readonly instanceClass?: pulumi.Input<string>;
    /**
     * Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
     */
    readonly livenessCheck: pulumi.Input<inputs.appengine.FlexibleAppVersionLivenessCheck>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the
     * state of its memory over time.
     */
    readonly manualScaling?: pulumi.Input<inputs.appengine.FlexibleAppVersionManualScaling>;
    /**
     * Extra network settings
     */
    readonly network?: pulumi.Input<inputs.appengine.FlexibleAppVersionNetwork>;
    /**
     * Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
     */
    readonly nobuildFilesRegex?: pulumi.Input<string>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    readonly noopOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic
     * rotation.
     */
    readonly readinessCheck: pulumi.Input<inputs.appengine.FlexibleAppVersionReadinessCheck>;
    /**
     * Machine resources for a version.
     */
    readonly resources?: pulumi.Input<inputs.appengine.FlexibleAppVersionResources>;
    /**
     * Desired runtime. Example python27.
     */
    readonly runtime: pulumi.Input<string>;
    /**
     * The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at
     * https://cloud.google.com/appengine/docs/standard//config/appref
     */
    readonly runtimeApiVersion?: pulumi.Input<string>;
    /**
     * The channel of the runtime to use. Only available for some runtimes.
     */
    readonly runtimeChannel?: pulumi.Input<string>;
    /**
     * The path or name of the app's main executable.
     */
    readonly runtimeMainExecutablePath?: pulumi.Input<string>;
    /**
     * AppEngine service resource
     */
    readonly service?: pulumi.Input<string>;
    /**
     * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
     * Defaults to SERVING.
     */
    readonly servingStatus?: pulumi.Input<string>;
    /**
     * Relative name of the version within the service. For example, 'v1'. Version names can contain only lowercase
     * letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    readonly versionId?: pulumi.Input<string>;
    /**
     * Enables VPC connectivity for standard apps.
     */
    readonly vpcAccessConnector?: pulumi.Input<inputs.appengine.FlexibleAppVersionVpcAccessConnector>;
}