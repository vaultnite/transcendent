package buildinfo

import (
    "fmt"
    "time"
)

var (
    Branch  = "unknown"
    Commit  = "unknown"
    Version = "dev"

    buildDateRaw = "unknown"
)

func BuildDate() (t time.Time) {
    parsed, err := time.Parse(time.RFC3339, buildDateRaw)
    if err != nil {
        return time.Time{}
    }
    return parsed
}

func String() string {
    return fmt.Sprintf("branch=%s commit=%s version=%s built=%s\n", Branch, Commit, Version, BuildDate())
}
