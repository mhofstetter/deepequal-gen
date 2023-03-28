//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

// Code generated by deepequal-gen. DO NOT EDIT.

package wholepkg

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ManualStructAlias) DeepEqual(other *ManualStructAlias) bool {
	if other == nil {
		return false
	}

	if in.StringField != other.StringField {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructB) DeepEqual(other *StructB) bool {
	if other == nil {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedInt) DeepEqual(other *StructEmbedInt) bool {
	if other == nil {
		return false
	}

	if in.int != other.int {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedManualStruct) DeepEqual(other *StructEmbedManualStruct) bool {
	if other == nil {
		return false
	}

	if in.ManualStruct != other.ManualStruct {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedPointer) DeepEqual(other *StructEmbedPointer) bool {
	if other == nil {
		return false
	}

	if (in.int == nil) != (other.int == nil) {
		return false
	} else if in.int != nil {
		if *in.int != *other.int {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedStructPrimitivePointers) DeepEqual(other *StructEmbedStructPrimitivePointers) bool {
	if other == nil {
		return false
	}

	if !in.StructPrimitivePointers.DeepEqual(&other.StructPrimitivePointers) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedStructPrimitives) DeepEqual(other *StructEmbedStructPrimitives) bool {
	if other == nil {
		return false
	}

	if in.StructPrimitives != other.StructPrimitives {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmbedStructSlices) DeepEqual(other *StructEmbedStructSlices) bool {
	if other == nil {
		return false
	}

	if !in.StructSlices.DeepEqual(&other.StructSlices) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEmpty) DeepEqual(other *StructEmpty) bool {
	if other == nil {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructEverything) DeepEqual(other *StructEverything) bool {
	if other == nil {
		return false
	}

	if in.BoolField != other.BoolField {
		return false
	}
	if in.IntField != other.IntField {
		return false
	}
	if in.StringField != other.StringField {
		return false
	}
	if in.FloatField != other.FloatField {
		return false
	}
	if in.StructField != other.StructField {
		return false
	}

	if in.EmptyStructField != other.EmptyStructField {
		return false
	}

	if in.ManualStructField != other.ManualStructField {
		return false
	}

	if in.ManualStructAliasField != other.ManualStructAliasField {
		return false
	}

	if (in.BoolPtrField == nil) != (other.BoolPtrField == nil) {
		return false
	} else if in.BoolPtrField != nil {
		if *in.BoolPtrField != *other.BoolPtrField {
			return false
		}
	}

	if (in.IntPtrField == nil) != (other.IntPtrField == nil) {
		return false
	} else if in.IntPtrField != nil {
		if *in.IntPtrField != *other.IntPtrField {
			return false
		}
	}

	if (in.StringPtrField == nil) != (other.StringPtrField == nil) {
		return false
	} else if in.StringPtrField != nil {
		if *in.StringPtrField != *other.StringPtrField {
			return false
		}
	}

	if (in.FloatPtrField == nil) != (other.FloatPtrField == nil) {
		return false
	} else if in.FloatPtrField != nil {
		if *in.FloatPtrField != *other.FloatPtrField {
			return false
		}
	}

	if !in.PrimitivePointersField.DeepEqual(&other.PrimitivePointersField) {
		return false
	}

	if (in.ManualStructPtrField == nil) != (other.ManualStructPtrField == nil) {
		return false
	} else if in.ManualStructPtrField != nil {
		if !in.ManualStructPtrField.DeepEqual(other.ManualStructPtrField) {
			return false
		}
	}

	if (in.ManualStructAliasPtrField == nil) != (other.ManualStructAliasPtrField == nil) {
		return false
	} else if in.ManualStructAliasPtrField != nil {
		if !in.ManualStructAliasPtrField.DeepEqual(other.ManualStructAliasPtrField) {
			return false
		}
	}

	if ((in.SliceBoolField != nil) && (other.SliceBoolField != nil)) || ((in.SliceBoolField == nil) != (other.SliceBoolField == nil)) {
		in, other := &in.SliceBoolField, &other.SliceBoolField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceByteField != nil) && (other.SliceByteField != nil)) || ((in.SliceByteField == nil) != (other.SliceByteField == nil)) {
		in, other := &in.SliceByteField, &other.SliceByteField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceIntField != nil) && (other.SliceIntField != nil)) || ((in.SliceIntField == nil) != (other.SliceIntField == nil)) {
		in, other := &in.SliceIntField, &other.SliceIntField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceStringField != nil) && (other.SliceStringField != nil)) || ((in.SliceStringField == nil) != (other.SliceStringField == nil)) {
		in, other := &in.SliceStringField, &other.SliceStringField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceFloatField != nil) && (other.SliceFloatField != nil)) || ((in.SliceFloatField == nil) != (other.SliceFloatField == nil)) {
		in, other := &in.SliceFloatField, &other.SliceFloatField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if !in.SlicesField.DeepEqual(&other.SlicesField) {
		return false
	}

	if ((in.SliceManualStructField != nil) && (other.SliceManualStructField != nil)) || ((in.SliceManualStructField == nil) != (other.SliceManualStructField == nil)) {
		in, other := &in.SliceManualStructField, &other.SliceManualStructField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.ManualSliceField != nil) && (other.ManualSliceField != nil)) || ((in.ManualSliceField == nil) != (other.ManualSliceField == nil)) {
		in, other := &in.ManualSliceField, &other.ManualSliceField
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructExplicitObject) DeepEqual(other *StructExplicitObject) bool {
	if other == nil {
		return false
	}

	if in.x != other.x {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructExplicitSelectorExplicitObject) DeepEqual(other *StructExplicitSelectorExplicitObject) bool {
	if other == nil {
		return false
	}

	if in.StructTypeMeta != other.StructTypeMeta {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructNonPointerExplicitObject) DeepEqual(other *StructNonPointerExplicitObject) bool {
	if other == nil {
		return false
	}

	if in.x != other.x {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructObjectAndList) DeepEqual(other *StructObjectAndList) bool {
	if other == nil {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructObjectAndObject) DeepEqual(other *StructObjectAndObject) bool {
	if other == nil {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructPrimitivePointers) DeepEqual(other *StructPrimitivePointers) bool {
	if other == nil {
		return false
	}

	if (in.BoolPtrField == nil) != (other.BoolPtrField == nil) {
		return false
	} else if in.BoolPtrField != nil {
		if *in.BoolPtrField != *other.BoolPtrField {
			return false
		}
	}

	if (in.IntPtrField == nil) != (other.IntPtrField == nil) {
		return false
	} else if in.IntPtrField != nil {
		if *in.IntPtrField != *other.IntPtrField {
			return false
		}
	}

	if (in.StringPtrField == nil) != (other.StringPtrField == nil) {
		return false
	} else if in.StringPtrField != nil {
		if *in.StringPtrField != *other.StringPtrField {
			return false
		}
	}

	if (in.FloatPtrField == nil) != (other.FloatPtrField == nil) {
		return false
	} else if in.FloatPtrField != nil {
		if *in.FloatPtrField != *other.FloatPtrField {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructPrimitivePointersAlias) DeepEqual(other *StructPrimitivePointersAlias) bool {
	if other == nil {
		return false
	}

	if (in.BoolPtrField == nil) != (other.BoolPtrField == nil) {
		return false
	} else if in.BoolPtrField != nil {
		if *in.BoolPtrField != *other.BoolPtrField {
			return false
		}
	}

	if (in.IntPtrField == nil) != (other.IntPtrField == nil) {
		return false
	} else if in.IntPtrField != nil {
		if *in.IntPtrField != *other.IntPtrField {
			return false
		}
	}

	if (in.StringPtrField == nil) != (other.StringPtrField == nil) {
		return false
	} else if in.StringPtrField != nil {
		if *in.StringPtrField != *other.StringPtrField {
			return false
		}
	}

	if (in.FloatPtrField == nil) != (other.FloatPtrField == nil) {
		return false
	} else if in.FloatPtrField != nil {
		if *in.FloatPtrField != *other.FloatPtrField {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructPrimitives) DeepEqual(other *StructPrimitives) bool {
	if other == nil {
		return false
	}

	if in.BoolField != other.BoolField {
		return false
	}
	if in.IntField != other.IntField {
		return false
	}
	if in.StringField != other.StringField {
		return false
	}
	if in.FloatField != other.FloatField {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructPrimitivesAlias) DeepEqual(other *StructPrimitivesAlias) bool {
	if other == nil {
		return false
	}

	if in.BoolField != other.BoolField {
		return false
	}
	if in.IntField != other.IntField {
		return false
	}
	if in.StringField != other.StringField {
		return false
	}
	if in.FloatField != other.FloatField {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructSlices) DeepEqual(other *StructSlices) bool {
	if other == nil {
		return false
	}

	if ((in.SliceBoolField != nil) && (other.SliceBoolField != nil)) || ((in.SliceBoolField == nil) != (other.SliceBoolField == nil)) {
		in, other := &in.SliceBoolField, &other.SliceBoolField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceByteField != nil) && (other.SliceByteField != nil)) || ((in.SliceByteField == nil) != (other.SliceByteField == nil)) {
		in, other := &in.SliceByteField, &other.SliceByteField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceIntField != nil) && (other.SliceIntField != nil)) || ((in.SliceIntField == nil) != (other.SliceIntField == nil)) {
		in, other := &in.SliceIntField, &other.SliceIntField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceStringField != nil) && (other.SliceStringField != nil)) || ((in.SliceStringField == nil) != (other.SliceStringField == nil)) {
		in, other := &in.SliceStringField, &other.SliceStringField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceFloatField != nil) && (other.SliceFloatField != nil)) || ((in.SliceFloatField == nil) != (other.SliceFloatField == nil)) {
		in, other := &in.SliceFloatField, &other.SliceFloatField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivesField != nil) && (other.SliceStructPrimitivesField != nil)) || ((in.SliceStructPrimitivesField == nil) != (other.SliceStructPrimitivesField == nil)) {
		in, other := &in.SliceStructPrimitivesField, &other.SliceStructPrimitivesField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivesAliasField != nil) && (other.SliceStructPrimitivesAliasField != nil)) || ((in.SliceStructPrimitivesAliasField == nil) != (other.SliceStructPrimitivesAliasField == nil)) {
		in, other := &in.SliceStructPrimitivesAliasField, &other.SliceStructPrimitivesAliasField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivePointersField != nil) && (other.SliceStructPrimitivePointersField != nil)) || ((in.SliceStructPrimitivePointersField == nil) != (other.SliceStructPrimitivePointersField == nil)) {
		in, other := &in.SliceStructPrimitivePointersField, &other.SliceStructPrimitivePointersField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivePointersAliasField != nil) && (other.SliceStructPrimitivePointersAliasField != nil)) || ((in.SliceStructPrimitivePointersAliasField == nil) != (other.SliceStructPrimitivePointersAliasField == nil)) {
		in, other := &in.SliceStructPrimitivePointersAliasField, &other.SliceStructPrimitivePointersAliasField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceManualStructField != nil) && (other.SliceManualStructField != nil)) || ((in.SliceManualStructField == nil) != (other.SliceManualStructField == nil)) {
		in, other := &in.SliceManualStructField, &other.SliceManualStructField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.ManualSliceField != nil) && (other.ManualSliceField != nil)) || ((in.ManualSliceField == nil) != (other.ManualSliceField == nil)) {
		in, other := &in.ManualSliceField, &other.ManualSliceField
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructSlicesAlias) DeepEqual(other *StructSlicesAlias) bool {
	if other == nil {
		return false
	}

	if ((in.SliceBoolField != nil) && (other.SliceBoolField != nil)) || ((in.SliceBoolField == nil) != (other.SliceBoolField == nil)) {
		in, other := &in.SliceBoolField, &other.SliceBoolField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceByteField != nil) && (other.SliceByteField != nil)) || ((in.SliceByteField == nil) != (other.SliceByteField == nil)) {
		in, other := &in.SliceByteField, &other.SliceByteField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceIntField != nil) && (other.SliceIntField != nil)) || ((in.SliceIntField == nil) != (other.SliceIntField == nil)) {
		in, other := &in.SliceIntField, &other.SliceIntField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceStringField != nil) && (other.SliceStringField != nil)) || ((in.SliceStringField == nil) != (other.SliceStringField == nil)) {
		in, other := &in.SliceStringField, &other.SliceStringField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceFloatField != nil) && (other.SliceFloatField != nil)) || ((in.SliceFloatField == nil) != (other.SliceFloatField == nil)) {
		in, other := &in.SliceFloatField, &other.SliceFloatField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivesField != nil) && (other.SliceStructPrimitivesField != nil)) || ((in.SliceStructPrimitivesField == nil) != (other.SliceStructPrimitivesField == nil)) {
		in, other := &in.SliceStructPrimitivesField, &other.SliceStructPrimitivesField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivesAliasField != nil) && (other.SliceStructPrimitivesAliasField != nil)) || ((in.SliceStructPrimitivesAliasField == nil) != (other.SliceStructPrimitivesAliasField == nil)) {
		in, other := &in.SliceStructPrimitivesAliasField, &other.SliceStructPrimitivesAliasField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivePointersField != nil) && (other.SliceStructPrimitivePointersField != nil)) || ((in.SliceStructPrimitivePointersField == nil) != (other.SliceStructPrimitivePointersField == nil)) {
		in, other := &in.SliceStructPrimitivePointersField, &other.SliceStructPrimitivePointersField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceStructPrimitivePointersAliasField != nil) && (other.SliceStructPrimitivePointersAliasField != nil)) || ((in.SliceStructPrimitivePointersAliasField == nil) != (other.SliceStructPrimitivePointersAliasField == nil)) {
		in, other := &in.SliceStructPrimitivePointersAliasField, &other.SliceStructPrimitivePointersAliasField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.SliceManualStructField != nil) && (other.SliceManualStructField != nil)) || ((in.SliceManualStructField == nil) != (other.SliceManualStructField == nil)) {
		in, other := &in.SliceManualStructField, &other.SliceManualStructField
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.ManualSliceField != nil) && (other.ManualSliceField != nil)) || ((in.ManualSliceField == nil) != (other.ManualSliceField == nil)) {
		in, other := &in.ManualSliceField, &other.ManualSliceField
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructStructPrimitivePointers) DeepEqual(other *StructStructPrimitivePointers) bool {
	if other == nil {
		return false
	}

	if !in.StructField.DeepEqual(&other.StructField) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructStructPrimitives) DeepEqual(other *StructStructPrimitives) bool {
	if other == nil {
		return false
	}

	if in.StructField != other.StructField {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *StructStructSlices) DeepEqual(other *StructStructSlices) bool {
	if other == nil {
		return false
	}

	if !in.StructField.DeepEqual(&other.StructField) {
		return false
	}

	return true
}
