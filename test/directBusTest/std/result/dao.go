/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/15 12:11
 */

package result

import (
	. "github.com/UritMedical/qf2/define"
	. "github.com/UritMedical/qf2/test/directBusTest/std/qDefine"
)

type dao struct {
	BaseDao
}

func (d *dao) deleteWithTube(tubeId uint64) QError {
	result := d.DB().Where("TubeId = ?", tubeId).Delete(&Tube{})
	if result.Error != nil {
		return Error(ErrorCodeDeleteFailure, result.Error.Error())
	}
	if result.RowsAffected > 0 {
		return nil
	}
	return Error(ErrorCodeDeleteFailure, "record not found")
}
