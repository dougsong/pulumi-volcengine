// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class CloudServerScheduleStrategy
    {
        /// <summary>
        /// The network strategy.
        /// </summary>
        public readonly string NetworkStrategy;
        /// <summary>
        /// The price strategy. The value can be `high_priority` or `low_priority`.
        /// </summary>
        public readonly string PriceStrategy;
        /// <summary>
        /// The type of schedule strategy. The value can be `dispersion` or `concentration`.
        /// </summary>
        public readonly string ScheduleStrategy;

        [OutputConstructor]
        private CloudServerScheduleStrategy(
            string networkStrategy,

            string priceStrategy,

            string scheduleStrategy)
        {
            NetworkStrategy = networkStrategy;
            PriceStrategy = priceStrategy;
            ScheduleStrategy = scheduleStrategy;
        }
    }
}
