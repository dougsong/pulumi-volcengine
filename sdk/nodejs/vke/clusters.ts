// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vke clusters
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vke.Clusters({
 *     podsConfigPodNetworkMode: "VpcCniShared",
 *     statuses: [{
 *         conditionsType: "Progressing",
 *         phase: "Creating",
 *     }],
 * });
 * ```
 */
export function clusters(args?: ClustersArgs, opts?: pulumi.InvokeOptions): Promise<ClustersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vke/clusters:Clusters", {
        "createClientToken": args.createClientToken,
        "deleteProtectionEnabled": args.deleteProtectionEnabled,
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "podsConfigPodNetworkMode": args.podsConfigPodNetworkMode,
        "statuses": args.statuses,
        "tags": args.tags,
        "updateClientToken": args.updateClientToken,
    }, opts);
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersArgs {
    /**
     * ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
     */
    createClientToken?: string;
    /**
     * The delete protection of the cluster, the value is `true` or `false`.
     */
    deleteProtectionEnabled?: boolean;
    /**
     * A list of Cluster IDs.
     */
    ids?: string[];
    /**
     * The name of the cluster.
     */
    name?: string;
    /**
     * A Name Regex of Cluster.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The page number of clusters query.
     */
    pageNumber?: number;
    /**
     * The page size of clusters query.
     */
    pageSize?: number;
    /**
     * The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
     */
    podsConfigPodNetworkMode?: string;
    /**
     * Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
     */
    statuses?: inputs.vke.ClustersStatus[];
    /**
     * Tags.
     */
    tags?: inputs.vke.ClustersTag[];
    /**
     * The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
     */
    updateClientToken?: string;
}

/**
 * A collection of values returned by Clusters.
 */
export interface ClustersResult {
    /**
     * The collection of VkeCluster query.
     */
    readonly clusters: outputs.vke.ClustersCluster[];
    readonly createClientToken?: string;
    /**
     * The delete protection of the cluster, the value is `true` or `false`.
     */
    readonly deleteProtectionEnabled?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of the cluster.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly pageNumber: number;
    readonly pageSize: number;
    readonly podsConfigPodNetworkMode?: string;
    readonly statuses?: outputs.vke.ClustersStatus[];
    /**
     * Tags of the Cluster.
     */
    readonly tags?: outputs.vke.ClustersTag[];
    /**
     * The total count of Cluster query.
     */
    readonly totalCount: number;
    readonly updateClientToken?: string;
}
/**
 * Use this data source to query detailed information of vke clusters
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vke.Clusters({
 *     podsConfigPodNetworkMode: "VpcCniShared",
 *     statuses: [{
 *         conditionsType: "Progressing",
 *         phase: "Creating",
 *     }],
 * });
 * ```
 */
export function clustersOutput(args?: ClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ClustersResult> {
    return pulumi.output(args).apply((a: any) => clusters(a, opts))
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersOutputArgs {
    /**
     * ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
     */
    createClientToken?: pulumi.Input<string>;
    /**
     * The delete protection of the cluster, the value is `true` or `false`.
     */
    deleteProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * A list of Cluster IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * A Name Regex of Cluster.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The page number of clusters query.
     */
    pageNumber?: pulumi.Input<number>;
    /**
     * The page size of clusters query.
     */
    pageSize?: pulumi.Input<number>;
    /**
     * The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
     */
    podsConfigPodNetworkMode?: pulumi.Input<string>;
    /**
     * Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.vke.ClustersStatusArgs>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vke.ClustersTagArgs>[]>;
    /**
     * The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
     */
    updateClientToken?: pulumi.Input<string>;
}
