package config

type Config struct {
	StoreSrv      []string `toml:"store_srv"`
	SectionNumber []int32  `toml:"section_number"`
}
