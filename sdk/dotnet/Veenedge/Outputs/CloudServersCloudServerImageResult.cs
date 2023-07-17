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
    public sealed class CloudServersCloudServerImageResult
    {
        /// <summary>
        /// The id of image.
        /// </summary>
        public readonly string ImageIdentity;
        /// <summary>
        /// The name of image.
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// The property of system.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// The arch of system.
        /// </summary>
        public readonly string SystemArch;
        /// <summary>
        /// The bit of system.
        /// </summary>
        public readonly string SystemBit;
        /// <summary>
        /// The type of system.
        /// </summary>
        public readonly string SystemType;
        /// <summary>
        /// The version of system.
        /// </summary>
        public readonly string SystemVersion;

        [OutputConstructor]
        private CloudServersCloudServerImageResult(
            string imageIdentity,

            string imageName,

            string property,

            string systemArch,

            string systemBit,

            string systemType,

            string systemVersion)
        {
            ImageIdentity = imageIdentity;
            ImageName = imageName;
            Property = property;
            SystemArch = systemArch;
            SystemBit = systemBit;
            SystemType = systemType;
            SystemVersion = systemVersion;
        }
    }
}
