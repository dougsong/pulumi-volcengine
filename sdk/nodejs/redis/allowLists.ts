// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis allow lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.AllowLists({
 *     regionId: "cn-beijing",
 * });
 * ```
 */
export function allowLists(args: AllowListsArgs, opts?: pulumi.InvokeOptions): Promise<AllowListsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:redis/allowLists:AllowLists", {
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "regionId": args.regionId,
    }, opts);
}

/**
 * A collection of arguments for invoking AllowLists.
 */
export interface AllowListsArgs {
    /**
     * The Id of instance.
     */
    instanceId?: string;
    /**
     * A Name Regex of Allow List.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The Id of region.
     */
    regionId: string;
}

/**
 * A collection of values returned by AllowLists.
 */
export interface AllowListsResult {
    /**
     * Information of list of allow list.
     */
    readonly allowLists: outputs.redis.AllowListsAllowList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Id of instance.
     */
    readonly instanceId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly regionId: string;
    /**
     * The total count of allow list query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of redis allow lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.AllowLists({
 *     regionId: "cn-beijing",
 * });
 * ```
 */
export function allowListsOutput(args: AllowListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AllowListsResult> {
    return pulumi.output(args).apply((a: any) => allowLists(a, opts))
}

/**
 * A collection of arguments for invoking AllowLists.
 */
export interface AllowListsOutputArgs {
    /**
     * The Id of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * A Name Regex of Allow List.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Id of region.
     */
    regionId: pulumi.Input<string>;
}
