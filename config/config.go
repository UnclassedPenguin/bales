package config

// Struct for configuration.
type Configuration struct {
  DatabaseDir  string`yaml:"DatabaseDir"`
  RealDatabase string`yaml:"RealDatabase"`
  TestDatabase string`yaml:"TestDatabase"`
}
