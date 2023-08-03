// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of listeners
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.Zones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooClb, err := clb.NewClb(ctx, "fooClb", &clb.ClbArgs{
//				Type:             pulumi.String("public"),
//				SubnetId:         fooSubnet.ID(),
//				LoadBalancerSpec: pulumi.String("small_1"),
//				Description:      pulumi.String("acc0Demo"),
//				LoadBalancerName: pulumi.String("acc-test-create"),
//				EipBillingConfig: &clb.ClbEipBillingConfigArgs{
//					Isp:            pulumi.String("BGP"),
//					EipBillingType: pulumi.String("PostPaidByBandwidth"),
//					Bandwidth:      pulumi.Int(1),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooServerGroup, err := clb.NewServerGroup(ctx, "fooServerGroup", &clb.ServerGroupArgs{
//				LoadBalancerId:  fooClb.ID(),
//				ServerGroupName: pulumi.String("acc-test-create"),
//				Description:     pulumi.String("hello demo11"),
//			})
//			if err != nil {
//				return err
//			}
//			fooListener, err := clb.NewListener(ctx, "fooListener", &clb.ListenerArgs{
//				LoadBalancerId: fooClb.ID(),
//				ListenerName:   pulumi.String("acc-test-listener"),
//				Protocol:       pulumi.String("HTTP"),
//				Port:           pulumi.Int(90),
//				ServerGroupId:  fooServerGroup.ID(),
//				HealthCheck: &clb.ListenerHealthCheckArgs{
//					Enabled:            pulumi.String("on"),
//					Interval:           pulumi.Int(10),
//					Timeout:            pulumi.Int(3),
//					HealthyThreshold:   pulumi.Int(5),
//					UnHealthyThreshold: pulumi.Int(2),
//					Domain:             pulumi.String("volcengine.com"),
//					HttpCode:           pulumi.String("http_2xx"),
//					Method:             pulumi.String("GET"),
//					Uri:                pulumi.String("/"),
//				},
//				Enabled: pulumi.String("on"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = clb.ListenersOutput(ctx, clb.ListenersOutputArgs{
//				Ids: pulumi.StringArray{
//					fooListener.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func Listeners(ctx *pulumi.Context, args *ListenersArgs, opts ...pulumi.InvokeOption) (*ListenersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListenersResult
	err := ctx.Invoke("volcengine:clb/listeners:Listeners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Listeners.
type ListenersArgs struct {
	// A list of Listener IDs.
	Ids []string `pulumi:"ids"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The id of the Clb.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A Name Regex of Listener.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Listeners.
type ListenersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The collection of Listener query.
	Listeners      []ListenersListener `pulumi:"listeners"`
	LoadBalancerId *string             `pulumi:"loadBalancerId"`
	NameRegex      *string             `pulumi:"nameRegex"`
	OutputFile     *string             `pulumi:"outputFile"`
	// The total count of Listener query.
	TotalCount int `pulumi:"totalCount"`
}

func ListenersOutput(ctx *pulumi.Context, args ListenersOutputArgs, opts ...pulumi.InvokeOption) ListenersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListenersResult, error) {
			args := v.(ListenersArgs)
			r, err := Listeners(ctx, &args, opts...)
			var s ListenersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListenersResultOutput)
}

// A collection of arguments for invoking Listeners.
type ListenersOutputArgs struct {
	// A list of Listener IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of the Listener.
	ListenerName pulumi.StringPtrInput `pulumi:"listenerName"`
	// The id of the Clb.
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	// A Name Regex of Listener.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (ListenersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenersArgs)(nil)).Elem()
}

// A collection of values returned by Listeners.
type ListenersResultOutput struct{ *pulumi.OutputState }

func (ListenersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenersResult)(nil)).Elem()
}

func (o ListenersResultOutput) ToListenersResultOutput() ListenersResultOutput {
	return o
}

func (o ListenersResultOutput) ToListenersResultOutputWithContext(ctx context.Context) ListenersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ListenersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListenersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListenersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListenersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of the Listener.
func (o ListenersResultOutput) ListenerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.ListenerName }).(pulumi.StringPtrOutput)
}

// The collection of Listener query.
func (o ListenersResultOutput) Listeners() ListenersListenerArrayOutput {
	return o.ApplyT(func(v ListenersResult) []ListenersListener { return v.Listeners }).(ListenersListenerArrayOutput)
}

func (o ListenersResultOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.LoadBalancerId }).(pulumi.StringPtrOutput)
}

func (o ListenersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ListenersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of Listener query.
func (o ListenersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ListenersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenersResultOutput{})
}
