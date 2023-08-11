// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis backups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.Backups({
 *     backupStrategyLists: ["ManualBackup"],
 *     instanceId: "redis-cnlfvrv4qye6u4lpa",
 * });
 * ```
 */
export function backups(args: BackupsArgs, opts?: pulumi.InvokeOptions): Promise<BackupsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:redis/backups:Backups", {
        "backupStrategyLists": args.backupStrategyLists,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsArgs {
    /**
     * The list of backup strategy, support AutomatedBackup and ManualBackup.
     */
    backupStrategyLists?: string[];
    /**
     * Query end time.
     */
    endTime?: string;
    /**
     * Id of instance.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Query start time.
     */
    startTime?: string;
}

/**
 * A collection of values returned by Backups.
 */
export interface BackupsResult {
    readonly backupStrategyLists?: string[];
    /**
     * Information of backups.
     */
    readonly backups: outputs.redis.BackupsBackup[];
    /**
     * End time of backup.
     */
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Id of instance.
     */
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * Start time of backup.
     */
    readonly startTime?: string;
    /**
     * The total count of backup query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of redis backups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.Backups({
 *     backupStrategyLists: ["ManualBackup"],
 *     instanceId: "redis-cnlfvrv4qye6u4lpa",
 * });
 * ```
 */
export function backupsOutput(args: BackupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BackupsResult> {
    return pulumi.output(args).apply((a: any) => backups(a, opts))
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsOutputArgs {
    /**
     * The list of backup strategy, support AutomatedBackup and ManualBackup.
     */
    backupStrategyLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Query end time.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Id of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Query start time.
     */
    startTime?: pulumi.Input<string>;
}
