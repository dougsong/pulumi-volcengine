// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of volumes
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultVolumes = pulumi.output(volcengine.ebs.Volumes({
 *     ids: ["vol-3tzg6y5imn3b9fchkedb"],
 * }));
 * ```
 */
export function volumes(args?: VolumesArgs, opts?: pulumi.InvokeOptions): Promise<VolumesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:ebs/volumes:Volumes", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "kind": args.kind,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "volumeName": args.volumeName,
        "volumeStatus": args.volumeStatus,
        "volumeType": args.volumeType,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking Volumes.
 */
export interface VolumesArgs {
    /**
     * A list of Volume IDs.
     */
    ids?: string[];
    /**
     * The Id of instance.
     */
    instanceId?: string;
    /**
     * The Kind of Volume.
     */
    kind?: string;
    /**
     * A Name Regex of Volume.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of Volume.
     */
    volumeName?: string;
    /**
     * The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
     */
    volumeStatus?: string;
    /**
     * The type of Volume.
     */
    volumeType?: string;
    /**
     * The Id of Zone.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by Volumes.
 */
export interface VolumesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly instanceId?: string;
    readonly kind?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of Volume query.
     */
    readonly totalCount: number;
    readonly volumeName?: string;
    readonly volumeStatus?: string;
    readonly volumeType?: string;
    /**
     * The collection of Volume query.
     */
    readonly volumes: outputs.ebs.VolumesVolume[];
    readonly zoneId?: string;
}

export function volumesOutput(args?: VolumesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VolumesResult> {
    return pulumi.output(args).apply(a => volumes(a, opts))
}

/**
 * A collection of arguments for invoking Volumes.
 */
export interface VolumesOutputArgs {
    /**
     * A list of Volume IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Id of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The Kind of Volume.
     */
    kind?: pulumi.Input<string>;
    /**
     * A Name Regex of Volume.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of Volume.
     */
    volumeName?: pulumi.Input<string>;
    /**
     * The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
     */
    volumeStatus?: pulumi.Input<string>;
    /**
     * The type of Volume.
     */
    volumeType?: pulumi.Input<string>;
    /**
     * The Id of Zone.
     */
    zoneId?: pulumi.Input<string>;
}
