// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb
{
    public static class ClbRules
    {
        /// <summary>
        /// Use this data source to query detailed information of clb rules
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
        ///         var @default = Output.Create(Volcengine.Clb.ClbRules.InvokeAsync(new Volcengine.Clb.ClbRulesArgs
        ///         {
        ///             ListenerId = "lsn-273ywvnmiu70g7fap8u2xzg9d",
        ///             Ids = 
        ///             {
        ///                 "rule-273z9jo9v3mrk7fap8sq8v5x7",
        ///             },
        ///         }));
        ///         this.Data = @default;
        ///     }
        /// 
        ///     [Output("data")]
        ///     public Output&lt;string&gt; Data { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ClbRulesResult> InvokeAsync(ClbRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ClbRulesResult>("volcengine:Clb/clbRules:ClbRules", args ?? new ClbRulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb rules
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
        ///         var @default = Output.Create(Volcengine.Clb.ClbRules.InvokeAsync(new Volcengine.Clb.ClbRulesArgs
        ///         {
        ///             ListenerId = "lsn-273ywvnmiu70g7fap8u2xzg9d",
        ///             Ids = 
        ///             {
        ///                 "rule-273z9jo9v3mrk7fap8sq8v5x7",
        ///             },
        ///         }));
        ///         this.Data = @default;
        ///     }
        /// 
        ///     [Output("data")]
        ///     public Output&lt;string&gt; Data { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ClbRulesResult> Invoke(ClbRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ClbRulesResult>("volcengine:Clb/clbRules:ClbRules", args ?? new ClbRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ClbRulesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Id of listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public ClbRulesArgs()
        {
        }
    }

    public sealed class ClbRulesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Id of listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public ClbRulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ClbRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string ListenerId;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of Rule query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClbRulesRuleResult> Rules;

        [OutputConstructor]
        private ClbRulesResult(
            string id,

            ImmutableArray<string> ids,

            string listenerId,

            string? outputFile,

            ImmutableArray<Outputs.ClbRulesRuleResult> rules)
        {
            Id = id;
            Ids = ids;
            ListenerId = listenerId;
            OutputFile = outputFile;
            Rules = rules;
        }
    }
}
