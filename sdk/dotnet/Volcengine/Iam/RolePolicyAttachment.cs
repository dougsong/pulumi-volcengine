// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Iam
{
    /// <summary>
    /// Provides a resource to manage iam role policy attachment
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
    ///     var role = new Volcengine.Iam.Role("role", new()
    ///     {
    ///         RoleName = "TerraformTestRole",
    ///         DisplayName = "terraform role",
    ///         TrustPolicyDocument = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}",
    ///         Description = "created by terraform",
    ///         MaxSessionDuration = 43200,
    ///     });
    /// 
    ///     var policy = new Volcengine.Iam.Policy("policy", new()
    ///     {
    ///         PolicyName = "TerraformResourceTest1",
    ///         Description = "created by terraform 1",
    ///         PolicyDocument = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}",
    ///     });
    /// 
    ///     var foo = new Volcengine.Iam.RolePolicyAttachment("foo", new()
    ///     {
    ///         RoleName = role.Id,
    ///         PolicyName = policy.Id,
    ///         PolicyType = policy.PolicyType,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iam role policy attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:iam/rolePolicyAttachment:RolePolicyAttachment default TerraformTestRole:TerraformTestPolicy:Custom
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:iam/rolePolicyAttachment:RolePolicyAttachment")]
    public partial class RolePolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The type of the Policy.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;


        /// <summary>
        /// Create a RolePolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RolePolicyAttachment(string name, RolePolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("volcengine:iam/rolePolicyAttachment:RolePolicyAttachment", name, args ?? new RolePolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RolePolicyAttachment(string name, Input<string> id, RolePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:iam/rolePolicyAttachment:RolePolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RolePolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RolePolicyAttachment Get(string name, Input<string> id, RolePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new RolePolicyAttachment(name, id, state, options);
        }
    }

    public sealed class RolePolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The type of the Policy.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<string> PolicyType { get; set; } = null!;

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public RolePolicyAttachmentArgs()
        {
        }
        public static new RolePolicyAttachmentArgs Empty => new RolePolicyAttachmentArgs();
    }

    public sealed class RolePolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Policy.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The type of the Policy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// The name of the Role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        public RolePolicyAttachmentState()
        {
        }
        public static new RolePolicyAttachmentState Empty => new RolePolicyAttachmentState();
    }
}
