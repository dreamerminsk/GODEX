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

type AccessFlags uint64

const (
	ACC_PUBLIC       AccessFlags = 1 << iota
	ACC_PRIVATE      AccessFlags = 1 << iota
	ACC_PROTECTED    AccessFlags = 1 << iota
	ACC_STATIC       AccessFlags = 1 << iota
	ACC_FINAL        AccessFlags = 1 << iota
	ACC_SYNCHRONIZED AccessFlags = 1 << iota
	ACC_VOLATILE     AccessFlags = 1 << iota
	ACC_TRANSIENT    AccessFlags = 1 << iota
	ACC_NATIVE       AccessFlags = 1 << iota
	ACC_INTERFACE    AccessFlags = 1 << iota
	ACC_ABSTRACT     AccessFlags = 1 << iota
	ACC_STRICT       AccessFlags = 1 << iota
	ACC_SYNTHETIC    AccessFlags = 1 << iota
	ACC_ANNOTATION   AccessFlags = 1 << iota
	ACC_ENUM         AccessFlags = 1 << iota
	_
	ACC_CONSTRUCTOR           AccessFlags = 1 << iota
	ACC_DECLARED_SYNCHRONIZED AccessFlags = 1 << iota
)

const (
	ACC_BRIDGE  AccessFlags = ACC_VOLATILE
	ACC_VARARGS AccessFlags = ACC_TRANSIENT
)
