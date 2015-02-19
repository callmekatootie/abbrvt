package abbrvt

import "testing"

func TestGet(t *testing.T) {
    result, err := Get("Indian Space Research Organisation")
    if err != nil {
        t.Error("Expected ISRO got ", err)
    } else if result != "ISRO" {
        t.Error("Expected ISRO got ", result)
    }

    result, err = Get("Germany")
    if err != nil {
        t.Error("Expected Grmny got ", err)
    } else if result != "Grmny" {
        t.Error("Expected Grmny got ", result)
    }

    result, err = Get("National Aeronautics and Space Administration")
    if err != nil {
        t.Error("Expected NASA got ", err)
    } else if result != "NASA" {
        t.Error("Expected NASA got ", result)
    }
}
