// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

type ProgressToken struct {
	Name   string
	Number float64
}

type PartialResultParams struct {
	// PartialResultToken an optional token that a server can use to report partial results (e.g. streaming) to the client.
	PartialResultToken ProgressToken `json:"partialResultToken,omitempty"`
}

// WorkDoneProgressClientCapabilities
type WorkDoneProgressClientCapabilities struct {
	// Window specific client capabilities.
	Window *Window `json:"window,omitempty"`
}

// WindowClientCapabilities
type WindowClientCapabilities struct {
	// Window specific client capabilities.
	Window interface{} `json:"window,omitempty"`

	// WorkDoneProgress whether client supports handling progress notifications. If set servers are allowed to
	// report in `workDoneProgress` property in the request specific server capabilities.
	//
	// Since 3.15.0.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// Window Window specific client capabilities.
type Window struct {
	// WorkDoneProgress whether client supports handling progress notifications. If set servers are allowed to
	// report in `workDoneProgress` property in the request specific server capabilities.
	//
	// Since 3.15.0.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// WorkDoneProgressBegin
type WorkDoneProgressBegin struct {
	Kind string `json:"kind"`

	// Title mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	//
	// Examples: "Indexing" or "Linking dependencies".
	Title string `json:"title"`

	// Cancellable controls if a cancel button should show to allow the user to cancel the
	// long running operation. Clients that don't support cancellation are allowed
	// to ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message is optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage is optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage float64 `json:"percentage,omitempty"`
}

// WorkDoneProgressCancelNotificationHandlerSignature
type WorkDoneProgressCancelNotificationHandlerSignature struct{}

// WorkDoneProgressCreateRequestHandlerSignature
type WorkDoneProgressCreateRequestHandlerSignature struct{}

// WorkDoneProgressReport
type WorkDoneProgressReport struct {
	Kind string `json:"kind"`

	// Cancellable controls enablement state of a cancel button.
	//
	// Clients that don't support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable bool `json:"cancellable,omitempty"`

	// Message is optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Percentage is optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage float64 `json:"percentage,omitempty"`
}

// WorkDoneProgressEnd
type WorkDoneProgressEnd struct {
	Kind string `json:"kind"`

	// Message is optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message string `json:"message,omitempty"`
}

type WorkDoneProgressParams struct {
	// WorkDoneToken an optional token that a server can use to report work done progress.
	WorkDoneToken ProgressToken `json:"workDoneToken,omitempty"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// WorkDoneProgressCreateParams
type WorkDoneProgressCreateParams struct {
	// The token to be used to report progress.
	Token interface{} `json:"token"`
}

// WorkDoneProgressCancelParams
type WorkDoneProgressCancelParams struct {
	// The token to be used to report progress.
	Token interface{} `json:"token"`
}

// // WindowClientCapabilities represents a WindowClientCapabilities specific client capabilities.
// type WindowClientCapabilities struct {
// 	// Window specific client capabilities.
// 	Window interface{} `json:"window,omitempty"`
//
// 	// WorkDoneProgress whether client supports handling progress notifications.
// 	// If set servers are allowed to report in `workDoneProgress` property in the request specific server capabilities.
// 	//
// 	// Since 3.15.0.
// 	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
// }
//
// type WorkDoneProgressBegin struct {
// 	Kind string `json:"kind"`
//
// 	// Title mandatory title of the progress operation. Used to briefly inform about
// 	// the kind of operation being performed.
// 	//
// 	// Examples: "Indexing" or "Linking dependencies".
// 	Title string `json:"title"`
//
// 	// Cancellable controls if a cancel button should show to allow the user to cancel the
// 	// long running operation. Clients that don't support cancellation are allowed
// 	// to ignore the setting.
// 	Cancellable bool `json:"cancellable,omitempty"`
//
// 	// Message is optional, more detailed associated progress message. Contains
// 	// complementary information to the `title`.
// 	//
// 	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
// 	// If unset, the previous progress message (if any) is still valid.
// 	Message string `json:"message,omitempty"`
//
// 	// Percentage is optional progress percentage to display (value 100 is considered 100%).
// 	// If not provided infinite progress is assumed and clients are allowed
// 	// to ignore the `percentage` value in subsequent in report notifications.
// 	//
// 	// The value should be steadily rising. Clients are free to ignore values
// 	// that are not following this rule.
// 	Percentage float64 `json:"percentage,omitempty"`
// }
//
// type WorkDoneProgressReport struct {
// 	Kind string `json:"kind"`
//
// 	// Cancellable controls enablement state of a cancel button.
// 	//
// 	// Clients that don't support cancellation or don't support controlling the button's enablement state are allowed to ignore the property.
// 	Cancellable bool `json:"cancellable,omitempty"`
//
// 	// Message is optional, more detailed associated progress message. Contains
// 	// complementary information to the `title`.
// 	//
// 	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
// 	// If unset, the previous progress message (if any) is still valid.
// 	Message string `json:"message,omitempty"`
//
// 	// Percentage is optional progress percentage to display (value 100 is considered 100%).
// 	// If not provided infinite progress is assumed and clients are allowed
// 	// to ignore the `percentage` value in subsequent in report notifications.
// 	//
// 	// The value should be steadily rising. Clients are free to ignore values
// 	// that are not following this rule.
// 	Percentage float64 `json:"percentage,omitempty"`
// }
//
// type WorkDoneProgressParams struct {
// 	// WorkDoneToken an optional token that a server can use to report work done progress.
// 	WorkDoneToken ProgressToken `json:"workDoneToken,omitempty"`
// }
//
// type WorkDoneProgressEnd struct {
// 	Kind string `json:"kind"`
//
// 	// Message is optional, a final message indicating to for example indicate the outcome of the operation.
// 	Message string `json:"message,omitempty"`
// }
//
// type WorkDoneProgressOptions struct {
// 	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
// }
//
// type WorkDoneProgress interface {
// 	Kind() string
// }
//
// type ProgressParams struct {
// 	// Token is the progress token provided by the client or server.
// 	Token ProgressToken `json:"token"`
//
// 	// Value is the progress data.
// 	Value interface{} `json:"value"`
// }
//
// type WorkDoneProgressCreateParams struct {
// 	Token ProgressToken `json:"token"`
// }
//
// type WorkDoneProgressCancelParams struct {
// 	Token ProgressToken `json:"token"`
// }
