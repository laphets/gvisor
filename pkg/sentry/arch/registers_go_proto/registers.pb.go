// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: pkg/sentry/arch/registers.proto

package registers_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AMD64Registers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rax     uint64 `protobuf:"varint,1,opt,name=rax,proto3" json:"rax,omitempty"`
	Rbx     uint64 `protobuf:"varint,2,opt,name=rbx,proto3" json:"rbx,omitempty"`
	Rcx     uint64 `protobuf:"varint,3,opt,name=rcx,proto3" json:"rcx,omitempty"`
	Rdx     uint64 `protobuf:"varint,4,opt,name=rdx,proto3" json:"rdx,omitempty"`
	Rsi     uint64 `protobuf:"varint,5,opt,name=rsi,proto3" json:"rsi,omitempty"`
	Rdi     uint64 `protobuf:"varint,6,opt,name=rdi,proto3" json:"rdi,omitempty"`
	Rsp     uint64 `protobuf:"varint,7,opt,name=rsp,proto3" json:"rsp,omitempty"`
	Rbp     uint64 `protobuf:"varint,8,opt,name=rbp,proto3" json:"rbp,omitempty"`
	R8      uint64 `protobuf:"varint,9,opt,name=r8,proto3" json:"r8,omitempty"`
	R9      uint64 `protobuf:"varint,10,opt,name=r9,proto3" json:"r9,omitempty"`
	R10     uint64 `protobuf:"varint,11,opt,name=r10,proto3" json:"r10,omitempty"`
	R11     uint64 `protobuf:"varint,12,opt,name=r11,proto3" json:"r11,omitempty"`
	R12     uint64 `protobuf:"varint,13,opt,name=r12,proto3" json:"r12,omitempty"`
	R13     uint64 `protobuf:"varint,14,opt,name=r13,proto3" json:"r13,omitempty"`
	R14     uint64 `protobuf:"varint,15,opt,name=r14,proto3" json:"r14,omitempty"`
	R15     uint64 `protobuf:"varint,16,opt,name=r15,proto3" json:"r15,omitempty"`
	Rip     uint64 `protobuf:"varint,17,opt,name=rip,proto3" json:"rip,omitempty"`
	Rflags  uint64 `protobuf:"varint,18,opt,name=rflags,proto3" json:"rflags,omitempty"`
	OrigRax uint64 `protobuf:"varint,19,opt,name=orig_rax,json=origRax,proto3" json:"orig_rax,omitempty"`
	Cs      uint64 `protobuf:"varint,20,opt,name=cs,proto3" json:"cs,omitempty"`
	Ds      uint64 `protobuf:"varint,21,opt,name=ds,proto3" json:"ds,omitempty"`
	Es      uint64 `protobuf:"varint,22,opt,name=es,proto3" json:"es,omitempty"`
	Fs      uint64 `protobuf:"varint,23,opt,name=fs,proto3" json:"fs,omitempty"`
	Gs      uint64 `protobuf:"varint,24,opt,name=gs,proto3" json:"gs,omitempty"`
	Ss      uint64 `protobuf:"varint,25,opt,name=ss,proto3" json:"ss,omitempty"`
	FsBase  uint64 `protobuf:"varint,26,opt,name=fs_base,json=fsBase,proto3" json:"fs_base,omitempty"`
	GsBase  uint64 `protobuf:"varint,27,opt,name=gs_base,json=gsBase,proto3" json:"gs_base,omitempty"`
}

func (x *AMD64Registers) Reset() {
	*x = AMD64Registers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_arch_registers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AMD64Registers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AMD64Registers) ProtoMessage() {}

func (x *AMD64Registers) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_arch_registers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AMD64Registers.ProtoReflect.Descriptor instead.
func (*AMD64Registers) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_arch_registers_proto_rawDescGZIP(), []int{0}
}

func (x *AMD64Registers) GetRax() uint64 {
	if x != nil {
		return x.Rax
	}
	return 0
}

func (x *AMD64Registers) GetRbx() uint64 {
	if x != nil {
		return x.Rbx
	}
	return 0
}

func (x *AMD64Registers) GetRcx() uint64 {
	if x != nil {
		return x.Rcx
	}
	return 0
}

func (x *AMD64Registers) GetRdx() uint64 {
	if x != nil {
		return x.Rdx
	}
	return 0
}

