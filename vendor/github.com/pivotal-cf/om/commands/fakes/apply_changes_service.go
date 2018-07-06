// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ApplyChangesService struct {
	CreateInstallationStub        func(bool, bool) (api.InstallationsServiceOutput, error)
	createInstallationMutex       sync.RWMutex
	createInstallationArgsForCall []struct {
		arg1 bool
		arg2 bool
	}
	createInstallationReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	createInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	GetInstallationStub        func(id int) (api.InstallationsServiceOutput, error)
	getInstallationMutex       sync.RWMutex
	getInstallationArgsForCall []struct {
		id int
	}
	getInstallationReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	getInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	GetInstallationLogsStub        func(id int) (api.InstallationsServiceOutput, error)
	getInstallationLogsMutex       sync.RWMutex
	getInstallationLogsArgsForCall []struct {
		id int
	}
	getInstallationLogsReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	getInstallationLogsReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	RunningInstallationStub        func() (api.InstallationsServiceOutput, error)
	runningInstallationMutex       sync.RWMutex
	runningInstallationArgsForCall []struct{}
	runningInstallationReturns     struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	runningInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	ListInstallationsStub        func() ([]api.InstallationsServiceOutput, error)
	listInstallationsMutex       sync.RWMutex
	listInstallationsArgsForCall []struct{}
	listInstallationsReturns     struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}
	listInstallationsReturnsOnCall map[int]struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ApplyChangesService) CreateInstallation(arg1 bool, arg2 bool) (api.InstallationsServiceOutput, error) {
	fake.createInstallationMutex.Lock()
	ret, specificReturn := fake.createInstallationReturnsOnCall[len(fake.createInstallationArgsForCall)]
	fake.createInstallationArgsForCall = append(fake.createInstallationArgsForCall, struct {
		arg1 bool
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("CreateInstallation", []interface{}{arg1, arg2})
	fake.createInstallationMutex.Unlock()
	if fake.CreateInstallationStub != nil {
		return fake.CreateInstallationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createInstallationReturns.result1, fake.createInstallationReturns.result2
}

func (fake *ApplyChangesService) CreateInstallationCallCount() int {
	fake.createInstallationMutex.RLock()
	defer fake.createInstallationMutex.RUnlock()
	return len(fake.createInstallationArgsForCall)
}

func (fake *ApplyChangesService) CreateInstallationArgsForCall(i int) (bool, bool) {
	fake.createInstallationMutex.RLock()
	defer fake.createInstallationMutex.RUnlock()
	return fake.createInstallationArgsForCall[i].arg1, fake.createInstallationArgsForCall[i].arg2
}

func (fake *ApplyChangesService) CreateInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.CreateInstallationStub = nil
	fake.createInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) CreateInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.CreateInstallationStub = nil
	if fake.createInstallationReturnsOnCall == nil {
		fake.createInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.createInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) GetInstallation(id int) (api.InstallationsServiceOutput, error) {
	fake.getInstallationMutex.Lock()
	ret, specificReturn := fake.getInstallationReturnsOnCall[len(fake.getInstallationArgsForCall)]
	fake.getInstallationArgsForCall = append(fake.getInstallationArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("GetInstallation", []interface{}{id})
	fake.getInstallationMutex.Unlock()
	if fake.GetInstallationStub != nil {
		return fake.GetInstallationStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInstallationReturns.result1, fake.getInstallationReturns.result2
}

func (fake *ApplyChangesService) GetInstallationCallCount() int {
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	return len(fake.getInstallationArgsForCall)
}

func (fake *ApplyChangesService) GetInstallationArgsForCall(i int) int {
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	return fake.getInstallationArgsForCall[i].id
}

func (fake *ApplyChangesService) GetInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.GetInstallationStub = nil
	fake.getInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) GetInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.GetInstallationStub = nil
	if fake.getInstallationReturnsOnCall == nil {
		fake.getInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.getInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) GetInstallationLogs(id int) (api.InstallationsServiceOutput, error) {
	fake.getInstallationLogsMutex.Lock()
	ret, specificReturn := fake.getInstallationLogsReturnsOnCall[len(fake.getInstallationLogsArgsForCall)]
	fake.getInstallationLogsArgsForCall = append(fake.getInstallationLogsArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("GetInstallationLogs", []interface{}{id})
	fake.getInstallationLogsMutex.Unlock()
	if fake.GetInstallationLogsStub != nil {
		return fake.GetInstallationLogsStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInstallationLogsReturns.result1, fake.getInstallationLogsReturns.result2
}

func (fake *ApplyChangesService) GetInstallationLogsCallCount() int {
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	return len(fake.getInstallationLogsArgsForCall)
}

func (fake *ApplyChangesService) GetInstallationLogsArgsForCall(i int) int {
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	return fake.getInstallationLogsArgsForCall[i].id
}

func (fake *ApplyChangesService) GetInstallationLogsReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.GetInstallationLogsStub = nil
	fake.getInstallationLogsReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) GetInstallationLogsReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.GetInstallationLogsStub = nil
	if fake.getInstallationLogsReturnsOnCall == nil {
		fake.getInstallationLogsReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.getInstallationLogsReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) RunningInstallation() (api.InstallationsServiceOutput, error) {
	fake.runningInstallationMutex.Lock()
	ret, specificReturn := fake.runningInstallationReturnsOnCall[len(fake.runningInstallationArgsForCall)]
	fake.runningInstallationArgsForCall = append(fake.runningInstallationArgsForCall, struct{}{})
	fake.recordInvocation("RunningInstallation", []interface{}{})
	fake.runningInstallationMutex.Unlock()
	if fake.RunningInstallationStub != nil {
		return fake.RunningInstallationStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runningInstallationReturns.result1, fake.runningInstallationReturns.result2
}

func (fake *ApplyChangesService) RunningInstallationCallCount() int {
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	return len(fake.runningInstallationArgsForCall)
}

func (fake *ApplyChangesService) RunningInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.RunningInstallationStub = nil
	fake.runningInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) RunningInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.RunningInstallationStub = nil
	if fake.runningInstallationReturnsOnCall == nil {
		fake.runningInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.runningInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) ListInstallations() ([]api.InstallationsServiceOutput, error) {
	fake.listInstallationsMutex.Lock()
	ret, specificReturn := fake.listInstallationsReturnsOnCall[len(fake.listInstallationsArgsForCall)]
	fake.listInstallationsArgsForCall = append(fake.listInstallationsArgsForCall, struct{}{})
	fake.recordInvocation("ListInstallations", []interface{}{})
	fake.listInstallationsMutex.Unlock()
	if fake.ListInstallationsStub != nil {
		return fake.ListInstallationsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listInstallationsReturns.result1, fake.listInstallationsReturns.result2
}

func (fake *ApplyChangesService) ListInstallationsCallCount() int {
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	return len(fake.listInstallationsArgsForCall)
}

func (fake *ApplyChangesService) ListInstallationsReturns(result1 []api.InstallationsServiceOutput, result2 error) {
	fake.ListInstallationsStub = nil
	fake.listInstallationsReturns = struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) ListInstallationsReturnsOnCall(i int, result1 []api.InstallationsServiceOutput, result2 error) {
	fake.ListInstallationsStub = nil
	if fake.listInstallationsReturnsOnCall == nil {
		fake.listInstallationsReturnsOnCall = make(map[int]struct {
			result1 []api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.listInstallationsReturnsOnCall[i] = struct {
		result1 []api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *ApplyChangesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createInstallationMutex.RLock()
	defer fake.createInstallationMutex.RUnlock()
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ApplyChangesService) recordInvocation(key string, args []interface{}) {
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
