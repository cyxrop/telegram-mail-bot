package service

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i gitlab.ozon.dev/cyxrop/homework-2/internal/app/service/user.Cryptographer -o ./internal/app/service/user/cryptographer_mock_test.go -n CryptographerMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// CryptographerMock implements Cryptographer
type CryptographerMock struct {
	t minimock.Tester

	funcDecrypt          func(s1 string) (s2 string, err error)
	inspectFuncDecrypt   func(s1 string)
	afterDecryptCounter  uint64
	beforeDecryptCounter uint64
	DecryptMock          mCryptographerMockDecrypt

	funcEncrypt          func(s1 string) (s2 string, err error)
	inspectFuncEncrypt   func(s1 string)
	afterEncryptCounter  uint64
	beforeEncryptCounter uint64
	EncryptMock          mCryptographerMockEncrypt
}

// NewCryptographerMock returns a mock for Cryptographer
func NewCryptographerMock(t minimock.Tester) *CryptographerMock {
	m := &CryptographerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DecryptMock = mCryptographerMockDecrypt{mock: m}
	m.DecryptMock.callArgs = []*CryptographerMockDecryptParams{}

	m.EncryptMock = mCryptographerMockEncrypt{mock: m}
	m.EncryptMock.callArgs = []*CryptographerMockEncryptParams{}

	return m
}

type mCryptographerMockDecrypt struct {
	mock               *CryptographerMock
	defaultExpectation *CryptographerMockDecryptExpectation
	expectations       []*CryptographerMockDecryptExpectation

	callArgs []*CryptographerMockDecryptParams
	mutex    sync.RWMutex
}

// CryptographerMockDecryptExpectation specifies expectation struct of the Cryptographer.Decrypt
type CryptographerMockDecryptExpectation struct {
	mock    *CryptographerMock
	params  *CryptographerMockDecryptParams
	results *CryptographerMockDecryptResults
	Counter uint64
}

// CryptographerMockDecryptParams contains parameters of the Cryptographer.Decrypt
type CryptographerMockDecryptParams struct {
	s1 string
}

// CryptographerMockDecryptResults contains results of the Cryptographer.Decrypt
type CryptographerMockDecryptResults struct {
	s2  string
	err error
}

// Expect sets up expected params for Cryptographer.Decrypt
func (mmDecrypt *mCryptographerMockDecrypt) Expect(s1 string) *mCryptographerMockDecrypt {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("CryptographerMock.Decrypt mock is already set by Set")
	}

	if mmDecrypt.defaultExpectation == nil {
		mmDecrypt.defaultExpectation = &CryptographerMockDecryptExpectation{}
	}

	mmDecrypt.defaultExpectation.params = &CryptographerMockDecryptParams{s1}
	for _, e := range mmDecrypt.expectations {
		if minimock.Equal(e.params, mmDecrypt.defaultExpectation.params) {
			mmDecrypt.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDecrypt.defaultExpectation.params)
		}
	}

	return mmDecrypt
}

// Inspect accepts an inspector function that has same arguments as the Cryptographer.Decrypt
func (mmDecrypt *mCryptographerMockDecrypt) Inspect(f func(s1 string)) *mCryptographerMockDecrypt {
	if mmDecrypt.mock.inspectFuncDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("Inspect function is already set for CryptographerMock.Decrypt")
	}

	mmDecrypt.mock.inspectFuncDecrypt = f

	return mmDecrypt
}

// Return sets up results that will be returned by Cryptographer.Decrypt
func (mmDecrypt *mCryptographerMockDecrypt) Return(s2 string, err error) *CryptographerMock {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("CryptographerMock.Decrypt mock is already set by Set")
	}

	if mmDecrypt.defaultExpectation == nil {
		mmDecrypt.defaultExpectation = &CryptographerMockDecryptExpectation{mock: mmDecrypt.mock}
	}
	mmDecrypt.defaultExpectation.results = &CryptographerMockDecryptResults{s2, err}
	return mmDecrypt.mock
}

