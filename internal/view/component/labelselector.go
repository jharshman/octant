package component

import "encoding/json"

// LabelSelector is a component for a single label within a selector
type LabelSelector struct {
	Metadata Metadata            `json:"metadata"`
	Config   LabelSelectorConfig `json:"config"`
}

// LabelSelectorConfig is the contents of LabelSelector
type LabelSelectorConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewLabelSelector creates a labelSelector component
func NewLabelSelector(k, v string) *LabelSelector {
	return &LabelSelector{
		Metadata: Metadata{
			Type: "labelSelector",
		},
		Config: LabelSelectorConfig{
			Key:   k,
			Value: v,
		},
	}
}

// GetMetadata accesses the components metadata. Implements ViewComponent.
func (t *LabelSelector) GetMetadata() Metadata {
	return t.Metadata
}

// IsEmpty specifes whether the component is considered empty. Implements ViewComponent.
func (t *LabelSelector) IsEmpty() bool {
	return t.Config.Key == "" && t.Config.Value == ""
}

// IsSelector marks the component as selector flavor. Implements Selector.
func (t *LabelSelector) IsSelector() {
}

type labelSelectorMarshal LabelSelector

// MarshalJSON implements json.Marshaler
func (t *LabelSelector) MarshalJSON() ([]byte, error) {
	m := labelSelectorMarshal(*t)
	m.Metadata.Type = "labelSelector"
	return json.Marshal(&m)
}
