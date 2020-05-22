// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"
	"log"
	"strings"

	"regexp"

	syntax "github.com/grafeas/grafeas/cel"
	"github.com/grafeas/grafeas/go/filtering/common"
	"github.com/grafeas/grafeas/go/filtering/operators"
	"github.com/grafeas/grafeas/go/filtering/parser"
)

//	var fs filterSql
//	fmt.Println("\nsql:", fs.makeSql(expr))

type MysqlFilterSql struct {
	selects     int
	data_prefix string
}

func (fs *MysqlFilterSql) sqlFromCall(func_name string, args []*syntax.Expr) string {

	var sql_op string
	switch func_name {
	case operators.Equals:
		sql_op = "="
	case operators.Greater:
		sql_op = ">"
	case operators.GreaterEquals:
		sql_op = ">="
	case operators.Less:
		sql_op = "<"
	case operators.LessEquals:
		sql_op = "<="
	case operators.NotEquals:
		sql_op = "!="
	case operators.LogicalAnd:
		sql_op = "AND"
	case operators.LogicalOr:
		sql_op = "OR"
	case operators.Index:
		sql_op = "["
	default:
		sql_op = ""
	}
	var arg_names []string
	for _, arg := range args {
		arg_names = append(arg_names, fs.makeSql(arg))
	}
	if sql_op == "[" {
		return fmt.Sprintf("%s[%s]", arg_names[0], arg_names[1])
	} else if sql_op != "" {
		return fmt.Sprintf("(%s %s %s)", arg_names[0], sql_op, arg_names[1])
	} else {
		return fmt.Sprintf("%s(%s)", func_name, strings.Join(arg_names, ", "))
	}
}

func (fs *MysqlFilterSql) sqlFromSelect(select_node syntax.Expr_Select) string {
	operand := fs.makeSql(select_node.GetOperand())
	field := select_node.GetField()
	return fmt.Sprintf("%s.%s", operand, field)
}

func (fs *MysqlFilterSql) getConstantValue(const_expr syntax.Constant) string {

	switch const_expr.GetConstantKind().(type) {
	case *syntax.Constant_Int64Value:
		return fmt.Sprintf("%d", const_expr.GetInt64Value())
	case *syntax.Constant_Uint64Value:
		return fmt.Sprintf("%d", const_expr.GetUint64Value())
	case *syntax.Constant_DoubleValue:
		return fmt.Sprintf("%f", const_expr.GetDoubleValue())
	case *syntax.Constant_StringValue:
		return fmt.Sprintf("\"%s\"", const_expr.GetStringValue())
	}
	return "NO CONST"
}

func (fs *MysqlFilterSql) makeSql(node *syntax.Expr) string {
	//log.Println("depth: ", fs.depth, "node:", node)
	switch node.GetExprKind().(type) {
	case *syntax.Expr_CallExpr:
		func_node := *node.GetCallExpr()
		return fs.sqlFromCall(func_node.Function, func_node.Args)
	case *syntax.Expr_SelectExpr:
		select_node := *node.GetSelectExpr()
		fs.selects++
		ret_str := fs.sqlFromSelect(select_node)
		fs.selects--
		if fs.selects == 0 {
			return fs.data_prefix + "data->'$." + ret_str + "'"
		} else {
			return ret_str
		}
	case *syntax.Expr_IdentExpr:
		i_expr := *node.GetIdentExpr()
		// I'm not entirely sure this is the right thing here.
		// We'll see though.
		if fs.selects > 0 {
			return i_expr.Name
		} else {
			return fs.data_prefix + "data->'$." + i_expr.Name + "'"
		}
	case *syntax.Expr_ConstExpr:
		c_expr := *node.GetConstExpr()
		return fs.getConstantValue(c_expr)
	}

	return "NO SQL"

}

func (fs *MysqlFilterSql) ParseFilter(filter string, data_prefix string) string {

	fs.data_prefix = data_prefix

	// log.Println(filter)
	// replace string value for kind with integer
	kindPattern := `kind=".*?"`
	re := regexp.MustCompile(kindPattern)
	indices := re.FindStringIndex(filter)
	if indices != nil {
		int_val := 0
		switch filter[indices[0]+6 : indices[1]-1] {
		case "VULNERABILITY":
			int_val = 1
		case "BUILD":
			int_val = 2
		case "IMAGE":
			int_val = 3
		case "PACKAGE":
			int_val = 4
		case "DEPLOYMENT":
			int_val = 5
		case "DISCOVERY":
			int_val = 6
		case "ATTESTATION":
			int_val = 7
		case "INTOTO":
			int_val = 8
		}
		filter = filter[:indices[0]+5] + fmt.Sprintf("%d", int_val) + filter[indices[1]:]
		// log.Println(filter)
	}
	s := common.NewStringSource(filter, "urlParam") // function
	result, err := parser.Parse(s)
	if err != nil {
		log.Println(err)
		return ""
	}
	// log.Println(result.Expr.GetCallExpr().Args[1].GetCallExpr().Args[0].GetIdentExpr().Name)
	// log.Println(result.Expr)
	sql := fs.makeSql(result.Expr)
	//log.Println(sql)
	return sql
}
