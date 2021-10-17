package constvar

import "time"

type Time time.Time

func (c Time) Val() time.Time { return time.Time(c) }
