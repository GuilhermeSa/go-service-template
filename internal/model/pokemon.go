package model

// Pokemon is a struct that represents a pokemon
type Pokemon struct {
	Abilities              []Ability     `json:"abilities"`
	BaseExperience         int64         `json:"base_experience"`
	Forms                  []Species     `json:"forms"`
	GameIndices            []GameIndex   `json:"game_indices"`
	Height                 int64         `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int64         `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Move        `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int64         `json:"order"`
	PastTypes              []interface{} `json:"past_types"`
	Species                Species       `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stat        `json:"stats"`
	Types                  []Type        `json:"types"`
	Weight                 int64         `json:"weight"`
}

type Ability struct {
	Ability  Species `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int64   `json:"slot"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GameIndex struct {
	GameIndex int64   `json:"game_index"`
	Version   Species `json:"version"`
}

type Move struct {
	Move                Species              `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int64   `json:"level_learned_at"`
	MoveLearnMethod Species `json:"move_learn_method"`
	VersionGroup    Species `json:"version_group"`
}

type GenerationV struct {
	BlackWhite Sprites `json:"black-white"`
}

type GenerationIv struct {
	DiamondPearl        Sprites `json:"diamond-pearl"`
	HeartgoldSoulsilver Sprites `json:"heartgold-soulsilver"`
	Platinum            Sprites `json:"platinum"`
}

type Versions struct {
	GenerationI    GenerationI     `json:"generation-i"`
	GenerationIi   GenerationIi    `json:"generation-ii"`
	GenerationIii  GenerationIii   `json:"generation-iii"`
	GenerationIv   GenerationIv    `json:"generation-iv"`
	GenerationV    GenerationV     `json:"generation-v"`
	GenerationVi   map[string]Home `json:"generation-vi"`
	GenerationVii  GenerationVii   `json:"generation-vii"`
	GenerationViii GenerationViii  `json:"generation-viii"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
	Other            *Other      `json:"other,omitempty"`
	Versions         *Versions   `json:"versions,omitempty"`
	Animated         *Sprites    `json:"animated,omitempty"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  RedBlue `json:"yellow"`
}

type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Gold    `json:"silver"`
}

type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}

type Gold struct {
	BackDefault      string  `json:"back_default"`
	BackShiny        string  `json:"back_shiny"`
	FrontDefault     string  `json:"front_default"`
	FrontShiny       string  `json:"front_shiny"`
	FrontTransparent *string `json:"front_transparent,omitempty"`
}

type GenerationIii struct {
	Emerald          OfficialArtwork `json:"emerald"`
	FireredLeafgreen Gold            `json:"firered-leafgreen"`
	RubySapphire     Gold            `json:"ruby-sapphire"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Home struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type GenerationVii struct {
	Icons             DreamWorld `json:"icons"`
	UltraSunUltraMoon Home       `json:"ultra-sun-ultra-moon"`
}

type DreamWorld struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type GenerationViii struct {
	Icons DreamWorld `json:"icons"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}

type Stat struct {
	BaseStat int64   `json:"base_stat"`
	Effort   int64   `json:"effort"`
	Stat     Species `json:"stat"`
}

type Type struct {
	Slot int64   `json:"slot"`
	Type Species `json:"type"`
}
