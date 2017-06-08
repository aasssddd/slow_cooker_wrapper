package slowcooker

import "testing"

func TestWrapper(T *testing.T) {
	var i Wrapper = Parameters{
		Parameter{"step 1a", "http://localhost:8080", 10, 1, "GET", "1s", false, false, false, "", 10, nil, "", "", 1, 0, 0, "http://localhost:8086", "influxdb", "", "", ""},
		Parameter{"step 2", "http://localhost:8080", 10, 10, "GET", "1s", false, false, false, "", 100, nil, "", "", 2, 0, 0, "http://localhost:8086", "influxdb", "", "", ""},
		Parameter{"step 1b", "http://localhost:8080", 10, 1, "GET", "1s", false, false, false, "", 10, nil, "", "", 1, 0, 0, "http://localhost:8086", "influxdb", "", "", ""},
	}
	i.Runs()
}
