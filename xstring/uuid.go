package xstring


/*
@version v1.0
@author nickzydeng
@copyright Copyright (c) 2018 Tencent Corporation, All Rights Reserved
@license http://opensource.org/licenses/gpl-license.php GNU Public License

You may not use this file except in compliance with the License.

Most recent version can be found at:
http://git.code.oa.com/going_proj/going_proj

Please see README.md for more information.
*/

import (
	"strings"

	"github.com/google/uuid"
)

// GetUUID generate uuid
func GetUUID() string {
	uuidStr := uuid.New().String()
	return strings.Replace(uuidStr, "-", "", -1)
}

