// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Autoscaling
{
    public static class ScalingConfigurations
    {
        /// <summary>
        /// Use this data source to query detailed information of scaling configurations
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
        ///         var @default = Output.Create(Volcengine.Autoscaling.ScalingConfigurations.InvokeAsync(new Volcengine.Autoscaling.ScalingConfigurationsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "scc-ybrurj4uw6gh9zecj327",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ScalingConfigurationsResult> InvokeAsync(ScalingConfigurationsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ScalingConfigurationsResult>("volcengine:autoscaling/scalingConfigurations:ScalingConfigurations", args ?? new ScalingConfigurationsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scaling configurations
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
        ///         var @default = Output.Create(Volcengine.Autoscaling.ScalingConfigurations.InvokeAsync(new Volcengine.Autoscaling.ScalingConfigurationsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "scc-ybrurj4uw6gh9zecj327",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ScalingConfigurationsResult> Invoke(ScalingConfigurationsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ScalingConfigurationsResult>("volcengine:autoscaling/scalingConfigurations:ScalingConfigurations", args ?? new ScalingConfigurationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ScalingConfigurationsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of scaling configuration ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of scaling configuration.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("scalingConfigurationNames")]
        private List<string>? _scalingConfigurationNames;

        /// <summary>
        /// A list of scaling configuration names.
        /// </summary>
        public List<string> ScalingConfigurationNames
        {
            get => _scalingConfigurationNames ?? (_scalingConfigurationNames = new List<string>());
            set => _scalingConfigurationNames = value;
        }

        /// <summary>
        /// An id of scaling group.
        /// </summary>
        [Input("scalingGroupId")]
        public string? ScalingGroupId { get; set; }

        public ScalingConfigurationsArgs()
        {
        }
    }

    public sealed class ScalingConfigurationsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of scaling configuration ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of scaling configuration.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("scalingConfigurationNames")]
        private InputList<string>? _scalingConfigurationNames;

        /// <summary>
        /// A list of scaling configuration names.
        /// </summary>
        public InputList<string> ScalingConfigurationNames
        {
            get => _scalingConfigurationNames ?? (_scalingConfigurationNames = new InputList<string>());
            set => _scalingConfigurationNames = value;
        }

        /// <summary>
        /// An id of scaling group.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public ScalingConfigurationsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ScalingConfigurationsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly ImmutableArray<string> ScalingConfigurationNames;
        /// <summary>
        /// The collection of scaling configuration query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingConfigurationsScalingConfigurationResult> ScalingConfigurations;
        /// <summary>
        /// The id of the scaling group to which the scaling configuration belongs.
        /// </summary>
        public readonly string? ScalingGroupId;
        /// <summary>
        /// The total count of scaling configuration query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private ScalingConfigurationsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<string> scalingConfigurationNames,

            ImmutableArray<Outputs.ScalingConfigurationsScalingConfigurationResult> scalingConfigurations,

            string? scalingGroupId,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ScalingConfigurationNames = scalingConfigurationNames;
            ScalingConfigurations = scalingConfigurations;
            ScalingGroupId = scalingGroupId;
            TotalCount = totalCount;
        }
    }
}
