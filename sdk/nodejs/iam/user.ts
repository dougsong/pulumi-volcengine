// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage iam user
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.iam.User("foo", {
 *     description: "test",
 *     displayName: "name",
 *     userName: "tf-test",
 * });
 * ```
 *
 * ## Import
 *
 * Iam user can be imported using the UserName, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:iam/user:User default user_name
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:iam/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The account id of the user.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The create date of the user.
     */
    public /*out*/ readonly createDate!: pulumi.Output<string>;
    /**
     * The description of the user.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the user.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The email of the user.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * The mobile phone of the user.
     */
    public readonly mobilePhone!: pulumi.Output<string | undefined>;
    /**
     * The trn of the user.
     */
    public /*out*/ readonly trn!: pulumi.Output<string>;
    /**
     * The update date of the user.
     */
    public /*out*/ readonly updateDate!: pulumi.Output<string>;
    /**
     * The name of the user.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["createDate"] = state ? state.createDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["mobilePhone"] = state ? state.mobilePhone : undefined;
            resourceInputs["trn"] = state ? state.trn : undefined;
            resourceInputs["updateDate"] = state ? state.updateDate : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["mobilePhone"] = args ? args.mobilePhone : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["createDate"] = undefined /*out*/;
            resourceInputs["trn"] = undefined /*out*/;
            resourceInputs["updateDate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The account id of the user.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The create date of the user.
     */
    createDate?: pulumi.Input<string>;
    /**
     * The description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the user.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The email of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * The mobile phone of the user.
     */
    mobilePhone?: pulumi.Input<string>;
    /**
     * The trn of the user.
     */
    trn?: pulumi.Input<string>;
    /**
     * The update date of the user.
     */
    updateDate?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the user.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The email of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * The mobile phone of the user.
     */
    mobilePhone?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    userName: pulumi.Input<string>;
}
