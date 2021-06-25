/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmlog

import (
    "log"
)


func LogError(data ...interface{}) {
    log.Println("error:", data)
}

func LogDebug(data ...interface{}) {
    log.Println("debug:", data)
}
//EOF
