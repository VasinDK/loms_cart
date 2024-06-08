// Code generated by http://github.com/gojuno/minimock (v3.3.11). DO NOT EDIT.

package mock

//go:generate minimock -i route256/cart/internal/service/list/get_cart.Repository -o repository_mock.go -n RepositoryMock -p mock

import (
	"route256/cart/internal/model"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements get_cart.Repository
type RepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCheckSKU          func(i1 int64) (pp1 *model.Product, err error)
	inspectFuncCheckSKU   func(i1 int64)
	afterCheckSKUCounter  uint64
	beforeCheckSKUCounter uint64
	CheckSKUMock          mRepositoryMockCheckSKU

	funcGetCart          func(i1 int64) (m1 map[int64]*model.Product, err error)
	inspectFuncGetCart   func(i1 int64)
	afterGetCartCounter  uint64
	beforeGetCartCounter uint64
	GetCartMock          mRepositoryMockGetCart
}

// NewRepositoryMock returns a mock for get_cart.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckSKUMock = mRepositoryMockCheckSKU{mock: m}
	m.CheckSKUMock.callArgs = []*RepositoryMockCheckSKUParams{}

	m.GetCartMock = mRepositoryMockGetCart{mock: m}
	m.GetCartMock.callArgs = []*RepositoryMockGetCartParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRepositoryMockCheckSKU struct {
	optional           bool
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockCheckSKUExpectation
	expectations       []*RepositoryMockCheckSKUExpectation

	callArgs []*RepositoryMockCheckSKUParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// RepositoryMockCheckSKUExpectation specifies expectation struct of the Repository.CheckSKU
type RepositoryMockCheckSKUExpectation struct {
	mock      *RepositoryMock
	params    *RepositoryMockCheckSKUParams
	paramPtrs *RepositoryMockCheckSKUParamPtrs
	results   *RepositoryMockCheckSKUResults
	Counter   uint64
}

// RepositoryMockCheckSKUParams contains parameters of the Repository.CheckSKU
type RepositoryMockCheckSKUParams struct {
	i1 int64
}

// RepositoryMockCheckSKUParamPtrs contains pointers to parameters of the Repository.CheckSKU
type RepositoryMockCheckSKUParamPtrs struct {
	i1 *int64
}

// RepositoryMockCheckSKUResults contains results of the Repository.CheckSKU
type RepositoryMockCheckSKUResults struct {
	pp1 *model.Product
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCheckSKU *mRepositoryMockCheckSKU) Optional() *mRepositoryMockCheckSKU {
	mmCheckSKU.optional = true
	return mmCheckSKU
}

// Expect sets up expected params for Repository.CheckSKU
func (mmCheckSKU *mRepositoryMockCheckSKU) Expect(i1 int64) *mRepositoryMockCheckSKU {
	if mmCheckSKU.mock.funcCheckSKU != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by Set")
	}

	if mmCheckSKU.defaultExpectation == nil {
		mmCheckSKU.defaultExpectation = &RepositoryMockCheckSKUExpectation{}
	}

	if mmCheckSKU.defaultExpectation.paramPtrs != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by ExpectParams functions")
	}

	mmCheckSKU.defaultExpectation.params = &RepositoryMockCheckSKUParams{i1}
	for _, e := range mmCheckSKU.expectations {
		if minimock.Equal(e.params, mmCheckSKU.defaultExpectation.params) {
			mmCheckSKU.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheckSKU.defaultExpectation.params)
		}
	}

	return mmCheckSKU
}

// ExpectI1Param1 sets up expected param i1 for Repository.CheckSKU
func (mmCheckSKU *mRepositoryMockCheckSKU) ExpectI1Param1(i1 int64) *mRepositoryMockCheckSKU {
	if mmCheckSKU.mock.funcCheckSKU != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by Set")
	}

	if mmCheckSKU.defaultExpectation == nil {
		mmCheckSKU.defaultExpectation = &RepositoryMockCheckSKUExpectation{}
	}

	if mmCheckSKU.defaultExpectation.params != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by Expect")
	}

	if mmCheckSKU.defaultExpectation.paramPtrs == nil {
		mmCheckSKU.defaultExpectation.paramPtrs = &RepositoryMockCheckSKUParamPtrs{}
	}
	mmCheckSKU.defaultExpectation.paramPtrs.i1 = &i1

	return mmCheckSKU
}

