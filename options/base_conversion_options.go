package options

type BaseConversionOptions struct {
	ZeroPadding *bool
}

func BaseConversion() *BaseConversionOptions {
	return new(BaseConversionOptions)
}

func (o *BaseConversionOptions) SetZeroPadding(b bool) *BaseConversionOptions {
	o.ZeroPadding = &b
	return o
}

func MergeBaseConversionOptions(opts ...*BaseConversionOptions) *BaseConversionOptions {
	o := BaseConversion()

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.ZeroPadding != nil {
			o.ZeroPadding = opt.ZeroPadding
		}
	}

	return o
}
