package jolt

import (
	"bytes"
)

type AggregatedInfluence struct {
	influence []Influence
}

func NewAggregatedInfluence() *AggregatedInfluence {
	return &AggregatedInfluence{}
}

func (ai *AggregatedInfluence) AddInfluence(influence Influence) {
	ai.influence = append(ai.influence, influence)
}

func (ai *AggregatedInfluence) Digest() (Digest, error) {
	var data bytes.Buffer

	for _, influence := range ai.influence {
		digest, err := influence.Digest()
		if err != nil {
			return "", err
		}
		data.WriteString(string(digest))
	}

	return Sha1(&data)
}
