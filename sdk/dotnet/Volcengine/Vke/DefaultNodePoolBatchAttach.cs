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
    /// Provides a resource to manage vke default node pool batch attach
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
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-project1",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
    ///     {
    ///         SubnetName = "acc-subnet-test-2",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = "cn-beijing-a",
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SecurityGroupName = "acc-test-security-group2",
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Ecs.Instance("fooInstance", new()
    ///     {
    ///         ImageId = "image-ybqi99s7yq8rx7mnk44b",
    ///         InstanceType = "ecs.g1ie.large",
    ///         InstanceName = "acc-test-ecs-name2",
    ///         Password = "93f0cb0614Aab12",
    ///         InstanceChargeType = "PostPaid",
    ///         SystemVolumeType = "ESSD_PL0",
    ///         SystemVolumeSize = 40,
    ///         SubnetId = fooSubnet.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var fooCluster = new Volcengine.Vke.Cluster("fooCluster", new()
    ///     {
    ///         Description = "created by terraform",
    ///         DeleteProtectionEnabled = false,
    ///         ClusterConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigArgs
    ///         {
    ///             SubnetIds = new[]
    ///             {
    ///                 fooSubnet.Id,
    ///             },
    ///             ApiServerPublicAccessEnabled = true,
    ///             ApiServerPublicAccessConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
    ///             {
    ///                 PublicAccessNetworkConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
    ///                 {
    ///                     BillingType = "PostPaidByBandwidth",
    ///                     Bandwidth = 1,
    ///                 },
    ///             },
    ///             ResourcePublicAccessDefaultEnabled = true,
    ///         },
    ///         PodsConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigArgs
    ///         {
    ///             PodNetworkMode = "VpcCniShared",
    ///             VpcCniConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigVpcCniConfigArgs
    ///             {
    ///                 SubnetIds = new[]
    ///                 {
    ///                     fooSubnet.Id,
    ///                 },
    ///             },
    ///         },
    ///         ServicesConfig = new Volcengine.Vke.Inputs.ClusterServicesConfigArgs
    ///         {
    ///             ServiceCidrsv4s = new[]
    ///             {
    ///                 "172.30.0.0/18",
    ///             },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Vke.Inputs.ClusterTagArgs
    ///             {
    ///                 Key = "tf-k1",
    ///                 Value = "tf-v1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultDefaultNodePool = new Volcengine.Vke.DefaultNodePool("defaultDefaultNodePool", new()
    ///     {
    ///         ClusterId = fooCluster.Id,
    ///         NodeConfig = new Volcengine.Vke.Inputs.DefaultNodePoolNodeConfigArgs
    ///         {
    ///             Security = new Volcengine.Vke.Inputs.DefaultNodePoolNodeConfigSecurityArgs
    ///             {
    ///                 Login = new Volcengine.Vke.Inputs.DefaultNodePoolNodeConfigSecurityLoginArgs
    ///                 {
    ///                     Password = "amw4WTdVcTRJVVFsUXpVTw==",
    ///                 },
    ///                 SecurityGroupIds = new[]
    ///                 {
    ///                     fooSecurityGroup.Id,
    ///                 },
    ///                 SecurityStrategies = new[]
    ///                 {
    ///                     "Hids",
    ///                 },
    ///             },
    ///             InitializeScript = "ISMvYmluL2Jhc2gKZWNobyAx",
    ///         },
    ///         KubernetesConfig = new Volcengine.Vke.Inputs.DefaultNodePoolKubernetesConfigArgs
    ///         {
    ///             Labels = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "tf-key1",
    ///                     Value = "tf-value1",
    ///                 },
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "tf-key2",
    ///                     Value = "tf-value2",
    ///                 },
    ///             },
    ///             Taints = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolKubernetesConfigTaintArgs
    ///                 {
    ///                     Key = "tf-key3",
    ///                     Value = "tf-value3",
    ///                     Effect = "NoSchedule",
    ///                 },
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolKubernetesConfigTaintArgs
    ///                 {
    ///                     Key = "tf-key4",
    ///                     Value = "tf-value4",
    ///                     Effect = "NoSchedule",
    ///                 },
    ///             },
    ///             Cordon = true,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Vke.Inputs.DefaultNodePoolTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultDefaultNodePoolBatchAttach = new Volcengine.Vke.DefaultNodePoolBatchAttach("defaultDefaultNodePoolBatchAttach", new()
    ///     {
    ///         ClusterId = fooCluster.Id,
    ///         DefaultNodePoolId = defaultDefaultNodePool.Id,
    ///         Instances = new[]
    ///         {
    ///             new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachInstanceArgs
    ///             {
    ///                 InstanceId = fooInstance.Id,
    ///                 KeepInstanceName = true,
    ///                 AdditionalContainerStorageEnabled = false,
    ///             },
    ///         },
    ///         KubernetesConfig = new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachKubernetesConfigArgs
    ///         {
    ///             Labels = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "tf-key1",
    ///                     Value = "tf-value1",
    ///                 },
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs
    ///                 {
    ///                     Key = "tf-key2",
    ///                     Value = "tf-value2",
    ///                 },
    ///             },
    ///             Taints = new[]
    ///             {
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs
    ///                 {
    ///                     Key = "tf-key3",
    ///                     Value = "tf-value3",
    ///                     Effect = "NoSchedule",
    ///                 },
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs
    ///                 {
    ///                     Key = "tf-key4",
    ///                     Value = "tf-value4",
    ///                     Effect = "NoSchedule",
    ///                 },
    ///             },
    ///             Cordon = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach")]
    public partial class DefaultNodePoolBatchAttach : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Output("defaultNodePoolId")]
        public Output<string> DefaultNodePoolId { get; private set; } = null!;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachInstance>> Instances { get; private set; } = null!;

        /// <summary>
        /// Is import of the DefaultNodePool. It only works when imported, set to true.
        /// </summary>
        [Output("isImport")]
        public Output<bool> IsImport { get; private set; } = null!;

        /// <summary>
        /// The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        /// </summary>
        [Output("kubernetesConfig")]
        public Output<Outputs.DefaultNodePoolBatchAttachKubernetesConfig?> KubernetesConfig { get; private set; } = null!;

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        [Output("nodeConfigs")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachNodeConfig>> NodeConfigs { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultNodePoolBatchAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultNodePoolBatchAttach(string name, DefaultNodePoolBatchAttachArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, args ?? new DefaultNodePoolBatchAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultNodePoolBatchAttach(string name, Input<string> id, DefaultNodePoolBatchAttachState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultNodePoolBatchAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultNodePoolBatchAttach Get(string name, Input<string> id, DefaultNodePoolBatchAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultNodePoolBatchAttach(name, id, state, options);
        }
    }

    public sealed class DefaultNodePoolBatchAttachArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Input("defaultNodePoolId", required: true)]
        public Input<string> DefaultNodePoolId { get; set; } = null!;

        [Input("instances")]
        private InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs>? _instances;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        /// </summary>
        [Input("kubernetesConfig")]
        public Input<Inputs.DefaultNodePoolBatchAttachKubernetesConfigArgs>? KubernetesConfig { get; set; }

        public DefaultNodePoolBatchAttachArgs()
        {
        }
        public static new DefaultNodePoolBatchAttachArgs Empty => new DefaultNodePoolBatchAttachArgs();
    }

    public sealed class DefaultNodePoolBatchAttachState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Input("defaultNodePoolId")]
        public Input<string>? DefaultNodePoolId { get; set; }

        [Input("instances")]
        private InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs>? _instances;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Is import of the DefaultNodePool. It only works when imported, set to true.
        /// </summary>
        [Input("isImport")]
        public Input<bool>? IsImport { get; set; }

        /// <summary>
        /// The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        /// </summary>
        [Input("kubernetesConfig")]
        public Input<Inputs.DefaultNodePoolBatchAttachKubernetesConfigGetArgs>? KubernetesConfig { get; set; }

        [Input("nodeConfigs")]
        private InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs>? _nodeConfigs;

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs> NodeConfigs
        {
            get => _nodeConfigs ?? (_nodeConfigs = new InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs>());
            set => _nodeConfigs = value;
        }

        [Input("tags")]
        private InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs>());
            set => _tags = value;
        }

        public DefaultNodePoolBatchAttachState()
        {
        }
        public static new DefaultNodePoolBatchAttachState Empty => new DefaultNodePoolBatchAttachState();
    }
}
