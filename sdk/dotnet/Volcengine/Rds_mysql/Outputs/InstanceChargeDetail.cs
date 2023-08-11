// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class InstanceChargeDetail
    {
        /// <summary>
        /// Whether to automatically renew in prepaid scenarios.
        /// </summary>
        public readonly bool? AutoRenew;
        /// <summary>
        /// Billing expiry time (yearly and monthly only).
        /// </summary>
        public readonly string? ChargeEndTime;
        /// <summary>
        /// Billing start time (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        public readonly string? ChargeStartTime;
        /// <summary>
        /// Pay status. Value:
        /// normal - normal
        /// overdue - overdue
        /// .
        /// </summary>
        public readonly string? ChargeStatus;
        /// <summary>
        /// Payment type. Value:
        /// PostPaid - Pay-As-You-Go
        /// PrePaid - Yearly and monthly (default).
        /// </summary>
        public readonly string? ChargeType;
        /// <summary>
        /// Estimated release time when arrears are closed (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        public readonly string? OverdueReclaimTime;
        /// <summary>
        /// Shutdown time in arrears (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        public readonly string? OverdueTime;
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
        private InstanceChargeDetail(
            bool? autoRenew,

            string? chargeEndTime,

            string? chargeStartTime,

            string? chargeStatus,

            string? chargeType,

            string? overdueReclaimTime,

            string? overdueTime,

            int? period,

            string? periodUnit)
        {
            AutoRenew = autoRenew;
            ChargeEndTime = chargeEndTime;
            ChargeStartTime = chargeStartTime;
            ChargeStatus = chargeStatus;
            ChargeType = chargeType;
            OverdueReclaimTime = overdueReclaimTime;
            OverdueTime = overdueTime;
            Period = period;
            PeriodUnit = periodUnit;
        }
    }
}
