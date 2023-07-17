// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceParamResult
    {
        /// <summary>
        /// Current value of the configuration parameter.
        /// </summary>
        public readonly string CurrentValue;
        /// <summary>
        /// Default value of the configuration parameter.
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// The description of this option item.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the current redis instance supports editing this parameter.
        /// </summary>
        public readonly bool EditableForInstance;
        /// <summary>
        /// Whether need to reboot the redis instance when modifying this parameter.
        /// </summary>
        public readonly bool NeedReboot;
        /// <summary>
        /// The list of options. Valid when the configuration parameter type is `Radio`.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceParamOptionResult> Options;
        /// <summary>
        /// The name of the configuration parameter.
        /// </summary>
        public readonly string ParamName;
        /// <summary>
        /// The valid value range of the numeric type configuration parameter.
        /// </summary>
        public readonly string Range;
        /// <summary>
        /// The type of the configuration parameter.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unit of the numeric type configuration parameter.
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private InstancesInstanceParamResult(
            string currentValue,

            string defaultValue,

            string description,

            bool editableForInstance,

            bool needReboot,

            ImmutableArray<Outputs.InstancesInstanceParamOptionResult> options,

            string paramName,

            string range,

            string type,

            string unit)
        {
            CurrentValue = currentValue;
            DefaultValue = defaultValue;
            Description = description;
            EditableForInstance = editableForInstance;
            NeedReboot = needReboot;
            Options = options;
            ParamName = paramName;
            Range = range;
            Type = type;
            Unit = unit;
        }
    }
}
