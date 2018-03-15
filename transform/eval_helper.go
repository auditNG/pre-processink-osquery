package transform

import (
	// "fmt"
	"github.com/Knetic/govaluate"
	"net"
)

var functions = map[string]govaluate.ExpressionFunction{
	"IPInSubnet": func(args ...interface{}) (interface{}, error) {
		ipstr := args[0].(string)
		subnet := args[1].(string)

		//Parse IP of the ip address in question
		ip := net.ParseIP(ipstr)
		if ip == nil {
			return false, nil
		}

		//Get the subnet to check against
		_, cidrnet, err := net.ParseCIDR(subnet)
		if err != nil {
			return false, err
		}

		contains := cidrnet.Contains(ip)
		return contains, nil
	},
}

type EvalHelper struct {
}

func (e EvalHelper) EvaluateFunction(expString string, fields map[string]interface{}) bool {
	expression, _ := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)
	result, _ := expression.Evaluate(fields)
	return result.(bool)
}
