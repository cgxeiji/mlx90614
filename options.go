package mlx90614

// Option defines a functional option for the device.
type Option func(d *Device) (Option, error)

// Options sets different configuration options and returns the previous value
// of the last option passed.
func (d *Device) Options(options ...Option) (last Option, err error) {
	for _, opt := range options {
		last, err = opt(d)
		if err != nil {
			return nil, err
		}
	}

	return last, nil
}
