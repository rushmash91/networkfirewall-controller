// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/networkfirewall"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/networkfirewall-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.NetworkFirewall{}
	_ = &svcapitypes.Firewall{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeFirewallOutput
	resp, err = rm.sdkapi.DescribeFirewallWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeFirewall", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Firewall != nil {
		f0 := &svcapitypes.Firewall_SDK{}
		if resp.Firewall.DeleteProtection != nil {
			f0.DeleteProtection = resp.Firewall.DeleteProtection
		}
		if resp.Firewall.Description != nil {
			f0.Description = resp.Firewall.Description
		}
		if resp.Firewall.EncryptionConfiguration != nil {
			f0f2 := &svcapitypes.EncryptionConfiguration{}
			if resp.Firewall.EncryptionConfiguration.KeyId != nil {
				f0f2.KeyID = resp.Firewall.EncryptionConfiguration.KeyId
			}
			if resp.Firewall.EncryptionConfiguration.Type != nil {
				f0f2.Type = resp.Firewall.EncryptionConfiguration.Type
			}
			f0.EncryptionConfiguration = f0f2
		}
		if resp.Firewall.FirewallArn != nil {
			f0.FirewallARN = resp.Firewall.FirewallArn
		}
		if resp.Firewall.FirewallId != nil {
			f0.FirewallID = resp.Firewall.FirewallId
		}
		if resp.Firewall.FirewallName != nil {
			f0.FirewallName = resp.Firewall.FirewallName
		}
		if resp.Firewall.FirewallPolicyArn != nil {
			f0.FirewallPolicyARN = resp.Firewall.FirewallPolicyArn
		}
		if resp.Firewall.FirewallPolicyChangeProtection != nil {
			f0.FirewallPolicyChangeProtection = resp.Firewall.FirewallPolicyChangeProtection
		}
		if resp.Firewall.SubnetChangeProtection != nil {
			f0.SubnetChangeProtection = resp.Firewall.SubnetChangeProtection
		}
		if resp.Firewall.SubnetMappings != nil {
			f0f9 := []*svcapitypes.SubnetMapping{}
			for _, f0f9iter := range resp.Firewall.SubnetMappings {
				f0f9elem := &svcapitypes.SubnetMapping{}
				if f0f9iter.IPAddressType != nil {
					f0f9elem.IPAddressType = f0f9iter.IPAddressType
				}
				if f0f9iter.SubnetId != nil {
					f0f9elem.SubnetID = f0f9iter.SubnetId
				}
				f0f9 = append(f0f9, f0f9elem)
			}
			f0.SubnetMappings = f0f9
		}
		if resp.Firewall.Tags != nil {
			f0f10 := []*svcapitypes.Tag{}
			for _, f0f10iter := range resp.Firewall.Tags {
				f0f10elem := &svcapitypes.Tag{}
				if f0f10iter.Key != nil {
					f0f10elem.Key = f0f10iter.Key
				}
				if f0f10iter.Value != nil {
					f0f10elem.Value = f0f10iter.Value
				}
				f0f10 = append(f0f10, f0f10elem)
			}
			f0.Tags = f0f10
		}
		if resp.Firewall.VpcId != nil {
			f0.VPCID = resp.Firewall.VpcId
		}
		ko.Status.Firewall = f0
	} else {
		ko.Status.Firewall = nil
	}
	if resp.FirewallStatus != nil {
		f1 := &svcapitypes.FirewallStatus_SDK{}
		if resp.FirewallStatus.CapacityUsageSummary != nil {
			f1f0 := &svcapitypes.CapacityUsageSummary{}
			if resp.FirewallStatus.CapacityUsageSummary.CIDRs != nil {
				f1f0f0 := &svcapitypes.CIDRSummary{}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.AvailableCIDRCount != nil {
					f1f0f0.AvailableCIDRCount = resp.FirewallStatus.CapacityUsageSummary.CIDRs.AvailableCIDRCount
				}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.IPSetReferences != nil {
					f1f0f0f1 := map[string]*svcapitypes.IPSetMetadata{}
					for f1f0f0f1key, f1f0f0f1valiter := range resp.FirewallStatus.CapacityUsageSummary.CIDRs.IPSetReferences {
						f1f0f0f1val := &svcapitypes.IPSetMetadata{}
						if f1f0f0f1valiter.ResolvedCIDRCount != nil {
							f1f0f0f1val.ResolvedCIDRCount = f1f0f0f1valiter.ResolvedCIDRCount
						}
						f1f0f0f1[f1f0f0f1key] = f1f0f0f1val
					}
					f1f0f0.IPSetReferences = f1f0f0f1
				}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.UtilizedCIDRCount != nil {
					f1f0f0.UtilizedCIDRCount = resp.FirewallStatus.CapacityUsageSummary.CIDRs.UtilizedCIDRCount
				}
				f1f0.CIDRs = f1f0f0
			}
			f1.CapacityUsageSummary = f1f0
		}
		if resp.FirewallStatus.ConfigurationSyncStateSummary != nil {
			f1.ConfigurationSyncStateSummary = resp.FirewallStatus.ConfigurationSyncStateSummary
		}
		if resp.FirewallStatus.Status != nil {
			f1.Status = resp.FirewallStatus.Status
		}
		if resp.FirewallStatus.SyncStates != nil {
			f1f3 := map[string]*svcapitypes.SyncState{}
			for f1f3key, f1f3valiter := range resp.FirewallStatus.SyncStates {
				f1f3val := &svcapitypes.SyncState{}
				if f1f3valiter.Attachment != nil {
					f1f3valf0 := &svcapitypes.Attachment{}
					if f1f3valiter.Attachment.EndpointId != nil {
						f1f3valf0.EndpointID = f1f3valiter.Attachment.EndpointId
					}
					if f1f3valiter.Attachment.Status != nil {
						f1f3valf0.Status = f1f3valiter.Attachment.Status
					}
					if f1f3valiter.Attachment.StatusMessage != nil {
						f1f3valf0.StatusMessage = f1f3valiter.Attachment.StatusMessage
					}
					if f1f3valiter.Attachment.SubnetId != nil {
						f1f3valf0.SubnetID = f1f3valiter.Attachment.SubnetId
					}
					f1f3val.Attachment = f1f3valf0
				}
				if f1f3valiter.Config != nil {
					f1f3valf1 := map[string]*svcapitypes.PerObjectStatus{}
					for f1f3valf1key, f1f3valf1valiter := range f1f3valiter.Config {
						f1f3valf1val := &svcapitypes.PerObjectStatus{}
						if f1f3valf1valiter.SyncStatus != nil {
							f1f3valf1val.SyncStatus = f1f3valf1valiter.SyncStatus
						}
						if f1f3valf1valiter.UpdateToken != nil {
							f1f3valf1val.UpdateToken = f1f3valf1valiter.UpdateToken
						}
						f1f3valf1[f1f3valf1key] = f1f3valf1val
					}
					f1f3val.Config = f1f3valf1
				}
				f1f3[f1f3key] = f1f3val
			}
			f1.SyncStates = f1f3
		}
		ko.Status.FirewallStatus = f1
	} else {
		ko.Status.FirewallStatus = nil
	}

	rm.setStatusDefaults(ko)

	if resp.Firewall != nil {
		if resp.Firewall.VpcId != nil {
			ko.Spec.VPCID = resp.Firewall.VpcId
		}

		if resp.Firewall.FirewallPolicyArn != nil {
			ko.Spec.FirewallPolicyARN = resp.Firewall.FirewallPolicyArn
		}

		if resp.Firewall.SubnetMappings != nil {
			subnetMappings := []*svcapitypes.SubnetMapping{}
			for _, subnetMapping := range resp.Firewall.SubnetMappings {
				subnetMap := &svcapitypes.SubnetMapping{}
				if subnetMapping.SubnetId != nil {
					subnetMap.SubnetID = subnetMapping.SubnetId
				}
				subnetMappings = append(subnetMappings, subnetMap)
			}
			ko.Spec.SubnetMappings = subnetMappings
		}
	}

	if err := rm.addLoggingConfigToSpec(ctx, r, ko); err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return false
}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeFirewallInput, error) {
	res := &svcsdk.DescribeFirewallInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetFirewallArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.FirewallName != nil {
		res.SetFirewallName(*r.ko.Spec.FirewallName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateFirewallOutput
	_ = resp
	resp, err = rm.sdkapi.CreateFirewallWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateFirewall", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.Firewall != nil {
		f0 := &svcapitypes.Firewall_SDK{}
		if resp.Firewall.DeleteProtection != nil {
			f0.DeleteProtection = resp.Firewall.DeleteProtection
		}
		if resp.Firewall.Description != nil {
			f0.Description = resp.Firewall.Description
		}
		if resp.Firewall.EncryptionConfiguration != nil {
			f0f2 := &svcapitypes.EncryptionConfiguration{}
			if resp.Firewall.EncryptionConfiguration.KeyId != nil {
				f0f2.KeyID = resp.Firewall.EncryptionConfiguration.KeyId
			}
			if resp.Firewall.EncryptionConfiguration.Type != nil {
				f0f2.Type = resp.Firewall.EncryptionConfiguration.Type
			}
			f0.EncryptionConfiguration = f0f2
		}
		if resp.Firewall.FirewallArn != nil {
			f0.FirewallARN = resp.Firewall.FirewallArn
		}
		if resp.Firewall.FirewallId != nil {
			f0.FirewallID = resp.Firewall.FirewallId
		}
		if resp.Firewall.FirewallName != nil {
			f0.FirewallName = resp.Firewall.FirewallName
		}
		if resp.Firewall.FirewallPolicyArn != nil {
			f0.FirewallPolicyARN = resp.Firewall.FirewallPolicyArn
		}
		if resp.Firewall.FirewallPolicyChangeProtection != nil {
			f0.FirewallPolicyChangeProtection = resp.Firewall.FirewallPolicyChangeProtection
		}
		if resp.Firewall.SubnetChangeProtection != nil {
			f0.SubnetChangeProtection = resp.Firewall.SubnetChangeProtection
		}
		if resp.Firewall.SubnetMappings != nil {
			f0f9 := []*svcapitypes.SubnetMapping{}
			for _, f0f9iter := range resp.Firewall.SubnetMappings {
				f0f9elem := &svcapitypes.SubnetMapping{}
				if f0f9iter.IPAddressType != nil {
					f0f9elem.IPAddressType = f0f9iter.IPAddressType
				}
				if f0f9iter.SubnetId != nil {
					f0f9elem.SubnetID = f0f9iter.SubnetId
				}
				f0f9 = append(f0f9, f0f9elem)
			}
			f0.SubnetMappings = f0f9
		}
		if resp.Firewall.Tags != nil {
			f0f10 := []*svcapitypes.Tag{}
			for _, f0f10iter := range resp.Firewall.Tags {
				f0f10elem := &svcapitypes.Tag{}
				if f0f10iter.Key != nil {
					f0f10elem.Key = f0f10iter.Key
				}
				if f0f10iter.Value != nil {
					f0f10elem.Value = f0f10iter.Value
				}
				f0f10 = append(f0f10, f0f10elem)
			}
			f0.Tags = f0f10
		}
		if resp.Firewall.VpcId != nil {
			f0.VPCID = resp.Firewall.VpcId
		}
		ko.Status.Firewall = f0
	} else {
		ko.Status.Firewall = nil
	}
	if resp.FirewallStatus != nil {
		f1 := &svcapitypes.FirewallStatus_SDK{}
		if resp.FirewallStatus.CapacityUsageSummary != nil {
			f1f0 := &svcapitypes.CapacityUsageSummary{}
			if resp.FirewallStatus.CapacityUsageSummary.CIDRs != nil {
				f1f0f0 := &svcapitypes.CIDRSummary{}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.AvailableCIDRCount != nil {
					f1f0f0.AvailableCIDRCount = resp.FirewallStatus.CapacityUsageSummary.CIDRs.AvailableCIDRCount
				}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.IPSetReferences != nil {
					f1f0f0f1 := map[string]*svcapitypes.IPSetMetadata{}
					for f1f0f0f1key, f1f0f0f1valiter := range resp.FirewallStatus.CapacityUsageSummary.CIDRs.IPSetReferences {
						f1f0f0f1val := &svcapitypes.IPSetMetadata{}
						if f1f0f0f1valiter.ResolvedCIDRCount != nil {
							f1f0f0f1val.ResolvedCIDRCount = f1f0f0f1valiter.ResolvedCIDRCount
						}
						f1f0f0f1[f1f0f0f1key] = f1f0f0f1val
					}
					f1f0f0.IPSetReferences = f1f0f0f1
				}
				if resp.FirewallStatus.CapacityUsageSummary.CIDRs.UtilizedCIDRCount != nil {
					f1f0f0.UtilizedCIDRCount = resp.FirewallStatus.CapacityUsageSummary.CIDRs.UtilizedCIDRCount
				}
				f1f0.CIDRs = f1f0f0
			}
			f1.CapacityUsageSummary = f1f0
		}
		if resp.FirewallStatus.ConfigurationSyncStateSummary != nil {
			f1.ConfigurationSyncStateSummary = resp.FirewallStatus.ConfigurationSyncStateSummary
		}
		if resp.FirewallStatus.Status != nil {
			f1.Status = resp.FirewallStatus.Status
		}
		if resp.FirewallStatus.SyncStates != nil {
			f1f3 := map[string]*svcapitypes.SyncState{}
			for f1f3key, f1f3valiter := range resp.FirewallStatus.SyncStates {
				f1f3val := &svcapitypes.SyncState{}
				if f1f3valiter.Attachment != nil {
					f1f3valf0 := &svcapitypes.Attachment{}
					if f1f3valiter.Attachment.EndpointId != nil {
						f1f3valf0.EndpointID = f1f3valiter.Attachment.EndpointId
					}
					if f1f3valiter.Attachment.Status != nil {
						f1f3valf0.Status = f1f3valiter.Attachment.Status
					}
					if f1f3valiter.Attachment.StatusMessage != nil {
						f1f3valf0.StatusMessage = f1f3valiter.Attachment.StatusMessage
					}
					if f1f3valiter.Attachment.SubnetId != nil {
						f1f3valf0.SubnetID = f1f3valiter.Attachment.SubnetId
					}
					f1f3val.Attachment = f1f3valf0
				}
				if f1f3valiter.Config != nil {
					f1f3valf1 := map[string]*svcapitypes.PerObjectStatus{}
					for f1f3valf1key, f1f3valf1valiter := range f1f3valiter.Config {
						f1f3valf1val := &svcapitypes.PerObjectStatus{}
						if f1f3valf1valiter.SyncStatus != nil {
							f1f3valf1val.SyncStatus = f1f3valf1valiter.SyncStatus
						}
						if f1f3valf1valiter.UpdateToken != nil {
							f1f3valf1val.UpdateToken = f1f3valf1valiter.UpdateToken
						}
						f1f3valf1[f1f3valf1key] = f1f3valf1val
					}
					f1f3val.Config = f1f3valf1
				}
				f1f3[f1f3key] = f1f3val
			}
			f1.SyncStates = f1f3
		}
		ko.Status.FirewallStatus = f1
	} else {
		ko.Status.FirewallStatus = nil
	}

	rm.setStatusDefaults(ko)
	if err := rm.createLoggingConfig(ctx, desired); err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateFirewallInput, error) {
	res := &svcsdk.CreateFirewallInput{}

	if r.ko.Spec.DeleteProtection != nil {
		res.SetDeleteProtection(*r.ko.Spec.DeleteProtection)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EncryptionConfiguration != nil {
		f2 := &svcsdk.EncryptionConfiguration{}
		if r.ko.Spec.EncryptionConfiguration.KeyID != nil {
			f2.SetKeyId(*r.ko.Spec.EncryptionConfiguration.KeyID)
		}
		if r.ko.Spec.EncryptionConfiguration.Type != nil {
			f2.SetType(*r.ko.Spec.EncryptionConfiguration.Type)
		}
		res.SetEncryptionConfiguration(f2)
	}
	if r.ko.Spec.FirewallName != nil {
		res.SetFirewallName(*r.ko.Spec.FirewallName)
	}
	if r.ko.Spec.FirewallPolicyARN != nil {
		res.SetFirewallPolicyArn(*r.ko.Spec.FirewallPolicyARN)
	}
	if r.ko.Spec.FirewallPolicyChangeProtection != nil {
		res.SetFirewallPolicyChangeProtection(*r.ko.Spec.FirewallPolicyChangeProtection)
	}
	if r.ko.Spec.SubnetChangeProtection != nil {
		res.SetSubnetChangeProtection(*r.ko.Spec.SubnetChangeProtection)
	}
	if r.ko.Spec.SubnetMappings != nil {
		f7 := []*svcsdk.SubnetMapping{}
		for _, f7iter := range r.ko.Spec.SubnetMappings {
			f7elem := &svcsdk.SubnetMapping{}
			if f7iter.IPAddressType != nil {
				f7elem.SetIPAddressType(*f7iter.IPAddressType)
			}
			if f7iter.SubnetID != nil {
				f7elem.SetSubnetId(*f7iter.SubnetID)
			}
			f7 = append(f7, f7elem)
		}
		res.SetSubnetMappings(f7)
	}
	if r.ko.Spec.Tags != nil {
		f8 := []*svcsdk.Tag{}
		for _, f8iter := range r.ko.Spec.Tags {
			f8elem := &svcsdk.Tag{}
			if f8iter.Key != nil {
				f8elem.SetKey(*f8iter.Key)
			}
			if f8iter.Value != nil {
				f8elem.SetValue(*f8iter.Value)
			}
			f8 = append(f8, f8elem)
		}
		res.SetTags(f8)
	}
	if r.ko.Spec.VPCID != nil {
		res.SetVpcId(*r.ko.Spec.VPCID)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateFirewall(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	if err := rm.deleteLoggingConfig(ctx, r); err != nil {
		return nil, err
	}
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteFirewallOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteFirewallWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFirewall", err)
	if err == nil {
		if observed, err := rm.sdkFind(ctx, r); err != ackerr.NotFound {
			if err != nil {
				return nil, err
			}
			r.SetStatus(observed)
			return r, requeueWaitWhileDeleting
		}
	}
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFirewallInput, error) {
	res := &svcsdk.DeleteFirewallInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetFirewallArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.FirewallName != nil {
		res.SetFirewallName(*r.ko.Spec.FirewallName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Firewall,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidRequestException":
		return true
	default:
		return false
	}
}

// newLoggingConfiguration returns a LoggingConfiguration object
// with each the field set by the resource's corresponding spec field.
func (rm *resourceManager) newLoggingConfiguration(
	r *resource,
) *svcsdk.LoggingConfiguration {
	res := &svcsdk.LoggingConfiguration{}

	if r.ko.Spec.LoggingConfiguration.LogDestinationConfigs != nil {
		resf0 := []*svcsdk.LogDestinationConfig{}
		for _, resf0iter := range r.ko.Spec.LoggingConfiguration.LogDestinationConfigs {
			resf0elem := &svcsdk.LogDestinationConfig{}
			if resf0iter.LogDestination != nil {
				resf0elemf0 := map[string]*string{}
				for resf0elemf0key, resf0elemf0valiter := range resf0iter.LogDestination {
					var resf0elemf0val string
					resf0elemf0val = *resf0elemf0valiter
					resf0elemf0[resf0elemf0key] = &resf0elemf0val
				}
				resf0elem.SetLogDestination(resf0elemf0)
			}
			if resf0iter.LogDestinationType != nil {
				resf0elem.SetLogDestinationType(*resf0iter.LogDestinationType)
			}
			if resf0iter.LogType != nil {
				resf0elem.SetLogType(*resf0iter.LogType)
			}
			resf0 = append(resf0, resf0elem)
		}
		res.SetLogDestinationConfigs(resf0)
	}

	return res
}

// setResourceLoggingConfiguration sets the `LoggingConfiguration` spec field
// given the output of a `DescribeLoggingConfiguration` operation.
func (rm *resourceManager) setResourceLoggingConfiguration(
	r *resource,
	resp *svcsdk.DescribeLoggingConfigurationOutput,
) *svcapitypes.LoggingConfiguration {
	res := &svcapitypes.LoggingConfiguration{}

	if resp.LoggingConfiguration.LogDestinationConfigs != nil {
		resf0 := []*svcapitypes.LogDestinationConfig{}
		for _, resf0iter := range resp.LoggingConfiguration.LogDestinationConfigs {
			resf0elem := &svcapitypes.LogDestinationConfig{}
			if resf0iter.LogDestination != nil {
				resf0elemf0 := map[string]*string{}
				for resf0elemf0key, resf0elemf0valiter := range resf0iter.LogDestination {
					var resf0elemf0val string
					resf0elemf0val = *resf0elemf0valiter
					resf0elemf0[resf0elemf0key] = &resf0elemf0val
				}
				resf0elem.LogDestination = resf0elemf0
			}
			if resf0iter.LogDestinationType != nil {
				resf0elem.LogDestinationType = resf0iter.LogDestinationType
			}
			if resf0iter.LogType != nil {
				resf0elem.LogType = resf0iter.LogType
			}
			resf0 = append(resf0, resf0elem)
		}
		res.LogDestinationConfigs = resf0
	}

	return res
}
