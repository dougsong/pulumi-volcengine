// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Autoscaling
{
    /// <summary>
    /// Provides a resource to manage scaling policy
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Autoscaling.ScalingPolicy("foo", new()
    ///     {
    ///         Active = false,
    ///         AdjustmentType = "QuantityChangeInCapacity",
    ///         AdjustmentValue = 100,
    ///         AlarmPolicyConditionComparisonOperator = "=",
    ///         AlarmPolicyConditionMetricName = "Instance_CpuBusy_Avg",
    ///         AlarmPolicyConditionMetricUnit = "Percent",
    ///         AlarmPolicyConditionThreshold = "100",
    ///         AlarmPolicyEvaluationCount = 1,
    ///         AlarmPolicyRuleType = "Static",
    ///         Cooldown = 10,
    ///         ScalingGroupId = "scg-ybqm0b6kcigh9zu9ce6t",
    ///         ScalingPolicyName = "tf-test",
    ///         ScalingPolicyType = "Alarm",
    ///         ScheduledPolicyLaunchTime = "2022-07-09T09:59Z",
    ///         ScheduledPolicyRecurrenceEndTime = "2022-07-24T09:25Z",
    ///         ScheduledPolicyRecurrenceType = "Daily",
    ///         ScheduledPolicyRecurrenceValue = "10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ScalingPolicy can be imported using the ScalingGroupId:ScalingPolicyId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:autoscaling/scalingPolicy:ScalingPolicy default scg-yblfbfhy7agh9zn72iaz:sp-yblf9l4fvcl8j1prohsp
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:autoscaling/scalingPolicy:ScalingPolicy")]
    public partial class ScalingPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The active flag of the scaling policy. [Warning] the scaling policy can be active only when the scaling group be active otherwise will fail.
        /// </summary>
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        /// <summary>
        /// The adjustment type of the scaling policy. Valid values: QuantityChangeInCapacity, PercentChangeInCapacity, TotalCapacity.
        /// </summary>
        [Output("adjustmentType")]
        public Output<string> AdjustmentType { get; private set; } = null!;

        /// <summary>
        /// The adjustment value of the scaling policy. When the value of the `AdjustmentType` parameter is `QuantityChangeInCapacity`: -100 ~ 100, 0 is not allowed, unit: piece. When the value of the `AdjustmentType` parameter is `PercentChangeInCapacity`: -100 ~ 10000, 0 is not allowed, unit: %. When the value of the `AdjustmentType` parameter is `TotalCapacity`: the default is 0 to 100, unit: piece.
        /// </summary>
        [Output("adjustmentValue")]
        public Output<int> AdjustmentValue { get; private set; } = null!;

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. Valid values: `&gt;`, `&lt;`, `=`. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Output("alarmPolicyConditionComparisonOperator")]
        public Output<string?> AlarmPolicyConditionComparisonOperator { get; private set; } = null!;

        /// <summary>
        /// The metric name of the alarm policy condition of the scaling policy. Valid values: CpuTotal_Max, CpuTotal_Min, CpuTotal_Avg, MemoryUsedUtilization_Max, MemoryUsedUtilization_Min, MemoryUsedUtilization_Avg, Instance_CpuBusy_Max, Instance_CpuBusy_Min, Instance_CpuBusy_Avg.
        /// </summary>
        [Output("alarmPolicyConditionMetricName")]
        public Output<string?> AlarmPolicyConditionMetricName { get; private set; } = null!;

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Output("alarmPolicyConditionMetricUnit")]
        public Output<string?> AlarmPolicyConditionMetricUnit { get; private set; } = null!;

        /// <summary>
        /// The threshold of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Output("alarmPolicyConditionThreshold")]
        public Output<string?> AlarmPolicyConditionThreshold { get; private set; } = null!;

        /// <summary>
        /// The evaluation count of the alarm policy of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Output("alarmPolicyEvaluationCount")]
        public Output<int?> AlarmPolicyEvaluationCount { get; private set; } = null!;

        /// <summary>
        /// The rule type of the alarm policy of the scaling policy. Valid value: Static. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Output("alarmPolicyRuleType")]
        public Output<string?> AlarmPolicyRuleType { get; private set; } = null!;

        /// <summary>
        /// The cooldown of the scaling policy. Default value is the cooldown time of the scaling group. Value: 0~86400, unit: second, if left blank, the cooling time of the scaling group will be used by default.
        /// </summary>
        [Output("cooldown")]
        public Output<int> Cooldown { get; private set; } = null!;

        /// <summary>
        /// The id of the scaling group to which the scaling policy belongs.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the scaling policy.
        /// </summary>
        [Output("scalingPolicyName")]
        public Output<string> ScalingPolicyName { get; private set; } = null!;

        /// <summary>
        /// The type of scaling policy. Valid values: Scheduled, Recurrence, Alarm.
        /// </summary>
        [Output("scalingPolicyType")]
        public Output<string> ScalingPolicyType { get; private set; } = null!;

        /// <summary>
        /// The launch time of the scheduled policy of the scaling policy.
        /// When the value of `ScalingPolicyType` is `Scheduled`, it means that the trigger time of the scheduled task must be greater than the current time.
        /// When the value of `ScalingPolicyType` is `Recurrence`: If `ScheduledPolicy.RecurrenceType` is not specified, it means to execute only once according to the date and time specified here.
        /// If `ScheduledPolicy.RecurrenceType` is specified, it indicates the start time of the periodic task. Only the time within 90 days from the date of creation/modification is supported.
        /// When the value of `ScalingPolicyType` is `Alarm`, this parameter is invalid.
        /// </summary>
        [Output("scheduledPolicyLaunchTime")]
        public Output<string> ScheduledPolicyLaunchTime { get; private set; } = null!;

        /// <summary>
        /// The recurrence end time of the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. If not configured, it will default to the day/week/month after this moment according to the recurrence period (ScheduledPolicy.RecurrenceType).
        /// </summary>
        [Output("scheduledPolicyRecurrenceEndTime")]
        public Output<string?> ScheduledPolicyRecurrenceEndTime { get; private set; } = null!;

        /// <summary>
        /// The recurrence type the scheduled policy of the scaling policy. Valid values: Daily, Weekly, Monthly, Cron.
        /// </summary>
        [Output("scheduledPolicyRecurrenceType")]
        public Output<string?> ScheduledPolicyRecurrenceType { get; private set; } = null!;

        /// <summary>
        /// The recurrence value the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. When the value of the ScheduledPolicy.RecurrenceType parameter is Daily, only one value can be filled in, ranging from 1 to 31.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Weekly, you can enter multiple values separated by commas (,). The values from Monday to Sunday are: 1,2,3,4,5,6,7.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Monthly, the format is A-B. The value ranges of A and B are both 1~31, and B must be greater than or equal to A.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Cron, it means UTC+8 time, supports 5-field expressions of minutes, hours, days, months, and weeks, and supports wildcard English commas (,), English question marks (?), and conjunctions ( -), asterisk (*), pound sign (#), slash (/), L, and W.
        /// </summary>
        [Output("scheduledPolicyRecurrenceValue")]
        public Output<string?> ScheduledPolicyRecurrenceValue { get; private set; } = null!;

        /// <summary>
        /// The status of the scaling policy. Valid values: Active, InActive.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ScalingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScalingPolicy(string name, ScalingPolicyArgs args, CustomResourceOptions? options = null)
            : base("volcengine:autoscaling/scalingPolicy:ScalingPolicy", name, args ?? new ScalingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScalingPolicy(string name, Input<string> id, ScalingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:autoscaling/scalingPolicy:ScalingPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ScalingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScalingPolicy Get(string name, Input<string> id, ScalingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ScalingPolicy(name, id, state, options);
        }
    }

    public sealed class ScalingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The active flag of the scaling policy. [Warning] the scaling policy can be active only when the scaling group be active otherwise will fail.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The adjustment type of the scaling policy. Valid values: QuantityChangeInCapacity, PercentChangeInCapacity, TotalCapacity.
        /// </summary>
        [Input("adjustmentType", required: true)]
        public Input<string> AdjustmentType { get; set; } = null!;

        /// <summary>
        /// The adjustment value of the scaling policy. When the value of the `AdjustmentType` parameter is `QuantityChangeInCapacity`: -100 ~ 100, 0 is not allowed, unit: piece. When the value of the `AdjustmentType` parameter is `PercentChangeInCapacity`: -100 ~ 10000, 0 is not allowed, unit: %. When the value of the `AdjustmentType` parameter is `TotalCapacity`: the default is 0 to 100, unit: piece.
        /// </summary>
        [Input("adjustmentValue", required: true)]
        public Input<int> AdjustmentValue { get; set; } = null!;

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. Valid values: `&gt;`, `&lt;`, `=`. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionComparisonOperator")]
        public Input<string>? AlarmPolicyConditionComparisonOperator { get; set; }

        /// <summary>
        /// The metric name of the alarm policy condition of the scaling policy. Valid values: CpuTotal_Max, CpuTotal_Min, CpuTotal_Avg, MemoryUsedUtilization_Max, MemoryUsedUtilization_Min, MemoryUsedUtilization_Avg, Instance_CpuBusy_Max, Instance_CpuBusy_Min, Instance_CpuBusy_Avg.
        /// </summary>
        [Input("alarmPolicyConditionMetricName")]
        public Input<string>? AlarmPolicyConditionMetricName { get; set; }

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionMetricUnit")]
        public Input<string>? AlarmPolicyConditionMetricUnit { get; set; }

        /// <summary>
        /// The threshold of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionThreshold")]
        public Input<string>? AlarmPolicyConditionThreshold { get; set; }

        /// <summary>
        /// The evaluation count of the alarm policy of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyEvaluationCount")]
        public Input<int>? AlarmPolicyEvaluationCount { get; set; }

        /// <summary>
        /// The rule type of the alarm policy of the scaling policy. Valid value: Static. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyRuleType")]
        public Input<string>? AlarmPolicyRuleType { get; set; }

        /// <summary>
        /// The cooldown of the scaling policy. Default value is the cooldown time of the scaling group. Value: 0~86400, unit: second, if left blank, the cooling time of the scaling group will be used by default.
        /// </summary>
        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        /// <summary>
        /// The id of the scaling group to which the scaling policy belongs.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        /// <summary>
        /// The name of the scaling policy.
        /// </summary>
        [Input("scalingPolicyName", required: true)]
        public Input<string> ScalingPolicyName { get; set; } = null!;

        /// <summary>
        /// The type of scaling policy. Valid values: Scheduled, Recurrence, Alarm.
        /// </summary>
        [Input("scalingPolicyType", required: true)]
        public Input<string> ScalingPolicyType { get; set; } = null!;

        /// <summary>
        /// The launch time of the scheduled policy of the scaling policy.
        /// When the value of `ScalingPolicyType` is `Scheduled`, it means that the trigger time of the scheduled task must be greater than the current time.
        /// When the value of `ScalingPolicyType` is `Recurrence`: If `ScheduledPolicy.RecurrenceType` is not specified, it means to execute only once according to the date and time specified here.
        /// If `ScheduledPolicy.RecurrenceType` is specified, it indicates the start time of the periodic task. Only the time within 90 days from the date of creation/modification is supported.
        /// When the value of `ScalingPolicyType` is `Alarm`, this parameter is invalid.
        /// </summary>
        [Input("scheduledPolicyLaunchTime")]
        public Input<string>? ScheduledPolicyLaunchTime { get; set; }

        /// <summary>
        /// The recurrence end time of the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. If not configured, it will default to the day/week/month after this moment according to the recurrence period (ScheduledPolicy.RecurrenceType).
        /// </summary>
        [Input("scheduledPolicyRecurrenceEndTime")]
        public Input<string>? ScheduledPolicyRecurrenceEndTime { get; set; }

        /// <summary>
        /// The recurrence type the scheduled policy of the scaling policy. Valid values: Daily, Weekly, Monthly, Cron.
        /// </summary>
        [Input("scheduledPolicyRecurrenceType")]
        public Input<string>? ScheduledPolicyRecurrenceType { get; set; }

        /// <summary>
        /// The recurrence value the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. When the value of the ScheduledPolicy.RecurrenceType parameter is Daily, only one value can be filled in, ranging from 1 to 31.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Weekly, you can enter multiple values separated by commas (,). The values from Monday to Sunday are: 1,2,3,4,5,6,7.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Monthly, the format is A-B. The value ranges of A and B are both 1~31, and B must be greater than or equal to A.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Cron, it means UTC+8 time, supports 5-field expressions of minutes, hours, days, months, and weeks, and supports wildcard English commas (,), English question marks (?), and conjunctions ( -), asterisk (*), pound sign (#), slash (/), L, and W.
        /// </summary>
        [Input("scheduledPolicyRecurrenceValue")]
        public Input<string>? ScheduledPolicyRecurrenceValue { get; set; }

        public ScalingPolicyArgs()
        {
        }
        public static new ScalingPolicyArgs Empty => new ScalingPolicyArgs();
    }

    public sealed class ScalingPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The active flag of the scaling policy. [Warning] the scaling policy can be active only when the scaling group be active otherwise will fail.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The adjustment type of the scaling policy. Valid values: QuantityChangeInCapacity, PercentChangeInCapacity, TotalCapacity.
        /// </summary>
        [Input("adjustmentType")]
        public Input<string>? AdjustmentType { get; set; }

        /// <summary>
        /// The adjustment value of the scaling policy. When the value of the `AdjustmentType` parameter is `QuantityChangeInCapacity`: -100 ~ 100, 0 is not allowed, unit: piece. When the value of the `AdjustmentType` parameter is `PercentChangeInCapacity`: -100 ~ 10000, 0 is not allowed, unit: %. When the value of the `AdjustmentType` parameter is `TotalCapacity`: the default is 0 to 100, unit: piece.
        /// </summary>
        [Input("adjustmentValue")]
        public Input<int>? AdjustmentValue { get; set; }

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. Valid values: `&gt;`, `&lt;`, `=`. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionComparisonOperator")]
        public Input<string>? AlarmPolicyConditionComparisonOperator { get; set; }

        /// <summary>
        /// The metric name of the alarm policy condition of the scaling policy. Valid values: CpuTotal_Max, CpuTotal_Min, CpuTotal_Avg, MemoryUsedUtilization_Max, MemoryUsedUtilization_Min, MemoryUsedUtilization_Avg, Instance_CpuBusy_Max, Instance_CpuBusy_Min, Instance_CpuBusy_Avg.
        /// </summary>
        [Input("alarmPolicyConditionMetricName")]
        public Input<string>? AlarmPolicyConditionMetricName { get; set; }

        /// <summary>
        /// The comparison operator of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionMetricUnit")]
        public Input<string>? AlarmPolicyConditionMetricUnit { get; set; }

        /// <summary>
        /// The threshold of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyConditionThreshold")]
        public Input<string>? AlarmPolicyConditionThreshold { get; set; }

        /// <summary>
        /// The evaluation count of the alarm policy of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyEvaluationCount")]
        public Input<int>? AlarmPolicyEvaluationCount { get; set; }

        /// <summary>
        /// The rule type of the alarm policy of the scaling policy. Valid value: Static. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.
        /// </summary>
        [Input("alarmPolicyRuleType")]
        public Input<string>? AlarmPolicyRuleType { get; set; }

        /// <summary>
        /// The cooldown of the scaling policy. Default value is the cooldown time of the scaling group. Value: 0~86400, unit: second, if left blank, the cooling time of the scaling group will be used by default.
        /// </summary>
        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        /// <summary>
        /// The id of the scaling group to which the scaling policy belongs.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        /// <summary>
        /// The name of the scaling policy.
        /// </summary>
        [Input("scalingPolicyName")]
        public Input<string>? ScalingPolicyName { get; set; }

        /// <summary>
        /// The type of scaling policy. Valid values: Scheduled, Recurrence, Alarm.
        /// </summary>
        [Input("scalingPolicyType")]
        public Input<string>? ScalingPolicyType { get; set; }

        /// <summary>
        /// The launch time of the scheduled policy of the scaling policy.
        /// When the value of `ScalingPolicyType` is `Scheduled`, it means that the trigger time of the scheduled task must be greater than the current time.
        /// When the value of `ScalingPolicyType` is `Recurrence`: If `ScheduledPolicy.RecurrenceType` is not specified, it means to execute only once according to the date and time specified here.
        /// If `ScheduledPolicy.RecurrenceType` is specified, it indicates the start time of the periodic task. Only the time within 90 days from the date of creation/modification is supported.
        /// When the value of `ScalingPolicyType` is `Alarm`, this parameter is invalid.
        /// </summary>
        [Input("scheduledPolicyLaunchTime")]
        public Input<string>? ScheduledPolicyLaunchTime { get; set; }

        /// <summary>
        /// The recurrence end time of the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. If not configured, it will default to the day/week/month after this moment according to the recurrence period (ScheduledPolicy.RecurrenceType).
        /// </summary>
        [Input("scheduledPolicyRecurrenceEndTime")]
        public Input<string>? ScheduledPolicyRecurrenceEndTime { get; set; }

        /// <summary>
        /// The recurrence type the scheduled policy of the scaling policy. Valid values: Daily, Weekly, Monthly, Cron.
        /// </summary>
        [Input("scheduledPolicyRecurrenceType")]
        public Input<string>? ScheduledPolicyRecurrenceType { get; set; }

        /// <summary>
        /// The recurrence value the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. When the value of the ScheduledPolicy.RecurrenceType parameter is Daily, only one value can be filled in, ranging from 1 to 31.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Weekly, you can enter multiple values separated by commas (,). The values from Monday to Sunday are: 1,2,3,4,5,6,7.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Monthly, the format is A-B. The value ranges of A and B are both 1~31, and B must be greater than or equal to A.
        /// When the value of the ScheduledPolicy.RecurrenceType parameter is Cron, it means UTC+8 time, supports 5-field expressions of minutes, hours, days, months, and weeks, and supports wildcard English commas (,), English question marks (?), and conjunctions ( -), asterisk (*), pound sign (#), slash (/), L, and W.
        /// </summary>
        [Input("scheduledPolicyRecurrenceValue")]
        public Input<string>? ScheduledPolicyRecurrenceValue { get; set; }

        /// <summary>
        /// The status of the scaling policy. Valid values: Active, InActive.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ScalingPolicyState()
        {
        }
        public static new ScalingPolicyState Empty => new ScalingPolicyState();
    }
}
