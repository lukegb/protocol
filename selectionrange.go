// Copyright 2020 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

// SelectionRange represents a selection range represents a part of a selection hierarchy.
//
// A selection range may have a parent selection range that contains it.
type SelectionRange struct {
	// Range is the Range of this selection range.
	Range Range `json:"range"`

	// Parent is the parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent,omitempty"`
}

type SelectionRangeClientCapabilities struct {
	// DynamicRegistration whether implementation supports dynamic registration for selection range providers. If this is set to `true`
	// the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type SelectionRangeOptions struct {
	WorkDoneProgressOptions
}

// SelectionRangeParams represents a parameter literal used in selection range requests.
type SelectionRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// TextDocument is the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Positions is the positions inside the text document.
	Positions []Position `json:"positions"`
}

type SelectionRangeRegistrationOptions struct {
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}