func (x *AMD64Registers) GetRsi() uint64 {
	if x != nil {
		return x.Rsi
	}
	return 0
}

func (x *AMD64Registers) GetRdi() uint64 {
	if x != nil {
		return x.Rdi
	}
	return 0
}

func (x *AMD64Registers) GetRsp() uint64 {
	if x != nil {
		return x.Rsp
	}
	return 0
}

func (x *AMD64Registers) GetRbp() uint64 {
	if x != nil {
		return x.Rbp
	}
	return 0
}

func (x *AMD64Registers) GetR8() uint64 {
	if x != nil {
		return x.R8
	}
	return 0
}

func (x *AMD64Registers) GetR9() uint64 {
	if x != nil {
		return x.R9
	}
	return 0
}

func (x *AMD64Registers) GetR10() uint64 {
	if x != nil {
		return x.R10
	}
	return 0
}

func (x *AMD64Registers) GetR11() uint64 {
	if x != nil {
		return x.R11
	}
	return 0
}

func (x *AMD64Registers) GetR12() uint64 {
	if x != nil {
		return x.R12
	}
	return 0
}

func (x *AMD64Registers) GetR13() uint64 {
	if x != nil {
		return x.R13
	}
	return 0
}

func (x *AMD64Registers) GetR14() uint64 {
	if x != nil {
		return x.R14
	}
	return 0
}

func (x *AMD64Registers) GetR15() uint64 {
	if x != nil {
		return x.R15
	}
	return 0
}

func (x *AMD64Registers) GetRip() uint64 {
	if x != nil {
		return x.Rip
	}
	return 0
}

func (x *AMD64Registers) GetRflags() uint64 {
	if x != nil {
		return x.Rflags
	}
	return 0
}

func (x *AMD64Registers) GetOrigRax() uint64 {
	if x != nil {
		return x.OrigRax
	}
	return 0
}

func (x *AMD64Registers) GetCs() uint64 {
	if x != nil {
		return x.Cs
	}
	return 0
}

func (x *AMD64Registers) GetDs() uint64 {
	if x != nil {
		return x.Ds
	}
	return 0
}

func (x *AMD64Registers) GetEs() uint64 {
	if x != nil {
		return x.Es
	}
	return 0
}

func (x *AMD64Registers) GetFs() uint64 {
	if x != nil {
		return x.Fs
	}
	return 0
}

func (x *AMD64Registers) GetGs() uint64 {
	if x != nil {
		return x.Gs
	}
	return 0
}

func (x *AMD64Registers) GetSs() uint64 {
	if x != nil {
		return x.Ss
	}
	return 0
}

func (x *AMD64Registers) GetFsBase() uint64 {
	if x != nil {
		return x.FsBase
	}
	return 0
}

func (x *AMD64Registers) GetGsBase() uint64 {
	if x != nil {
		return x.GsBase
	}
	return 0
}

