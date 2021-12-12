package main

import (
	"database/sql"
	"errors"
	"fmt"
)
//问：
//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//答：
//应该wrap一下，因为业务层面解决问题的时候需要知道是哪个程序，哪条sql，或者哪个条件没有查询。不然排查问题很困难

import (
	"database/sql"
	"errors"
	"fmt"
)

func foo() error{
	return fmt.Errorf("foo failed: %w ",sql.ErrNoRows)
	//return sql.ErrNoRows
}
func  bar() error{
	if err:= foo();err!=nil{
		return fmt.Errorf("bar failed: %w ",foo())
	}
	return nil
}
func main(){
	err := bar()
	if errors.Is(err,sql.ErrNoRows){
		fmt.Printf("data not found,  %+v \n",err)
		return
	}
}