//Set uses given function f to mock the Cryptographer.Decrypt method
func (mmDecrypt *mCryptographerMockDecrypt) Set(f func(s1 string) (s2 string, err error)) *CryptographerMock {
	if mmDecrypt.defaultExpectation != nil {
		mmDecrypt.mock.t.Fatalf("Default expectation is already set for the Cryptographer.Decrypt method")
	}

	if len(mmDecrypt.expectations) > 0 {
		mmDecrypt.mock.t.Fatalf("Some expectations are already set for the Cryptographer.Decrypt method")
	}

	mmDecrypt.mock.funcDecrypt = f
	return mmDecrypt.mock
}

// When sets expectation for the Cryptographer.Decrypt which will trigger the result defined by the following
// Then helper
func (mmDecrypt *mCryptographerMockDecrypt) When(s1 string) *CryptographerMockDecryptExpectation {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("CryptographerMock.Decrypt mock is already set by Set")
	}

	expectation := &CryptographerMockDecryptExpectation{
		mock:   mmDecrypt.mock,
		params: &CryptographerMockDecryptParams{s1},
	}
	mmDecrypt.expectations = append(mmDecrypt.expectations, expectation)
	return expectation
}

// Then sets up Cryptographer.Decrypt return parameters for the expectation previously defined by the When method
func (e *CryptographerMockDecryptExpectation) Then(s2 string, err error) *CryptographerMock {
	e.results = &CryptographerMockDecryptResults{s2, err}
	return e.mock
}

// Decrypt implements Cryptographer
func (mmDecrypt *CryptographerMock) Decrypt(s1 string) (s2 string, err error) {
	mm_atomic.AddUint64(&mmDecrypt.beforeDecryptCounter, 1)
	defer mm_atomic.AddUint64(&mmDecrypt.afterDecryptCounter, 1)

	if mmDecrypt.inspectFuncDecrypt != nil {
		mmDecrypt.inspectFuncDecrypt(s1)
	}

	mm_params := &CryptographerMockDecryptParams{s1}

	// Record call args
	mmDecrypt.DecryptMock.mutex.Lock()
	mmDecrypt.DecryptMock.callArgs = append(mmDecrypt.DecryptMock.callArgs, mm_params)
	mmDecrypt.DecryptMock.mutex.Unlock()

	for _, e := range mmDecrypt.DecryptMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2, e.results.err
		}
	}

	if mmDecrypt.DecryptMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDecrypt.DecryptMock.defaultExpectation.Counter, 1)
		mm_want := mmDecrypt.DecryptMock.defaultExpectation.params
		mm_got := CryptographerMockDecryptParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDecrypt.t.Errorf("CryptographerMock.Decrypt got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDecrypt.DecryptMock.defaultExpectation.results
		if mm_results == nil {
			mmDecrypt.t.Fatal("No results are set for the CryptographerMock.Decrypt")
		}
		return (*mm_results).s2, (*mm_results).err
	}
	if mmDecrypt.funcDecrypt != nil {
		return mmDecrypt.funcDecrypt(s1)
	}
	mmDecrypt.t.Fatalf("Unexpected call to CryptographerMock.Decrypt. %v", s1)
	return
}

// DecryptAfterCounter returns a count of finished CryptographerMock.Decrypt invocations
func (mmDecrypt *CryptographerMock) DecryptAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecrypt.afterDecryptCounter)
}

// DecryptBeforeCounter returns a count of CryptographerMock.Decrypt invocations
func (mmDecrypt *CryptographerMock) DecryptBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecrypt.beforeDecryptCounter)
}

// Calls returns a list of arguments used in each call to CryptographerMock.Decrypt.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDecrypt *mCryptographerMockDecrypt) Calls() []*CryptographerMockDecryptParams {
	mmDecrypt.mutex.RLock()

	argCopy := make([]*CryptographerMockDecryptParams, len(mmDecrypt.callArgs))
	copy(argCopy, mmDecrypt.callArgs)

	mmDecrypt.mutex.RUnlock()

	return argCopy
}

// MinimockDecryptDone returns true if the count of the Decrypt invocations corresponds
// the number of defined expectations
func (m *CryptographerMock) MinimockDecryptDone() bool {
	for _, e := range m.DecryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecrypt != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		return false
	}
	return true
}

