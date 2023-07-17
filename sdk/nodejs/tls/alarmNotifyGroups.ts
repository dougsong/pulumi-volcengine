// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tls alarm notify groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultAlarmNotifyGroups = pulumi.output(volcengine.tls.AlarmNotifyGroups());
 * ```
 */
export function alarmNotifyGroups(args?: AlarmNotifyGroupsArgs, opts?: pulumi.InvokeOptions): Promise<AlarmNotifyGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:tls/alarmNotifyGroups:AlarmNotifyGroups", {
        "alarmNotifyGroupId": args.alarmNotifyGroupId,
        "alarmNotifyGroupName": args.alarmNotifyGroupName,
        "iamProjectName": args.iamProjectName,
        "outputFile": args.outputFile,
        "receiverName": args.receiverName,
    }, opts);
}

/**
 * A collection of arguments for invoking AlarmNotifyGroups.
 */
export interface AlarmNotifyGroupsArgs {
    /**
     * The id of the alarm notify group.
     */
    alarmNotifyGroupId?: string;
    /**
     * The name of the alarm notify group.
     */
    alarmNotifyGroupName?: string;
    /**
     * The name of the iam project.
     */
    iamProjectName?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of the receiver.
     */
    receiverName?: string;
}

/**
 * A collection of values returned by AlarmNotifyGroups.
 */
export interface AlarmNotifyGroupsResult {
    /**
     * The id of the notify group.
     */
    readonly alarmNotifyGroupId?: string;
    /**
     * Name of the notification group.
     */
    readonly alarmNotifyGroupName?: string;
    /**
     * The list of the notify groups.
     */
    readonly groups: outputs.tls.AlarmNotifyGroupsGroup[];
    /**
     * The iam project name.
     */
    readonly iamProjectName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    readonly receiverName?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}

export function alarmNotifyGroupsOutput(args?: AlarmNotifyGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AlarmNotifyGroupsResult> {
    return pulumi.output(args).apply(a => alarmNotifyGroups(a, opts))
}

/**
 * A collection of arguments for invoking AlarmNotifyGroups.
 */
export interface AlarmNotifyGroupsOutputArgs {
    /**
     * The id of the alarm notify group.
     */
    alarmNotifyGroupId?: pulumi.Input<string>;
    /**
     * The name of the alarm notify group.
     */
    alarmNotifyGroupName?: pulumi.Input<string>;
    /**
     * The name of the iam project.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of the receiver.
     */
    receiverName?: pulumi.Input<string>;
}
