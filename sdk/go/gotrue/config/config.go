// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetHeaders(ctx *pulumi.Context) string {
	return config.Get(ctx, "gotrue:headers")
}
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "gotrue:url")
}
