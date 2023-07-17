// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds.Outputs
{

    [OutputType]
    public sealed class DatabasesRdsDatabaseResult
    {
        /// <summary>
        /// The account names of the RDS database.
        /// </summary>
        public readonly string AccountNames;
        /// <summary>
        /// The character set of the RDS database.
        /// </summary>
        public readonly string CharacterSetName;
        /// <summary>
        /// The name of the RDS database.
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// The status of the RDS database.
        /// </summary>
        public readonly string DbStatus;
        /// <summary>
        /// The ID of the RDS database.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private DatabasesRdsDatabaseResult(
            string accountNames,

            string characterSetName,

            string dbName,

            string dbStatus,

            string id)
        {
            AccountNames = accountNames;
            CharacterSetName = characterSetName;
            DbName = dbName;
            DbStatus = dbStatus;
            Id = id;
        }
    }
}
