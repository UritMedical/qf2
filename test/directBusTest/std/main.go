/**
 * @Author: Joey
 * @Description:使用代码生成器构造的项目模板
 * @Create Date: 2023/8/15 16:22
 */

package main

import (
	"fmt"
	qf "github.com/UritMedical/qf2"
	. "github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/test/directBusTest/std/api"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qdefine"
	"github.com/UritMedical/qf2/test/directBusTest/std/result"
	"github.com/UritMedical/qf2/test/directBusTest/std/tube"
)

func main() {
	db := NewSqlite("lis.db")
	api := api.New()
	plugins := []QPlugin{
		tube.New(db),
		result.New(db),
		api,
	}
	bus := qf.Bus.NewDirect(qf.Logger.NewFmt())
	qf.Run(bus, plugins) // 运行插件
	err := api.TubeInsert(Tube{
		BaseModel: BaseModel{
			Id:       1,
			FullInfo: "Tube1 Info",
			Summary:  "Tube1 Summary",
		},
	})

	fmt.Println("TubeInsert", err)
	err = api.TubeDelete(1)
	fmt.Println("TubeDelete", err)
	bus.Logger().Error("123")
	a := "a"
	_, e := fmt.Scan(&a)
	if e != nil {
		return
	}
}
