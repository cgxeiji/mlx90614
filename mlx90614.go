package mlx90614

import (
	"fmt"

	"github.com/cgxeiji/serial"
)

// Device defines a MLX90614 device.
type Device struct {
	i2c *serial.I2C
}

// New returns a new MLX90614 device. Pass busName="" and addr=0 to use the
// default values.
func New(busName string, addr uint16) (*Device, error) {
	if addr == 0 {
		addr = Addr
	}

	i2c, err := serial.NewI2C(busName, addr)
	if err != nil {
		return nil, fmt.Errorf("mlx90614: could not initialize I2C: %w", err)
	}

	d := &Device{
		i2c: i2c,
	}

	// default options
	d.Options()

	return d, nil
}

// Close closes the device and cleans after itself.
func (d *Device) Close() {
	d.i2c.Close()
}

// TAmbient returns the ambient temperature detected by the sensor.
func (d *Device) TAmbient() (Temperature, error) {
	v, err := d.getT(Tamb)
	if err != nil {
		return 0, fmt.Errorf("mlx90614: could not read ambient temperature: %w", err)
	}

	return v, nil
}

// TObject1 returns the temperature of the object 1 detected by the sensor.
func (d *Device) TObject1() (Temperature, error) {
	v, err := d.getT(Tobj1)
	if err != nil {
		return 0, fmt.Errorf("mlx90614: could not read object 1 temperature: %w", err)
	}

	return v, nil
}

// TObject2 returns the temperature of the object 2 detected by the sensor.
func (d *Device) TObject2() (Temperature, error) {
	v, err := d.getT(Tobj2)
	if err != nil {
		return 0, fmt.Errorf("mlx90614: could not read object 2 temperature: %w", err)
	}

	return v, nil
}

func (d *Device) getT(reg byte) (Temperature, error) {
	value, err := d.ReadBytes(reg, 2)
	if err != nil {
		return 0, err
	}

	v := uint16(value[1])<<8 | uint16(value[0])

	return Temperature(v), nil
}

// Read reads a single byte from a register.
func (d *Device) Read(reg byte) (byte, error) {
	return d.i2c.Read(reg)
}

//ReadBytes reads n bytes from a register.
func (d *Device) ReadBytes(reg byte, n int) ([]byte, error) {
	return d.i2c.ReadBytes(reg, n)
}

// Write writes a byte or bytes to a register.
func (d *Device) Write(reg byte, data ...byte) error {
	return d.i2c.Write(reg, data...)
}
