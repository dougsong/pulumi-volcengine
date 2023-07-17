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
    /// Provides a resource to manage iam user policy attachment
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
    ///         var user = new Volcengine.Iam.User("user", new Volcengine.Iam.UserArgs
    ///         {
    ///             UserName = "TfTest",
    ///             Description = "test",
    ///         });
    ///         var policy = new Volcengine.Iam.Policy("policy", new Volcengine.Iam.PolicyArgs
    ///         {
    ///             PolicyName = "TerraformResourceTest1",
    ///             Description = "created by terraform 1",
    ///             PolicyDocument = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}",
    ///         });
    ///         var foo = new Volcengine.Iam.UserPolicyAttachment("foo", new Volcengine.Iam.UserPolicyAttachmentArgs
    ///         {
    ///             UserName = user.UserName,
    ///             PolicyName = policy.PolicyName,
    ///             PolicyType = policy.PolicyType,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iam user policy attachment can be imported using the UserName:PolicyName:PolicyType, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:iam/userPolicyAttachment:UserPolicyAttachment default TerraformTestUser:TerraformTestPolicy:Custom
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:iam/userPolicyAttachment:UserPolicyAttachment")]
    public partial class UserPolicyAttachment : Pulumi.CustomResource
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
        /// The name of the user.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a UserPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPolicyAttachment(string name, UserPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("volcengine:iam/userPolicyAttachment:UserPolicyAttachment", name, args ?? new UserPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPolicyAttachment(string name, Input<string> id, UserPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:iam/userPolicyAttachment:UserPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPolicyAttachment Get(string name, Input<string> id, UserPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new UserPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class UserPolicyAttachmentArgs : Pulumi.ResourceArgs
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
        /// The name of the user.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserPolicyAttachmentArgs()
        {
        }
    }

    public sealed class UserPolicyAttachmentState : Pulumi.ResourceArgs
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
        /// The name of the user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UserPolicyAttachmentState()
        {
        }
    }
}
