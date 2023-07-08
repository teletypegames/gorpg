package konstructor

type PlatformTypeMapKey string

type PlatformTypeMap map[PlatformTypeMapKey]PlatformType

type PlatformType struct {
	ID     string
	Render Render
}

type Platform struct {
	ID       string
	Type     PlatformType
	Position Position
}

// Helper, which return with the path of a specific platform type
func GetPlatformTypeImagePath(name PlatformTypeMapKey) string {
	return GetObjectDirectory(PlatformObjectType) + string(name) + ".png"
}
