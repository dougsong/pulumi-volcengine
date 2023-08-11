// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds account privilege
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const appName = new volcengine.rds.Account("appName", {
 *     instanceId: "mysql-0fdd3bab2e7c",
 *     accountName: "terraform-test-app",
 *     accountPassword: "Aatest123",
 *     accountType: "Normal",
 * });
 * const foo = new volcengine.rds.AccountPrivilege("foo", {
 *     instanceId: "mysql-0fdd3bab2e7c",
 *     accountName: appName.accountName,
 *     dbPrivileges: [
 *         {
 *             dbName: "foo",
 *             accountPrivilege: "Custom",
 *             accountPrivilegeStr: "ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES",
 *         },
 *         {
 *             dbName: "bar",
 *             accountPrivilege: "DDLOnly",
 *         },
 *         {
 *             dbName: "demo",
 *             accountPrivilege: "ReadWrite",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * RDS account privilege can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:rds/accountPrivilege:AccountPrivilege default mysql-42b38c769c4b:account_name
 * ```
 */
export class AccountPrivilege extends pulumi.CustomResource {
    /**
     * Get an existing AccountPrivilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountPrivilegeState, opts?: pulumi.CustomResourceOptions): AccountPrivilege {
        return new AccountPrivilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds/accountPrivilege:AccountPrivilege';

    /**
     * Returns true if the given object is an instance of AccountPrivilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountPrivilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountPrivilege.__pulumiType;
    }

    /**
     * Database account name. The rules are as follows:
     * Unique name.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, or underscores (_).
     * The length is 2~32 characters.
     * The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The privileges of the account.
     */
    public readonly dbPrivileges!: pulumi.Output<outputs.rds.AccountPrivilegeDbPrivilege[]>;
    /**
     * The ID of the RDS instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a AccountPrivilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountPrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountPrivilegeArgs | AccountPrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountPrivilegeState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["dbPrivileges"] = state ? state.dbPrivileges : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as AccountPrivilegeArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.dbPrivileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbPrivileges'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["dbPrivileges"] = args ? args.dbPrivileges : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountPrivilege.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountPrivilege resources.
 */
export interface AccountPrivilegeState {
    /**
     * Database account name. The rules are as follows:
     * Unique name.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, or underscores (_).
     * The length is 2~32 characters.
     * The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The privileges of the account.
     */
    dbPrivileges?: pulumi.Input<pulumi.Input<inputs.rds.AccountPrivilegeDbPrivilege>[]>;
    /**
     * The ID of the RDS instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountPrivilege resource.
 */
export interface AccountPrivilegeArgs {
    /**
     * Database account name. The rules are as follows:
     * Unique name.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, or underscores (_).
     * The length is 2~32 characters.
     * The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
     */
    accountName: pulumi.Input<string>;
    /**
     * The privileges of the account.
     */
    dbPrivileges: pulumi.Input<pulumi.Input<inputs.rds.AccountPrivilegeDbPrivilege>[]>;
    /**
     * The ID of the RDS instance.
     */
    instanceId: pulumi.Input<string>;
}
