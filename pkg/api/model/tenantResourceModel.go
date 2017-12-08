// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model


type TenantResList []*TenantResource

type PagedTenantResList struct {
	List []*TenantResource `json:"list"`
	Length int `json:"length"`
}
//TenantResource path参数
//swagger:parameters getVolumes getDepVolumes
type TenantResource struct {
	AllocatedCPU  int    `json:"alloc_cpu"`
	AllocatedMEM  int    `json:"alloc_memory"`

	UsedCPU  int    `json:"used_cpu"`
	UsedMEM  int    `json:"used_memory"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
	EID  string `json:"eid"`
}
func (list TenantResList) Len() int {
	return len(list)
}

func (list TenantResList) Less(i, j int) bool {
	if list[i].UsedMEM > list[j].UsedMEM {
		return true
	} else if list[i].UsedMEM < list[j].UsedMEM {
		return false
	} else {
		return list[i].UsedCPU > list[j].UsedCPU
	}
}

func (list TenantResList) Swap(i, j int) {
	var temp *TenantResource = list[i]
	list[i] = list[j]
	list[j] = temp
}