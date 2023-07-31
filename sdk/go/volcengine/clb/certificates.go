// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of certificates
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.Certificates(ctx, &clb.CertificatesArgs{
//				Ids: []string{
//					"cert-274scdwqufwg07fap8u5fu8pi",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Certificates(ctx *pulumi.Context, args *CertificatesArgs, opts ...pulumi.InvokeOption) (*CertificatesResult, error) {
	var rv CertificatesResult
	err := ctx.Invoke("volcengine:clb/certificates:Certificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Certificates.
type CertificatesArgs struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The list of Certificate IDs.
	Ids []string `pulumi:"ids"`
	// The Name Regex of Certificate.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ProjectName of Certificate.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []CertificatesTag `pulumi:"tags"`
}

// A collection of values returned by Certificates.
type CertificatesResult struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The collection of Certificate query.
	Certificates []CertificatesCertificate `pulumi:"certificates"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ProjectName of the Certificate.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []CertificatesTag `pulumi:"tags"`
	// The total count of Certificate query.
	TotalCount int `pulumi:"totalCount"`
}

func CertificatesOutput(ctx *pulumi.Context, args CertificatesOutputArgs, opts ...pulumi.InvokeOption) CertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CertificatesResult, error) {
			args := v.(CertificatesArgs)
			r, err := Certificates(ctx, &args, opts...)
			var s CertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CertificatesResultOutput)
}

// A collection of arguments for invoking Certificates.
type CertificatesOutputArgs struct {
	// The name of the Certificate.
	CertificateName pulumi.StringPtrInput `pulumi:"certificateName"`
	// The list of Certificate IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name Regex of Certificate.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ProjectName of Certificate.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags CertificatesTagArrayInput `pulumi:"tags"`
}

func (CertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatesArgs)(nil)).Elem()
}

// A collection of values returned by Certificates.
type CertificatesResultOutput struct{ *pulumi.OutputState }

func (CertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatesResult)(nil)).Elem()
}

func (o CertificatesResultOutput) ToCertificatesResultOutput() CertificatesResultOutput {
	return o
}

func (o CertificatesResultOutput) ToCertificatesResultOutputWithContext(ctx context.Context) CertificatesResultOutput {
	return o
}

// The name of the Certificate.
func (o CertificatesResultOutput) CertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatesResult) *string { return v.CertificateName }).(pulumi.StringPtrOutput)
}

// The collection of Certificate query.
func (o CertificatesResultOutput) Certificates() CertificatesCertificateArrayOutput {
	return o.ApplyT(func(v CertificatesResult) []CertificatesCertificate { return v.Certificates }).(CertificatesCertificateArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o CertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o CertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o CertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o CertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ProjectName of the Certificate.
func (o CertificatesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o CertificatesResultOutput) Tags() CertificatesTagArrayOutput {
	return o.ApplyT(func(v CertificatesResult) []CertificatesTag { return v.Tags }).(CertificatesTagArrayOutput)
}

// The total count of Certificate query.
func (o CertificatesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v CertificatesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificatesResultOutput{})
}