// MinimockDecryptInspect logs each unmet expectation
func (m *CryptographerMock) MinimockDecryptInspect() {
	for _, e := range m.DecryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CryptographerMock.Decrypt with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		if m.DecryptMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CryptographerMock.Decrypt")
		} else {
			m.t.Errorf("Expected call to CryptographerMock.Decrypt with params: %#v", *m.DecryptMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecrypt != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		m.t.Error("Expected call to CryptographerMock.Decrypt")
	}
}

type mCryptographerMockEncrypt struct {
	mock               *CryptographerMock
	defaultExpectation *CryptographerMockEncryptExpectation
	expectations       []*CryptographerMockEncryptExpectation

	callArgs []*CryptographerMockEncryptParams
	mutex    sync.RWMutex
}

// CryptographerMockEncryptExpectation specifies expectation struct of the Cryptographer.Encrypt
type CryptographerMockEncryptExpectation struct {
	mock    *CryptographerMock
	params  *CryptographerMockEncryptParams
	results *CryptographerMockEncryptResults
	Counter uint64
}

// CryptographerMockEncryptParams contains parameters of the Cryptographer.Encrypt
type CryptographerMockEncryptParams struct {
	s1 string
}

// CryptographerMockEncryptResults contains results of the Cryptographer.Encrypt
type CryptographerMockEncryptResults struct {
	s2  string
	err error
}

// Expect sets up expected params for Cryptographer.Encrypt
func (mmEncrypt *mCryptographerMockEncrypt) Expect(s1 string) *mCryptographerMockEncrypt {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("CryptographerMock.Encrypt mock is already set by Set")
	}

	if mmEncrypt.defaultExpectation == nil {
		mmEncrypt.defaultExpectation = &CryptographerMockEncryptExpectation{}
	}

	mmEncrypt.defaultExpectation.params = &CryptographerMockEncryptParams{s1}
	for _, e := range mmEncrypt.expectations {
		if minimock.Equal(e.params, mmEncrypt.defaultExpectation.params) {
			mmEncrypt.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmEncrypt.defaultExpectation.params)
		}
	}

	return mmEncrypt
}

// Inspect accepts an inspector function that has same arguments as the Cryptographer.Encrypt
func (mmEncrypt *mCryptographerMockEncrypt) Inspect(f func(s1 string)) *mCryptographerMockEncrypt {
	if mmEncrypt.mock.inspectFuncEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("Inspect function is already set for CryptographerMock.Encrypt")
	}

	mmEncrypt.mock.inspectFuncEncrypt = f

	return mmEncrypt
}

// Return sets up results that will be returned by Cryptographer.Encrypt
func (mmEncrypt *mCryptographerMockEncrypt) Return(s2 string, err error) *CryptographerMock {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("CryptographerMock.Encrypt mock is already set by Set")
	}

	if mmEncrypt.defaultExpectation == nil {
		mmEncrypt.defaultExpectation = &CryptographerMockEncryptExpectation{mock: mmEncrypt.mock}
	}
	mmEncrypt.defaultExpectation.results = &CryptographerMockEncryptResults{s2, err}
	return mmEncrypt.mock
}

//Set uses given function f to mock the Cryptographer.Encrypt method
func (mmEncrypt *mCryptographerMockEncrypt) Set(f func(s1 string) (s2 string, err error)) *CryptographerMock {
	if mmEncrypt.defaultExpectation != nil {
		mmEncrypt.mock.t.Fatalf("Default expectation is already set for the Cryptographer.Encrypt method")
	}

	if len(mmEncrypt.expectations) > 0 {
		mmEncrypt.mock.t.Fatalf("Some expectations are already set for the Cryptographer.Encrypt method")
	}

	mmEncrypt.mock.funcEncrypt = f
	return mmEncrypt.mock
}

