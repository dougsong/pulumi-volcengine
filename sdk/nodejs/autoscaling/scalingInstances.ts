// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of scaling instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultScalingInstances = pulumi.output(volcengine.autoscaling.ScalingInstances({
 *     ids: [
 *         "i-ybzl4u5uogl8j07hgcbg",
 *         "i-ybyncrcpzpgh9zmlct0w",
 *         "i-ybyncrcpzogh9z4ax9tv",
 *     ],
 *     scalingConfigurationId: "scc-ybtawzucw95pkgon0wst",
 *     scalingGroupId: "scg-ybtawtznszgh9yv8agcp",
 *     status: "InService",
 * }));
 * ```
 */
export function scalingInstances(args: ScalingInstancesArgs, opts?: pulumi.InvokeOptions): Promise<ScalingInstancesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:autoscaling/scalingInstances:ScalingInstances", {
        "creationType": args.creationType,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "scalingConfigurationId": args.scalingConfigurationId,
        "scalingGroupId": args.scalingGroupId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking ScalingInstances.
 */
export interface ScalingInstancesArgs {
    /**
     * The creation type of the instances. Valid values: AutoCreated, Attached.
     */
    creationType?: string;
    /**
     * A list of instance ids.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of the scaling configuration id.
     */
    scalingConfigurationId?: string;
    /**
     * The id of the scaling group.
     */
    scalingGroupId: string;
    /**
     * The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
     */
    status?: string;
}

/**
 * A collection of values returned by ScalingInstances.
 */
export interface ScalingInstancesResult {
    /**
     * The creation type of the instance. Valid values: AutoCreated, Attached.
     */
    readonly creationType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The id of the scaling configuration.
     */
    readonly scalingConfigurationId?: string;
    /**
     * The id of the scaling group.
     */
    readonly scalingGroupId: string;
    /**
     * The collection of scaling instances.
     */
    readonly scalingInstances: outputs.autoscaling.ScalingInstancesScalingInstance[];
    /**
     * The status of instances.
     */
    readonly status?: string;
    /**
     * The total count of scaling instances query.
     */
    readonly totalCount: number;
}

export function scalingInstancesOutput(args: ScalingInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ScalingInstancesResult> {
    return pulumi.output(args).apply(a => scalingInstances(a, opts))
}

/**
 * A collection of arguments for invoking ScalingInstances.
 */
export interface ScalingInstancesOutputArgs {
    /**
     * The creation type of the instances. Valid values: AutoCreated, Attached.
     */
    creationType?: pulumi.Input<string>;
    /**
     * A list of instance ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of the scaling configuration id.
     */
    scalingConfigurationId?: pulumi.Input<string>;
    /**
     * The id of the scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
    /**
     * The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
     */
    status?: pulumi.Input<string>;
}
