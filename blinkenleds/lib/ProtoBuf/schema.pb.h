/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.7 */

#ifndef PB_SCHEMA_PB_H_INCLUDED
#define PB_SCHEMA_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

/* Enum definitions */
typedef enum _Button {
    Button_BUTTON_1 = 0,
    Button_BUTTON_2 = 1,
    Button_BUTTON_3 = 2,
    Button_BUTTON_4 = 3,
    Button_BUTTON_5 = 4,
    Button_BUTTON_6 = 5,
    Button_BUTTON_7 = 6,
    Button_BUTTON_8 = 7,
    Button_BUTTON_9 = 8,
    Button_BUTTON_10 = 9,
    Button_DIRECTION_1_UP = 10,
    Button_DIRECTION_1_DOWN = 11,
    Button_DIRECTION_1_LEFT = 12,
    Button_DIRECTION_1_RIGHT = 13,
    Button_DIRECTION_2_UP = 14,
    Button_DIRECTION_2_DOWN = 15,
    Button_DIRECTION_2_LEFT = 16,
    Button_DIRECTION_2_RIGHT = 17
} Button;

typedef enum _EasingMode {
    EasingMode_LINEAR = 0,
    EasingMode_EASE_IN_QUAD = 1,
    EasingMode_EASE_OUT_QUAD = 2,
    EasingMode_EASE_IN_OUT_QUAD = 3,
    EasingMode_EASE_IN_CUBIC = 4,
    EasingMode_EASE_OUT_CUBIC = 5,
    EasingMode_EASE_IN_OUT_CUBIC = 6,
    EasingMode_EASE_IN_QUART = 7,
    EasingMode_EASE_OUT_QUART = 8,
    EasingMode_EASE_IN_OUT_QUART = 9,
    EasingMode_EASE_IN_QUINT = 10,
    EasingMode_EASE_OUT_QUINT = 11,
    EasingMode_EASE_IN_OUT_QUINT = 12,
    EasingMode_EASE_IN_EXPO = 13,
    EasingMode_EASE_OUT_EXPO = 14,
    EasingMode_EASE_IN_OUT_EXPO = 15
} EasingMode;

/* Struct definitions */
typedef PB_BYTES_ARRAY_T(640) Frame_data_t;
typedef PB_BYTES_ARRAY_T(192) Frame_palette_t;
/* Frame with one byte per pixel and an RGB palette that defines the colors. */
typedef struct _Frame {
    Frame_data_t data; /* Selects pixel colors from the palette. First pixel is top left. One panel after the other. */
    Frame_palette_t palette; /* Series of RGB values. 8bit per color. */
    uint32_t easing_interval; /* Optional. In milliseconds. Fade this frame over the previous one with an easing curve. Colors are blended in gamma corrected RGB space. */
} Frame;

typedef PB_BYTES_ARRAY_T(640) WFrame_data_t;
typedef PB_BYTES_ARRAY_T(256) WFrame_palette_t;
/* The same as the frame but with access to the white component in the palette (RGBW). */
typedef struct _WFrame {
    WFrame_data_t data; /* Selects pixel colors from the palette. First pixel is top left. One panel after the other. */
    WFrame_palette_t palette; /* Series of RGBW values. 8bit per color. */
    uint32_t easing_interval; /* Optional. In milliseconds. Fade this frame over the previous one with an easing curve. Colors are blended in gamma corrected RGB space. */
} WFrame;

typedef PB_BYTES_ARRAY_T(1920) RGBFrame_data_t;
/* Frame with 3 bytes per pixel (RGB) and no color palette
 not yet implemented */
typedef struct _RGBFrame {
    RGBFrame_data_t data; /* Series of RGB values. 8bit per color. First pixel is top left. One panel after the other. */
    uint32_t easing_interval; /* Optional. In milliseconds. Fade this frame over the previous one with an easing curve. Colors are blended in gamma corrected RGB space. */
} RGBFrame;

typedef struct _InputEvent {
    Button button;
    bool pressed;
} InputEvent;

typedef struct _FirmwareConfig {
    uint32_t luminance;
    EasingMode easing_mode;
    bool show_test_frame;
    uint32_t config_phash;
    bool enable_calibration;
} FirmwareConfig;

/* wrapper for all messages */
typedef struct _Packet {
    pb_size_t which_content;
    union {
        /* Configures firmware, for internal use only */
        FirmwareConfig firmware_config;
        /* Frames with pixel data. */
        Frame frame;
        WFrame w_frame;
        RGBFrame rgb_frame;
        /* Events from the input controllers */
        InputEvent input_event;
    } content;
} Packet;

typedef struct _FirmwareInfo {
    char hostname[21];
    char build_time[21];
    uint32_t panel_index;
    uint32_t fps;
    uint32_t config_phash;
} FirmwareInfo;

typedef struct _RemoteLog {
    char message[101];
} RemoteLog;

/* From Firmware to Octopus, internal use only */
typedef struct _FirmwarePacket {
    pb_size_t which_content;
    union {
        FirmwareInfo firmware_info;
        RemoteLog remote_log;
    } content;
} FirmwarePacket;


