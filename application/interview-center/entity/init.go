package entity

import (
	"Ai-HireSphere/common/coze"
)

func New() *Entity {
	cozeApi, err := coze.NewCozeApi(coze.Config{
		DatasetID: "7454197247615008820",
		Iss:       "1166857449405",
		Kid:       "zJoGPJI-sKd24NP2phvI9i_SNkNKIb4GByhDd1yo_G4",
		PKeyPath:  "/Users/bkwang/items/Ai-HireSphere/private/private_key.pem",
	})
	if err != nil {
		panic(err)
	}
	return &Entity{
		cozeApi: cozeApi,
	}
}
