package beyarkey

import "time"

type ChangesFile struct {
	Changes []struct {
		Stage  int       `yaml:"stage"`
		Start  time.Time `yaml:"start"`
		Finsh  time.Time `yanl:"finsh"`
		Source string    `yanl:"source"`
	} `yaml:"changes"`
	HistoricalChanges []interface{} `yaml:"historical_changes"`
}
