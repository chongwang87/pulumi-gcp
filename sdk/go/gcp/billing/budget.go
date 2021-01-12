// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Budget configuration for a billing account.
//
// To get more information about Budget, see:
//
// * [API documentation](https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets)
// * How-to Guides
//     * [Creating a budget](https://cloud.google.com/billing/docs/how-to/budgets)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the Billing Budgets API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Billing Budget Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/billing"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "000000-0000000-0000000-000000"
// 		account, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = billing.NewBudget(ctx, "budget", &billing.BudgetArgs{
// 			BillingAccount: pulumi.String(account.Id),
// 			DisplayName:    pulumi.String("Example Billing Budget"),
// 			Amount: &billing.BudgetAmountArgs{
// 				SpecifiedAmount: &billing.BudgetAmountSpecifiedAmountArgs{
// 					CurrencyCode: pulumi.String("USD"),
// 					Units:        pulumi.String("100000"),
// 				},
// 			},
// 			ThresholdRules: billing.BudgetThresholdRuleArray{
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(0.5),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Billing Budget Lastperiod
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/billing"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "000000-0000000-0000000-000000"
// 		account, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = billing.NewBudget(ctx, "budget", &billing.BudgetArgs{
// 			BillingAccount: pulumi.String(account.Id),
// 			DisplayName:    pulumi.String("Example Billing Budget"),
// 			BudgetFilter: &billing.BudgetBudgetFilterArgs{
// 				Projects: pulumi.StringArray{
// 					pulumi.String(fmt.Sprintf("%v%v", "projects/", project.Number)),
// 				},
// 			},
// 			Amount: &billing.BudgetAmountArgs{
// 				LastPeriodAmount: pulumi.Bool(true),
// 			},
// 			ThresholdRules: billing.BudgetThresholdRuleArray{
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(10),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Billing Budget Filter
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/billing"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "000000-0000000-0000000-000000"
// 		account, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = billing.NewBudget(ctx, "budget", &billing.BudgetArgs{
// 			BillingAccount: pulumi.String(account.Id),
// 			DisplayName:    pulumi.String("Example Billing Budget"),
// 			BudgetFilter: &billing.BudgetBudgetFilterArgs{
// 				Projects: pulumi.StringArray{
// 					pulumi.String(fmt.Sprintf("%v%v", "projects/", project.Number)),
// 				},
// 				CreditTypesTreatment: pulumi.String("EXCLUDE_ALL_CREDITS"),
// 				Services: pulumi.StringArray{
// 					pulumi.String("services/24E6-581D-38E5"),
// 				},
// 			},
// 			Amount: &billing.BudgetAmountArgs{
// 				SpecifiedAmount: &billing.BudgetAmountSpecifiedAmountArgs{
// 					CurrencyCode: pulumi.String("USD"),
// 					Units:        pulumi.String("100000"),
// 				},
// 			},
// 			ThresholdRules: billing.BudgetThresholdRuleArray{
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(0.5),
// 				},
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(0.9),
// 					SpendBasis:       pulumi.String("FORECASTED_SPEND"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Billing Budget Notify
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/billing"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/monitoring"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "000000-0000000-0000000-000000"
// 		account, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		notificationChannel, err := monitoring.NewNotificationChannel(ctx, "notificationChannel", &monitoring.NotificationChannelArgs{
// 			DisplayName: pulumi.String("Example Notification Channel"),
// 			Type:        pulumi.String("email"),
// 			Labels: pulumi.StringMap{
// 				"email_address": pulumi.String("address@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = billing.NewBudget(ctx, "budget", &billing.BudgetArgs{
// 			BillingAccount: pulumi.String(account.Id),
// 			DisplayName:    pulumi.String("Example Billing Budget"),
// 			BudgetFilter: &billing.BudgetBudgetFilterArgs{
// 				Projects: pulumi.StringArray{
// 					pulumi.String(fmt.Sprintf("%v%v", "projects/", project.Number)),
// 				},
// 			},
// 			Amount: &billing.BudgetAmountArgs{
// 				SpecifiedAmount: &billing.BudgetAmountSpecifiedAmountArgs{
// 					CurrencyCode: pulumi.String("USD"),
// 					Units:        pulumi.String("100000"),
// 				},
// 			},
// 			ThresholdRules: billing.BudgetThresholdRuleArray{
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(1),
// 				},
// 				&billing.BudgetThresholdRuleArgs{
// 					ThresholdPercent: pulumi.Float64(1),
// 					SpendBasis:       pulumi.String("FORECASTED_SPEND"),
// 				},
// 			},
// 			AllUpdatesRule: &billing.BudgetAllUpdatesRuleArgs{
// 				MonitoringNotificationChannels: pulumi.StringArray{
// 					notificationChannel.ID(),
// 				},
// 				DisableDefaultIamRecipients: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource does not support import.
type Budget struct {
	pulumi.CustomResourceState

	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// Structure is documented below.
	AllUpdatesRule BudgetAllUpdatesRulePtrOutput `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	// Structure is documented below.
	Amount BudgetAmountOutput `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// Structure is documented below.
	BudgetFilter BudgetBudgetFilterOutput `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// Structure is documented below.
	ThresholdRules BudgetThresholdRuleArrayOutput `pulumi:"thresholdRules"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	if args.ThresholdRules == nil {
		return nil, errors.New("invalid value for required argument 'ThresholdRules'")
	}
	var resource Budget
	err := ctx.RegisterResource("gcp:billing/budget:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("gcp:billing/budget:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Budget resources.
type budgetState struct {
	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// Structure is documented below.
	AllUpdatesRule *BudgetAllUpdatesRule `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	// Structure is documented below.
	Amount *BudgetAmount `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount *string `pulumi:"billingAccount"`
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// Structure is documented below.
	BudgetFilter *BudgetBudgetFilter `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name *string `pulumi:"name"`
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// Structure is documented below.
	ThresholdRules []BudgetThresholdRule `pulumi:"thresholdRules"`
}

type BudgetState struct {
	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// Structure is documented below.
	AllUpdatesRule BudgetAllUpdatesRulePtrInput
	// The budgeted amount for each usage period.
	// Structure is documented below.
	Amount BudgetAmountPtrInput
	// ID of the billing account to set a budget on.
	BillingAccount pulumi.StringPtrInput
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// Structure is documented below.
	BudgetFilter BudgetBudgetFilterPtrInput
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name pulumi.StringPtrInput
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// Structure is documented below.
	ThresholdRules BudgetThresholdRuleArrayInput
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// Structure is documented below.
	AllUpdatesRule *BudgetAllUpdatesRule `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	// Structure is documented below.
	Amount BudgetAmount `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount string `pulumi:"billingAccount"`
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// Structure is documented below.
	BudgetFilter *BudgetBudgetFilter `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// Structure is documented below.
	ThresholdRules []BudgetThresholdRule `pulumi:"thresholdRules"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// Structure is documented below.
	AllUpdatesRule BudgetAllUpdatesRulePtrInput
	// The budgeted amount for each usage period.
	// Structure is documented below.
	Amount BudgetAmountInput
	// ID of the billing account to set a budget on.
	BillingAccount pulumi.StringInput
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// Structure is documented below.
	BudgetFilter BudgetBudgetFilterPtrInput
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	// Structure is documented below.
	ThresholdRules BudgetThresholdRuleArrayInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (Budget) ElementType() reflect.Type {
	return reflect.TypeOf((*Budget)(nil)).Elem()
}

func (i Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

type BudgetOutput struct {
	*pulumi.OutputState
}

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetOutput)(nil)).Elem()
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BudgetOutput{})
}