type ARM64Registers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R0     uint64 `protobuf:"varint,1,opt,name=r0,proto3" json:"r0,omitempty"`
	R1     uint64 `protobuf:"varint,2,opt,name=r1,proto3" json:"r1,omitempty"`
	R2     uint64 `protobuf:"varint,3,opt,name=r2,proto3" json:"r2,omitempty"`
	R3     uint64 `protobuf:"varint,4,opt,name=r3,proto3" json:"r3,omitempty"`
	R4     uint64 `protobuf:"varint,5,opt,name=r4,proto3" json:"r4,omitempty"`
	R5     uint64 `protobuf:"varint,6,opt,name=r5,proto3" json:"r5,omitempty"`
	R6     uint64 `protobuf:"varint,7,opt,name=r6,proto3" json:"r6,omitempty"`
	R7     uint64 `protobuf:"varint,8,opt,name=r7,proto3" json:"r7,omitempty"`
	R8     uint64 `protobuf:"varint,9,opt,name=r8,proto3" json:"r8,omitempty"`
	R9     uint64 `protobuf:"varint,10,opt,name=r9,proto3" json:"r9,omitempty"`
	R10    uint64 `protobuf:"varint,11,opt,name=r10,proto3" json:"r10,omitempty"`
	R11    uint64 `protobuf:"varint,12,opt,name=r11,proto3" json:"r11,omitempty"`
	R12    uint64 `protobuf:"varint,13,opt,name=r12,proto3" json:"r12,omitempty"`
	R13    uint64 `protobuf:"varint,14,opt,name=r13,proto3" json:"r13,omitempty"`
	R14    uint64 `protobuf:"varint,15,opt,name=r14,proto3" json:"r14,omitempty"`
	R15    uint64 `protobuf:"varint,16,opt,name=r15,proto3" json:"r15,omitempty"`
	R16    uint64 `protobuf:"varint,17,opt,name=r16,proto3" json:"r16,omitempty"`
	R17    uint64 `protobuf:"varint,18,opt,name=r17,proto3" json:"r17,omitempty"`
	R18    uint64 `protobuf:"varint,19,opt,name=r18,proto3" json:"r18,omitempty"`
	R19    uint64 `protobuf:"varint,20,opt,name=r19,proto3" json:"r19,omitempty"`
	R20    uint64 `protobuf:"varint,21,opt,name=r20,proto3" json:"r20,omitempty"`
	R21    uint64 `protobuf:"varint,22,opt,name=r21,proto3" json:"r21,omitempty"`
	R22    uint64 `protobuf:"varint,23,opt,name=r22,proto3" json:"r22,omitempty"`
	R23    uint64 `protobuf:"varint,24,opt,name=r23,proto3" json:"r23,omitempty"`
	R24    uint64 `protobuf:"varint,25,opt,name=r24,proto3" json:"r24,omitempty"`
	R25    uint64 `protobuf:"varint,26,opt,name=r25,proto3" json:"r25,omitempty"`
	R26    uint64 `protobuf:"varint,27,opt,name=r26,proto3" json:"r26,omitempty"`
	R27    uint64 `protobuf:"varint,28,opt,name=r27,proto3" json:"r27,omitempty"`
	R28    uint64 `protobuf:"varint,29,opt,name=r28,proto3" json:"r28,omitempty"`
	R29    uint64 `protobuf:"varint,30,opt,name=r29,proto3" json:"r29,omitempty"`
	R30    uint64 `protobuf:"varint,31,opt,name=r30,proto3" json:"r30,omitempty"`
	Sp     uint64 `protobuf:"varint,32,opt,name=sp,proto3" json:"sp,omitempty"`
	Pc     uint64 `protobuf:"varint,33,opt,name=pc,proto3" json:"pc,omitempty"`
	Pstate uint64 `protobuf:"varint,34,opt,name=pstate,proto3" json:"pstate,omitempty"`
	Tls    uint64 `protobuf:"varint,35,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *ARM64Registers) Reset() {
	*x = ARM64Registers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_arch_registers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ARM64Registers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ARM64Registers) ProtoMessage() {}

func (x *ARM64Registers) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_arch_registers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ARM64Registers.ProtoReflect.Descriptor instead.
func (*ARM64Registers) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_arch_registers_proto_rawDescGZIP(), []int{1}
}

func (x *ARM64Registers) GetR0() uint64 {
	if x != nil {
		return x.R0
	}
	return 0
}

func (x *ARM64Registers) GetR1() uint64 {
	if x != nil {
		return x.R1
	}
	return 0
}

func (x *ARM64Registers) GetR2() uint64 {
	if x != nil {
		return x.R2
	}
	return 0
}

func (x *ARM64Registers) GetR3() uint64 {
	if x != nil {
		return x.R3
	}
	return 0
}

func (x *ARM64Registers) GetR4() uint64 {
	if x != nil {
		return x.R4
	}
	return 0
}

func (x *ARM64Registers) GetR5() uint64 {
	if x != nil {
		return x.R5
	}
	return 0
}

func (x *ARM64Registers) GetR6() uint64 {
	if x != nil {
		return x.R6
	}
	return 0
}

func (x *ARM64Registers) GetR7() uint64 {
	if x != nil {
		return x.R7
	}
	return 0
}

func (x *ARM64Registers) GetR8() uint64 {
	if x != nil {
		return x.R8
	}
	return 0
}

func (x *ARM64Registers) GetR9() uint64 {
	if x != nil {
		return x.R9
	}
	return 0
}

func (x *ARM64Registers) GetR10() uint64 {
	if x != nil {
		return x.R10
	}
	return 0
}

func (x *ARM64Registers) GetR11() uint64 {
	if x != nil {
		return x.R11
	}
	return 0
}

func (x *ARM64Registers) GetR12() uint64 {
	if x != nil {
		return x.R12
	}
	return 0
}

func (x *ARM64Registers) GetR13() uint64 {
	if x != nil {
		return x.R13
	}
	return 0
}

func (x *ARM64Registers) GetR14() uint64 {
	if x != nil {
		return x.R14
	}
	return 0
}

func (x *ARM64Registers) GetR15() uint64 {
	if x != nil {
		return x.R15
	}
	return 0
}

func (x *ARM64Registers) GetR16() uint64 {
	if x != nil {
		return x.R16
	}
	return 0
}

func (x *ARM64Registers) GetR17() uint64 {
	if x != nil {
		return x.R17
	}
	return 0
}

func (x *ARM64Registers) GetR18() uint64 {
	if x != nil {
		return x.R18
	}
	return 0
}

func (x *ARM64Registers) GetR19() uint64 {
	if x != nil {
		return x.R19
	}
	return 0
}

func (x *ARM64Registers) GetR20() uint64 {
	if x != nil {
		return x.R20
	}
	return 0
}

func (x *ARM64Registers) GetR21() uint64 {
	if x != nil {
		return x.R21
	}
	return 0
}

func (x *ARM64Registers) GetR22() uint64 {
	if x != nil {
		return x.R22
	}
	return 0
}

func (x *ARM64Registers) GetR23() uint64 {
	if x != nil {
		return x.R23
	}
	return 0
}

func (x *ARM64Registers) GetR24() uint64 {
	if x != nil {
		return x.R24
	}
	return 0
}

func (x *ARM64Registers) GetR25() uint64 {
	if x != nil {
		return x.R25
	}
	return 0
}

func (x *ARM64Registers) GetR26() uint64 {
	if x != nil {
		return x.R26
	}
	return 0
}

func (x *ARM64Registers) GetR27() uint64 {
	if x != nil {
		return x.R27
	}
	return 0
}

func (x *ARM64Registers) GetR28() uint64 {
	if x != nil {
		return x.R28
	}
	return 0
}

func (x *ARM64Registers) GetR29() uint64 {
	if x != nil {
		return x.R29
	}
	return 0
}

func (x *ARM64Registers) GetR30() uint64 {
	if x != nil {
		return x.R30
	}
	return 0
}

func (x *ARM64Registers) GetSp() uint64 {
	if x != nil {
		return x.Sp
	}
	return 0
}

func (x *ARM64Registers) GetPc() uint64 {
	if x != nil {
		return x.Pc
	}
	return 0
}

func (x *ARM64Registers) GetPstate() uint64 {
	if x != nil {
		return x.Pstate
	}
	return 0
}

func (x *ARM64Registers) GetTls() uint64 {
	if x != nil {
		return x.Tls
	}
	return 0
}

type Registers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Arch:
	//	*Registers_Amd64
	//	*Registers_Arm64
	Arch isRegisters_Arch `protobuf_oneof:"arch"`
}

func (x *Registers) Reset() {
	*x = Registers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_arch_registers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registers) ProtoMessage() {}

func (x *Registers) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_arch_registers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registers.ProtoReflect.Descriptor instead.
func (*Registers) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_arch_registers_proto_rawDescGZIP(), []int{2}
}

func (m *Registers) GetArch() isRegisters_Arch {
	if m != nil {
		return m.Arch
	}
	return nil
}

func (x *Registers) GetAmd64() *AMD64Registers {
	if x, ok := x.GetArch().(*Registers_Amd64); ok {
		return x.Amd64
	}
	return nil
}

func (x *Registers) GetArm64() *ARM64Registers {
	if x, ok := x.GetArch().(*Registers_Arm64); ok {
		return x.Arm64
	}
	return nil
}

type isRegisters_Arch interface {
	isRegisters_Arch()
}

type Registers_Amd64 struct {
	Amd64 *AMD64Registers `protobuf:"bytes,1,opt,name=amd64,proto3,oneof"`
}

type Registers_Arm64 struct {
	Arm64 *ARM64Registers `protobuf:"bytes,2,opt,name=arm64,proto3,oneof"`
}

func (*Registers_Amd64) isRegisters_Arch() {}

func (*Registers_Arm64) isRegisters_Arch() {}

var File_pkg_sentry_arch_registers_proto protoreflect.FileDescriptor

var file_pkg_sentry_arch_registers_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x22, 0x83, 0x04, 0x0a, 0x0e, 0x41, 0x4d,
	0x44, 0x36, 0x34, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x61, 0x78, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x62, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x62, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x63, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72,
	0x63, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x64, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x72, 0x64, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x73, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x73, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x64, 0x69, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x64, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x73, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x62,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x62, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x72, 0x38, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x72, 0x38, 0x12, 0x0e, 0x0a, 0x02,
	0x72, 0x39, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x72, 0x39, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x31, 0x30, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x30, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x31, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x31,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72,
	0x31, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x33, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x72, 0x31, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x34, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x31, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x35, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x70, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x5f, 0x72, 0x61, 0x78, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x52, 0x61, 0x78, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x63, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x64, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x64, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x65, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x66, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x66, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x67, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x67, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x73, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x73, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x66, 0x73, 0x42, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x73, 0x42, 0x61, 0x73, 0x65, 0x22,
	0xf4, 0x04, 0x0a, 0x0e, 0x41, 0x52, 0x4d, 0x36, 0x34, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x36, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x37, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x37, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x38, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x39, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x72, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x30, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x72, 0x31, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x32, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x33, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31,
	0x34, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x34, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x31, 0x35, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x35, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x31, 0x36, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x36,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x37, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72,
	0x31, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x38, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x72, 0x31, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x39, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x31, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x30, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x31, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32,
	0x32, 0x18, 0x17, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x32, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x32, 0x33, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x33, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x32, 0x34, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x34,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x35, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72,
	0x32, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x36, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x72, 0x32, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x37, 0x18, 0x1c, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x32, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x38, 0x18, 0x1d, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x32, 0x39, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x32, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x33,
	0x30, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x72, 0x33, 0x30, 0x12, 0x0e, 0x0a, 0x02,
	0x73, 0x70, 0x18, 0x20, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x70, 0x63, 0x18, 0x21, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x23, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x22, 0x73, 0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x6d, 0x64, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x41, 0x4d, 0x44, 0x36,
	0x34, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6d,
	0x64, 0x36, 0x34, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x41, 0x52, 0x4d, 0x36,
	0x34, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x05, 0x61, 0x72,
	0x6d, 0x36, 0x34, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_sentry_arch_registers_proto_rawDescOnce sync.Once
	file_pkg_sentry_arch_registers_proto_rawDescData = file_pkg_sentry_arch_registers_proto_rawDesc
)

func file_pkg_sentry_arch_registers_proto_rawDescGZIP() []byte {
	file_pkg_sentry_arch_registers_proto_rawDescOnce.Do(func() {
		file_pkg_sentry_arch_registers_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sentry_arch_registers_proto_rawDescData)
	})
	return file_pkg_sentry_arch_registers_proto_rawDescData
}

var file_pkg_sentry_arch_registers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_sentry_arch_registers_proto_goTypes = []interface{}{
	(*AMD64Registers)(nil), // 0: gvisor.AMD64Registers
	(*ARM64Registers)(nil), // 1: gvisor.ARM64Registers
	(*Registers)(nil),      // 2: gvisor.Registers
}
var file_pkg_sentry_arch_registers_proto_depIdxs = []int32{
	0, // 0: gvisor.Registers.amd64:type_name -> gvisor.AMD64Registers
	1, // 1: gvisor.Registers.arm64:type_name -> gvisor.ARM64Registers
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_sentry_arch_registers_proto_init() }
func file_pkg_sentry_arch_registers_proto_init() {
	if File_pkg_sentry_arch_registers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sentry_arch_registers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AMD64Registers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sentry_arch_registers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ARM64Registers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_sentry_arch_registers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_sentry_arch_registers_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Registers_Amd64)(nil),
		(*Registers_Arm64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_sentry_arch_registers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sentry_arch_registers_proto_goTypes,
		DependencyIndexes: file_pkg_sentry_arch_registers_proto_depIdxs,
		MessageInfos:      file_pkg_sentry_arch_registers_proto_msgTypes,
	}.Build()
	File_pkg_sentry_arch_registers_proto = out.File
	file_pkg_sentry_arch_registers_proto_rawDesc = nil
	file_pkg_sentry_arch_registers_proto_goTypes = nil
	file_pkg_sentry_arch_registers_proto_depIdxs = nil
}
