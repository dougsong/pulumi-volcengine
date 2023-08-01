// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vke
{
    /// <summary>
    /// Provides a resource to manage vke cluster
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
    ///         var foo = new Volcengine.Vke.Cluster("foo", new Volcengine.Vke.ClusterArgs
    ///         {
    ///             ClusterConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigArgs
    ///             {
    ///                 ApiServerPublicAccessConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
    ///                 {
    ///                     PublicAccessNetworkConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
    ///                     {
    ///                         Bandwidth = 1,
    ///                         BillingType = "PostPaidByBandwidth",
    ///                     },
    ///                 },
    ///                 ApiServerPublicAccessEnabled = true,
    ///                 ResourcePublicAccessDefaultEnabled = true,
    ///                 SubnetIds = 
    ///                 {
    ///                     "subnet-rrqvkt2nq1hcv0x57ccqf3x",
    ///                 },
    ///             },
    ///             DeleteProtectionEnabled = false,
    ///             Description = "created by terraform",
    ///             LoggingConfig = new Volcengine.Vke.Inputs.ClusterLoggingConfigArgs
    ///             {
    ///                 LogSetups = 
    ///                 {
    ///                     new Volcengine.Vke.Inputs.ClusterLoggingConfigLogSetupArgs
    ///                     {
    ///                         Enabled = false,
    ///                         LogTtl = 30,
    ///                         LogType = "Audit",
    ///                     },
    ///                 },
    ///             },
    ///             PodsConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigArgs
    ///             {
    ///                 PodNetworkMode = "VpcCniShared",
    ///                 VpcCniConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigVpcCniConfigArgs
    ///                 {
    ///                     SubnetIds = 
    ///                     {
    ///                         "subnet-rrqvkt2nq1hcv0x57ccqf3x",
    ///                         "subnet-miklcqh75vcw5smt1amo4ik5",
    ///                         "subnet-13g0x0ytpm0hs3n6nu5j591lv",
    ///                     },
    ///                 },
    ///             },
    ///             ServicesConfig = new Volcengine.Vke.Inputs.ClusterServicesConfigArgs
    ///             {
    ///                 ServiceCidrsv4s = 
    ///                 {
    ///                     "172.30.0.0/18",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 new Volcengine.Vke.Inputs.ClusterTagArgs
    ///                 {
    ///                     Key = "k1",
    ///                     Value = "v1",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VkeCluster can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vke/cluster:Cluster default cc9l74mvqtofjnoj5****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vke/cluster:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
        /// </summary>
        [Output("clientToken")]
        public Output<string?> ClientToken { get; private set; } = null!;

        /// <summary>
        /// The config of the cluster.
        /// </summary>
        [Output("clusterConfig")]
        public Output<Outputs.ClusterClusterConfig> ClusterConfig { get; private set; } = null!;

        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        [Output("deleteProtectionEnabled")]
        public Output<bool?> DeleteProtectionEnabled { get; private set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Eip allocation Id.
        /// </summary>
        [Output("eipAllocationId")]
        public Output<string> EipAllocationId { get; private set; } = null!;

        /// <summary>
        /// Kubeconfig data with private network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
        /// </summary>
        [Output("kubeconfigPrivate")]
        public Output<string> KubeconfigPrivate { get; private set; } = null!;

        /// <summary>
        /// Kubeconfig data with public network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
        /// </summary>
        [Output("kubeconfigPublic")]
        public Output<string> KubeconfigPublic { get; private set; } = null!;

        /// <summary>
        /// The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.
        /// </summary>
        [Output("kubernetesVersion")]
        public Output<string> KubernetesVersion { get; private set; } = null!;

        /// <summary>
        /// Cluster log configuration information.
        /// </summary>
        [Output("loggingConfig")]
        public Output<Outputs.ClusterLoggingConfig?> LoggingConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The config of the pods.
        /// </summary>
        [Output("podsConfig")]
        public Output<Outputs.ClusterPodsConfig> PodsConfig { get; private set; } = null!;

        /// <summary>
        /// The config of the services.
        /// </summary>
        [Output("servicesConfig")]
        public Output<Outputs.ClusterServicesConfig> ServicesConfig { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ClusterTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vke/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vke/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// The config of the cluster.
        /// </summary>
        [Input("clusterConfig", required: true)]
        public Input<Inputs.ClusterClusterConfigArgs> ClusterConfig { get; set; } = null!;

        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public Input<bool>? DeleteProtectionEnabled { get; set; }

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.
        /// </summary>
        [Input("kubernetesVersion")]
        public Input<string>? KubernetesVersion { get; set; }

        /// <summary>
        /// Cluster log configuration information.
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.ClusterLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The config of the pods.
        /// </summary>
        [Input("podsConfig", required: true)]
        public Input<Inputs.ClusterPodsConfigArgs> PodsConfig { get; set; } = null!;

        /// <summary>
        /// The config of the services.
        /// </summary>
        [Input("servicesConfig", required: true)]
        public Input<Inputs.ClusterServicesConfigArgs> ServicesConfig { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ClusterTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClusterTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClusterTagArgs>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// The config of the cluster.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.ClusterClusterConfigGetArgs>? ClusterConfig { get; set; }

        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public Input<bool>? DeleteProtectionEnabled { get; set; }

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Eip allocation Id.
        /// </summary>
        [Input("eipAllocationId")]
        public Input<string>? EipAllocationId { get; set; }

        /// <summary>
        /// Kubeconfig data with private network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
        /// </summary>
        [Input("kubeconfigPrivate")]
        public Input<string>? KubeconfigPrivate { get; set; }

        /// <summary>
        /// Kubeconfig data with public network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
        /// </summary>
        [Input("kubeconfigPublic")]
        public Input<string>? KubeconfigPublic { get; set; }

        /// <summary>
        /// The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.
        /// </summary>
        [Input("kubernetesVersion")]
        public Input<string>? KubernetesVersion { get; set; }

        /// <summary>
        /// Cluster log configuration information.
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.ClusterLoggingConfigGetArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The config of the pods.
        /// </summary>
        [Input("podsConfig")]
        public Input<Inputs.ClusterPodsConfigGetArgs>? PodsConfig { get; set; }

        /// <summary>
        /// The config of the services.
        /// </summary>
        [Input("servicesConfig")]
        public Input<Inputs.ClusterServicesConfigGetArgs>? ServicesConfig { get; set; }

        [Input("tags")]
        private InputList<Inputs.ClusterTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClusterTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClusterTagGetArgs>());
            set => _tags = value;
        }

        public ClusterState()
        {
        }
    }
}
