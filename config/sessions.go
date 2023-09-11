package config

// store 就是前面创建的存储引擎
// store.Options(sessions.Options{
//     Secure:   true,
//     SameSite: 4,
//     Path:     "/",
//     MaxAge:   m.MaxAge,
//   })

type sessions struct {
	Domain   string `yaml:"domain"`
	Secure   bool   `yaml:"secure"`
	SameSite int    `yaml:"sameSite"`
	Path     string `yaml:"path"`
	MaxAge   int    `yaml:"maxAge"`
}