// Inspect accepts an inspector function that has same arguments as the Repository.CheckSKU
func (mmCheckSKU *mRepositoryMockCheckSKU) Inspect(f func(i1 int64)) *mRepositoryMockCheckSKU {
	if mmCheckSKU.mock.inspectFuncCheckSKU != nil {
		mmCheckSKU.mock.t.Fatalf("Inspect function is already set for RepositoryMock.CheckSKU")
	}

	mmCheckSKU.mock.inspectFuncCheckSKU = f

	return mmCheckSKU
}

// Return sets up results that will be returned by Repository.CheckSKU
func (mmCheckSKU *mRepositoryMockCheckSKU) Return(pp1 *model.Product, err error) *RepositoryMock {
	if mmCheckSKU.mock.funcCheckSKU != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by Set")
	}

	if mmCheckSKU.defaultExpectation == nil {
		mmCheckSKU.defaultExpectation = &RepositoryMockCheckSKUExpectation{mock: mmCheckSKU.mock}
	}
	mmCheckSKU.defaultExpectation.results = &RepositoryMockCheckSKUResults{pp1, err}
	return mmCheckSKU.mock
}

// Set uses given function f to mock the Repository.CheckSKU method
func (mmCheckSKU *mRepositoryMockCheckSKU) Set(f func(i1 int64) (pp1 *model.Product, err error)) *RepositoryMock {
	if mmCheckSKU.defaultExpectation != nil {
		mmCheckSKU.mock.t.Fatalf("Default expectation is already set for the Repository.CheckSKU method")
	}

	if len(mmCheckSKU.expectations) > 0 {
		mmCheckSKU.mock.t.Fatalf("Some expectations are already set for the Repository.CheckSKU method")
	}

	mmCheckSKU.mock.funcCheckSKU = f
	return mmCheckSKU.mock
}

// When sets expectation for the Repository.CheckSKU which will trigger the result defined by the following
// Then helper
func (mmCheckSKU *mRepositoryMockCheckSKU) When(i1 int64) *RepositoryMockCheckSKUExpectation {
	if mmCheckSKU.mock.funcCheckSKU != nil {
		mmCheckSKU.mock.t.Fatalf("RepositoryMock.CheckSKU mock is already set by Set")
	}

	expectation := &RepositoryMockCheckSKUExpectation{
		mock:   mmCheckSKU.mock,
		params: &RepositoryMockCheckSKUParams{i1},
	}
	mmCheckSKU.expectations = append(mmCheckSKU.expectations, expectation)
	return expectation
}

// Then sets up Repository.CheckSKU return parameters for the expectation previously defined by the When method
func (e *RepositoryMockCheckSKUExpectation) Then(pp1 *model.Product, err error) *RepositoryMock {
	e.results = &RepositoryMockCheckSKUResults{pp1, err}
	return e.mock
}

