// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type RbacSnapshot struct {
	ServiceRoles        ServiceRolesByNamespace
	ServiceRoleBindings ServiceRoleBindingsByNamespace
	RbacConfigs         RbacConfigsByNamespace
}

func (s RbacSnapshot) Clone() RbacSnapshot {
	return RbacSnapshot{
		ServiceRoles:        s.ServiceRoles.Clone(),
		ServiceRoleBindings: s.ServiceRoleBindings.Clone(),
		RbacConfigs:         s.RbacConfigs.Clone(),
	}
}

func (s RbacSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashServiceRoles(),
		s.hashServiceRoleBindings(),
		s.hashRbacConfigs(),
	)
}

func (s RbacSnapshot) hashServiceRoles() uint64 {
	return hashutils.HashAll(s.ServiceRoles.List().AsInterfaces()...)
}

func (s RbacSnapshot) hashServiceRoleBindings() uint64 {
	return hashutils.HashAll(s.ServiceRoleBindings.List().AsInterfaces()...)
}

func (s RbacSnapshot) hashRbacConfigs() uint64 {
	return hashutils.HashAll(s.RbacConfigs.List().AsInterfaces()...)
}

func (s RbacSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("serviceRoles", s.hashServiceRoles()))
	fields = append(fields, zap.Uint64("serviceRoleBindings", s.hashServiceRoleBindings()))
	fields = append(fields, zap.Uint64("rbacConfigs", s.hashRbacConfigs()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}
