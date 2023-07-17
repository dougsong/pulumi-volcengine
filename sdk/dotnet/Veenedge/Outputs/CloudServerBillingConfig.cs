// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class CloudServerBillingConfig
    {
        /// <summary>
        /// The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.
        /// </summary>
        public readonly string BandwidthBillingMethod;
        /// <summary>
        /// The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.
        /// </summary>
        public readonly string ComputingBillingMethod;

        [OutputConstructor]
        private CloudServerBillingConfig(
            string bandwidthBillingMethod,

            string computingBillingMethod)
        {
            BandwidthBillingMethod = bandwidthBillingMethod;
            ComputingBillingMethod = computingBillingMethod;
        }
    }
}
