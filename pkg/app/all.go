package app

// 初始化全局APP

func InitAllApp() error {
	for _, app := range Apps {
		if err := app.Config(); err != nil {
			return err
		}
	}
	return nil
}
