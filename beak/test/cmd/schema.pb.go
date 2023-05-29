// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: schema.proto

package cmd

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

type EasingMode int32

const (
	EasingMode_LINEAR            EasingMode = 0
	EasingMode_EASE_IN_QUAD      EasingMode = 1
	EasingMode_EASE_OUT_QUAD     EasingMode = 2
	EasingMode_EASE_IN_OUT_QUAD  EasingMode = 3
	EasingMode_EASE_IN_CUBIC     EasingMode = 4
	EasingMode_EASE_OUT_CUBIC    EasingMode = 5
	EasingMode_EASE_IN_OUT_CUBIC EasingMode = 6
	EasingMode_EASE_IN_QUART     EasingMode = 7
	EasingMode_EASE_OUT_QUART    EasingMode = 8
	EasingMode_EASE_IN_OUT_QUART EasingMode = 9
	EasingMode_EASE_IN_QUINT     EasingMode = 10
	EasingMode_EASE_OUT_QUINT    EasingMode = 11
	EasingMode_EASE_IN_OUT_QUINT EasingMode = 12
	EasingMode_EASE_IN_EXPO      EasingMode = 13
	EasingMode_EASE_OUT_EXPO     EasingMode = 14
	EasingMode_EASE_IN_OUT_EXPO  EasingMode = 15
)

// Enum value maps for EasingMode.
var (
	EasingMode_name = map[int32]string{
		0:  "LINEAR",
		1:  "EASE_IN_QUAD",
		2:  "EASE_OUT_QUAD",
		3:  "EASE_IN_OUT_QUAD",
		4:  "EASE_IN_CUBIC",
		5:  "EASE_OUT_CUBIC",
		6:  "EASE_IN_OUT_CUBIC",
		7:  "EASE_IN_QUART",
		8:  "EASE_OUT_QUART",
		9:  "EASE_IN_OUT_QUART",
		10: "EASE_IN_QUINT",
		11: "EASE_OUT_QUINT",
		12: "EASE_IN_OUT_QUINT",
		13: "EASE_IN_EXPO",
		14: "EASE_OUT_EXPO",
		15: "EASE_IN_OUT_EXPO",
	}
	EasingMode_value = map[string]int32{
		"LINEAR":            0,
		"EASE_IN_QUAD":      1,
		"EASE_OUT_QUAD":     2,
		"EASE_IN_OUT_QUAD":  3,
		"EASE_IN_CUBIC":     4,
		"EASE_OUT_CUBIC":    5,
		"EASE_IN_OUT_CUBIC": 6,
		"EASE_IN_QUART":     7,
		"EASE_OUT_QUART":    8,
		"EASE_IN_OUT_QUART": 9,
		"EASE_IN_QUINT":     10,
		"EASE_OUT_QUINT":    11,
		"EASE_IN_OUT_QUINT": 12,
		"EASE_IN_EXPO":      13,
		"EASE_OUT_EXPO":     14,
		"EASE_IN_OUT_EXPO":  15,
	}
)

func (x EasingMode) Enum() *EasingMode {
	p := new(EasingMode)
	*p = x
	return p
}

func (x EasingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EasingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[0].Descriptor()
}

func (EasingMode) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[0]
}

