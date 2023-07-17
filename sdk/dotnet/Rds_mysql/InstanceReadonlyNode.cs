// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql
{
    /// <summary>
    /// Provides a resource to manage rds mysql instance readonly node
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
    ///         var foo = new Volcengine.Rds_mysql.InstanceReadonlyNode("foo", new Volcengine.Rds_mysql.InstanceReadonlyNodeArgs
    ///         {
    ///             InstanceId = "mysql-b3fca7f571d6",
    ///             NodeSpec = "rds.mysql.1c2g",
    ///             ZoneId = "cn-guilin-b",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode")]
    public partial class InstanceReadonlyNode : Pulumi.CustomResource
    {
        /// <summary>
        /// The RDS mysql instance id of the readonly node.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The id of the readonly node.
        /// </summary>
        [Output("nodeId")]
        public Output<string> NodeId { get; private set; } = null!;

        /// <summary>
        /// The specification of readonly node.
        /// </summary>
        [Output("nodeSpec")]
        public Output<string> NodeSpec { get; private set; } = null!;

        /// <summary>
        /// The available zone of readonly node.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceReadonlyNode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceReadonlyNode(string name, InstanceReadonlyNodeArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode", name, args ?? new InstanceReadonlyNodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceReadonlyNode(string name, Input<string> id, InstanceReadonlyNodeState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceReadonlyNode resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceReadonlyNode Get(string name, Input<string> id, InstanceReadonlyNodeState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceReadonlyNode(name, id, state, options);
        }
    }

    public sealed class InstanceReadonlyNodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The RDS mysql instance id of the readonly node.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The specification of readonly node.
        /// </summary>
        [Input("nodeSpec", required: true)]
        public Input<string> NodeSpec { get; set; } = null!;

        /// <summary>
        /// The available zone of readonly node.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public InstanceReadonlyNodeArgs()
        {
        }
    }

    public sealed class InstanceReadonlyNodeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The RDS mysql instance id of the readonly node.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The id of the readonly node.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// The specification of readonly node.
        /// </summary>
        [Input("nodeSpec")]
        public Input<string>? NodeSpec { get; set; }

        /// <summary>
        /// The available zone of readonly node.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceReadonlyNodeState()
        {
        }
    }
}
