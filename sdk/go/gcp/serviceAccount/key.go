// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates and manages service account key-pairs, which allow the user to establish identity of a service account outside of GCP. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys).
//
// ## Example Usage
// ### Creating A New Key Pair
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myaccount, err := serviceAccount.NewAccount(ctx, "myaccount", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("myaccount"),
// 			DisplayName: pulumi.String("My Service Account"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewKey(ctx, "mykey", &serviceAccount.KeyArgs{
// 			ServiceAccountId: myaccount.Name,
// 			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Key struct {
	pulumi.CustomResourceState

	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrOutput `pulumi:"keyAlgorithm"`
	// The name used for this key pair
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrOutput `pulumi:"privateKeyType"`
	// The public key, base64 encoded
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The output format of the public key requested. X509_PEM is the default output format.
	PublicKeyType pulumi.StringPtrOutput `pulumi:"publicKeyType"`
	// The Service account id of the Key Pair. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
	// unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter pulumi.StringOutput `pulumi:"validAfter"`
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore pulumi.StringOutput `pulumi:"validBefore"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	if args == nil {
		args = &KeyArgs{}
	}
	var resource Key
	err := ctx.RegisterResource("gcp:serviceAccount/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("gcp:serviceAccount/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// The name used for this key pair
	Name *string `pulumi:"name"`
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey *string `pulumi:"privateKey"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The public key, base64 encoded
	PublicKey *string `pulumi:"publicKey"`
	// The output format of the public key requested. X509_PEM is the default output format.
	PublicKeyType *string `pulumi:"publicKeyType"`
	// The Service account id of the Key Pair. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
	// unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter *string `pulumi:"validAfter"`
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore *string `pulumi:"validBefore"`
}

type KeyState struct {
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrInput
	// The name used for this key pair
	Name pulumi.StringPtrInput
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey pulumi.StringPtrInput
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrInput
	// The public key, base64 encoded
	PublicKey pulumi.StringPtrInput
	// The output format of the public key requested. X509_PEM is the default output format.
	PublicKeyType pulumi.StringPtrInput
	// The Service account id of the Key Pair. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
	// unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
	ServiceAccountId pulumi.StringPtrInput
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter pulumi.StringPtrInput
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The output format of the public key requested. X509_PEM is the default output format.
	PublicKeyType *string `pulumi:"publicKeyType"`
	// The Service account id of the Key Pair. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
	// unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrInput
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrInput
	// The output format of the public key requested. X509_PEM is the default output format.
	PublicKeyType pulumi.StringPtrInput
	// The Service account id of the Key Pair. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
	// unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
	ServiceAccountId pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}