// Times sets number of times Repository.CheckSKU should be invoked
func (mmCheckSKU *mRepositoryMockCheckSKU) Times(n uint64) *mRepositoryMockCheckSKU {
	if n == 0 {
		mmCheckSKU.mock.t.Fatalf("Times of RepositoryMock.CheckSKU mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCheckSKU.expectedInvocations, n)
	return mmCheckSKU
}

func (mmCheckSKU *mRepositoryMockCheckSKU) invocationsDone() bool {
	if len(mmCheckSKU.expectations) == 0 && mmCheckSKU.defaultExpectation == nil && mmCheckSKU.mock.funcCheckSKU == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCheckSKU.mock.afterCheckSKUCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCheckSKU.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CheckSKU implements get_cart.Repository
func (mmCheckSKU *RepositoryMock) CheckSKU(i1 int64) (pp1 *model.Product, err error) {
	mm_atomic.AddUint64(&mmCheckSKU.beforeCheckSKUCounter, 1)
	defer mm_atomic.AddUint64(&mmCheckSKU.afterCheckSKUCounter, 1)

	if mmCheckSKU.inspectFuncCheckSKU != nil {
		mmCheckSKU.inspectFuncCheckSKU(i1)
	}

	mm_params := RepositoryMockCheckSKUParams{i1}

	// Record call args
	mmCheckSKU.CheckSKUMock.mutex.Lock()
	mmCheckSKU.CheckSKUMock.callArgs = append(mmCheckSKU.CheckSKUMock.callArgs, &mm_params)
	mmCheckSKU.CheckSKUMock.mutex.Unlock()

	for _, e := range mmCheckSKU.CheckSKUMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp1, e.results.err
		}
	}

	if mmCheckSKU.CheckSKUMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheckSKU.CheckSKUMock.defaultExpectation.Counter, 1)
		mm_want := mmCheckSKU.CheckSKUMock.defaultExpectation.params
		mm_want_ptrs := mmCheckSKU.CheckSKUMock.defaultExpectation.paramPtrs

		mm_got := RepositoryMockCheckSKUParams{i1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.i1 != nil && !minimock.Equal(*mm_want_ptrs.i1, mm_got.i1) {
				mmCheckSKU.t.Errorf("RepositoryMock.CheckSKU got unexpected parameter i1, want: %#v, got: %#v%s\n", *mm_want_ptrs.i1, mm_got.i1, minimock.Diff(*mm_want_ptrs.i1, mm_got.i1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheckSKU.t.Errorf("RepositoryMock.CheckSKU got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheckSKU.CheckSKUMock.defaultExpectation.results
		if mm_results == nil {
			mmCheckSKU.t.Fatal("No results are set for the RepositoryMock.CheckSKU")
		}
		return (*mm_results).pp1, (*mm_results).err
	}
	if mmCheckSKU.funcCheckSKU != nil {
		return mmCheckSKU.funcCheckSKU(i1)
	}
	mmCheckSKU.t.Fatalf("Unexpected call to RepositoryMock.CheckSKU. %v", i1)
	return
}

// CheckSKUAfterCounter returns a count of finished RepositoryMock.CheckSKU invocations
func (mmCheckSKU *RepositoryMock) CheckSKUAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckSKU.afterCheckSKUCounter)
}

// CheckSKUBeforeCounter returns a count of RepositoryMock.CheckSKU invocations
func (mmCheckSKU *RepositoryMock) CheckSKUBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckSKU.beforeCheckSKUCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.CheckSKU.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheckSKU *mRepositoryMockCheckSKU) Calls() []*RepositoryMockCheckSKUParams {
	mmCheckSKU.mutex.RLock()

	argCopy := make([]*RepositoryMockCheckSKUParams, len(mmCheckSKU.callArgs))
	copy(argCopy, mmCheckSKU.callArgs)

	mmCheckSKU.mutex.RUnlock()

	return argCopy
}

// MinimockCheckSKUDone returns true if the count of the CheckSKU invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockCheckSKUDone() bool {
	if m.CheckSKUMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CheckSKUMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CheckSKUMock.invocationsDone()
}

// MinimockCheckSKUInspect logs each unmet expectation
func (m *RepositoryMock) MinimockCheckSKUInspect() {
	for _, e := range m.CheckSKUMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.CheckSKU with params: %#v", *e.params)
		}
	}

	afterCheckSKUCounter := mm_atomic.LoadUint64(&m.afterCheckSKUCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CheckSKUMock.defaultExpectation != nil && afterCheckSKUCounter < 1 {
		if m.CheckSKUMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.CheckSKU")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.CheckSKU with params: %#v", *m.CheckSKUMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckSKU != nil && afterCheckSKUCounter < 1 {
		m.t.Error("Expected call to RepositoryMock.CheckSKU")
	}

	if !m.CheckSKUMock.invocationsDone() && afterCheckSKUCounter > 0 {
		m.t.Errorf("Expected %d calls to RepositoryMock.CheckSKU but found %d calls",
			mm_atomic.LoadUint64(&m.CheckSKUMock.expectedInvocations), afterCheckSKUCounter)
	}
}

type mRepositoryMockGetCart struct {
	optional           bool
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetCartExpectation
	expectations       []*RepositoryMockGetCartExpectation

	callArgs []*RepositoryMockGetCartParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// RepositoryMockGetCartExpectation specifies expectation struct of the Repository.GetCart
type RepositoryMockGetCartExpectation struct {
	mock      *RepositoryMock
	params    *RepositoryMockGetCartParams
	paramPtrs *RepositoryMockGetCartParamPtrs
	results   *RepositoryMockGetCartResults
	Counter   uint64
}

// RepositoryMockGetCartParams contains parameters of the Repository.GetCart
type RepositoryMockGetCartParams struct {
	i1 int64
}

// RepositoryMockGetCartParamPtrs contains pointers to parameters of the Repository.GetCart
type RepositoryMockGetCartParamPtrs struct {
	i1 *int64
}

// RepositoryMockGetCartResults contains results of the Repository.GetCart
type RepositoryMockGetCartResults struct {
	m1  map[int64]*model.Product
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option by default unless you really need it, as it helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetCart *mRepositoryMockGetCart) Optional() *mRepositoryMockGetCart {
	mmGetCart.optional = true
	return mmGetCart
}

// Expect sets up expected params for Repository.GetCart
func (mmGetCart *mRepositoryMockGetCart) Expect(i1 int64) *mRepositoryMockGetCart {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by Set")
	}

	if mmGetCart.defaultExpectation == nil {
		mmGetCart.defaultExpectation = &RepositoryMockGetCartExpectation{}
	}

	if mmGetCart.defaultExpectation.paramPtrs != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by ExpectParams functions")
	}

	mmGetCart.defaultExpectation.params = &RepositoryMockGetCartParams{i1}
	for _, e := range mmGetCart.expectations {
		if minimock.Equal(e.params, mmGetCart.defaultExpectation.params) {
			mmGetCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCart.defaultExpectation.params)
		}
	}

	return mmGetCart
}

