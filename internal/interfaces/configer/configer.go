package configer

import "GoldPriceGetter/internal/entities/config"

type Configer interface {
	GetConfig() (config.Config, error)
}
