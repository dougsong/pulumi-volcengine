# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AccountsResult',
    'AwaitableAccountsResult',
    'accounts',
    'accounts_output',
]

@pulumi.output_type
class AccountsResult:
    """
    A collection of values returned by Accounts.
    """
    def __init__(__self__, account_name=None, accounts=None, id=None, instance_id=None, output_file=None, total_count=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if accounts and not isinstance(accounts, list):
            raise TypeError("Expected argument 'accounts' to be a list")
        pulumi.set(__self__, "accounts", accounts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[str]:
        """
        The name of account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter
    def accounts(self) -> Sequence['outputs.AccountsAccountResult']:
        """
        The collection of accounts query.
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of accounts query.
        """
        return pulumi.get(self, "total_count")


class AwaitableAccountsResult(AccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AccountsResult(
            account_name=self.account_name,
            accounts=self.accounts,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            total_count=self.total_count)


def accounts(account_name: Optional[str] = None,
             instance_id: Optional[str] = None,
             output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAccountsResult:
    """
    Use this data source to query detailed information of mongodb accounts
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.mongodb.accounts(instance_id="mongo-replica-xxx")
    ```


    :param str account_name: The name of account, current support only `root`.
    :param str instance_id: Target query mongo instance id.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:mongodb/accounts:Accounts', __args__, opts=opts, typ=AccountsResult).value

    return AwaitableAccountsResult(
        account_name=__ret__.account_name,
        accounts=__ret__.accounts,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(accounts)
def accounts_output(account_name: Optional[pulumi.Input[Optional[str]]] = None,
                    instance_id: Optional[pulumi.Input[str]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AccountsResult]:
    """
    Use this data source to query detailed information of mongodb accounts
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.mongodb.accounts(instance_id="mongo-replica-xxx")
    ```


    :param str account_name: The name of account, current support only `root`.
    :param str instance_id: Target query mongo instance id.
    :param str output_file: File name where to save data source results.
    """
    ...
