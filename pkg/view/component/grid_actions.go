/*
 *  Copyright (c) 2020 the Octant contributors. All Rights Reserved.
 *  SPDX-License-Identifier: Apache-2.0
 *
 */

package component

import (
	"encoding/json"

	"github.com/vmware-tanzu/octant/pkg/action"
)

const (
	// GridActionKey is the key for grid action in a table row.
	GridActionKey = "_action"
)

// GridAction is an action that can be performed on a data grid row.
type GridAction struct {
	// Name is the name of action. It will be shown to the user.
	Name string `json:"name"`
	// ActionPath is the path of the action.
	ActionPath string `json:"actionPath"`
	// Payload is the payload that will be submitted with the action is invoked.
	Payload action.Payload `json:"payload"`
	// Confirmation is a confirmation that will be show to the user before the
	// action is invoked. It is optional.
	Confirmation *Confirmation `json:"confirmation"`
}

// GridActions add the ability to have specific actions for rows. This will allow for dynamic injection of actions
// that could be dependent on the content of a grid row.
type GridActions struct {
	base

	Config GridActionsConfig `json:"config"`
}

var _ Component = &GridActions{}

// NewGridActions creates an instance of GridActions.
func NewGridActions() *GridActions {
	a := GridActions{
		base: newBase(typeGridActions, nil),
	}

	return &a
}

// AddAction adds an action to GridActions.
func (a *GridActions) AddAction(name, actionPath string, payload action.Payload, confirmation *Confirmation) {
	gridAction := GridAction{
		Name:         name,
		ActionPath:   actionPath,
		Payload:      payload,
		Confirmation: confirmation,
	}

	a.AddGridAction(gridAction)
}

// AddGridAction adds a GridAction to GridActions.
func (a *GridActions) AddGridAction(gridAction GridAction) {
	a.Config.Actions = append(a.Config.Actions, gridAction)
}

type gridActionsMarshal GridActions

// MarshalJSON converts the GridActions to a JSON.
func (a GridActions) MarshalJSON() ([]byte, error) {
	m := gridActionsMarshal{
		base:   a.base,
		Config: a.Config,
	}

	m.Metadata.Type = typeGridActions
	return json.Marshal(&m)
}

// GridActionsConfig is configuration items for GridActions.
type GridActionsConfig struct {
	// Actions is a slice that contains actions that can be performed by the user.
	Actions []GridAction `json:"actions"`
}
