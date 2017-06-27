package drivers

type DriverType int

const (
  HOST
  OPENEBS
)

var driverTypes = []string{
        "host",
        "openebs",

}

func (dt DriverType) String() string {
        return driverTypes[dt]

