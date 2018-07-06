// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type CreateCertificateAuthorityService struct {
	CreateCertificateAuthorityStub        func(api.CertificateAuthorityInput) (api.CA, error)
	createCertificateAuthorityMutex       sync.RWMutex
	createCertificateAuthorityArgsForCall []struct {
		arg1 api.CertificateAuthorityInput
	}
	createCertificateAuthorityReturns struct {
		result1 api.CA
		result2 error
	}
	createCertificateAuthorityReturnsOnCall map[int]struct {
		result1 api.CA
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CreateCertificateAuthorityService) CreateCertificateAuthority(arg1 api.CertificateAuthorityInput) (api.CA, error) {
	fake.createCertificateAuthorityMutex.Lock()
	ret, specificReturn := fake.createCertificateAuthorityReturnsOnCall[len(fake.createCertificateAuthorityArgsForCall)]
	fake.createCertificateAuthorityArgsForCall = append(fake.createCertificateAuthorityArgsForCall, struct {
		arg1 api.CertificateAuthorityInput
	}{arg1})
	fake.recordInvocation("CreateCertificateAuthority", []interface{}{arg1})
	fake.createCertificateAuthorityMutex.Unlock()
	if fake.CreateCertificateAuthorityStub != nil {
		return fake.CreateCertificateAuthorityStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCertificateAuthorityReturns.result1, fake.createCertificateAuthorityReturns.result2
}

func (fake *CreateCertificateAuthorityService) CreateCertificateAuthorityCallCount() int {
	fake.createCertificateAuthorityMutex.RLock()
	defer fake.createCertificateAuthorityMutex.RUnlock()
	return len(fake.createCertificateAuthorityArgsForCall)
}

func (fake *CreateCertificateAuthorityService) CreateCertificateAuthorityArgsForCall(i int) api.CertificateAuthorityInput {
	fake.createCertificateAuthorityMutex.RLock()
	defer fake.createCertificateAuthorityMutex.RUnlock()
	return fake.createCertificateAuthorityArgsForCall[i].arg1
}

func (fake *CreateCertificateAuthorityService) CreateCertificateAuthorityReturns(result1 api.CA, result2 error) {
	fake.CreateCertificateAuthorityStub = nil
	fake.createCertificateAuthorityReturns = struct {
		result1 api.CA
		result2 error
	}{result1, result2}
}

func (fake *CreateCertificateAuthorityService) CreateCertificateAuthorityReturnsOnCall(i int, result1 api.CA, result2 error) {
	fake.CreateCertificateAuthorityStub = nil
	if fake.createCertificateAuthorityReturnsOnCall == nil {
		fake.createCertificateAuthorityReturnsOnCall = make(map[int]struct {
			result1 api.CA
			result2 error
		})
	}
	fake.createCertificateAuthorityReturnsOnCall[i] = struct {
		result1 api.CA
		result2 error
	}{result1, result2}
}

func (fake *CreateCertificateAuthorityService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCertificateAuthorityMutex.RLock()
	defer fake.createCertificateAuthorityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CreateCertificateAuthorityService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
