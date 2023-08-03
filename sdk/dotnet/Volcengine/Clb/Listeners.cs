// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Clb
{
    public static class Listeners
    {
        /// <summary>
        /// Use this data source to query detailed information of listeners
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var fooZones = Output.Create(Volcengine.Ecs.Zones.InvokeAsync());
        ///         var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new Volcengine.Vpc.VpcArgs
        ///         {
        ///             VpcName = "acc-test-vpc",
        ///             CidrBlock = "172.16.0.0/16",
        ///         });
        ///         var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new Volcengine.Vpc.SubnetArgs
        ///         {
        ///             SubnetName = "acc-test-subnet",
        ///             CidrBlock = "172.16.0.0/24",
        ///             ZoneId = fooZones.Apply(fooZones =&gt; fooZones.Zones?[0]?.Id),
        ///             VpcId = fooVpc.Id,
        ///         });
        ///         var fooClb = new Volcengine.Clb.Clb("fooClb", new Volcengine.Clb.ClbArgs
        ///         {
        ///             Type = "public",
        ///             SubnetId = fooSubnet.Id,
        ///             LoadBalancerSpec = "small_1",
        ///             Description = "acc0Demo",
        ///             LoadBalancerName = "acc-test-create",
        ///             EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///             {
        ///                 Isp = "BGP",
        ///                 EipBillingType = "PostPaidByBandwidth",
        ///                 Bandwidth = 1,
        ///             },
        ///         });
        ///         var fooServerGroup = new Volcengine.Clb.ServerGroup("fooServerGroup", new Volcengine.Clb.ServerGroupArgs
        ///         {
        ///             LoadBalancerId = fooClb.Id,
        ///             ServerGroupName = "acc-test-create",
        ///             Description = "hello demo11",
        ///         });
        ///         var fooListener = new Volcengine.Clb.Listener("fooListener", new Volcengine.Clb.ListenerArgs
        ///         {
        ///             LoadBalancerId = fooClb.Id,
        ///             ListenerName = "acc-test-listener",
        ///             Protocol = "HTTP",
        ///             Port = 90,
        ///             ServerGroupId = fooServerGroup.Id,
        ///             HealthCheck = new Volcengine.Clb.Inputs.ListenerHealthCheckArgs
        ///             {
        ///                 Enabled = "on",
        ///                 Interval = 10,
        ///                 Timeout = 3,
        ///                 HealthyThreshold = 5,
        ///                 UnHealthyThreshold = 2,
        ///                 Domain = "volcengine.com",
        ///                 HttpCode = "http_2xx",
        ///                 Method = "GET",
        ///                 Uri = "/",
        ///             },
        ///             Enabled = "on",
        ///         });
        ///         var fooListeners = Volcengine.Clb.Listeners.Invoke(new Volcengine.Clb.ListenersInvokeArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 fooListener.Id,
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ListenersResult> InvokeAsync(ListenersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListenersResult>("volcengine:clb/listeners:Listeners", args ?? new ListenersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of listeners
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var fooZones = Output.Create(Volcengine.Ecs.Zones.InvokeAsync());
        ///         var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new Volcengine.Vpc.VpcArgs
        ///         {
        ///             VpcName = "acc-test-vpc",
        ///             CidrBlock = "172.16.0.0/16",
        ///         });
        ///         var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new Volcengine.Vpc.SubnetArgs
        ///         {
        ///             SubnetName = "acc-test-subnet",
        ///             CidrBlock = "172.16.0.0/24",
        ///             ZoneId = fooZones.Apply(fooZones =&gt; fooZones.Zones?[0]?.Id),
        ///             VpcId = fooVpc.Id,
        ///         });
        ///         var fooClb = new Volcengine.Clb.Clb("fooClb", new Volcengine.Clb.ClbArgs
        ///         {
        ///             Type = "public",
        ///             SubnetId = fooSubnet.Id,
        ///             LoadBalancerSpec = "small_1",
        ///             Description = "acc0Demo",
        ///             LoadBalancerName = "acc-test-create",
        ///             EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///             {
        ///                 Isp = "BGP",
        ///                 EipBillingType = "PostPaidByBandwidth",
        ///                 Bandwidth = 1,
        ///             },
        ///         });
        ///         var fooServerGroup = new Volcengine.Clb.ServerGroup("fooServerGroup", new Volcengine.Clb.ServerGroupArgs
        ///         {
        ///             LoadBalancerId = fooClb.Id,
        ///             ServerGroupName = "acc-test-create",
        ///             Description = "hello demo11",
        ///         });
        ///         var fooListener = new Volcengine.Clb.Listener("fooListener", new Volcengine.Clb.ListenerArgs
        ///         {
        ///             LoadBalancerId = fooClb.Id,
        ///             ListenerName = "acc-test-listener",
        ///             Protocol = "HTTP",
        ///             Port = 90,
        ///             ServerGroupId = fooServerGroup.Id,
        ///             HealthCheck = new Volcengine.Clb.Inputs.ListenerHealthCheckArgs
        ///             {
        ///                 Enabled = "on",
        ///                 Interval = 10,
        ///                 Timeout = 3,
        ///                 HealthyThreshold = 5,
        ///                 UnHealthyThreshold = 2,
        ///                 Domain = "volcengine.com",
        ///                 HttpCode = "http_2xx",
        ///                 Method = "GET",
        ///                 Uri = "/",
        ///             },
        ///             Enabled = "on",
        ///         });
        ///         var fooListeners = Volcengine.Clb.Listeners.Invoke(new Volcengine.Clb.ListenersInvokeArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 fooListener.Id,
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ListenersResult> Invoke(ListenersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ListenersResult>("volcengine:clb/listeners:Listeners", args ?? new ListenersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListenersArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the Listener.
        /// </summary>
        [Input("listenerName")]
        public string? ListenerName { get; set; }

        /// <summary>
        /// The id of the Clb.
        /// </summary>
        [Input("loadBalancerId")]
        public string? LoadBalancerId { get; set; }

        /// <summary>
        /// A Name Regex of Listener.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public ListenersArgs()
        {
        }
    }

    public sealed class ListenersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the Listener.
        /// </summary>
        [Input("listenerName")]
        public Input<string>? ListenerName { get; set; }

        /// <summary>
        /// The id of the Clb.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// A Name Regex of Listener.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public ListenersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListenersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of the Listener.
        /// </summary>
        public readonly string? ListenerName;
        /// <summary>
        /// The collection of Listener query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ListenersListenerResult> Listeners;
        public readonly string? LoadBalancerId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of Listener query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private ListenersResult(
            string id,

            ImmutableArray<string> ids,

            string? listenerName,

            ImmutableArray<Outputs.ListenersListenerResult> listeners,

            string? loadBalancerId,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            ListenerName = listenerName;
            Listeners = listeners;
            LoadBalancerId = loadBalancerId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
