// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of scaling lifecycle hooks
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultScalingLifecycleHooks = pulumi.output(volcengine.autoscaling.ScalingLifecycleHooks({
 *     scalingGroupId: "scg-ybru8pazhgl8j1di4tyd",
 * }));
 * ```
 */
export function scalingLifecycleHooks(args: ScalingLifecycleHooksArgs, opts?: pulumi.InvokeOptions): Promise<ScalingLifecycleHooksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:autoscaling/scalingLifecycleHooks:ScalingLifecycleHooks", {
        "ids": args.ids,
        "lifecycleHookNames": args.lifecycleHookNames,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scalingGroupId": args.scalingGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking ScalingLifecycleHooks.
 */
export interface ScalingLifecycleHooksArgs {
    /**
     * A list of lifecycle hook ids.
     */
    ids?: string[];
    /**
     * A list of lifecycle hook names.
     */
    lifecycleHookNames?: string[];
    /**
     * A Name Regex of lifecycle hook.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * An id of scaling group id.
     */
    scalingGroupId: string;
}

/**
 * A collection of values returned by ScalingLifecycleHooks.
 */
export interface ScalingLifecycleHooksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly lifecycleHookNames?: string[];
    /**
     * The collection of lifecycle hook query.
     */
    readonly lifecycleHooks: outputs.autoscaling.ScalingLifecycleHooksLifecycleHook[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The id of the scaling group.
     */
    readonly scalingGroupId: string;
    /**
     * The total count of lifecycle hook query.
     */
    readonly totalCount: number;
}

export function scalingLifecycleHooksOutput(args: ScalingLifecycleHooksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ScalingLifecycleHooksResult> {
    return pulumi.output(args).apply(a => scalingLifecycleHooks(a, opts))
}

/**
 * A collection of arguments for invoking ScalingLifecycleHooks.
 */
export interface ScalingLifecycleHooksOutputArgs {
    /**
     * A list of lifecycle hook ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of lifecycle hook names.
     */
    lifecycleHookNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of lifecycle hook.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * An id of scaling group id.
     */
    scalingGroupId: pulumi.Input<string>;
}
