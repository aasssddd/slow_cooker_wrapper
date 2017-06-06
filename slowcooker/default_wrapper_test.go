package slowcooker

import "testing"

func TestWrapper(T *testing.T) {
	var i Wrapper = Parameters{
		Parameter{"step 1", "http://localhost:8080", 10, 10, "GET", "1s", false, false, false, "", 1000, nil, "", "", 1, 0, 0, "", false, false, "", "", ""},
		Parameter{"step 2", "http://localhost:8080", 10, 10, "GET", "1s", false, false, false, "", 1000, nil, "", "", 2, 0, 0, "", false, false, "", "", ""},
		Parameter{"step 1", "http://localhost:8080", 10, 10, "GET", "1s", false, false, false, "", 1000, nil, "", "", 1, 0, 0, "", false, false, "", "", ""},
	}
	i.Runs()
}