// ExpectI1Param1 sets up expected param i1 for Repository.GetCart
func (mmGetCart *mRepositoryMockGetCart) ExpectI1Param1(i1 int64) *mRepositoryMockGetCart {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by Set")
	}

	if mmGetCart.defaultExpectation == nil {
		mmGetCart.defaultExpectation = &RepositoryMockGetCartExpectation{}
	}

	if mmGetCart.defaultExpectation.params != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by Expect")
	}

	if mmGetCart.defaultExpectation.paramPtrs == nil {
		mmGetCart.defaultExpectation.paramPtrs = &RepositoryMockGetCartParamPtrs{}
	}
	mmGetCart.defaultExpectation.paramPtrs.i1 = &i1

	return mmGetCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetCart
func (mmGetCart *mRepositoryMockGetCart) Inspect(f func(i1 int64)) *mRepositoryMockGetCart {
	if mmGetCart.mock.inspectFuncGetCart != nil {
		mmGetCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetCart")
	}

	mmGetCart.mock.inspectFuncGetCart = f

	return mmGetCart
}

// Return sets up results that will be returned by Repository.GetCart
func (mmGetCart *mRepositoryMockGetCart) Return(m1 map[int64]*model.Product, err error) *RepositoryMock {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by Set")
	}

	if mmGetCart.defaultExpectation == nil {
		mmGetCart.defaultExpectation = &RepositoryMockGetCartExpectation{mock: mmGetCart.mock}
	}
	mmGetCart.defaultExpectation.results = &RepositoryMockGetCartResults{m1, err}
	return mmGetCart.mock
}

// Set uses given function f to mock the Repository.GetCart method
func (mmGetCart *mRepositoryMockGetCart) Set(f func(i1 int64) (m1 map[int64]*model.Product, err error)) *RepositoryMock {
	if mmGetCart.defaultExpectation != nil {
		mmGetCart.mock.t.Fatalf("Default expectation is already set for the Repository.GetCart method")
	}

	if len(mmGetCart.expectations) > 0 {
		mmGetCart.mock.t.Fatalf("Some expectations are already set for the Repository.GetCart method")
	}

	mmGetCart.mock.funcGetCart = f
	return mmGetCart.mock
}

// When sets expectation for the Repository.GetCart which will trigger the result defined by the following
// Then helper
func (mmGetCart *mRepositoryMockGetCart) When(i1 int64) *RepositoryMockGetCartExpectation {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("RepositoryMock.GetCart mock is already set by Set")
	}

	expectation := &RepositoryMockGetCartExpectation{
		mock:   mmGetCart.mock,
		params: &RepositoryMockGetCartParams{i1},
	}
	mmGetCart.expectations = append(mmGetCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetCartExpectation) Then(m1 map[int64]*model.Product, err error) *RepositoryMock {
	e.results = &RepositoryMockGetCartResults{m1, err}
	return e.mock
}

