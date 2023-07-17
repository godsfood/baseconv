package options

import (
	"testing"
)

func TestOptions(t *testing.T) {
	t.Run("Zero padding 1", func(t *testing.T) {
		o := BaseConversion()

		if o.ZeroPadding != nil {
			t.Fatal()
		}
	})

	t.Run("Zero padding 2", func(t *testing.T) {
		o := BaseConversion()
		o.SetZeroPadding(false)

		if o.ZeroPadding == nil || *o.ZeroPadding != false {
			t.Fatal()
		}
	})

	t.Run("Zero padding 3", func(t *testing.T) {
		o := BaseConversion()
		o.SetZeroPadding(true)

		if o.ZeroPadding == nil || *o.ZeroPadding != true {
			t.Fatal()
		}
	})
}

func TestMergeOptions(t *testing.T) {
	t.Run("Merge options 1", func(t *testing.T) {
		var o1, o2, o3 *BaseConversionOptions

		o := MergeBaseConversionOptions(o1, o2, o3)

		if o.ZeroPadding != nil {
			t.Fatal()
		}
	})

	t.Run("Merge options 2", func(t *testing.T) {
		var o1, o2, o3 *BaseConversionOptions = BaseConversion(), BaseConversion(), nil
		o1.SetZeroPadding(true)
		o2.SetZeroPadding(false)

		o := MergeBaseConversionOptions(o1, o2, o3)

		if o.ZeroPadding == nil || *o.ZeroPadding != false {
			t.Fatal()
		}
	})

	t.Run("Merge options 3", func(t *testing.T) {
		var o1, o2, o3 *BaseConversionOptions = BaseConversion(), BaseConversion(), BaseConversion()
		o1.SetZeroPadding(false)
		o2.SetZeroPadding(true)

		o := MergeBaseConversionOptions(o1, o2, o3)

		if o.ZeroPadding == nil || *o.ZeroPadding != true {
			t.Fatal()
		}
	})
}
