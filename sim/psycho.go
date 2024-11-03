package main

const (
	OPENNESS = iota + 1
	CONSCIENTIOUSNESS
	EXTRAVERSION
	AGREEABLENESS
	NEUROTICISM
)

type PsychoTraits struct {
	Openness          float32 `json:"openness"`
	Conscientiousness float32 `json:"conscientiousness"`
	Extraversion      float32 `json:"extraversion"`
	Agreeableness     float32 `json:"agreeableness"`
	Neuroticism       float32 `json:"neuroticism"`
}