// Times sets number of times Repository.GetCart should be invoked
func (mmGetCart *mRepositoryMockGetCart) Times(n uint64) *mRepositoryMockGetCart {
	if n == 0 {
		mmGetCart.mock.t.Fatalf("Times of RepositoryMock.GetCart mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetCart.expectedInvocations, n)
	return mmGetCart
}

func (mmGetCart *mRepositoryMockGetCart) invocationsDone() bool {
	if len(mmGetCart.expectations) == 0 && mmGetCart.defaultExpectation == nil && mmGetCart.mock.funcGetCart == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetCart.mock.afterGetCartCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetCart.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetCart implements get_cart.Repository
func (mmGetCart *RepositoryMock) GetCart(i1 int64) (m1 map[int64]*model.Product, err error) {
	mm_atomic.AddUint64(&mmGetCart.beforeGetCartCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCart.afterGetCartCounter, 1)

	if mmGetCart.inspectFuncGetCart != nil {
		mmGetCart.inspectFuncGetCart(i1)
	}

	mm_params := RepositoryMockGetCartParams{i1}

	// Record call args
	mmGetCart.GetCartMock.mutex.Lock()
	mmGetCart.GetCartMock.callArgs = append(mmGetCart.GetCartMock.callArgs, &mm_params)
	mmGetCart.GetCartMock.mutex.Unlock()

	for _, e := range mmGetCart.GetCartMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmGetCart.GetCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCart.GetCartMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCart.GetCartMock.defaultExpectation.params
		mm_want_ptrs := mmGetCart.GetCartMock.defaultExpectation.paramPtrs

		mm_got := RepositoryMockGetCartParams{i1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.i1 != nil && !minimock.Equal(*mm_want_ptrs.i1, mm_got.i1) {
				mmGetCart.t.Errorf("RepositoryMock.GetCart got unexpected parameter i1, want: %#v, got: %#v%s\n", *mm_want_ptrs.i1, mm_got.i1, minimock.Diff(*mm_want_ptrs.i1, mm_got.i1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCart.t.Errorf("RepositoryMock.GetCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCart.GetCartMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCart.t.Fatal("No results are set for the RepositoryMock.GetCart")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmGetCart.funcGetCart != nil {
		return mmGetCart.funcGetCart(i1)
	}
	mmGetCart.t.Fatalf("Unexpected call to RepositoryMock.GetCart. %v", i1)
	return
}

// GetCartAfterCounter returns a count of finished RepositoryMock.GetCart invocations
func (mmGetCart *RepositoryMock) GetCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCart.afterGetCartCounter)
}

// GetCartBeforeCounter returns a count of RepositoryMock.GetCart invocations
func (mmGetCart *RepositoryMock) GetCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCart.beforeGetCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCart *mRepositoryMockGetCart) Calls() []*RepositoryMockGetCartParams {
	mmGetCart.mutex.RLock()

	argCopy := make([]*RepositoryMockGetCartParams, len(mmGetCart.callArgs))
	copy(argCopy, mmGetCart.callArgs)

	mmGetCart.mutex.RUnlock()

	return argCopy
}

// MinimockGetCartDone returns true if the count of the GetCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetCartDone() bool {
	if m.GetCartMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetCartMock.invocationsDone()
}

// MinimockGetCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetCartInspect() {
	for _, e := range m.GetCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetCart with params: %#v", *e.params)
		}
	}

	afterGetCartCounter := mm_atomic.LoadUint64(&m.afterGetCartCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetCartMock.defaultExpectation != nil && afterGetCartCounter < 1 {
		if m.GetCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetCart with params: %#v", *m.GetCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCart != nil && afterGetCartCounter < 1 {
		m.t.Error("Expected call to RepositoryMock.GetCart")
	}

	if !m.GetCartMock.invocationsDone() && afterGetCartCounter > 0 {
		m.t.Errorf("Expected %d calls to RepositoryMock.GetCart but found %d calls",
			mm_atomic.LoadUint64(&m.GetCartMock.expectedInvocations), afterGetCartCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCheckSKUInspect()

			m.MinimockGetCartInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCheckSKUDone() &&
		m.MinimockGetCartDone()
}