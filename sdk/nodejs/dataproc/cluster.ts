// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Cloud Dataproc cluster resource within GCP. For more information see
 * [the official dataproc documentation](https://cloud.google.com/dataproc/).
 *
 * !> **Warning:** Due to limitations of the API, all arguments except
 * `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing others will cause recreation of the
 * whole cluster!
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const simplecluster = new gcp.dataproc.Cluster("simplecluster", {
 *     region: "us-central1",
 * });
 * ```
 * ### Advanced
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.serviceAccount.Account("default", {
 *     accountId: "service-account-id",
 *     displayName: "Service Account",
 * });
 * const mycluster = new gcp.dataproc.Cluster("mycluster", {
 *     region: "us-central1",
 *     gracefulDecommissionTimeout: "120s",
 *     labels: {
 *         foo: "bar",
 *     },
 *     clusterConfig: {
 *         stagingBucket: "dataproc-staging-bucket",
 *         masterConfig: {
 *             numInstances: 1,
 *             machineType: "e2-medium",
 *             diskConfig: {
 *                 bootDiskType: "pd-ssd",
 *                 bootDiskSizeGb: 15,
 *             },
 *         },
 *         workerConfig: {
 *             numInstances: 2,
 *             machineType: "e2-medium",
 *             minCpuPlatform: "Intel Skylake",
 *             diskConfig: {
 *                 bootDiskSizeGb: 15,
 *                 numLocalSsds: 1,
 *             },
 *         },
 *         preemptibleWorkerConfig: {
 *             numInstances: 0,
 *         },
 *         softwareConfig: {
 *             imageVersion: "1.3.7-deb9",
 *             overrideProperties: {
 *                 "dataproc:dataproc.allow.zero.workers": "true",
 *             },
 *         },
 *         gceClusterConfig: {
 *             tags: [
 *                 "foo",
 *                 "bar",
 *             ],
 *             serviceAccount: _default.email,
 *             serviceAccountScopes: ["cloud-platform"],
 *         },
 *         initializationActions: [{
 *             script: "gs://dataproc-initialization-actions/stackdriver/stackdriver.sh",
 *             timeoutSec: 500,
 *         }],
 *     },
 * });
 * ```
 * ### Using A GPU Accelerator
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const acceleratedCluster = new gcp.dataproc.Cluster("accelerated_cluster", {
 *     clusterConfig: {
 *         gceClusterConfig: {
 *             zone: "us-central1-a",
 *         },
 *         masterConfig: {
 *             accelerators: [{
 *                 acceleratorCount: 1,
 *                 acceleratorType: "nvidia-tesla-k80",
 *             }],
 *         },
 *     },
 *     region: "us-central1",
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    public readonly clusterConfig!: pulumi.Output<outputs.dataproc.ClusterClusterConfig>;
    /**
     * The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
     * terraform apply
     */
    public readonly gracefulDecommissionTimeout!: pulumi.Output<string | undefined>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    public readonly region!: pulumi.Output<string | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["clusterConfig"] = state ? state.clusterConfig : undefined;
            inputs["gracefulDecommissionTimeout"] = state ? state.gracefulDecommissionTimeout : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            inputs["gracefulDecommissionTimeout"] = args ? args.gracefulDecommissionTimeout : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<inputs.dataproc.ClusterClusterConfig>;
    /**
     * The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
     * terraform apply
     */
    readonly gracefulDecommissionTimeout?: pulumi.Input<string>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<inputs.dataproc.ClusterClusterConfig>;
    /**
     * The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
     * terraform apply
     */
    readonly gracefulDecommissionTimeout?: pulumi.Input<string>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}
