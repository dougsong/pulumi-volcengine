// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds_v2.Outputs
{

    [OutputType]
    public sealed class RdsInstanceV2ChargeInfo
    {
        /// <summary>
        /// Whether to automatically renew in prepaid scenarios.
        /// </summary>
        public readonly bool? AutoRenew;
        /// <summary>
        /// Payment type. Value:
        /// PostPaid - Pay-As-You-Go
        /// PrePaid - Yearly and monthly (default).
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// Purchase duration in prepaid scenarios. Default: 1.
        /// </summary>
        public readonly int? Period;
        /// <summary>
        /// The purchase cycle in the prepaid scenario.
        /// Month - monthly subscription (default)
        /// Year - Package year.
        /// </summary>
        public readonly string? PeriodUnit;

        [OutputConstructor]
        private RdsInstanceV2ChargeInfo(
            bool? autoRenew,

            string chargeType,

            int? period,

            string? periodUnit)
        {
            AutoRenew = autoRenew;
            ChargeType = chargeType;
            Period = period;
            PeriodUnit = periodUnit;
        }
    }
}
