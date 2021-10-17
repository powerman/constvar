package constvar

import "time"

func Time(c time.Time) func() time.Time { return func() time.Time { return c } }