func (x EasingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EasingMode.Descriptor instead.
func (EasingMode) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

type EventType int32

const (
	EventType_BUTTON_1  EventType = 0
	EventType_BUTTON_2  EventType = 1
	EventType_BUTTON_3  EventType = 2
	EventType_BUTTON_4  EventType = 3
	EventType_BUTTON_5  EventType = 4
	EventType_BUTTON_6  EventType = 5
	EventType_BUTTON_7  EventType = 6
	EventType_BUTTON_8  EventType = 7
	EventType_BUTTON_9  EventType = 8
	EventType_BUTTON_10 EventType = 9
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "BUTTON_1",
		1: "BUTTON_2",
		2: "BUTTON_3",
		3: "BUTTON_4",
		4: "BUTTON_5",
		5: "BUTTON_6",
		6: "BUTTON_7",
		7: "BUTTON_8",
		8: "BUTTON_9",
		9: "BUTTON_10",
	}
	EventType_value = map[string]int32{
		"BUTTON_1":  0,
		"BUTTON_2":  1,
		"BUTTON_3":  2,
		"BUTTON_4":  3,
		"BUTTON_5":  4,
		"BUTTON_6":  5,
		"BUTTON_7":  6,
		"BUTTON_8":  7,
		"BUTTON_9":  8,
		"BUTTON_10": 9,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[1].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[1]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*Packet_Frame
	//	*Packet_WFrame
	//	*Packet_RgbFrame
	//	*Packet_AudioFrame
	//	*Packet_InputEvent
	//	*Packet_Config
	Content isPacket_Content `protobuf_oneof:"content"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (m *Packet) GetContent() isPacket_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Packet) GetFrame() *Frame {
	if x, ok := x.GetContent().(*Packet_Frame); ok {
		return x.Frame
	}
	return nil
}

func (x *Packet) GetWFrame() *WFrame {
	if x, ok := x.GetContent().(*Packet_WFrame); ok {
		return x.WFrame
	}
	return nil
}

func (x *Packet) GetRgbFrame() *RGBFrame {
	if x, ok := x.GetContent().(*Packet_RgbFrame); ok {
		return x.RgbFrame
	}
	return nil
}

func (x *Packet) GetAudioFrame() *AudioFrame {
	if x, ok := x.GetContent().(*Packet_AudioFrame); ok {
		return x.AudioFrame
	}
	return nil
}

func (x *Packet) GetInputEvent() *InputEvent {
	if x, ok := x.GetContent().(*Packet_InputEvent); ok {
		return x.InputEvent
	}
	return nil
}

func (x *Packet) GetConfig() *Config {
	if x, ok := x.GetContent().(*Packet_Config); ok {
		return x.Config
	}
	return nil
}

type isPacket_Content interface {
	isPacket_Content()
}

type Packet_Frame struct {
	// From Apps/UDP to Octopus.Mixer to Firmware
	Frame *Frame `protobuf:"bytes,2,opt,name=frame,proto3,oneof"`
}

type Packet_WFrame struct {
	WFrame *WFrame `protobuf:"bytes,3,opt,name=w_frame,json=wFrame,proto3,oneof"`
}

type Packet_RgbFrame struct {
	RgbFrame *RGBFrame `protobuf:"bytes,4,opt,name=rgb_frame,json=rgbFrame,proto3,oneof"`
}

type Packet_AudioFrame struct {
	// From Apps/UDP to Octopus.Mixer to beak
	AudioFrame *AudioFrame `protobuf:"bytes,5,opt,name=audio_frame,json=audioFrame,proto3,oneof"`
}

type Packet_InputEvent struct {
	// From joystick controller to Octopus to  app/UDP
	InputEvent *InputEvent `protobuf:"bytes,6,opt,name=input_event,json=inputEvent,proto3,oneof"`
}

type Packet_Config struct {
	// Configures firmware, not forwarded by the mixer
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

func (*Packet_Frame) isPacket_Content() {}

func (*Packet_WFrame) isPacket_Content() {}

func (*Packet_RgbFrame) isPacket_Content() {}

func (*Packet_AudioFrame) isPacket_Content() {}

func (*Packet_InputEvent) isPacket_Content() {}

func (*Packet_Config) isPacket_Content() {}

// Frame with one byte per pixel and an RGB palette that defines the colors.
type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`       // Selects pixel colors from the palette. First pixel is top left. One panel after the other.
	Palette []byte `protobuf:"bytes,2,opt,name=palette,proto3" json:"palette,omitempty"` // Series of RGB values. 8bit per color.
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Frame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Frame) GetPalette() []byte {
	if x != nil {
		return x.Palette
	}
	return nil
}

// The same as the frame but with access to the white component in the palette (RGBW).
// Not yet implemented
type WFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`       // Selects pixel colors from the palette. First pixel is top left. One panel after the other.
	Palette []byte `protobuf:"bytes,2,opt,name=palette,proto3" json:"palette,omitempty"` // Series of RGBW values. 8bit per color.
}

func (x *WFrame) Reset() {
	*x = WFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WFrame) ProtoMessage() {}

func (x *WFrame) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WFrame.ProtoReflect.Descriptor instead.
func (*WFrame) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (x *WFrame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WFrame) GetPalette() []byte {
	if x != nil {
		return x.Palette
	}
	return nil
}

// Frame with 3 bytes per pixel (RGB).
// not yet implemented
type RGBFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Series of RGB values. 8bit per color. First pixel is top left. One panel after the other.
}

func (x *RGBFrame) Reset() {
	*x = RGBFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RGBFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RGBFrame) ProtoMessage() {}

func (x *RGBFrame) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RGBFrame.ProtoReflect.Descriptor instead.
func (*RGBFrame) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *RGBFrame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luminance         uint32     `protobuf:"varint,1,opt,name=luminance,proto3" json:"luminance,omitempty"`
	EasingIntervalMs  uint32     `protobuf:"varint,2,opt,name=easing_interval_ms,json=easingIntervalMs,proto3" json:"easing_interval_ms,omitempty"`
	EasingMode        EasingMode `protobuf:"varint,3,opt,name=easing_mode,json=easingMode,proto3,enum=EasingMode" json:"easing_mode,omitempty"`
	ShowTestFrame     bool       `protobuf:"varint,4,opt,name=show_test_frame,json=showTestFrame,proto3" json:"show_test_frame,omitempty"`
	ConfigPhash       uint32     `protobuf:"varint,5,opt,name=config_phash,json=configPhash,proto3" json:"config_phash,omitempty"`
	EnableCalibration bool       `protobuf:"varint,6,opt,name=enable_calibration,json=enableCalibration,proto3" json:"enable_calibration,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4}
}

