package godex

type DexFile struct {
	Header        *HeaderItem
	StringIds     *[]StringIdItem
	TypeIds       *[]TypeIdItem
	ProtoIds      *[]ProtoIdItem
	FieldIds      *[]FieldIdItem
	MethodIds     *[]MethodIdItem
	ClassDefs     *[]ClassDefItem
	CallSiteIds   *[]CallSiteIdItem
	MethodHandles *[]MethodHandleItem
	Data          *[]uint8
	LinkData      *[]uint8
}

type HeaderItem struct {
}

type StringIdItem struct {
}

type TypeIdItem struct {
}

type ProtoIdItem struct {
}

type FieldIdItem struct {
}

type MethodIdItem struct {
}

type ClassDefItem struct {
}

type CallSiteIdItem struct {
}

type MethodHandleItem struct {
}
