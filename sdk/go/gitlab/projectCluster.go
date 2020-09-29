// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_project\_cluster
//
// This resource allows you to create and manage project clusters for your GitLab projects.
// For further information on clusters, consult the [gitlab
// documentation](https://docs.gitlab.com/ce/user/project/clusters/index.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := gitlab.NewProject(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewProjectCluster(ctx, "bar", &gitlab.ProjectClusterArgs{
// 			Domain:                      pulumi.String("example.com"),
// 			Enabled:                     pulumi.Bool(true),
// 			EnvironmentScope:            pulumi.String("*"),
// 			KubernetesApiUrl:            pulumi.String("https://124.124.124"),
// 			KubernetesAuthorizationType: pulumi.String("rbac"),
// 			KubernetesCaCert:            pulumi.String("some-cert"),
// 			KubernetesNamespace:         pulumi.String("namespace"),
// 			KubernetesToken:             pulumi.String("some-token"),
// 			ManagementProjectId:         pulumi.String("123456"),
// 			Project:                     foo.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectCluster struct {
	pulumi.CustomResourceState

	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	CreatedAt   pulumi.StringOutput `pulumi:"createdAt"`
	// The base domain of the cluster.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrOutput `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringOutput `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrOutput `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrOutput `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrOutput `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringOutput `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrOutput `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrOutput `pulumi:"managementProjectId"`
	// The name of cluster.
	Name         pulumi.StringOutput `pulumi:"name"`
	PlatformType pulumi.StringOutput `pulumi:"platformType"`
	// The id of the project to add the cluster to.
	Project      pulumi.StringOutput `pulumi:"project"`
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
}

// NewProjectCluster registers a new resource with the given unique name, arguments, and options.
func NewProjectCluster(ctx *pulumi.Context,
	name string, args *ProjectClusterArgs, opts ...pulumi.ResourceOption) (*ProjectCluster, error) {
	if args == nil || args.KubernetesApiUrl == nil {
		return nil, errors.New("missing required argument 'KubernetesApiUrl'")
	}
	if args == nil || args.KubernetesToken == nil {
		return nil, errors.New("missing required argument 'KubernetesToken'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &ProjectClusterArgs{}
	}
	var resource ProjectCluster
	err := ctx.RegisterResource("gitlab:index/projectCluster:ProjectCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectCluster gets an existing ProjectCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectClusterState, opts ...pulumi.ResourceOption) (*ProjectCluster, error) {
	var resource ProjectCluster
	err := ctx.ReadResource("gitlab:index/projectCluster:ProjectCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectCluster resources.
type projectClusterState struct {
	ClusterType *string `pulumi:"clusterType"`
	CreatedAt   *string `pulumi:"createdAt"`
	// The base domain of the cluster.
	Domain *string `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled *bool `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl *string `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType *string `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the project.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken *string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name         *string `pulumi:"name"`
	PlatformType *string `pulumi:"platformType"`
	// The id of the project to add the cluster to.
	Project      *string `pulumi:"project"`
	ProviderType *string `pulumi:"providerType"`
}

type ProjectClusterState struct {
	ClusterType pulumi.StringPtrInput
	CreatedAt   pulumi.StringPtrInput
	// The base domain of the cluster.
	Domain pulumi.StringPtrInput
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrInput
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrInput
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringPtrInput
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrInput
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrInput
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringPtrInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name         pulumi.StringPtrInput
	PlatformType pulumi.StringPtrInput
	// The id of the project to add the cluster to.
	Project      pulumi.StringPtrInput
	ProviderType pulumi.StringPtrInput
}

func (ProjectClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectClusterState)(nil)).Elem()
}

type projectClusterArgs struct {
	// The base domain of the cluster.
	Domain *string `pulumi:"domain"`
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled *bool `pulumi:"enabled"`
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope *string `pulumi:"environmentScope"`
	// The URL to access the Kubernetes API.
	KubernetesApiUrl string `pulumi:"kubernetesApiUrl"`
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType *string `pulumi:"kubernetesAuthorizationType"`
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// The unique namespace related to the project.
	KubernetesNamespace *string `pulumi:"kubernetesNamespace"`
	// The token to authenticate against Kubernetes.
	KubernetesToken string `pulumi:"kubernetesToken"`
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed *bool `pulumi:"managed"`
	// The ID of the management project for the cluster.
	ManagementProjectId *string `pulumi:"managementProjectId"`
	// The name of cluster.
	Name *string `pulumi:"name"`
	// The id of the project to add the cluster to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectCluster resource.
type ProjectClusterArgs struct {
	// The base domain of the cluster.
	Domain pulumi.StringPtrInput
	// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
	Enabled pulumi.BoolPtrInput
	// The associated environment to the cluster. Defaults to `*`.
	EnvironmentScope pulumi.StringPtrInput
	// The URL to access the Kubernetes API.
	KubernetesApiUrl pulumi.StringInput
	// The cluster authorization type. Valid values are `rbac`, `abac`, `unknownAuthorization`. Defaults to `rbac`.
	KubernetesAuthorizationType pulumi.StringPtrInput
	// TLS certificate (needed if API is using a self-signed TLS certificate).
	KubernetesCaCert pulumi.StringPtrInput
	// The unique namespace related to the project.
	KubernetesNamespace pulumi.StringPtrInput
	// The token to authenticate against Kubernetes.
	KubernetesToken pulumi.StringInput
	// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
	Managed pulumi.BoolPtrInput
	// The ID of the management project for the cluster.
	ManagementProjectId pulumi.StringPtrInput
	// The name of cluster.
	Name pulumi.StringPtrInput
	// The id of the project to add the cluster to.
	Project pulumi.StringInput
}

func (ProjectClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectClusterArgs)(nil)).Elem()
}
