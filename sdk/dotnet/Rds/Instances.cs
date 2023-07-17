// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds
{
    public static class Instances
    {
        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Rds.Instances.InvokeAsync(new Volcengine.Rds.InstancesArgs
        ///         {
        ///             InstanceId = "mysql-0fdd3bab2e7c",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("volcengine:rds/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Rds.Instances.InvokeAsync(new Volcengine.Rds.InstancesArgs
        ///         {
        ///             InstanceId = "mysql-0fdd3bab2e7c",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesResult>("volcengine:rds/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of creating RDS instance.
        /// </summary>
        [Input("createEndTime")]
        public string? CreateEndTime { get; set; }

        /// <summary>
        /// The start time of creating RDS instance.
        /// </summary>
        [Input("createStartTime")]
        public string? CreateStartTime { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The status of the RDS instance.
        /// </summary>
        [Input("instanceStatus")]
        public string? InstanceStatus { get; set; }

        /// <summary>
        /// The type of the RDS instance.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// A Name Regex of RDS instance.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public InstancesArgs()
        {
        }
    }

    public sealed class InstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of creating RDS instance.
        /// </summary>
        [Input("createEndTime")]
        public Input<string>? CreateEndTime { get; set; }

        /// <summary>
        /// The start time of creating RDS instance.
        /// </summary>
        [Input("createStartTime")]
        public Input<string>? CreateStartTime { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The status of the RDS instance.
        /// </summary>
        [Input("instanceStatus")]
        public Input<string>? InstanceStatus { get; set; }

        /// <summary>
        /// The type of the RDS instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// A Name Regex of RDS instance.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesResult
    {
        public readonly string? CreateEndTime;
        public readonly string? CreateStartTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The status of the RDS instance.
        /// </summary>
        public readonly string? InstanceStatus;
        /// <summary>
        /// The type of the RDS instance.
        /// </summary>
        public readonly string? InstanceType;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of RDS instance query.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsInstanceResult> RdsInstances;
        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The total count of RDS instance query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private InstancesResult(
            string? createEndTime,

            string? createStartTime,

            string id,

            string? instanceId,

            string? instanceStatus,

            string? instanceType,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<Outputs.InstancesRdsInstanceResult> rdsInstances,

            string? region,

            int totalCount,

            string? zone)
        {
            CreateEndTime = createEndTime;
            CreateStartTime = createStartTime;
            Id = id;
            InstanceId = instanceId;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            RdsInstances = rdsInstances;
            Region = region;
            TotalCount = totalCount;
            Zone = zone;
        }
    }
}