// When sets expectation for the Cryptographer.Encrypt which will trigger the result defined by the following
// Then helper
func (mmEncrypt *mCryptographerMockEncrypt) When(s1 string) *CryptographerMockEncryptExpectation {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("CryptographerMock.Encrypt mock is already set by Set")
	}

	expectation := &CryptographerMockEncryptExpectation{
		mock:   mmEncrypt.mock,
		params: &CryptographerMockEncryptParams{s1},
	}
	mmEncrypt.expectations = append(mmEncrypt.expectations, expectation)
	return expectation
}

// Then sets up Cryptographer.Encrypt return parameters for the expectation previously defined by the When method
func (e *CryptographerMockEncryptExpectation) Then(s2 string, err error) *CryptographerMock {
	e.results = &CryptographerMockEncryptResults{s2, err}
	return e.mock
}

// Encrypt implements Cryptographer
func (mmEncrypt *CryptographerMock) Encrypt(s1 string) (s2 string, err error) {
	mm_atomic.AddUint64(&mmEncrypt.beforeEncryptCounter, 1)
	defer mm_atomic.AddUint64(&mmEncrypt.afterEncryptCounter, 1)

	if mmEncrypt.inspectFuncEncrypt != nil {
		mmEncrypt.inspectFuncEncrypt(s1)
	}

	mm_params := &CryptographerMockEncryptParams{s1}

	// Record call args
	mmEncrypt.EncryptMock.mutex.Lock()
	mmEncrypt.EncryptMock.callArgs = append(mmEncrypt.EncryptMock.callArgs, mm_params)
	mmEncrypt.EncryptMock.mutex.Unlock()

	for _, e := range mmEncrypt.EncryptMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2, e.results.err
		}
	}

	if mmEncrypt.EncryptMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEncrypt.EncryptMock.defaultExpectation.Counter, 1)
		mm_want := mmEncrypt.EncryptMock.defaultExpectation.params
		mm_got := CryptographerMockEncryptParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmEncrypt.t.Errorf("CryptographerMock.Encrypt got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmEncrypt.EncryptMock.defaultExpectation.results
		if mm_results == nil {
			mmEncrypt.t.Fatal("No results are set for the CryptographerMock.Encrypt")
		}
		return (*mm_results).s2, (*mm_results).err
	}
	if mmEncrypt.funcEncrypt != nil {
		return mmEncrypt.funcEncrypt(s1)
	}
	mmEncrypt.t.Fatalf("Unexpected call to CryptographerMock.Encrypt. %v", s1)
	return
}

// EncryptAfterCounter returns a count of finished CryptographerMock.Encrypt invocations
func (mmEncrypt *CryptographerMock) EncryptAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncrypt.afterEncryptCounter)
}

// EncryptBeforeCounter returns a count of CryptographerMock.Encrypt invocations
func (mmEncrypt *CryptographerMock) EncryptBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncrypt.beforeEncryptCounter)
}

// Calls returns a list of arguments used in each call to CryptographerMock.Encrypt.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmEncrypt *mCryptographerMockEncrypt) Calls() []*CryptographerMockEncryptParams {
	mmEncrypt.mutex.RLock()

	argCopy := make([]*CryptographerMockEncryptParams, len(mmEncrypt.callArgs))
	copy(argCopy, mmEncrypt.callArgs)

	mmEncrypt.mutex.RUnlock()

	return argCopy
}

// MinimockEncryptDone returns true if the count of the Encrypt invocations corresponds
// the number of defined expectations
func (m *CryptographerMock) MinimockEncryptDone() bool {
	for _, e := range m.EncryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncrypt != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		return false
	}
	return true
}

// MinimockEncryptInspect logs each unmet expectation
func (m *CryptographerMock) MinimockEncryptInspect() {
	for _, e := range m.EncryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CryptographerMock.Encrypt with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		if m.EncryptMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CryptographerMock.Encrypt")
		} else {
			m.t.Errorf("Expected call to CryptographerMock.Encrypt with params: %#v", *m.EncryptMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncrypt != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		m.t.Error("Expected call to CryptographerMock.Encrypt")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CryptographerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDecryptInspect()

		m.MinimockEncryptInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CryptographerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CryptographerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDecryptDone() &&
		m.MinimockEncryptDone()
}
