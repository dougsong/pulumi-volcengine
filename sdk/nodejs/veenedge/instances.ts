// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of veenedge instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.veenedge.Instances({
 *     ids: [
 *         "veen28*****21",
 *         "veen177110*****172",
 *     ],
 * });
 * ```
 */
export function instances(args?: InstancesArgs, opts?: pulumi.InvokeOptions): Promise<InstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:veenedge/instances:Instances", {
        "cloudServerIds": args.cloudServerIds,
        "ids": args.ids,
        "names": args.names,
        "outputFile": args.outputFile,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    /**
     * The list of cloud server ids.
     */
    cloudServerIds?: string[];
    /**
     * A list of instance IDs.
     */
    ids?: string[];
    /**
     * A list of instance names.
     */
    names?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The list of instance status. The value can be `opening` or `starting` or `running` or `stopping` or `stop` or `rebooting` or `terminating`.
     */
    statuses?: string[];
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    readonly cloudServerIds?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The collection of instance query.
     */
    readonly instances: outputs.veenedge.InstancesInstance[];
    readonly names?: string[];
    readonly outputFile?: string;
    readonly statuses?: string[];
    /**
     * The total count of instance query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of veenedge instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.veenedge.Instances({
 *     ids: [
 *         "veen28*****21",
 *         "veen177110*****172",
 *     ],
 * });
 * ```
 */
export function instancesOutput(args?: InstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstancesResult> {
    return pulumi.output(args).apply((a: any) => instances(a, opts))
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesOutputArgs {
    /**
     * The list of cloud server ids.
     */
    cloudServerIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of instance IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of instance names.
     */
    names?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The list of instance status. The value can be `opening` or `starting` or `running` or `stopping` or `stop` or `rebooting` or `terminating`.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
}
