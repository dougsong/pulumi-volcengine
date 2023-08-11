// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RuleAppliersRuleUserDefineRuleResult
    {
        /// <summary>
        /// LogCollector extension configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleAdvancedResult> Advanceds;
        /// <summary>
        /// Whether to upload raw logs.
        /// </summary>
        public readonly bool EnableRawLog;
        /// <summary>
        /// Add constant fields to logs.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Fields;
        /// <summary>
        /// Rules for parsing collection paths. After the rules are set, the fields in the collection path will be extracted through the regular expressions specified in the rules, and added to the log data as metadata.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleParsePathRuleResult> ParsePathRules;
        /// <summary>
        /// Plugin configuration. After the plugin configuration is enabled, one or more LogCollector processor plugins can be added to parse logs with complex or variable structures.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleAppliersRuleUserDefineRulePluginResult> Plugins;
        /// <summary>
        /// Rules for routing log partitions. Setting this parameter indicates that the HashKey routing shard mode is used when collecting logs, and Log Service will write the data to the shard containing the specified Key value.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleShardHashKeyResult> ShardHashKeys;
        /// <summary>
        /// LogCollector collection strategy, which specifies whether LogCollector collects incremental logs or full logs. The default is false, which means to collect all logs.
        /// </summary>
        public readonly bool TailFiles;

        [OutputConstructor]
        private RuleAppliersRuleUserDefineRuleResult(
            ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleAdvancedResult> advanceds,

            bool enableRawLog,

            ImmutableDictionary<string, object> fields,

            ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleParsePathRuleResult> parsePathRules,

            ImmutableArray<Outputs.RuleAppliersRuleUserDefineRulePluginResult> plugins,

            ImmutableArray<Outputs.RuleAppliersRuleUserDefineRuleShardHashKeyResult> shardHashKeys,

            bool tailFiles)
        {
            Advanceds = advanceds;
            EnableRawLog = enableRawLog;
            Fields = fields;
            ParsePathRules = parsePathRules;
            Plugins = plugins;
            ShardHashKeys = shardHashKeys;
            TailFiles = tailFiles;
        }
    }
}
