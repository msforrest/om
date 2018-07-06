// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StageProductService struct {
	StageStub        func(api.StageProductInput, string) error
	stageMutex       sync.RWMutex
	stageArgsForCall []struct {
		arg1 api.StageProductInput
		arg2 string
	}
	stageReturns struct {
		result1 error
	}
	stageReturnsOnCall map[int]struct {
		result1 error
	}
	ListDeployedProductsStub        func() ([]api.DeployedProductOutput, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct{}
	listDeployedProductsReturns     struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	CheckProductAvailabilityStub        func(productName string, productVersion string) (bool, error)
	checkProductAvailabilityMutex       sync.RWMutex
	checkProductAvailabilityArgsForCall []struct {
		productName    string
		productVersion string
	}
	checkProductAvailabilityReturns struct {
		result1 bool
		result2 error
	}
	checkProductAvailabilityReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetDiagnosticReportStub        func() (api.DiagnosticReport, error)
	getDiagnosticReportMutex       sync.RWMutex
	getDiagnosticReportArgsForCall []struct{}
	getDiagnosticReportReturns     struct {
		result1 api.DiagnosticReport
		result2 error
	}
	getDiagnosticReportReturnsOnCall map[int]struct {
		result1 api.DiagnosticReport
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StageProductService) Stage(arg1 api.StageProductInput, arg2 string) error {
	fake.stageMutex.Lock()
	ret, specificReturn := fake.stageReturnsOnCall[len(fake.stageArgsForCall)]
	fake.stageArgsForCall = append(fake.stageArgsForCall, struct {
		arg1 api.StageProductInput
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Stage", []interface{}{arg1, arg2})
	fake.stageMutex.Unlock()
	if fake.StageStub != nil {
		return fake.StageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stageReturns.result1
}

func (fake *StageProductService) StageCallCount() int {
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	return len(fake.stageArgsForCall)
}

func (fake *StageProductService) StageArgsForCall(i int) (api.StageProductInput, string) {
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	return fake.stageArgsForCall[i].arg1, fake.stageArgsForCall[i].arg2
}

func (fake *StageProductService) StageReturns(result1 error) {
	fake.StageStub = nil
	fake.stageReturns = struct {
		result1 error
	}{result1}
}

func (fake *StageProductService) StageReturnsOnCall(i int, result1 error) {
	fake.StageStub = nil
	if fake.stageReturnsOnCall == nil {
		fake.stageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *StageProductService) ListDeployedProducts() ([]api.DeployedProductOutput, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct{}{})
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if fake.ListDeployedProductsStub != nil {
		return fake.ListDeployedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductsReturns.result1, fake.listDeployedProductsReturns.result2
}

func (fake *StageProductService) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *StageProductService) ListDeployedProductsReturns(result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) ListDeployedProductsReturnsOnCall(i int, result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []api.DeployedProductOutput
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) CheckProductAvailability(productName string, productVersion string) (bool, error) {
	fake.checkProductAvailabilityMutex.Lock()
	ret, specificReturn := fake.checkProductAvailabilityReturnsOnCall[len(fake.checkProductAvailabilityArgsForCall)]
	fake.checkProductAvailabilityArgsForCall = append(fake.checkProductAvailabilityArgsForCall, struct {
		productName    string
		productVersion string
	}{productName, productVersion})
	fake.recordInvocation("CheckProductAvailability", []interface{}{productName, productVersion})
	fake.checkProductAvailabilityMutex.Unlock()
	if fake.CheckProductAvailabilityStub != nil {
		return fake.CheckProductAvailabilityStub(productName, productVersion)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkProductAvailabilityReturns.result1, fake.checkProductAvailabilityReturns.result2
}

func (fake *StageProductService) CheckProductAvailabilityCallCount() int {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return len(fake.checkProductAvailabilityArgsForCall)
}

func (fake *StageProductService) CheckProductAvailabilityArgsForCall(i int) (string, string) {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return fake.checkProductAvailabilityArgsForCall[i].productName, fake.checkProductAvailabilityArgsForCall[i].productVersion
}

func (fake *StageProductService) CheckProductAvailabilityReturns(result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	fake.checkProductAvailabilityReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) CheckProductAvailabilityReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	if fake.checkProductAvailabilityReturnsOnCall == nil {
		fake.checkProductAvailabilityReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkProductAvailabilityReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) GetDiagnosticReport() (api.DiagnosticReport, error) {
	fake.getDiagnosticReportMutex.Lock()
	ret, specificReturn := fake.getDiagnosticReportReturnsOnCall[len(fake.getDiagnosticReportArgsForCall)]
	fake.getDiagnosticReportArgsForCall = append(fake.getDiagnosticReportArgsForCall, struct{}{})
	fake.recordInvocation("GetDiagnosticReport", []interface{}{})
	fake.getDiagnosticReportMutex.Unlock()
	if fake.GetDiagnosticReportStub != nil {
		return fake.GetDiagnosticReportStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDiagnosticReportReturns.result1, fake.getDiagnosticReportReturns.result2
}

func (fake *StageProductService) GetDiagnosticReportCallCount() int {
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	return len(fake.getDiagnosticReportArgsForCall)
}

func (fake *StageProductService) GetDiagnosticReportReturns(result1 api.DiagnosticReport, result2 error) {
	fake.GetDiagnosticReportStub = nil
	fake.getDiagnosticReportReturns = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) GetDiagnosticReportReturnsOnCall(i int, result1 api.DiagnosticReport, result2 error) {
	fake.GetDiagnosticReportStub = nil
	if fake.getDiagnosticReportReturnsOnCall == nil {
		fake.getDiagnosticReportReturnsOnCall = make(map[int]struct {
			result1 api.DiagnosticReport
			result2 error
		})
	}
	fake.getDiagnosticReportReturnsOnCall[i] = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *StageProductService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stageMutex.RLock()
	defer fake.stageMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StageProductService) recordInvocation(key string, args []interface{}) {
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
