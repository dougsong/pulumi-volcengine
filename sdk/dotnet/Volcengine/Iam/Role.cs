// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Iam
{
    /// <summary>
    /// Provides a resource to manage iam role
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
    ///         var foo = new Volcengine.Iam.Role("foo", new Volcengine.Iam.RoleArgs
    ///         {
    ///             Description = "created by terraform",
    ///             DisplayName = "terraform role",
    ///             MaxSessionDuration = 43200,
    ///             RoleName = "TerraformTestRole",
    ///             TrustPolicyDocument = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iam role can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:iam/role:Role default TerraformTestRole
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:iam/role:Role")]
    public partial class Role : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the Role.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The max session duration of the Role.
        /// </summary>
        [Output("maxSessionDuration")]
        public Output<int?> MaxSessionDuration { get; private set; } = null!;

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Role.
        /// </summary>
        [Output("trn")]
        public Output<string> Trn { get; private set; } = null!;

        /// <summary>
        /// The trust policy document of the Role.
        /// </summary>
        [Output("trustPolicyDocument")]
        public Output<string> TrustPolicyDocument { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("volcengine:iam/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:iam/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Role.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The max session duration of the Role.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The trust policy document of the Role.
        /// </summary>
        [Input("trustPolicyDocument", required: true)]
        public Input<string> TrustPolicyDocument { get; set; } = null!;

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Role.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The max session duration of the Role.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The resource name of the Role.
        /// </summary>
        [Input("trn")]
        public Input<string>? Trn { get; set; }

        /// <summary>
        /// The trust policy document of the Role.
        /// </summary>
        [Input("trustPolicyDocument")]
        public Input<string>? TrustPolicyDocument { get; set; }

        public RoleState()
        {
        }
    }
}