func (x *Config) GetLuminance() uint32 {
	if x != nil {
		return x.Luminance
	}
	return 0
}

func (x *Config) GetEasingIntervalMs() uint32 {
	if x != nil {
		return x.EasingIntervalMs
	}
	return 0
}

func (x *Config) GetEasingMode() EasingMode {
	if x != nil {
		return x.EasingMode
	}
	return EasingMode_LINEAR
}

func (x *Config) GetShowTestFrame() bool {
	if x != nil {
		return x.ShowTestFrame
	}
	return false
}

func (x *Config) GetConfigPhash() uint32 {
	if x != nil {
		return x.ConfigPhash
	}
	return 0
}

func (x *Config) GetEnableCalibration() bool {
	if x != nil {
		return x.EnableCalibration
	}
	return false
}

// This is not ideal and might change
type InputEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  EventType `protobuf:"varint,1,opt,name=type,proto3,enum=EventType" json:"type,omitempty"`
	Value uint32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"` // 1: pressed, 0: released
}

func (x *InputEvent) Reset() {
	*x = InputEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputEvent) ProtoMessage() {}

func (x *InputEvent) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputEvent.ProtoReflect.Descriptor instead.
func (*InputEvent) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{5}
}

func (x *InputEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_BUTTON_1
}

func (x *InputEvent) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// From Firmware to Octopus, internal use only
type FirmwarePacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*FirmwarePacket_FirmwareInfo
	//	*FirmwarePacket_RemoteLog
	Content isFirmwarePacket_Content `protobuf_oneof:"content"`
}

func (x *FirmwarePacket) Reset() {
	*x = FirmwarePacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwarePacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwarePacket) ProtoMessage() {}

func (x *FirmwarePacket) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwarePacket.ProtoReflect.Descriptor instead.
func (*FirmwarePacket) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{6}
}

