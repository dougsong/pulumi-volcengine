// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    /// <summary>
    /// Provides a resource to manage tls host group
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Tls.HostGroup("foo", new Volcengine.Tls.HostGroupArgs
    ///         {
    ///             AutoUpdate = false,
    ///             HostGroupName = "tfgroup",
    ///             HostGroupType = "Label",
    ///             HostIdentifier = "tf-controller",
    ///             ServiceLogging = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Tls Host Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:tls/hostGroup:HostGroup default edf052s21s*******dc15
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tls/hostGroup:HostGroup")]
    public partial class HostGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The abnormal heartbeat status count of host.
        /// </summary>
        [Output("abnormalHeartbeatStatusCount")]
        public Output<int> AbnormalHeartbeatStatusCount { get; private set; } = null!;

        /// <summary>
        /// The latest version of log collector.
        /// </summary>
        [Output("agentLatestVersion")]
        public Output<string> AgentLatestVersion { get; private set; } = null!;

        /// <summary>
        /// Whether enable auto update.
        /// </summary>
        [Output("autoUpdate")]
        public Output<bool?> AutoUpdate { get; private set; } = null!;

        /// <summary>
        /// The create time of host group.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The count of host.
        /// </summary>
        [Output("hostCount")]
        public Output<int> HostCount { get; private set; } = null!;

        /// <summary>
        /// The name of host group.
        /// </summary>
        [Output("hostGroupName")]
        public Output<string> HostGroupName { get; private set; } = null!;

        /// <summary>
        /// The type of host group. The value can be IP or Label.
        /// </summary>
        [Output("hostGroupType")]
        public Output<string> HostGroupType { get; private set; } = null!;

        /// <summary>
        /// The identifier of host.
        /// </summary>
        [Output("hostIdentifier")]
        public Output<string?> HostIdentifier { get; private set; } = null!;

        /// <summary>
        /// The ip list of host group.
        /// </summary>
        [Output("hostIpLists")]
        public Output<ImmutableArray<string>> HostIpLists { get; private set; } = null!;

        /// <summary>
        /// The project name of iam.
        /// </summary>
        [Output("iamProjectName")]
        public Output<string?> IamProjectName { get; private set; } = null!;

        /// <summary>
        /// The modify time of host group.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// The normal heartbeat status count of host.
        /// </summary>
        [Output("normalHeartbeatStatusCount")]
        public Output<int> NormalHeartbeatStatusCount { get; private set; } = null!;

        /// <summary>
        /// The rule count of host.
        /// </summary>
        [Output("ruleCount")]
        public Output<int> RuleCount { get; private set; } = null!;

        /// <summary>
        /// Whether enable service logging.
        /// </summary>
        [Output("serviceLogging")]
        public Output<bool?> ServiceLogging { get; private set; } = null!;

        /// <summary>
        /// The update end time of log collector.
        /// </summary>
        [Output("updateEndTime")]
        public Output<string> UpdateEndTime { get; private set; } = null!;

        /// <summary>
        /// The update start time of log collector.
        /// </summary>
        [Output("updateStartTime")]
        public Output<string> UpdateStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a HostGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostGroup(string name, HostGroupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tls/hostGroup:HostGroup", name, args ?? new HostGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostGroup(string name, Input<string> id, HostGroupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tls/hostGroup:HostGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HostGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostGroup Get(string name, Input<string> id, HostGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new HostGroup(name, id, state, options);
        }
    }

    public sealed class HostGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether enable auto update.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// The name of host group.
        /// </summary>
        [Input("hostGroupName", required: true)]
        public Input<string> HostGroupName { get; set; } = null!;

        /// <summary>
        /// The type of host group. The value can be IP or Label.
        /// </summary>
        [Input("hostGroupType", required: true)]
        public Input<string> HostGroupType { get; set; } = null!;

        /// <summary>
        /// The identifier of host.
        /// </summary>
        [Input("hostIdentifier")]
        public Input<string>? HostIdentifier { get; set; }

        [Input("hostIpLists")]
        private InputList<string>? _hostIpLists;

        /// <summary>
        /// The ip list of host group.
        /// </summary>
        public InputList<string> HostIpLists
        {
            get => _hostIpLists ?? (_hostIpLists = new InputList<string>());
            set => _hostIpLists = value;
        }

        /// <summary>
        /// The project name of iam.
        /// </summary>
        [Input("iamProjectName")]
        public Input<string>? IamProjectName { get; set; }

        /// <summary>
        /// Whether enable service logging.
        /// </summary>
        [Input("serviceLogging")]
        public Input<bool>? ServiceLogging { get; set; }

        /// <summary>
        /// The update end time of log collector.
        /// </summary>
        [Input("updateEndTime")]
        public Input<string>? UpdateEndTime { get; set; }

        /// <summary>
        /// The update start time of log collector.
        /// </summary>
        [Input("updateStartTime")]
        public Input<string>? UpdateStartTime { get; set; }

        public HostGroupArgs()
        {
        }
    }

    public sealed class HostGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The abnormal heartbeat status count of host.
        /// </summary>
        [Input("abnormalHeartbeatStatusCount")]
        public Input<int>? AbnormalHeartbeatStatusCount { get; set; }

        /// <summary>
        /// The latest version of log collector.
        /// </summary>
        [Input("agentLatestVersion")]
        public Input<string>? AgentLatestVersion { get; set; }

        /// <summary>
        /// Whether enable auto update.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// The create time of host group.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The count of host.
        /// </summary>
        [Input("hostCount")]
        public Input<int>? HostCount { get; set; }

        /// <summary>
        /// The name of host group.
        /// </summary>
        [Input("hostGroupName")]
        public Input<string>? HostGroupName { get; set; }

        /// <summary>
        /// The type of host group. The value can be IP or Label.
        /// </summary>
        [Input("hostGroupType")]
        public Input<string>? HostGroupType { get; set; }

        /// <summary>
        /// The identifier of host.
        /// </summary>
        [Input("hostIdentifier")]
        public Input<string>? HostIdentifier { get; set; }

        [Input("hostIpLists")]
        private InputList<string>? _hostIpLists;

        /// <summary>
        /// The ip list of host group.
        /// </summary>
        public InputList<string> HostIpLists
        {
            get => _hostIpLists ?? (_hostIpLists = new InputList<string>());
            set => _hostIpLists = value;
        }

        /// <summary>
        /// The project name of iam.
        /// </summary>
        [Input("iamProjectName")]
        public Input<string>? IamProjectName { get; set; }

        /// <summary>
        /// The modify time of host group.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// The normal heartbeat status count of host.
        /// </summary>
        [Input("normalHeartbeatStatusCount")]
        public Input<int>? NormalHeartbeatStatusCount { get; set; }

        /// <summary>
        /// The rule count of host.
        /// </summary>
        [Input("ruleCount")]
        public Input<int>? RuleCount { get; set; }

        /// <summary>
        /// Whether enable service logging.
        /// </summary>
        [Input("serviceLogging")]
        public Input<bool>? ServiceLogging { get; set; }

        /// <summary>
        /// The update end time of log collector.
        /// </summary>
        [Input("updateEndTime")]
        public Input<string>? UpdateEndTime { get; set; }

        /// <summary>
        /// The update start time of log collector.
        /// </summary>
        [Input("updateStartTime")]
        public Input<string>? UpdateStartTime { get; set; }

        public HostGroupState()
        {
        }
    }
}