#ifdef __cplusplus
extern "C" {
#endif

/* Helper constants for enums */
#define _Button_MIN Button_BUTTON_1
#define _Button_MAX Button_DIRECTION_2_RIGHT
#define _Button_ARRAYSIZE ((Button)(Button_DIRECTION_2_RIGHT+1))

#define _EasingMode_MIN EasingMode_LINEAR
#define _EasingMode_MAX EasingMode_EASE_IN_OUT_EXPO
#define _EasingMode_ARRAYSIZE ((EasingMode)(EasingMode_EASE_IN_OUT_EXPO+1))





#define InputEvent_button_ENUMTYPE Button

#define FirmwareConfig_easing_mode_ENUMTYPE EasingMode





/* Initializer values for message structs */
#define Packet_init_default                      {0, {FirmwareConfig_init_default}}
#define Frame_init_default                       {{0, {0}}, {0, {0}}, 0}
#define WFrame_init_default                      {{0, {0}}, {0, {0}}, 0}
#define RGBFrame_init_default                    {{0, {0}}, 0}
#define InputEvent_init_default                  {_Button_MIN, 0}
#define FirmwareConfig_init_default              {0, _EasingMode_MIN, 0, 0, 0}
#define FirmwarePacket_init_default              {0, {FirmwareInfo_init_default}}
#define FirmwareInfo_init_default                {"", "", 0, 0, 0}
#define RemoteLog_init_default                   {""}
#define Packet_init_zero                         {0, {FirmwareConfig_init_zero}}
#define Frame_init_zero                          {{0, {0}}, {0, {0}}, 0}
#define WFrame_init_zero                         {{0, {0}}, {0, {0}}, 0}
#define RGBFrame_init_zero                       {{0, {0}}, 0}
#define InputEvent_init_zero                     {_Button_MIN, 0}
#define FirmwareConfig_init_zero                 {0, _EasingMode_MIN, 0, 0, 0}
#define FirmwarePacket_init_zero                 {0, {FirmwareInfo_init_zero}}
#define FirmwareInfo_init_zero                   {"", "", 0, 0, 0}
#define RemoteLog_init_zero                      {""}

/* Field tags (for use in manual encoding/decoding) */
#define Frame_data_tag                           1
#define Frame_palette_tag                        2
#define Frame_easing_interval_tag                3
#define WFrame_data_tag                          1
#define WFrame_palette_tag                       2
#define WFrame_easing_interval_tag               3
#define RGBFrame_data_tag                        1
#define RGBFrame_easing_interval_tag             2
#define InputEvent_button_tag                    1
#define InputEvent_pressed_tag                   2
#define FirmwareConfig_luminance_tag             1
#define FirmwareConfig_easing_mode_tag           2
#define FirmwareConfig_show_test_frame_tag       3
#define FirmwareConfig_config_phash_tag          4
#define FirmwareConfig_enable_calibration_tag    5
#define Packet_firmware_config_tag               1
#define Packet_frame_tag                         2
#define Packet_w_frame_tag                       3
#define Packet_rgb_frame_tag                     4
#define Packet_input_event_tag                   6
#define FirmwareInfo_hostname_tag                1
#define FirmwareInfo_build_time_tag              2
#define FirmwareInfo_panel_index_tag             3
#define FirmwareInfo_fps_tag                     4
#define FirmwareInfo_config_phash_tag            5
#define RemoteLog_message_tag                    1
#define FirmwarePacket_firmware_info_tag         1
#define FirmwarePacket_remote_log_tag            2

/* Struct field encoding specification for nanopb */
#define Packet_FIELDLIST(X, a) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,firmware_config,content.firmware_config),   1) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,frame,content.frame),   2) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,w_frame,content.w_frame),   3) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,rgb_frame,content.rgb_frame),   4) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,input_event,content.input_event),   6)
#define Packet_CALLBACK NULL
#define Packet_DEFAULT NULL
#define Packet_content_firmware_config_MSGTYPE FirmwareConfig
#define Packet_content_frame_MSGTYPE Frame
#define Packet_content_w_frame_MSGTYPE WFrame
#define Packet_content_rgb_frame_MSGTYPE RGBFrame
#define Packet_content_input_event_MSGTYPE InputEvent

#define Frame_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, BYTES,    data,              1) \
X(a, STATIC,   SINGULAR, BYTES,    palette,           2) \
X(a, STATIC,   SINGULAR, UINT32,   easing_interval,   3)
#define Frame_CALLBACK NULL
#define Frame_DEFAULT NULL

#define WFrame_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, BYTES,    data,              1) \
X(a, STATIC,   SINGULAR, BYTES,    palette,           2) \
X(a, STATIC,   SINGULAR, UINT32,   easing_interval,   3)
#define WFrame_CALLBACK NULL
#define WFrame_DEFAULT NULL

#define RGBFrame_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, BYTES,    data,              1) \
X(a, STATIC,   SINGULAR, UINT32,   easing_interval,   2)
#define RGBFrame_CALLBACK NULL
#define RGBFrame_DEFAULT NULL

#define InputEvent_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    button,            1) \
X(a, STATIC,   SINGULAR, BOOL,     pressed,           2)
#define InputEvent_CALLBACK NULL
#define InputEvent_DEFAULT NULL

#define FirmwareConfig_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UINT32,   luminance,         1) \
X(a, STATIC,   SINGULAR, UENUM,    easing_mode,       2) \
X(a, STATIC,   SINGULAR, BOOL,     show_test_frame,   3) \
X(a, STATIC,   SINGULAR, UINT32,   config_phash,      4) \
X(a, STATIC,   SINGULAR, BOOL,     enable_calibration,   5)
#define FirmwareConfig_CALLBACK NULL
#define FirmwareConfig_DEFAULT NULL

#define FirmwarePacket_FIELDLIST(X, a) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,firmware_info,content.firmware_info),   1) \
X(a, STATIC,   ONEOF,    MESSAGE,  (content,remote_log,content.remote_log),   2)
#define FirmwarePacket_CALLBACK NULL
#define FirmwarePacket_DEFAULT NULL
#define FirmwarePacket_content_firmware_info_MSGTYPE FirmwareInfo
#define FirmwarePacket_content_remote_log_MSGTYPE RemoteLog

#define FirmwareInfo_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, STRING,   hostname,          1) \
X(a, STATIC,   SINGULAR, STRING,   build_time,        2) \
X(a, STATIC,   SINGULAR, UINT32,   panel_index,       3) \
X(a, STATIC,   SINGULAR, UINT32,   fps,               4) \
X(a, STATIC,   SINGULAR, UINT32,   config_phash,      5)
#define FirmwareInfo_CALLBACK NULL
#define FirmwareInfo_DEFAULT NULL

#define RemoteLog_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, STRING,   message,           1)
#define RemoteLog_CALLBACK NULL
#define RemoteLog_DEFAULT NULL

extern const pb_msgdesc_t Packet_msg;
extern const pb_msgdesc_t Frame_msg;
extern const pb_msgdesc_t WFrame_msg;
extern const pb_msgdesc_t RGBFrame_msg;
extern const pb_msgdesc_t InputEvent_msg;
extern const pb_msgdesc_t FirmwareConfig_msg;
extern const pb_msgdesc_t FirmwarePacket_msg;
extern const pb_msgdesc_t FirmwareInfo_msg;
extern const pb_msgdesc_t RemoteLog_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define Packet_fields &Packet_msg
#define Frame_fields &Frame_msg
#define WFrame_fields &WFrame_msg
#define RGBFrame_fields &RGBFrame_msg
#define InputEvent_fields &InputEvent_msg
#define FirmwareConfig_fields &FirmwareConfig_msg
#define FirmwarePacket_fields &FirmwarePacket_msg
#define FirmwareInfo_fields &FirmwareInfo_msg
#define RemoteLog_fields &RemoteLog_msg

/* Maximum encoded size of messages (where known) */
#define FirmwareConfig_size                      18
#define FirmwareInfo_size                        62
#define FirmwarePacket_size                      104
#define Frame_size                               844
#define InputEvent_size                          4
#define Packet_size                              1932
#define RGBFrame_size                            1929
#define RemoteLog_size                           102
#define WFrame_size                              908

#ifdef __cplusplus
} /* extern "C" */
#endif

