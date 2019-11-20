package app

const ROUTER_DEFAULT_TESTER = "tester"

func NewDefaultAppTester() App {
	return NewAppGoCraft()
}
