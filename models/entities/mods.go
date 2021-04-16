package entities

type Drivetrain string

func ToTrackTags(strings []string) []TrackTag{
	var tags []TrackTag
	for _, s := range strings {
		tags = append(tags, TrackTag(s))
	}
	return tags
}

type GearType string

type TrackTag string

type LayoutType string

const (
	F1 TrackTag = "F1"
	Touge TrackTag = "Touge"
	NASCAR  TrackTag = "NASCAR"
	Rally TrackTag = "Rally"
	CityTrack TrackTag = "City Track"
	StreetTrack TrackTag = "Street Track"
	Fictional TrackTag = "Fictional"
	Endurance TrackTag = "Endurance"
	Drift TrackTag = "Drift"
	Historic TrackTag = "Historic"
	OpenWorld TrackTag = "Open World"
)

const (
	RoadCourse LayoutType = "Road Course"
	Oval LayoutType = "Oval"
	AToB LayoutType = "A to B"
)

const (
	RearWheelDrive Drivetrain = "RWD"
	FrontWheelDrive  Drivetrain = "FWD"
	AllWheelDrive Drivetrain = "AWD"
)

const (
	Sequential GearType = "SEQUENTIAL"
	Manual GearType = "MANUAL"
)

type Mod struct {
	DownloadLink string
	Premium bool
	Image string
}

type Car struct {
	Mod
	Brand      CarBrand
	ModelName  string
	Categories []CarCategory
	Year uint
	Drivetrain Drivetrain
	GearType GearType
	BHP uint
	TopSpeed uint
	Weight uint
	Torque uint

}

type CarCategory struct {
	Name string
}

type CarBrand struct {
	Name   string
	Nation Nation
}

type Nation struct {
	Name string
}

type Track struct {
	Mod
	Name     string
	Tags []TrackTag
	Layouts  []Layout
	Location string
	Nation   Nation
	Year uint
}

type Layout struct {
	Name     string
	LengthM  float32
	Category LayoutType
}

type User struct {
	Username string
	Password string
	IsAdmin bool
}