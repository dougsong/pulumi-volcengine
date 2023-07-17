// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis accounts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultAccounts = pulumi.output(volcengine.redis.Accounts({
 *     instanceId: "redis-cn0398aizj8cwmopx",
 * }));
 * ```
 */
export function accounts(args: AccountsArgs, opts?: pulumi.InvokeOptions): Promise<AccountsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:redis/accounts:Accounts", {
        "accountName": args.accountName,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsArgs {
    /**
     * The name of the redis account.
     */
    accountName?: string;
    /**
     * The id of the Redis instance.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Accounts.
 */
export interface AccountsResult {
    /**
     * The name of the redis account.
     */
    readonly accountName?: string;
    /**
     * The collection of redis instance account query.
     */
    readonly accounts: outputs.redis.AccountsAccount[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The id of instance.
     */
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The total count of redis accounts query.
     */
    readonly totalCount: number;
}

export function accountsOutput(args: AccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AccountsResult> {
    return pulumi.output(args).apply(a => accounts(a, opts))
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsOutputArgs {
    /**
     * The name of the redis account.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The id of the Redis instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
