package mdclasses

func (o AttributeTypedValue) IsNull() bool {
	return len(o.Value) == 0
}

func (o NumberQualifiers) IsNull() bool  {
	return len(o.Precision) == 0
}

func (o StringQualifiers) IsNull() bool  {
	return len(o.Length) == 0
}

func (o DateQualifiers) IsNull() bool  {
	return len(o.DateFractions) == 0
}

func (o ObjectTypeRef) IsNull() bool  {
	return len(o.TypeId.Value) == 0 && len(o.ValueTypeId.Value) == 0
}

func (o ObjectKeyValueType) IsNull() bool  {
	return len(o.Key) == 0 && len(o.Value) == 0
}
