// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bioos
{
    /// <summary>
    /// Provides a resource to manage bioos cluster bind
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
    ///         var example = new Volcengine.Bioos.ClusterBind("example", new Volcengine.Bioos.ClusterBindArgs
    ///         {
    ///             ClusterId = "ucfhp1nteig48u8ufv8s0",
    ///             Type = "workflow",
    ///             WorkspaceId = "wcfhp1vdeig48u8ufv8sg",
    ///         });
    ///         //必填
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cluster binder can be imported using the workspace id and cluster id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:bioos/clusterBind:ClusterBind default wc*****:uc***
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:bioos/clusterBind:ClusterBind")]
    public partial class ClusterBind : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The type of the cluster bind.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The id of the workspace.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterBind resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterBind(string name, ClusterBindArgs args, CustomResourceOptions? options = null)
            : base("volcengine:bioos/clusterBind:ClusterBind", name, args ?? new ClusterBindArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterBind(string name, Input<string> id, ClusterBindState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:bioos/clusterBind:ClusterBind", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterBind resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterBind Get(string name, Input<string> id, ClusterBindState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterBind(name, id, state, options);
        }
    }

    public sealed class ClusterBindArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The type of the cluster bind.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The id of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public ClusterBindArgs()
        {
        }
    }

    public sealed class ClusterBindState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The type of the cluster bind.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The id of the workspace.
        /// </summary>
        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public ClusterBindState()
        {
        }
    }
}
