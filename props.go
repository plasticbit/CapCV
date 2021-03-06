package CapCV

// VidCapProp https://docs.opencv.org/4.5.2/d4/d15/group__videoio__flags__base.html#gaeb8dd9c89c10a5c63c139bf7c4f5704d
type VidCapProp uint8

const (
	CAP_PROP_POS_MSEC VidCapProp = iota
	CAP_PROP_POS_FRAMES
	CAP_PROP_POS_AVI_RATIO
	CAP_PROP_FRAME_WIDTH
	CAP_PROP_FRAME_HEIGHT
	CAP_PROP_FPS
	CAP_PROP_FOURCC
	CAP_PROP_FRAME_COUNT
	CAP_PROP_FORMAT
	CAP_PROP_MODE
	CAP_PROP_BRIGHTNESS
	CAP_PROP_CONTRAST
	CAP_PROP_SATURATION
	CAP_PROP_HUE
	CAP_PROP_GAIN
	CAP_PROP_EXPOSURE
	CAP_PROP_CONVERT_RGB
	CAP_PROP_WHITE_BALANCE_BLUE_U
	CAP_PROP_RECTIFICATION
	CAP_PROP_MONOCHROME
	CAP_PROP_SHARPNESS
	CAP_PROP_AUTO_EXPOSURE
	CAP_PROP_GAMMA
	CAP_PROP_TEMPERATURE
	CAP_PROP_TRIGGER
	CAP_PROP_TRIGGER_DELAY
	CAP_PROP_WHITE_BALANCE_RED_V
	CAP_PROP_ZOOM
	CAP_PROP_FOCUS
	CAP_PROP_GUID
	CAP_PROP_ISO_SPEED
	CAP_PROP_BACKLIGHT VidCapProp = iota + 1
	CAP_PROP_PAN
	CAP_PROP_TILT
	CAP_PROP_ROLL
	CAP_PROP_IRIS
	CAP_PROP_SETTINGS
	CAP_PROP_BUFFERSIZE
	CAP_PROP_AUTOFOCUS
	CAP_PROP_SAR_NUM
	CAP_PROP_SAR_DEN
	CAP_PROP_BACKEND
	CAP_PROP_CHANNEL
	CAP_PROP_AUTO_WB
	CAP_PROP_WB_TEMPERATURE
	CAP_PROP_CODEC_PIXEL_FORMAT
	CAP_PROP_BITRATE
	CAP_PROP_ORIENTATION_META
	CAP_PROP_ORIENTATION_AUTO
	CAP_PROP_HW_ACCELERATION
	CAP_PROP_HW_DEVICE
)
