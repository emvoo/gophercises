package types

type Stories []struct {
	Intro struct {
		Story
	} `json:"intro"`
	NewYork struct {
		Story
	} `json:"new-york"`
	Debate struct {
		Story
	} `json:"debate"`
	SeanKelly struct {
		Story
	} `json:"sean-kelly"`
	MarkBates struct {
		Story
	} `json:"mark-bates"`
	Denver struct {
		Story
	} `json:"denver"`
	Home struct {
		Story
	} `json:"home"`
}

type Story struct {
	Title   string    `json:"title,omitempty"`
	Story   []string  `json:"story,omitempty"`
	Options []Options `json:"options,omitempty"`
}

type Options struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}