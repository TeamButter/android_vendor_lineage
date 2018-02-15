package android
type Product_variables struct {
	Allows_invalid_pthread struct {
		Cppflags []string
	}
	Egl_needs_handle struct {
		Cppflags []string
	}
	Egl_workaround_bug_10194508 struct {
		Cppflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_media_extensions struct {
		Cflags []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
}

type ProductVariables struct {
	Allows_invalid_pthread    *bool `json:",omitempty"`
	Egl_needs_handle    *bool `json:",omitempty"`
	Egl_workaround_bug_10194508    *bool `json:",omitempty"`
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Uses_media_extensions  *bool `json:",omitempty"`
	Uses_qcom_bsp_legacy  *bool `json:",omitempty"`
	Target_shim_libs  *string `json:",omitempty"`
}