#ifdef __cplusplus
/* Message descriptors for nanopb */
namespace nanopb {
template <>
struct MessageDescriptor<Packet> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 5;
    static inline const pb_msgdesc_t* fields() {
        return &Packet_msg;
    }
};
template <>
struct MessageDescriptor<Frame> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 3;
    static inline const pb_msgdesc_t* fields() {
        return &Frame_msg;
    }
};
template <>
struct MessageDescriptor<WFrame> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 3;
    static inline const pb_msgdesc_t* fields() {
        return &WFrame_msg;
    }
};
template <>
struct MessageDescriptor<RGBFrame> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 2;
    static inline const pb_msgdesc_t* fields() {
        return &RGBFrame_msg;
    }
};
template <>
struct MessageDescriptor<InputEvent> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 2;
    static inline const pb_msgdesc_t* fields() {
        return &InputEvent_msg;
    }
};
template <>
struct MessageDescriptor<FirmwareConfig> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 5;
    static inline const pb_msgdesc_t* fields() {
        return &FirmwareConfig_msg;
    }
};
template <>
struct MessageDescriptor<FirmwarePacket> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 2;
    static inline const pb_msgdesc_t* fields() {
        return &FirmwarePacket_msg;
    }
};
template <>
struct MessageDescriptor<FirmwareInfo> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 5;
    static inline const pb_msgdesc_t* fields() {
        return &FirmwareInfo_msg;
    }
};
template <>
struct MessageDescriptor<RemoteLog> {
    static PB_INLINE_CONSTEXPR const pb_size_t fields_array_length = 1;
    static inline const pb_msgdesc_t* fields() {
        return &RemoteLog_msg;
    }
};
}  // namespace nanopb

#endif  /* __cplusplus */


#endif
