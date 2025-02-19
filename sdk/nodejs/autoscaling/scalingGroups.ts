// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of scaling groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.autoscaling.ScalingGroups({
 *     ids: ["scg-ybru8pazhgl8j1di4tyd"],
 * });
 * ```
 */
export function scalingGroups(args?: ScalingGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ScalingGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:autoscaling/scalingGroups:ScalingGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scalingGroupNames": args.scalingGroupNames,
    }, opts);
}

/**
 * A collection of arguments for invoking ScalingGroups.
 */
export interface ScalingGroupsArgs {
    /**
     * A list of scaling group ids.
     */
    ids?: string[];
    /**
     * A Name Regex of scaling group.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A list of scaling group names.
     */
    scalingGroupNames?: string[];
}

/**
 * A collection of values returned by ScalingGroups.
 */
export interface ScalingGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly scalingGroupNames?: string[];
    /**
     * The collection of scaling group query.
     */
    readonly scalingGroups: outputs.autoscaling.ScalingGroupsScalingGroup[];
    /**
     * The total count of scaling group query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of scaling groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.autoscaling.ScalingGroups({
 *     ids: ["scg-ybru8pazhgl8j1di4tyd"],
 * });
 * ```
 */
export function scalingGroupsOutput(args?: ScalingGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ScalingGroupsResult> {
    return pulumi.output(args).apply((a: any) => scalingGroups(a, opts))
}

/**
 * A collection of arguments for invoking ScalingGroups.
 */
export interface ScalingGroupsOutputArgs {
    /**
     * A list of scaling group ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of scaling group.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A list of scaling group names.
     */
    scalingGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
}
