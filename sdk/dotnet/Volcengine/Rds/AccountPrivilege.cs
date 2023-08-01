// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Rds
{
    /// <summary>
    /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds account privilege
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var appName = new Volcengine.Rds.Account("appName", new Volcengine.Rds.AccountArgs
    ///         {
    ///             InstanceId = "mysql-0fdd3bab2e7c",
    ///             AccountName = "terraform-test-app",
    ///             AccountPassword = "Aatest123",
    ///             AccountType = "Normal",
    ///         });
    ///         var foo = new Volcengine.Rds.AccountPrivilege("foo", new Volcengine.Rds.AccountPrivilegeArgs
    ///         {
    ///             InstanceId = "mysql-0fdd3bab2e7c",
    ///             AccountName = appName.AccountName,
    ///             DbPrivileges = 
    ///             {
    ///                 new Volcengine.Rds.Inputs.AccountPrivilegeDbPrivilegeArgs
    ///                 {
    ///                     DbName = "foo",
    ///                     AccountPrivilege = "Custom",
    ///                     AccountPrivilegeStr = "ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES",
    ///                 },
    ///                 new Volcengine.Rds.Inputs.AccountPrivilegeDbPrivilegeArgs
    ///                 {
    ///                     DbName = "bar",
    ///                     AccountPrivilege = "DDLOnly",
    ///                 },
    ///                 new Volcengine.Rds.Inputs.AccountPrivilegeDbPrivilegeArgs
    ///                 {
    ///                     DbName = "demo",
    ///                     AccountPrivilege = "ReadWrite",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS account privilege can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:rds/accountPrivilege:AccountPrivilege default mysql-42b38c769c4b:account_name
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds/accountPrivilege:AccountPrivilege")]
    public partial class AccountPrivilege : Pulumi.CustomResource
    {
        /// <summary>
        /// Database account name. The rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The privileges of the account.
        /// </summary>
        [Output("dbPrivileges")]
        public Output<ImmutableArray<Outputs.AccountPrivilegeDbPrivilege>> DbPrivileges { get; private set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPrivilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPrivilege(string name, AccountPrivilegeArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds/accountPrivilege:AccountPrivilege", name, args ?? new AccountPrivilegeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPrivilege(string name, Input<string> id, AccountPrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds/accountPrivilege:AccountPrivilege", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccountPrivilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPrivilege Get(string name, Input<string> id, AccountPrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPrivilege(name, id, state, options);
        }
    }

    public sealed class AccountPrivilegeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database account name. The rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("dbPrivileges", required: true)]
        private InputList<Inputs.AccountPrivilegeDbPrivilegeArgs>? _dbPrivileges;

        /// <summary>
        /// The privileges of the account.
        /// </summary>
        public InputList<Inputs.AccountPrivilegeDbPrivilegeArgs> DbPrivileges
        {
            get => _dbPrivileges ?? (_dbPrivileges = new InputList<Inputs.AccountPrivilegeDbPrivilegeArgs>());
            set => _dbPrivileges = value;
        }

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AccountPrivilegeArgs()
        {
        }
    }

    public sealed class AccountPrivilegeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database account name. The rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("dbPrivileges")]
        private InputList<Inputs.AccountPrivilegeDbPrivilegeGetArgs>? _dbPrivileges;

        /// <summary>
        /// The privileges of the account.
        /// </summary>
        public InputList<Inputs.AccountPrivilegeDbPrivilegeGetArgs> DbPrivileges
        {
            get => _dbPrivileges ?? (_dbPrivileges = new InputList<Inputs.AccountPrivilegeDbPrivilegeGetArgs>());
            set => _dbPrivileges = value;
        }

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AccountPrivilegeState()
        {
        }
    }
}