func (m *FirmwarePacket) GetContent() isFirmwarePacket_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *FirmwarePacket) GetFirmwareInfo() *FirmwareInfo {
	if x, ok := x.GetContent().(*FirmwarePacket_FirmwareInfo); ok {
		return x.FirmwareInfo
	}
	return nil
}

func (x *FirmwarePacket) GetRemoteLog() *RemoteLog {
	if x, ok := x.GetContent().(*FirmwarePacket_RemoteLog); ok {
		return x.RemoteLog
	}
	return nil
}

type isFirmwarePacket_Content interface {
	isFirmwarePacket_Content()
}

type FirmwarePacket_FirmwareInfo struct {
	FirmwareInfo *FirmwareInfo `protobuf:"bytes,1,opt,name=firmware_info,json=firmwareInfo,proto3,oneof"`
}

type FirmwarePacket_RemoteLog struct {
	RemoteLog *RemoteLog `protobuf:"bytes,2,opt,name=remote_log,json=remoteLog,proto3,oneof"`
}

func (*FirmwarePacket_FirmwareInfo) isFirmwarePacket_Content() {}

func (*FirmwarePacket_RemoteLog) isFirmwarePacket_Content() {}

type FirmwareInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname    string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	BuildTime   string `protobuf:"bytes,2,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	PanelIndex  uint32 `protobuf:"varint,3,opt,name=panel_index,json=panelIndex,proto3" json:"panel_index,omitempty"`
	Fps         uint32 `protobuf:"varint,4,opt,name=fps,proto3" json:"fps,omitempty"`
	ConfigPhash uint32 `protobuf:"varint,5,opt,name=config_phash,json=configPhash,proto3" json:"config_phash,omitempty"`
}

func (x *FirmwareInfo) Reset() {
	*x = FirmwareInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareInfo) ProtoMessage() {}

func (x *FirmwareInfo) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareInfo.ProtoReflect.Descriptor instead.
func (*FirmwareInfo) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{7}
}

func (x *FirmwareInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *FirmwareInfo) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

func (x *FirmwareInfo) GetPanelIndex() uint32 {
	if x != nil {
		return x.PanelIndex
	}
	return 0
}

func (x *FirmwareInfo) GetFps() uint32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *FirmwareInfo) GetConfigPhash() uint32 {
	if x != nil {
		return x.ConfigPhash
	}
	return 0
}

type RemoteLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoteLog) Reset() {
	*x = RemoteLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteLog) ProtoMessage() {}

func (x *RemoteLog) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteLog.ProtoReflect.Descriptor instead.
func (*RemoteLog) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{8}
}

func (x *RemoteLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AudioFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri     string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Channel uint32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AudioFrame) Reset() {
	*x = AudioFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioFrame) ProtoMessage() {}

func (x *AudioFrame) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioFrame.ProtoReflect.Descriptor instead.
func (*AudioFrame) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{9}
}

func (x *AudioFrame) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *AudioFrame) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x6e, 0x61, 0x6e, 0x6f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a,
	0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00,
	0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x77, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x57, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x48, 0x00, 0x52, 0x06, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x72,
	0x67, 0x62, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x52, 0x47, 0x42, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x67, 0x62,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x08,
	0x80, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x07, 0x70, 0x61, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x08, 0xc0,
	0x01, 0x52, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x57, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x08, 0x80, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x08, 0x80, 0x02, 0x52, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x22, 0x26, 0x0a, 0x08, 0x52, 0x47, 0x42, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x92, 0x3f,
	0x03, 0x08, 0x80, 0x0f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xfc, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x65, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d,
	0x73, 0x12, 0x2c, 0x0a, 0x0b, 0x65, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x45, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x0a, 0x65, 0x61, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x65,
	0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x0a, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7e, 0x0a,
	0x0e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x34, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xad, 0x01,
	0x0a, 0x0c, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x05, 0x92, 0x3f, 0x02, 0x70, 0x14, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x70, 0x14, 0x52, 0x09, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2c, 0x0a,
	0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1f, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x92, 0x3f, 0x02,
	0x70, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2a, 0xc8, 0x02, 0x0a, 0x0a, 0x45, 0x61, 0x73, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x51, 0x55, 0x41, 0x44,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x51,
	0x55, 0x41, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e,
	0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x45,
	0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x55, 0x42, 0x49, 0x43, 0x10, 0x04, 0x12, 0x12,
	0x0a, 0x0e, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x55, 0x42, 0x49, 0x43,
	0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x4f, 0x55,
	0x54, 0x5f, 0x43, 0x55, 0x42, 0x49, 0x43, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x41, 0x53,
	0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e,
	0x45, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x10, 0x08,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x51, 0x55, 0x41, 0x52, 0x54, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x41, 0x53, 0x45, 0x5f,
	0x49, 0x4e, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x41,
	0x53, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x0b, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x51, 0x55,
	0x49, 0x4e, 0x54, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e,
	0x5f, 0x45, 0x58, 0x50, 0x4f, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x41, 0x53, 0x45, 0x5f,
	0x4f, 0x55, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x41,
	0x53, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x10, 0x0f,
	0x2a, 0x98, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x31, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x32, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55,
	0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x33, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x54, 0x54,
	0x4f, 0x4e, 0x5f, 0x34, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e,
	0x5f, 0x35, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x36,
	0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x37, 0x10, 0x06,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x38, 0x10, 0x07, 0x12, 0x0c,
	0x0a, 0x08, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x39, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09,
	0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x31, 0x30, 0x10, 0x09, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_schema_proto_goTypes = []interface{}{
	(EasingMode)(0),        // 0: EasingMode
	(EventType)(0),         // 1: EventType
	(*Packet)(nil),         // 2: Packet
	(*Frame)(nil),          // 3: Frame
	(*WFrame)(nil),         // 4: WFrame
	(*RGBFrame)(nil),       // 5: RGBFrame
	(*Config)(nil),         // 6: Config
	(*InputEvent)(nil),     // 7: InputEvent
	(*FirmwarePacket)(nil), // 8: FirmwarePacket
	(*FirmwareInfo)(nil),   // 9: FirmwareInfo
	(*RemoteLog)(nil),      // 10: RemoteLog
	(*AudioFrame)(nil),     // 11: AudioFrame
}
var file_schema_proto_depIdxs = []int32{
	3,  // 0: Packet.frame:type_name -> Frame
	4,  // 1: Packet.w_frame:type_name -> WFrame
	5,  // 2: Packet.rgb_frame:type_name -> RGBFrame
	11, // 3: Packet.audio_frame:type_name -> AudioFrame
	7,  // 4: Packet.input_event:type_name -> InputEvent
	6,  // 5: Packet.config:type_name -> Config
	0,  // 6: Config.easing_mode:type_name -> EasingMode
	1,  // 7: InputEvent.type:type_name -> EventType
	9,  // 8: FirmwarePacket.firmware_info:type_name -> FirmwareInfo
	10, // 9: FirmwarePacket.remote_log:type_name -> RemoteLog
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	file_nanopb_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WFrame); i {
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
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RGBFrame); i {
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
		file_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputEvent); i {
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
		file_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwarePacket); i {
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
		file_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareInfo); i {
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
		file_schema_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteLog); i {
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
		file_schema_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioFrame); i {
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
	file_schema_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Packet_Frame)(nil),
		(*Packet_WFrame)(nil),
		(*Packet_RgbFrame)(nil),
		(*Packet_AudioFrame)(nil),
		(*Packet_InputEvent)(nil),
		(*Packet_Config)(nil),
	}
	file_schema_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*FirmwarePacket_FirmwareInfo)(nil),
		(*FirmwarePacket_RemoteLog)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		EnumInfos:         file_schema_proto_enumTypes,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}
