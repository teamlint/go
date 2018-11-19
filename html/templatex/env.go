package templatex

import "os"

// OS:
func env(s string) string       { return os.Getenv(s) }
func expandEnv(s string) string { return os.ExpandEnv(s) }
