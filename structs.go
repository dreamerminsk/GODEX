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
	Magic         [8]uint8  //magic value
	Checksum      uint32    //adler32 checksum of the rest of the file (everything but Magic and this field)
	Signature     [20]uint8 //SHA-1 signature of the rest of the file (everything but Magic, Checksum, and this field)
	FileSize      uint32    //size of the entire file (including the header), in bytes
	HeaderSize    uint32    //size of the header (this entire section), in bytes = 0x70
	EndianTag     uint32    //endianness tag
	LinkSize      uint32    //size of the link section, or 0 if this file isn't statically linked
	LinkOff       uint32    //offset from the start of the file to the link section, or 0 if LinkSize == 0
	MapOff        uint32    //offset from the start of the file to the map item
	StringIdsSize uint32    //count of strings in the string ids list
	StringIdsOff  uint32    //offset from the start of the file to the string ids list, or 0 if StringIdsSize == 0
	TypeIdsSize   uint32    //count of elements in the type ids list, at most 65535
	TypeIdsOff    uint32    //offset from the start of the file to the type ids list, or 0 if TypeIdsSize == 0
	ProtoIdsSize  uint32    //count of elements in the prototype ids list, at most 65535
	ProtoIdsOff   uint32    //offset from the start of the file to the prototype ids list, or 0 if ProtoIdsSize == 0
	FieldIdsSize  uint32    //count of elements in the field ids list
	FieldIdsOff   uint32    //offset from the start of the file to the field ids list, or 0 if FieldIdsSize == 0
	MethodIdsSize uint32    //count of elements in the method ids list
	MethodIdsOff  uint32    //offset from the start of the file to the method identifiers list, or 0 if MethodIdsSize == 0
	ClassDefsSize uint32    //count of elements in the class defs list
	ClassDefsOff  uint32    //offset from the start of the file to the class defs list, or 0 if ClassDefsSize == 0
	DataSize      uint32    //Size of data section in bytes. Must be an even multiple of sizeof(uint).
	DataOff       uint32    //offset from the start of the file to the start of the data section.
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
