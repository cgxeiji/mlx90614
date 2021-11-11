# mlx90614
[![Version](https://img.shields.io/github/v/tag/cgxeiji/mlx90614?sort=semver)](https://github.com/cgxeiji/mlx90614/releases)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cgxeiji/mlx90614)](https://pkg.go.dev/github.com/cgxeiji/mlx90614)
[![License](https://img.shields.io/github/license/cgxeiji/mlx90614)](https://github.com/cgxeiji/mlx90614/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/cgxeiji/mlx90614)

A Go API to control an MLX90614 temperature sensor using a Raspberry Pi.  Uses
[periph.io](https://periph.io/) to handle I2C communication with the sensor.

## Example

```go
func main() {
    // Create a new connection to the sensor. An empty bus ("") uses the first
    // available I2C bus. A 0 address uses the default address of this sensor
    // (0x5A). You can also use mlx90614.Addr as the address.
    sensor, err := mlx90614.New("", 0)
    if err != nil {
        log.Fatal(err)
    }
    // Don't forget to close the sensor after using it.
    defer sensor.Close()

    // This function returns a type mlx90614.Temperature that can output the
    // value in t.Celsius(), t.Kelvin(), and t.Fahrenheit().
    tA, err := sensor.TAmbient()
    if err != nil {
        log.Fatal(err)
    }

    tO, err := sensor.TObject1()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Ambient temperature: %.2f C", tA.Celsius())
    fmt.Printf("Object temperature: %.2f C", tO.Celsius())
}

```
