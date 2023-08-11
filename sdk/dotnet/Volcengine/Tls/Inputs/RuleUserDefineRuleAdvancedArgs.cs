// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Inputs
{

    public sealed class RuleUserDefineRuleAdvancedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to release the log file handle after reading to the end of the log file. The default is false.
        /// </summary>
        [Input("closeEof")]
        public Input<bool>? CloseEof { get; set; }

        /// <summary>
        /// The wait time to release the log file handle. When the log file has not written a new log for more than the specified time, release the handle of the log file.
        /// </summary>
        [Input("closeInactive")]
        public Input<int>? CloseInactive { get; set; }

        /// <summary>
        /// After the log file is removed, whether to release the handle of the log file. The default is false.
        /// </summary>
        [Input("closeRemoved")]
        public Input<bool>? CloseRemoved { get; set; }

        /// <summary>
        /// After the log file is renamed, whether to release the handle of the log file. The default is false.
        /// </summary>
        [Input("closeRenamed")]
        public Input<bool>? CloseRenamed { get; set; }

        /// <summary>
        /// The maximum length of time that LogCollector monitors log files. The unit is seconds, and the default is 0 seconds, which means that there is no limit to the length of time LogCollector monitors log files.
        /// </summary>
        [Input("closeTimeout")]
        public Input<int>? CloseTimeout { get; set; }

        public RuleUserDefineRuleAdvancedArgs()
        {
        }
        public static new RuleUserDefineRuleAdvancedArgs Empty => new RuleUserDefineRuleAdvancedArgs();
    }
}
