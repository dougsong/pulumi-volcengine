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
    public sealed class AlarmNotifyGroupsGroupReceiverResult
    {
        /// <summary>
        /// The end time.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The list of the receiver channels.
        /// </summary>
        public readonly ImmutableArray<string> ReceiverChannels;
        /// <summary>
        /// List of the receiver names.
        /// </summary>
        public readonly ImmutableArray<string> ReceiverNames;
        /// <summary>
        /// The receiver type.
        /// </summary>
        public readonly string ReceiverType;
        /// <summary>
        /// The start time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private AlarmNotifyGroupsGroupReceiverResult(
            string endTime,

            ImmutableArray<string> receiverChannels,

            ImmutableArray<string> receiverNames,

            string receiverType,

            string startTime)
        {
            EndTime = endTime;
            ReceiverChannels = receiverChannels;
            ReceiverNames = receiverNames;
            ReceiverType = receiverType;
            StartTime = startTime;
        }
    }
}
