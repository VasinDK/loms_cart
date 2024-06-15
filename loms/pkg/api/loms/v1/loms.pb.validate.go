// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: loms.proto

package loms

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ItemRequestMultiError, or
// nil if none found.
func (m *ItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	// no validation rules for Count

	if len(errors) > 0 {
		return ItemRequestMultiError(errors)
	}

	return nil
}

// ItemRequestMultiError is an error wrapping multiple validation errors
// returned by ItemRequest.ValidateAll() if the designated constraints aren't met.
type ItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemRequestMultiError) AllErrors() []error { return m }

// ItemRequestValidationError is the validation error returned by
// ItemRequest.Validate if the designated constraints aren't met.
type ItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemRequestValidationError) ErrorName() string { return "ItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e ItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemRequestValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderCreateRequestMultiError, or nil if none found.
func (m *OrderCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrderCreateRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderCreateRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderCreateRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrderCreateRequestMultiError(errors)
	}

	return nil
}

// OrderCreateRequestMultiError is an error wrapping multiple validation errors
// returned by OrderCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderCreateRequestMultiError) AllErrors() []error { return m }

// OrderCreateRequestValidationError is the validation error returned by
// OrderCreateRequest.Validate if the designated constraints aren't met.
type OrderCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderCreateRequestValidationError) ErrorName() string {
	return "OrderCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrderCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderCreateRequestValidationError{}

// Validate checks the field values on OrderId with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderId with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OrderIdMultiError, or nil if none found.
func (m *OrderId) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return OrderIdMultiError(errors)
	}

	return nil
}

// OrderIdMultiError is an error wrapping multiple validation errors returned
// by OrderId.ValidateAll() if the designated constraints aren't met.
type OrderIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderIdMultiError) AllErrors() []error { return m }

// OrderIdValidationError is the validation error returned by OrderId.Validate
// if the designated constraints aren't met.
type OrderIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderIdValidationError) ErrorName() string { return "OrderIdValidationError" }

// Error satisfies the builtin error interface
func (e OrderIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderIdValidationError{}

// Validate checks the field values on OrderInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderInfoResponseMultiError, or nil if none found.
func (m *OrderInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for User

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrderInfoResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderInfoResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderInfoResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrderInfoResponseMultiError(errors)
	}

	return nil
}

// OrderInfoResponseMultiError is an error wrapping multiple validation errors
// returned by OrderInfoResponse.ValidateAll() if the designated constraints
// aren't met.
type OrderInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderInfoResponseMultiError) AllErrors() []error { return m }

// OrderInfoResponseValidationError is the validation error returned by
// OrderInfoResponse.Validate if the designated constraints aren't met.
type OrderInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderInfoResponseValidationError) ErrorName() string {
	return "OrderInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrderInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderInfoResponseValidationError{}

// Validate checks the field values on Sku with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Sku) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sku with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SkuMultiError, or nil if none found.
func (m *Sku) ValidateAll() error {
	return m.validate(true)
}

func (m *Sku) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	if len(errors) > 0 {
		return SkuMultiError(errors)
	}

	return nil
}

// SkuMultiError is an error wrapping multiple validation errors returned by
// Sku.ValidateAll() if the designated constraints aren't met.
type SkuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SkuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SkuMultiError) AllErrors() []error { return m }

// SkuValidationError is the validation error returned by Sku.Validate if the
// designated constraints aren't met.
type SkuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkuValidationError) ErrorName() string { return "SkuValidationError" }

// Error satisfies the builtin error interface
func (e SkuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSku.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkuValidationError{}

// Validate checks the field values on Count with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Count) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Count with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CountMultiError, or nil if none found.
func (m *Count) ValidateAll() error {
	return m.validate(true)
}

func (m *Count) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return CountMultiError(errors)
	}

	return nil
}

// CountMultiError is an error wrapping multiple validation errors returned by
// Count.ValidateAll() if the designated constraints aren't met.
type CountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountMultiError) AllErrors() []error { return m }

// CountValidationError is the validation error returned by Count.Validate if
// the designated constraints aren't met.
type CountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountValidationError) ErrorName() string { return "CountValidationError" }

// Error satisfies the builtin error interface
func (e CountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountValidationError{}