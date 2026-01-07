package entity
import(
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)
func EmployeesValidTest (t *testing.T){
	g := gomega.NewGomegaWithT(t)

	t.Run(`Employees is valid`,func(t *testing.T) {
		employees := Employees{
			Name: "Waraphon",
			Salary: 20000,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())

	})
}