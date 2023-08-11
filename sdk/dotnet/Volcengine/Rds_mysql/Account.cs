// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds_mysql
{
    /// <summary>
    /// Provides a resource to manage rds mysql account
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Volcengine.Rds_mysql.Account("default", new()
    ///     {
    ///         AccountName = "test",
    ///         AccountPassword = "xdjsuiahHUH@",
    ///         AccountType = "Normal",
    ///         InstanceId = "mysql-e9293705eed6",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS mysql account can be imported using the instance_id:account_name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:rds_mysql/account:Account default mysql-42b38c769c4b:test
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mysql/account:Account")]
    public partial class Account : global::Pulumi.CustomResource
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
        /// The password of the database account.
        /// Illustrate:
        /// Cannot start with `!` or `@`.
        /// The length is 8~32 characters.
        /// It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        /// The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// The privilege information of account.
        /// </summary>
        [Output("accountPrivileges")]
        public Output<ImmutableArray<Outputs.AccountAccountPrivilege>> AccountPrivileges { get; private set; } = null!;

        /// <summary>
        /// Database account type, value:
        /// Super: A high-privilege account. Only one database account can be created for an instance.
        /// Normal: An account with ordinary privileges.
        /// </summary>
        [Output("accountType")]
        public Output<string> AccountType { get; private set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/account:Account", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
                AdditionalSecretOutputs =
                {
                    "accountPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
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

        [Input("accountPassword", required: true)]
        private Input<string>? _accountPassword;

        /// <summary>
        /// The password of the database account.
        /// Illustrate:
        /// Cannot start with `!` or `@`.
        /// The length is 8~32 characters.
        /// It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        /// The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.AccountAccountPrivilegeArgs>? _accountPrivileges;

        /// <summary>
        /// The privilege information of account.
        /// </summary>
        public InputList<Inputs.AccountAccountPrivilegeArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.AccountAccountPrivilegeArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// Database account type, value:
        /// Super: A high-privilege account. Only one database account can be created for an instance.
        /// Normal: An account with ordinary privileges.
        /// </summary>
        [Input("accountType", required: true)]
        public Input<string> AccountType { get; set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }

    public sealed class AccountState : global::Pulumi.ResourceArgs
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

        [Input("accountPassword")]
        private Input<string>? _accountPassword;

        /// <summary>
        /// The password of the database account.
        /// Illustrate:
        /// Cannot start with `!` or `@`.
        /// The length is 8~32 characters.
        /// It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        /// The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.AccountAccountPrivilegeGetArgs>? _accountPrivileges;

        /// <summary>
        /// The privilege information of account.
        /// </summary>
        public InputList<Inputs.AccountAccountPrivilegeGetArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.AccountAccountPrivilegeGetArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// Database account type, value:
        /// Super: A high-privilege account. Only one database account can be created for an instance.
        /// Normal: An account with ordinary privileges.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AccountState()
        {
        }
        public static new AccountState Empty => new AccountState();
    }
}
