// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of network interfaces
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.NetworkInterfaces({
 *     ids: ["eni-2744htx2w0j5s7fap8t3ivwze"],
 * });
 * ```
 */
export function networkInterfaces(args?: NetworkInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<NetworkInterfacesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/networkInterfaces:NetworkInterfaces", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "networkInterfaceIds": args.networkInterfaceIds,
        "networkInterfaceName": args.networkInterfaceName,
        "outputFile": args.outputFile,
        "primaryIpAddresses": args.primaryIpAddresses,
        "privateIpAddresses": args.privateIpAddresses,
        "projectName": args.projectName,
        "securityGroupId": args.securityGroupId,
        "status": args.status,
        "subnetId": args.subnetId,
        "tags": args.tags,
        "type": args.type,
        "vpcId": args.vpcId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking NetworkInterfaces.
 */
export interface NetworkInterfacesArgs {
    /**
     * A list of ENI ids.
     */
    ids?: string[];
    /**
     * An id of the instance to which the ENI is bound.
     */
    instanceId?: string;
    /**
     * A list of network interface ids.
     */
    networkInterfaceIds?: string[];
    /**
     * A name of ENI.
     */
    networkInterfaceName?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A list of primary IP address of ENI.
     */
    primaryIpAddresses?: string[];
    /**
     * A list of private IP addresses.
     */
    privateIpAddresses?: string[];
    /**
     * The ProjectName of the ENI.
     */
    projectName?: string;
    /**
     * An id of the security group to which the secondary ENI belongs.
     */
    securityGroupId?: string;
    /**
     * A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
     */
    status?: string;
    /**
     * An id of the subnet to which the ENI is connected.
     */
    subnetId?: string;
    /**
     * Tags.
     */
    tags?: inputs.vpc.NetworkInterfacesTag[];
    /**
     * A type of ENI, Optional choice contains `primary`, `secondary`.
     */
    type?: string;
    /**
     * An id of the virtual private cloud (VPC) to which the ENI belongs.
     */
    vpcId?: string;
    /**
     * The zone ID.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by NetworkInterfaces.
 */
export interface NetworkInterfacesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly instanceId?: string;
    readonly networkInterfaceIds?: string[];
    /**
     * The name of the ENI.
     */
    readonly networkInterfaceName?: string;
    /**
     * The collection of ENI.
     */
    readonly networkInterfaces: outputs.vpc.NetworkInterfacesNetworkInterface[];
    readonly outputFile?: string;
    readonly primaryIpAddresses?: string[];
    readonly privateIpAddresses?: string[];
    /**
     * The ProjectName of the ENI.
     */
    readonly projectName?: string;
    readonly securityGroupId?: string;
    /**
     * The status of the ENI.
     */
    readonly status?: string;
    /**
     * The id of the subnet to which the ENI is connected.
     */
    readonly subnetId?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.vpc.NetworkInterfacesTag[];
    /**
     * The total count of ENI query.
     */
    readonly totalCount: number;
    /**
     * The type of the ENI.
     */
    readonly type?: string;
    /**
     * The id of the virtual private cloud (VPC) to which the ENI belongs.
     */
    readonly vpcId?: string;
    /**
     * The zone id of the ENI.
     */
    readonly zoneId?: string;
}
/**
 * Use this data source to query detailed information of network interfaces
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.NetworkInterfaces({
 *     ids: ["eni-2744htx2w0j5s7fap8t3ivwze"],
 * });
 * ```
 */
export function networkInterfacesOutput(args?: NetworkInterfacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<NetworkInterfacesResult> {
    return pulumi.output(args).apply((a: any) => networkInterfaces(a, opts))
}

/**
 * A collection of arguments for invoking NetworkInterfaces.
 */
export interface NetworkInterfacesOutputArgs {
    /**
     * A list of ENI ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An id of the instance to which the ENI is bound.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * A list of network interface ids.
     */
    networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A name of ENI.
     */
    networkInterfaceName?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A list of primary IP address of ENI.
     */
    primaryIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of private IP addresses.
     */
    privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ProjectName of the ENI.
     */
    projectName?: pulumi.Input<string>;
    /**
     * An id of the security group to which the secondary ENI belongs.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * An id of the subnet to which the ENI is connected.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkInterfacesTagArgs>[]>;
    /**
     * A type of ENI, Optional choice contains `primary`, `secondary`.
     */
    type?: pulumi.Input<string>;
    /**
     * An id of the virtual private cloud (VPC) to which the ENI belongs.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The zone ID.
     */
    zoneId?: pulumi.Input<string>;
}
