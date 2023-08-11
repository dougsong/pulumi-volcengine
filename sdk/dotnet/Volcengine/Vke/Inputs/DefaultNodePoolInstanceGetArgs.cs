// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Inputs
{

    public sealed class DefaultNodePoolInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.
        /// </summary>
        [Input("additionalContainerStorageEnabled")]
        public Input<bool>? AdditionalContainerStorageEnabled { get; set; }

        /// <summary>
        /// The container storage path.When additional_container_storage_enabled is `false` will ignore.
        /// </summary>
        [Input("containerStoragePath")]
        public Input<string>? ContainerStoragePath { get; set; }

        /// <summary>
        /// ID of the resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The Image Id to the ECS Instance.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The flag of keep instance name, the value is `true` or `false`.Default is `false`.
        /// </summary>
        [Input("keepInstanceName")]
        public Input<bool>? KeepInstanceName { get; set; }

        [Input("phase")]
        public Input<string>? Phase { get; set; }

        public DefaultNodePoolInstanceGetArgs()
        {
        }
        public static new DefaultNodePoolInstanceGetArgs Empty => new DefaultNodePoolInstanceGetArgs();
    }
}
