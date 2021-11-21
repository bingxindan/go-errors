package jk_error

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// 正常方式输出错误
func TestMyselfDriverNormal(t *testing.T) {
	ctx := context.Background()
	// NoRows:id=0;HaveRows:id=1
	// id := 1
	id := 0
	err := MyselfQueryNoRows(ctx, id)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	fmt.Println("succ")
}

// trace输出错误
func TestMyselfDriverTrace(t *testing.T) {
	ctx := context.Background()
	// NoRows:id=0;HaveRows:id=1
	// id := 1
	id := 0
	err := MyselfQueryNoRows(ctx, id)
	if err, ok := err.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
		return
	}
	fmt.Println("succ")
}
