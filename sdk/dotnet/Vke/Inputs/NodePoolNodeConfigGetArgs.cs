// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolNodeConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AdditionalContainerStorageEnabled of NodeConfig.
        /// </summary>
        [Input("additionalContainerStorageEnabled")]
        public Input<bool>? AdditionalContainerStorageEnabled { get; set; }

        /// <summary>
        /// Is AutoRenew of PrePaid instance of NodeConfig. Valid values: true, false. when InstanceChargeType is PrePaid, default value is true.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The AutoRenewPeriod of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 6, 12. Unit: month. when InstanceChargeType is PrePaid and AutoRenew enable, default value is 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        [Input("dataVolumes")]
        private InputList<Inputs.NodePoolNodeConfigDataVolumeGetArgs>? _dataVolumes;

        /// <summary>
        /// The DataVolumes of NodeConfig.
        /// </summary>
        public InputList<Inputs.NodePoolNodeConfigDataVolumeGetArgs> DataVolumes
        {
            get => _dataVolumes ?? (_dataVolumes = new InputList<Inputs.NodePoolNodeConfigDataVolumeGetArgs>());
            set => _dataVolumes = value;
        }

        [Input("ecsTags")]
        private InputList<Inputs.NodePoolNodeConfigEcsTagGetArgs>? _ecsTags;

        /// <summary>
        /// Tags for Ecs.
        /// </summary>
        public InputList<Inputs.NodePoolNodeConfigEcsTagGetArgs> EcsTags
        {
            get => _ecsTags ?? (_ecsTags = new InputList<Inputs.NodePoolNodeConfigEcsTagGetArgs>());
            set => _ecsTags = value;
        }

        [Input("hpcClusterIds")]
        private InputList<string>? _hpcClusterIds;

        /// <summary>
        /// The IDs of HpcCluster, only one ID is supported currently.
        /// </summary>
        public InputList<string> HpcClusterIds
        {
            get => _hpcClusterIds ?? (_hpcClusterIds = new InputList<string>());
            set => _hpcClusterIds = value;
        }

        /// <summary>
        /// The ImageId of NodeConfig.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The initializeScript of NodeConfig.
        /// </summary>
        [Input("initializeScript")]
        public Input<string>? InitializeScript { get; set; }

        /// <summary>
        /// The InstanceChargeType of PrePaid instance of NodeConfig. Valid values: PostPaid, PrePaid. Default value: PostPaid.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        [Input("instanceTypeIds", required: true)]
        private InputList<string>? _instanceTypeIds;

        /// <summary>
        /// The InstanceTypeIds of NodeConfig.
        /// </summary>
        public InputList<string> InstanceTypeIds
        {
            get => _instanceTypeIds ?? (_instanceTypeIds = new InputList<string>());
            set => _instanceTypeIds = value;
        }

        /// <summary>
        /// The NamePrefix of NodeConfig.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The Period of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36. Unit: month. when InstanceChargeType is PrePaid, default value is 12.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The Security of NodeConfig.
        /// </summary>
        [Input("security", required: true)]
        public Input<Inputs.NodePoolNodeConfigSecurityGetArgs> Security { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The SubnetIds of NodeConfig.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The SystemVolume of NodeConfig.
        /// </summary>
        [Input("systemVolume")]
        public Input<Inputs.NodePoolNodeConfigSystemVolumeGetArgs>? SystemVolume { get; set; }

        public NodePoolNodeConfigGetArgs()
        {
        }
    }
}
