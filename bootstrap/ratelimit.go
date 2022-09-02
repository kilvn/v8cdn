package bootstrap

import (
    "github.com/aoaostar/v8cdn_panel/config"
    "github.com/aoaostar/v8cdn_panel/pkg"
    "github.com/juju/ratelimit"
    "time"
)

func InitRateLimit() {
    fillInterval := time.Second * time.Duration(config.Env.RateLimit.FillInterval)
    pkg.RateLimitBucket = ratelimit.NewBucket(fillInterval, config.Env.RateLimit.Capacity)
}
