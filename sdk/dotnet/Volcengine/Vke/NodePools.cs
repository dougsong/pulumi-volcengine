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
    public static class NodePools
    {
        /// <summary>
        /// Use this data source to query detailed information of vke node pools
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vkeTest = Volcengine.Vke.NodePools.Invoke(new()
        ///     {
        ///         ClusterIds = new[]
        ///         {
        ///             "ccabe57fqtofgrbln3dog",
        ///         },
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<NodePoolsResult> InvokeAsync(NodePoolsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<NodePoolsResult>("volcengine:vke/nodePools:NodePools", args ?? new NodePoolsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vke node pools
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var vkeTest = Volcengine.Vke.NodePools.Invoke(new()
        ///     {
        ///         ClusterIds = new[]
        ///         {
        ///             "ccabe57fqtofgrbln3dog",
        ///         },
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<NodePoolsResult> Invoke(NodePoolsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<NodePoolsResult>("volcengine:vke/nodePools:NodePools", args ?? new NodePoolsInvokeArgs(), options.WithDefaults());
    }


    public sealed class NodePoolsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Is enabled of AutoScaling.
        /// </summary>
        [Input("autoScalingEnabled")]
        public bool? AutoScalingEnabled { get; set; }

        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("clusterIds")]
        private List<string>? _clusterIds;

        /// <summary>
        /// The ClusterIds of NodePool IDs.
        /// </summary>
        public List<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new List<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        [Input("createClientToken")]
        public string? CreateClientToken { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The IDs of NodePool.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A Name Regex of NodePool.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("statuses")]
        private List<Inputs.NodePoolsStatusArgs>? _statuses;

        /// <summary>
        /// The Status of NodePool.
        /// </summary>
        public List<Inputs.NodePoolsStatusArgs> Statuses
        {
            get => _statuses ?? (_statuses = new List<Inputs.NodePoolsStatusArgs>());
            set => _statuses = value;
        }

        [Input("tags")]
        private List<Inputs.NodePoolsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.NodePoolsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.NodePoolsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ClientToken when last update was successful.
        /// </summary>
        [Input("updateClientToken")]
        public string? UpdateClientToken { get; set; }

        public NodePoolsArgs()
        {
        }
        public static new NodePoolsArgs Empty => new NodePoolsArgs();
    }

    public sealed class NodePoolsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Is enabled of AutoScaling.
        /// </summary>
        [Input("autoScalingEnabled")]
        public Input<bool>? AutoScalingEnabled { get; set; }

        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("clusterIds")]
        private InputList<string>? _clusterIds;

        /// <summary>
        /// The ClusterIds of NodePool IDs.
        /// </summary>
        public InputList<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new InputList<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        [Input("createClientToken")]
        public Input<string>? CreateClientToken { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The IDs of NodePool.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A Name Regex of NodePool.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("statuses")]
        private InputList<Inputs.NodePoolsStatusInputArgs>? _statuses;

        /// <summary>
        /// The Status of NodePool.
        /// </summary>
        public InputList<Inputs.NodePoolsStatusInputArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.NodePoolsStatusInputArgs>());
            set => _statuses = value;
        }

        [Input("tags")]
        private InputList<Inputs.NodePoolsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.NodePoolsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NodePoolsTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ClientToken when last update was successful.
        /// </summary>
        [Input("updateClientToken")]
        public Input<string>? UpdateClientToken { get; set; }

        public NodePoolsInvokeArgs()
        {
        }
        public static new NodePoolsInvokeArgs Empty => new NodePoolsInvokeArgs();
    }


    [OutputType]
    public sealed class NodePoolsResult
    {
        public readonly bool? AutoScalingEnabled;
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        public readonly string? ClusterId;
        public readonly ImmutableArray<string> ClusterIds;
        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        public readonly string? CreateClientToken;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        /// <summary>
        /// The collection of NodePools query.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodePoolsNodePoolResult> NodePools;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.NodePoolsStatusResult> Statuses;
        /// <summary>
        /// Tags of the NodePool.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodePoolsTagResult> Tags;
        /// <summary>
        /// Returns the total amount of the data list.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The ClientToken when last update was successful.
        /// </summary>
        public readonly string? UpdateClientToken;

        [OutputConstructor]
        private NodePoolsResult(
            bool? autoScalingEnabled,

            string? clusterId,

            ImmutableArray<string> clusterIds,

            string? createClientToken,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            ImmutableArray<Outputs.NodePoolsNodePoolResult> nodePools,

            string? outputFile,

            ImmutableArray<Outputs.NodePoolsStatusResult> statuses,

            ImmutableArray<Outputs.NodePoolsTagResult> tags,

            int totalCount,

            string? updateClientToken)
        {
            AutoScalingEnabled = autoScalingEnabled;
            ClusterId = clusterId;
            ClusterIds = clusterIds;
            CreateClientToken = createClientToken;
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            NodePools = nodePools;
            OutputFile = outputFile;
            Statuses = statuses;
            Tags = tags;
            TotalCount = totalCount;
            UpdateClientToken = updateClientToken;
        }
    }
}
