package transform

type TransformConfig struct {
		Fim struct {
			AppGreylist []string
			FileList    []string
			SyscallList []string
			FieldLists  []struct {
				Pattern string
				Type    string
			}
		}
	}
