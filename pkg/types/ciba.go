package types

type TranslateResp struct {
	Word     string   `json:"word,omitempty"`
	CN       string   `json:"cn,omitempty"`
	EN       []string `json:"en,omitempty"`
	Examples []string `json:"examples,omitempty"`
	Synomyms []string `json:"synomyms,omitempty"`
}

type CibaResp struct {
	Status  int         `json:"status,omitempty"`
	Message CibaMessage `json:"message,omitempty"`
}

type CibaMessage struct {
	Bidec   Bidec     `json:"bidec,omitempty"`
	Synonym []Synonym `json:"synonym,omitempty"`
	Collins []Collins `json:"collins,omitempty"`
	EEMean  []EE      `json:"ee_mean,omitempty"`
}

// EE
type EE struct {
	PartName string    `json:"part_name,omitempty"`
	Means    []EEMeans `json:"means,omitempty"`
}

type EEMeans struct {
	WordMean  string   `json:"word_mean,omitempty"`
	Sentences []string `json:"sentences,omitempty"`
}

// collins

type Collins struct {
	Entry []CollinsEntry `json:"entry,omitempty"`
}

type CollinsEntry struct {
	Def     string           `json:"def,omitempty"`
	Tran    string           `json:"tran,omitempty"`
	Posp    string           `json:"posp,omitempty"`
	Example []CollinsExample `json:"example,omitempty"`
}

type CollinsExample struct {
	Ex   string `json:"ex,omitempty"`
	Tran string `json:"tran,omitempty"`
}

// type syn
type Synonym struct {
	PartName string         `json:"part_name,omitempty"`
	Means    []SynonymMeans `json:"means,omitempty"`
}

type SynonymMeans struct {
	WordMean string   `json:"word_mean,omitempty"`
	Cis      []string `json:"cis,omitempty"`
}

// bidec
type Bidec struct {
	WordName string      `json:"word_name,omitempty"`
	Parts    []BidecPart `json:"parts,omitempty"`
}

type BidecPart struct {
	PartId   string
	PartName string
	WordId   string
	Means    []BidecPartMeans
}

type BidecPartMeans struct {
	MeanId    string                   `json:"mean_id,omitempty"`
	PartId    string                   `json:"part_id,omitempty"`
	WordMean  string                   `json:"word_mean,omitempty"`
	Sentences []BidecPartMeansSentence `json:"sentences,omitempty"`
}

type BidecPartMeansSentence struct {
	En string `json:"en,omitempty"`
	Cn string `json:"cn,omitempty"`
}
