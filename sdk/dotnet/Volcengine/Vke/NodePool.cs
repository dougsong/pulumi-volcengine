// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke
{
    /// <summary>
    /// Provides a resource to manage vke node pool
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
    ///     var vkeTest = new Volcengine.Vke.NodePool("vkeTest", new()
    ///     {
    ///         AutoScaling = new Volcengine.Vke.Inputs.NodePoolAutoScalingArgs
    ///         {
    ///             Enabled = true,
    ///             SubnetPolicy = "ZoneBalance",
    ///         },
    ///         ClusterId = "ccgd6066rsfegs2dkhlog",
    ///         KubernetesConfig = new Volcengine.Vke.Inputs.NodePoolKubernetesConfigArgs
    ///         {
    ///             Cordon = false,
    ///             Labels = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.NodePoolKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "aa",
    ///                     Value = "bb",
    ///                 },
    ///                 new Volcengine.Vke.Inputs.NodePoolKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "cccc",
    ///                     Value = "dddd",
    ///                 },
    ///             },
    ///         },
    ///         NodeConfig = new Volcengine.Vke.Inputs.NodePoolNodeConfigArgs
    ///         {
    ///             DataVolumes = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.NodePoolNodeConfigDataVolumeArgs
    ///                 {
    ///                     Size = 60,
    ///                     Type = "ESSD_PL0",
    ///                 },
    ///             },
    ///             EcsTags = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.NodePoolNodeConfigEcsTagArgs
    ///                 {
    ///                     Key = "ecs_k1",
    ///                     Value = "ecs_v1",
    ///                 },
    ///             },
    ///             InstanceChargeType = "PostPaid",
    ///             InstanceTypeIds = new[]
    ///             {
    ///                 "ecs.g1ie.xlarge",
    ///             },
    ///             Period = 1,
    ///             Security = new Volcengine.Vke.Inputs.NodePoolNodeConfigSecurityArgs
    ///             {
    ///                 Login = new Volcengine.Vke.Inputs.NodePoolNodeConfigSecurityLoginArgs
    ///                 {
    ///                     Password = "UHdkMTIzNDU2",
    ///                 },
    ///                 SecurityGroupIds = new[]
    ///                 {
    ///                     "sg-13fbyz0sok3y83n6nu4hv1q10",
    ///                     "sg-mj1e9tbztgqo5smt1ah8l4bh",
    ///                 },
    ///             },
    ///             SubnetIds = new[]
    ///             {
    ///                 "subnet-mj1e9jgu96v45smt1a674x3h",
    ///             },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Vke.Inputs.NodePoolTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NodePool can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vke/nodePool:NodePool default pcabe57vqtofgrbln3dp0
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vke/nodePool:NodePool")]
    public partial class NodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The node pool elastic scaling configuration information.
        /// </summary>
        [Output("autoScaling")]
        public Output<Outputs.NodePoolAutoScaling> AutoScaling { get; private set; } = null!;

        /// <summary>
        /// The ClientToken of NodePool.
        /// </summary>
        [Output("clientToken")]
        public Output<string?> ClientToken { get; private set; } = null!;

        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Output("clusterId")]
        public Output<string?> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The KubernetesConfig of NodeConfig.
        /// </summary>
        [Output("kubernetesConfig")]
        public Output<Outputs.NodePoolKubernetesConfig> KubernetesConfig { get; private set; } = null!;

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.NodePoolNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.NodePoolTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodePool(string name, NodePoolArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vke/nodePool:NodePool", name, args ?? new NodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodePool(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vke/nodePool:NodePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodePool Get(string name, Input<string> id, NodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new NodePool(name, id, state, options);
        }
    }

    public sealed class NodePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The node pool elastic scaling configuration information.
        /// </summary>
        [Input("autoScaling")]
        public Input<Inputs.NodePoolAutoScalingArgs>? AutoScaling { get; set; }

        /// <summary>
        /// The ClientToken of NodePool.
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The KubernetesConfig of NodeConfig.
        /// </summary>
        [Input("kubernetesConfig", required: true)]
        public Input<Inputs.NodePoolKubernetesConfigArgs> KubernetesConfig { get; set; } = null!;

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        [Input("nodeConfig", required: true)]
        public Input<Inputs.NodePoolNodeConfigArgs> NodeConfig { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.NodePoolTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.NodePoolTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NodePoolTagArgs>());
            set => _tags = value;
        }

        public NodePoolArgs()
        {
        }
        public static new NodePoolArgs Empty => new NodePoolArgs();
    }

    public sealed class NodePoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The node pool elastic scaling configuration information.
        /// </summary>
        [Input("autoScaling")]
        public Input<Inputs.NodePoolAutoScalingGetArgs>? AutoScaling { get; set; }

        /// <summary>
        /// The ClientToken of NodePool.
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The KubernetesConfig of NodeConfig.
        /// </summary>
        [Input("kubernetesConfig")]
        public Input<Inputs.NodePoolKubernetesConfigGetArgs>? KubernetesConfig { get; set; }

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.NodePoolNodeConfigGetArgs>? NodeConfig { get; set; }

        [Input("tags")]
        private InputList<Inputs.NodePoolTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.NodePoolTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NodePoolTagGetArgs>());
            set => _tags = value;
        }

        public NodePoolState()
        {
        }
        public static new NodePoolState Empty => new NodePoolState();
    }
}
