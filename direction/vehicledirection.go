package direction

type VehicleDirection int64

const (
	VNorth VehicleDirection = iota
	VSouth
	VWest
	VEast

	Unknown = "unknown"
)

func (v VehicleDirection) String() string {
	switch v {
	case VNorth:
		return "N"
	case VSouth:
		return "S"
	case VWest:
		return "W"
	case VEast:
		return "E"
	default:
		return Unknown
	}
}

func (v VehicleDirection) IsValid() bool {
	return v.String() != Unknown
}
