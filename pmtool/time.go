/*
 * Copyright: Oleg Borodin <onborodin@gmail.com>
 */


package pmtool

import (
    "time"
)

func NewIsoTimestamp() string {
    return time.Now().Format(time.RFC3339)
}


//EOF

