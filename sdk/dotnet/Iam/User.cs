// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam
{
    /// <summary>
    /// Provides a resource to manage iam user
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Iam.User("foo", new Volcengine.Iam.UserArgs
    ///         {
    ///             Description = "test",
    ///             DisplayName = "name",
    ///             UserName = "tf-test",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iam user can be imported using the UserName, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:Iam/user:User default user_name
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:Iam/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// The account id of the user.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The create date of the user.
        /// </summary>
        [Output("createDate")]
        public Output<string> CreateDate { get; private set; } = null!;

        /// <summary>
        /// The description of the user.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the user.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// The mobile phone of the user.
        /// </summary>
        [Output("mobilePhone")]
        public Output<string?> MobilePhone { get; private set; } = null!;

        /// <summary>
        /// The trn of the user.
        /// </summary>
        [Output("trn")]
        public Output<string> Trn { get; private set; } = null!;

        /// <summary>
        /// The update date of the user.
        /// </summary>
        [Output("updateDate")]
        public Output<string> UpdateDate { get; private set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("volcengine:Iam/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:Iam/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The mobile phone of the user.
        /// </summary>
        [Input("mobilePhone")]
        public Input<string>? MobilePhone { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id of the user.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The create date of the user.
        /// </summary>
        [Input("createDate")]
        public Input<string>? CreateDate { get; set; }

        /// <summary>
        /// The description of the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The email of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The mobile phone of the user.
        /// </summary>
        [Input("mobilePhone")]
        public Input<string>? MobilePhone { get; set; }

        /// <summary>
        /// The trn of the user.
        /// </summary>
        [Input("trn")]
        public Input<string>? Trn { get; set; }

        /// <summary>
        /// The update date of the user.
        /// </summary>
        [Input("updateDate")]
        public Input<string>? UpdateDate { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UserState()
        {
        }
    }
}
