package android
type Product_variables struct {
	Egl_needs_handle struct {
		Cppflags []string
	}
	Egl_workaround_bug_10194508 struct {
		Cppflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
}

type ProductVariables struct {
	Egl_needs_handle    *bool `json:",omitempty"`
	Egl_workaround_bug_10194508    *bool `json:",omitempty"`
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Uses_qcom_bsp_legacy  *bool `json:",omitempty"`
}
