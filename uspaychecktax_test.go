package uspaychecktax
import "testing"
func TestFederal(t *testing.T) { if FederalIncomeTax(60000) <= 0 { t.Fail() } }
func TestFica(t *testing.T) { if Fica(50000) <= 0 { t.Fail() } }
