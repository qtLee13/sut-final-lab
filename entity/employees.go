package entity
import(
	"gorm.io/gorm"
)
type Employees struct { 
	gorm.Model 
	Name     string   // ชื่อต้องยาว 2-80 ตัวอักษร 
	Salary   float64  // เงินเดือนต้องอยู่ในช่วง 15000-200000	 
	EmployeeCode string   // รหัสพนักงานต้องเป็นอักษรอังกฤษตัวใหญ่ 2 ตัว ตามด้วย "-" และตัวเลข 4 ตัว (เช่น "HR-1024") 
} 