package domain

import "context"

type Checker interface {
	Check(ctx context.Context, target string) (CheckResult, error)
}

type HTTPChecker struct{}

type TCPChecker struct{}

type PingChecker struct{}