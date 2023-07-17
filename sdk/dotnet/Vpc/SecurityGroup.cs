// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage security group
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
    ///         var g1test1 = new Volcengine.Vpc.SecurityGroup("g1test1", new Volcengine.Vpc.SecurityGroupArgs
    ///         {
    ///             ProjectName = "yuwenhao",
    ///             VpcId = "vpc-2feppmy1ugt1c59gp688n1fld",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SecurityGroup can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpc/securityGroup:SecurityGroup default sg-273ycgql3ig3k7fap8t3dyvqx
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpc/securityGroup:SecurityGroup")]
    public partial class SecurityGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of SecurityGroup.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Description of SecurityGroup.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Name of SecurityGroup.
        /// </summary>
        [Output("securityGroupName")]
        public Output<string> SecurityGroupName { get; private set; } = null!;

        /// <summary>
        /// Status of SecurityGroup.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.SecurityGroupTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroup(string name, SecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpc/securityGroup:SecurityGroup", name, args ?? new SecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroup(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpc/securityGroup:SecurityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroup Get(string name, Input<string> id, SecurityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroup(name, id, state, options);
        }
    }

    public sealed class SecurityGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of SecurityGroup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Name of SecurityGroup.
        /// </summary>
        [Input("securityGroupName")]
        public Input<string>? SecurityGroupName { get; set; }

        [Input("tags")]
        private InputList<Inputs.SecurityGroupTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.SecurityGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.SecurityGroupTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Id of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public SecurityGroupArgs()
        {
        }
    }

    public sealed class SecurityGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time of SecurityGroup.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// Description of SecurityGroup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Name of SecurityGroup.
        /// </summary>
        [Input("securityGroupName")]
        public Input<string>? SecurityGroupName { get; set; }

        /// <summary>
        /// Status of SecurityGroup.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.SecurityGroupTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.SecurityGroupTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.SecurityGroupTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SecurityGroupState()
        {
        }
    }
}
