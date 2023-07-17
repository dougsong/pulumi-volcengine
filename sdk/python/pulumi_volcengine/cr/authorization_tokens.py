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
    'AuthorizationTokensResult',
    'AwaitableAuthorizationTokensResult',
    'authorization_tokens',
    'authorization_tokens_output',
]

@pulumi.output_type
class AuthorizationTokensResult:
    """
    A collection of values returned by AuthorizationTokens.
    """
    def __init__(__self__, id=None, output_file=None, registry=None, tokens=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if registry and not isinstance(registry, str):
            raise TypeError("Expected argument 'registry' to be a str")
        pulumi.set(__self__, "registry", registry)
        if tokens and not isinstance(tokens, list):
            raise TypeError("Expected argument 'tokens' to be a list")
        pulumi.set(__self__, "tokens", tokens)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def registry(self) -> str:
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def tokens(self) -> Sequence['outputs.AuthorizationTokensTokenResult']:
        """
        The collection of users.
        """
        return pulumi.get(self, "tokens")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of instance query.
        """
        return pulumi.get(self, "total_count")


class AwaitableAuthorizationTokensResult(AuthorizationTokensResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AuthorizationTokensResult(
            id=self.id,
            output_file=self.output_file,
            registry=self.registry,
            tokens=self.tokens,
            total_count=self.total_count)


def authorization_tokens(output_file: Optional[str] = None,
                         registry: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAuthorizationTokensResult:
    """
    Use this data source to query detailed information of cr authorization tokens
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cr.authorization_tokens(registry="tf-1")
    ```


    :param str output_file: File name where to save data source results.
    :param str registry: The cr instance name want to query.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    __args__['registry'] = registry
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:cr/authorizationTokens:AuthorizationTokens', __args__, opts=opts, typ=AuthorizationTokensResult).value

    return AwaitableAuthorizationTokensResult(
        id=__ret__.id,
        output_file=__ret__.output_file,
        registry=__ret__.registry,
        tokens=__ret__.tokens,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(authorization_tokens)
def authorization_tokens_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                registry: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AuthorizationTokensResult]:
    """
    Use this data source to query detailed information of cr authorization tokens
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cr.authorization_tokens(registry="tf-1")
    ```


    :param str output_file: File name where to save data source results.
    :param str registry: The cr instance name want to query.
    """
    ...
