package entity
import(
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)
func SalaryInvalidTest (t *testing.T){
	g := gomega.NewGomegaWithT(t)

	t.Run(`Employees is valid`,func(t *testing.T) {
		employees := Employees{
			Name: "Waraphon",
			Salary: 10000,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Salary must be between 15000 and 200000"))

	})
}