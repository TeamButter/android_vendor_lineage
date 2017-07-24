# Insert new variables inside the Lineage structure
lineage_soong:
	$(hide) mkdir -p $(dir $@)
	$(hide) (\
	echo '{'; \
	echo '"Lineage": {'; \
	echo '    "Allows_invalid_pthread": $(if $(filter true,$(TARGET_ALLOWS_INVALID_PTHREAD)),true,false),'; \
	echo '    "Egl_needs_handle": $(if $(filter true,$(BOARD_EGL_NEEDS_HANDLE_VALUE)),true,false),'; \
	echo '    "Egl_workaround_bug_10194508": $(if $(filter true,$(BOARD_EGL_WORKAROUND_BUG_10194508)),true,false),'; \
	echo '    "Has_legacy_camera_hal1": $(if $(filter true,$(TARGET_HAS_LEGACY_CAMERA_HAL1)),true,false),'; \
	echo '    "Uses_generic_camera_parameter_library": $(if $(TARGET_SPECIFIC_CAMERA_PARAMETER_LIBRARY),false,true),'; \
	echo '    "Specific_camera_parameter_library": "$(TARGET_SPECIFIC_CAMERA_PARAMETER_LIBRARY)",'; \
	echo '    "Needs_text_relocations": $(if $(filter true,$(TARGET_NEEDS_PLATFORM_TEXT_RELOCATIONS)),true,false),'; \
	echo '    "Uses_media_extensions": $(if $(filter true,$(TARGET_USES_MEDIA_EXTENSIONS)),true,false),'; \
	echo '    "Uses_qcom_bsp_legacy": $(if $(filter true,$(TARGET_USES_QCOM_BSP_LEGACY)),true,false)'; \
	echo '},'; \
	echo '') > $(SOONG_VARIABLES_TMP)
