// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tls hosts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultHosts = pulumi.output(volcengine.tls.Hosts({
 *     hostGroupId: "527102e2-1e4f-45f4-a990-751152125da7",
 * }));
 * ```
 */
export function hosts(args: HostsArgs, opts?: pulumi.InvokeOptions): Promise<HostsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:tls/hosts:Hosts", {
        "heartbeatStatus": args.heartbeatStatus,
        "hostGroupId": args.hostGroupId,
        "ip": args.ip,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Hosts.
 */
export interface HostsArgs {
    /**
     * The the heartbeat status.
     */
    heartbeatStatus?: number;
    /**
     * The id of host group.
     */
    hostGroupId: string;
    /**
     * The ip address.
     */
    ip?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Hosts.
 */
export interface HostsResult {
    /**
     * The the heartbeat status.
     */
    readonly heartbeatStatus?: number;
    /**
     * The id of host group.
     */
    readonly hostGroupId: string;
    /**
     * The collection of query.
     */
    readonly hostInfos: outputs.tls.HostsHostInfo[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ip address.
     */
    readonly ip?: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}

export function hostsOutput(args: HostsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<HostsResult> {
    return pulumi.output(args).apply(a => hosts(a, opts))
}

/**
 * A collection of arguments for invoking Hosts.
 */
export interface HostsOutputArgs {
    /**
     * The the heartbeat status.
     */
    heartbeatStatus?: pulumi.Input<number>;
    /**
     * The id of host group.
     */
    hostGroupId: pulumi.Input<string>;
    /**
     * The ip address.
     */
    ip?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
