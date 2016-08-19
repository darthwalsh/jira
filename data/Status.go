package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command: slipscheme -pkg jiradata -overwrite ../schemas/TransitionsMeta.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

type Status struct {
	Description    string          `json:"description,omitempty" yaml:"description,omitempty"`
	IconURL        string          `json:"iconUrl,omitempty" yaml:"iconUrl,omitempty"`
	ID             string          `json:"id,omitempty" yaml:"id,omitempty"`
	Name           string          `json:"name,omitempty" yaml:"name,omitempty"`
	Self           string          `json:"self,omitempty" yaml:"self,omitempty"`
	StatusCategory *StatusCategory `json:"statusCategory,omitempty" yaml:"statusCategory,omitempty"`
	StatusColor    string          `json:"statusColor,omitempty" yaml:"statusColor,omitempty"`
}