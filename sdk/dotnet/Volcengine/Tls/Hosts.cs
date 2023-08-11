// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls
{
    public static class Hosts
    {
        /// <summary>
        /// Use this data source to query detailed information of tls hosts
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
        ///     var @default = Volcengine.Tls.Hosts.Invoke(new()
        ///     {
        ///         HostGroupId = "527102e2-1e4f-45f4-a990-751152125da7",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<HostsResult> InvokeAsync(HostsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<HostsResult>("volcengine:tls/hosts:Hosts", args ?? new HostsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tls hosts
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
        ///     var @default = Volcengine.Tls.Hosts.Invoke(new()
        ///     {
        ///         HostGroupId = "527102e2-1e4f-45f4-a990-751152125da7",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<HostsResult> Invoke(HostsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<HostsResult>("volcengine:tls/hosts:Hosts", args ?? new HostsInvokeArgs(), options.WithDefaults());
    }


    public sealed class HostsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        [Input("heartbeatStatus")]
        public int? HeartbeatStatus { get; set; }

        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public string HostGroupId { get; set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public HostsArgs()
        {
        }
        public static new HostsArgs Empty => new HostsArgs();
    }

    public sealed class HostsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        [Input("heartbeatStatus")]
        public Input<int>? HeartbeatStatus { get; set; }

        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public Input<string> HostGroupId { get; set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public HostsInvokeArgs()
        {
        }
        public static new HostsInvokeArgs Empty => new HostsInvokeArgs();
    }


    [OutputType]
    public sealed class HostsResult
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        public readonly int? HeartbeatStatus;
        /// <summary>
        /// The id of host group.
        /// </summary>
        public readonly string HostGroupId;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.HostsHostInfoResult> HostInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ip address.
        /// </summary>
        public readonly string? Ip;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private HostsResult(
            int? heartbeatStatus,

            string hostGroupId,

            ImmutableArray<Outputs.HostsHostInfoResult> hostInfos,

            string id,

            string? ip,

            string? outputFile,

            int totalCount)
        {
            HeartbeatStatus = heartbeatStatus;
            HostGroupId = hostGroupId;
            HostInfos = hostInfos;
            Id = id;
            Ip = ip;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
