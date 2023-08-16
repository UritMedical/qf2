/**
 * @Author: QPluginBuilder
 * @Description:Model define
 * @Create Date: {{CreateTime}}
 */

package qDefine

import (
	. "github.com/UritMedical/qf2/define"
)

//
// Result
// @Description: 检验结果
//
type Result struct {
	//	//
    // 
    // @Description: 基础类
    //
	 BaseModel `gorm:""`

//	//
    // TubeId
    // @Description: 试管序号
    //
	TubeId uint64 `gorm:"index"`

}

//
// Tube
// @Description: 试管信息
//
type Tube struct {
	//	//
    // 
    // @Description: 基础类
    //
	 BaseModel `gorm:""`

}

