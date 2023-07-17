// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Autoscaling
{
    public static class ScalingPolicies
    {
        /// <summary>
        /// Use this data source to query detailed information of scaling policies
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
        ///         var @default = Output.Create(Volcengine.Autoscaling.ScalingPolicies.InvokeAsync(new Volcengine.Autoscaling.ScalingPoliciesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "sp-ybruzckr8bgh9zrxw29v",
        ///             },
        ///             ScalingGroupId = "scg-ybqm0b6kcigh9zu9ce6t",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ScalingPoliciesResult> InvokeAsync(ScalingPoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ScalingPoliciesResult>("volcengine:autoscaling/scalingPolicies:ScalingPolicies", args ?? new ScalingPoliciesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scaling policies
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
        ///         var @default = Output.Create(Volcengine.Autoscaling.ScalingPolicies.InvokeAsync(new Volcengine.Autoscaling.ScalingPoliciesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "sp-ybruzckr8bgh9zrxw29v",
        ///             },
        ///             ScalingGroupId = "scg-ybqm0b6kcigh9zu9ce6t",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ScalingPoliciesResult> Invoke(ScalingPoliciesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ScalingPoliciesResult>("volcengine:autoscaling/scalingPolicies:ScalingPolicies", args ?? new ScalingPoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ScalingPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of scaling policy ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of scaling policy.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// An id of the scaling group to which the scaling policy belongs.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public string ScalingGroupId { get; set; } = null!;

        [Input("scalingPolicyNames")]
        private List<string>? _scalingPolicyNames;

        /// <summary>
        /// A list of scaling policy names.
        /// </summary>
        public List<string> ScalingPolicyNames
        {
            get => _scalingPolicyNames ?? (_scalingPolicyNames = new List<string>());
            set => _scalingPolicyNames = value;
        }

        /// <summary>
        /// A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.
        /// </summary>
        [Input("scalingPolicyType")]
        public string? ScalingPolicyType { get; set; }

        public ScalingPoliciesArgs()
        {
        }
    }

    public sealed class ScalingPoliciesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of scaling policy ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of scaling policy.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// An id of the scaling group to which the scaling policy belongs.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        [Input("scalingPolicyNames")]
        private InputList<string>? _scalingPolicyNames;

        /// <summary>
        /// A list of scaling policy names.
        /// </summary>
        public InputList<string> ScalingPolicyNames
        {
            get => _scalingPolicyNames ?? (_scalingPolicyNames = new InputList<string>());
            set => _scalingPolicyNames = value;
        }

        /// <summary>
        /// A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.
        /// </summary>
        [Input("scalingPolicyType")]
        public Input<string>? ScalingPolicyType { get; set; }

        public ScalingPoliciesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ScalingPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The id of the scaling group to which the scaling policy belongs.
        /// </summary>
        public readonly string ScalingGroupId;
        /// <summary>
        /// The collection of scaling policy query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingPoliciesScalingPolicyResult> ScalingPolicies;
        public readonly ImmutableArray<string> ScalingPolicyNames;
        /// <summary>
        /// The type of the scaling policy.
        /// </summary>
        public readonly string? ScalingPolicyType;
        /// <summary>
        /// The total count of scaling policy query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private ScalingPoliciesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string scalingGroupId,

            ImmutableArray<Outputs.ScalingPoliciesScalingPolicyResult> scalingPolicies,

            ImmutableArray<string> scalingPolicyNames,

            string? scalingPolicyType,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ScalingGroupId = scalingGroupId;
            ScalingPolicies = scalingPolicies;
            ScalingPolicyNames = scalingPolicyNames;
            ScalingPolicyType = scalingPolicyType;
            TotalCount = totalCount;
        }
    }
}
