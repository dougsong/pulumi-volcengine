// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mongodb allow lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.mongodb.MongoAllowLists({
 *     instanceId: "mongo-replica-xxx",
 *     regionId: "cn-xxx",
 * });
 * ```
 */
export function mongoAllowLists(args: MongoAllowListsArgs, opts?: pulumi.InvokeOptions): Promise<MongoAllowListsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:mongodb/mongoAllowLists:MongoAllowLists", {
        "allowListIds": args.allowListIds,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "regionId": args.regionId,
    }, opts);
}

/**
 * A collection of arguments for invoking MongoAllowLists.
 */
export interface MongoAllowListsArgs {
    /**
     * The allow list IDs to query.
     */
    allowListIds?: string[];
    /**
     * The instance ID to query.
     */
    instanceId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The region ID.
     */
    regionId: string;
}

/**
 * A collection of values returned by MongoAllowLists.
 */
export interface MongoAllowListsResult {
    readonly allowListIds?: string[];
    /**
     * The collection of mongodb allow list query.
     */
    readonly allowLists: outputs.mongodb.MongoAllowListsAllowList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The instance id that bound to the allow list.
     */
    readonly instanceId?: string;
    readonly outputFile?: string;
    readonly regionId: string;
    /**
     * The total count of mongodb allow lists query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of mongodb allow lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.mongodb.MongoAllowLists({
 *     instanceId: "mongo-replica-xxx",
 *     regionId: "cn-xxx",
 * });
 * ```
 */
export function mongoAllowListsOutput(args: MongoAllowListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<MongoAllowListsResult> {
    return pulumi.output(args).apply((a: any) => mongoAllowLists(a, opts))
}

/**
 * A collection of arguments for invoking MongoAllowLists.
 */
export interface MongoAllowListsOutputArgs {
    /**
     * The allow list IDs to query.
     */
    allowListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance ID to query.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The region ID.
     */
    regionId: pulumi.Input<string>;
}
