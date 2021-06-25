/*
 * Copyright: Oleg Borodin <onborodin@gmail.com>
 */

package pmtool

import (
    "github.com/satori/go.uuid"
)

func NewUUID() string {
    id := uuid.NewV4()
    return id.String()
}

//EOF

