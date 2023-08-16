/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 12:11
 */

package qdefine

import (
	. "github.com/UritMedical/qf2/define"
)

type Tube struct {
	BaseModel
}
type Result struct {
	BaseModel
	TubeId uint64
}
