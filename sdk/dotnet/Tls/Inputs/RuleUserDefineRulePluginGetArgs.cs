// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class RuleUserDefineRulePluginGetArgs : Pulumi.ResourceArgs
    {
        [Input("processors", required: true)]
        private InputList<string>? _processors;

        /// <summary>
        /// LogCollector plugin.
        /// </summary>
        public InputList<string> Processors
        {
            get => _processors ?? (_processors = new InputList<string>());
            set => _processors = value;
        }

        public RuleUserDefineRulePluginGetArgs()
        {
        }
    }
}
