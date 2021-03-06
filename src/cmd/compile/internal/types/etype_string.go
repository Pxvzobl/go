// Code generated by "stringer -type EType -trimprefix T"; DO NOT EDIT.

package types

import "strconv"

const _EType_name = "xxxINT8UINT8INT16UINT16INT32UINT32INT64UINT64INTUINTUINTPTRCOMPLEX64COMPLEX128FLOAT32FLOAT64BOOLPTR32PTR64FUNCSLICEARRAYSTRUCTCHANMAPINTERFORWANYSTRINGUNSAFEPTRIDEALNILBLANKFUNCARGSCHANARGSDDDFIELDSSATUPLENTYPE"

var _EType_index = [...]uint8{0, 3, 7, 12, 17, 23, 28, 34, 39, 45, 48, 52, 59, 68, 78, 85, 92, 96, 101, 106, 110, 115, 120, 126, 130, 133, 138, 142, 145, 151, 160, 165, 168, 173, 181, 189, 197, 200, 205, 210}

func (i EType) String() string {
	if i >= EType(len(_EType_index)-1) {
		return "EType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EType_name[_EType_index[i]:_EType_index[i+1]]
}
