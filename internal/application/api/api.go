package api

type Application struct {
	sensor Sensor
}

func NewApplication(sensor Sensor) *Application {
	return &Application{sensor: sensor}
}

func (app Application) GetRead() (float64, error) {
	val, err := app.sensor.Read()
	if err != nil {
		return 0, err
	}
	return val, nil
}
