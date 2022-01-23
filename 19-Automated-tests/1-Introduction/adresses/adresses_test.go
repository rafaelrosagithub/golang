package adresses

import "testing"

type testScenario struct {
	addressEntered string
	expectedReturn string
}

func TestAddressType(t *testing.T) {
	t.Parallel()

	testScenarios := []testScenario{
		{"Street ABC", "Street"},
		{"Avenue ABC", "Avenue"},
		{"Road ABC", "Road"},
		{"Fountain square", "Invalid type"},
	}

	for _, scenario := range testScenarios {
		receivedAddressType := AddressType(scenario.addressEntered)
		if receivedAddressType == scenario.expectedReturn {
			t.Errorf("Received type is different from expected! Expected %s and received %s",
				receivedAddressType,
				scenario.expectedReturn,
			)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste broke!!!")
	}
}
