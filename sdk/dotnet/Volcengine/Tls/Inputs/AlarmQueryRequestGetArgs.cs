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

    public sealed class AlarmQueryRequestGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.
        /// </summary>
        [Input("endTimeOffset", required: true)]
        public Input<int> EndTimeOffset { get; set; } = null!;

        /// <summary>
        /// Alarm object sequence number; increments from 1.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        /// <summary>
        /// Query statement, the maximum supported length is 1024.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        /// <summary>
        /// The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.
        /// </summary>
        [Input("startTimeOffset", required: true)]
        public Input<int> StartTimeOffset { get; set; } = null!;

        /// <summary>
        /// The id of the topic.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public AlarmQueryRequestGetArgs()
        {
        }
        public static new AlarmQueryRequestGetArgs Empty => new AlarmQueryRequestGetArgs();
    }
}
