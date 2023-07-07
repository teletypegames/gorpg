package domain

import "game/konstructor"

const (
	TestPlatformType konstructor.PlatformTypeMapKey = "test"
)

func (d *Domain) InitPlatformType() {
	d.PlatformTypeMap = konstructor.PlatformTypeMap{
		TestPlatformType: konstructor.PlatformType{
			ID: "test_platform_type",
			Render: konstructor.Render{
				Image:  "assets/images/platforms/test_platform.png",
				Width:  30,
				Height: 30,
			},
		},
	}
}

func (d *Domain) GetPlatformTypeMap() konstructor.PlatformTypeMap {
	return d.PlatformTypeMap
}

func (d *Domain) GetPlatformType(name konstructor.PlatformTypeMapKey) konstructor.PlatformType {
	value, _ := d.PlatformTypeMap[name]
	return value
}
