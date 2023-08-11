// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds accounts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.Accounts({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function accounts(args: AccountsArgs, opts?: pulumi.InvokeOptions): Promise<AccountsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds/accounts:Accounts", {
        "accountName": args.accountName,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsArgs {
    /**
     * The name of the database account.
     */
    accountName?: string;
    /**
     * The id of the RDS instance.
     */
    instanceId: string;
    /**
     * A Name Regex of database account.
     */
    nameRegex?: string;
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
     * The name of the database account.
     */
    readonly accountName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of RDS instance account query.
     */
    readonly rdsAccounts: outputs.rds.AccountsRdsAccount[];
    /**
     * The total count of database account query.
     */
    readonly totalCount: number;
}
/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds accounts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.Accounts({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function accountsOutput(args: AccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AccountsResult> {
    return pulumi.output(args).apply((a: any) => accounts(a, opts))
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsOutputArgs {
    /**
     * The name of the database account.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The id of the RDS instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A Name Regex of database account.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
