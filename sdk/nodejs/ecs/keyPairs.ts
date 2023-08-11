// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ecs key pairs
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooKeyPair = new volcengine.ecs.KeyPair("fooKeyPair", {
 *     keyPairName: "acc-test-key-name",
 *     description: "acc-test",
 * });
 * const fooKeyPairs = volcengine.ecs.KeyPairsOutput({
 *     keyPairName: fooKeyPair.keyPairName,
 * });
 * ```
 */
export function keyPairs(args?: KeyPairsArgs, opts?: pulumi.InvokeOptions): Promise<KeyPairsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:ecs/keyPairs:KeyPairs", {
        "fingerPrint": args.fingerPrint,
        "keyPairIds": args.keyPairIds,
        "keyPairName": args.keyPairName,
        "keyPairNames": args.keyPairNames,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking KeyPairs.
 */
export interface KeyPairsArgs {
    /**
     * The finger print info.
     */
    fingerPrint?: string;
    /**
     * Ids of key pair.
     */
    keyPairIds?: string[];
    /**
     * Name of key pair.
     */
    keyPairName?: string;
    /**
     * Key pair names info.
     */
    keyPairNames?: string[];
    /**
     * A Name Regex of ECS key pairs.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by KeyPairs.
 */
export interface KeyPairsResult {
    /**
     * The finger print info.
     */
    readonly fingerPrint?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyPairIds?: string[];
    /**
     * The name of key pair.
     */
    readonly keyPairName?: string;
    readonly keyPairNames?: string[];
    /**
     * The target query key pairs info.
     */
    readonly keyPairs: outputs.ecs.KeyPairsKeyPair[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of ECS key pair query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of ecs key pairs
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooKeyPair = new volcengine.ecs.KeyPair("fooKeyPair", {
 *     keyPairName: "acc-test-key-name",
 *     description: "acc-test",
 * });
 * const fooKeyPairs = volcengine.ecs.KeyPairsOutput({
 *     keyPairName: fooKeyPair.keyPairName,
 * });
 * ```
 */
export function keyPairsOutput(args?: KeyPairsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<KeyPairsResult> {
    return pulumi.output(args).apply((a: any) => keyPairs(a, opts))
}

/**
 * A collection of arguments for invoking KeyPairs.
 */
export interface KeyPairsOutputArgs {
    /**
     * The finger print info.
     */
    fingerPrint?: pulumi.Input<string>;
    /**
     * Ids of key pair.
     */
    keyPairIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of key pair.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * Key pair names info.
     */
    keyPairNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of ECS key pairs.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
