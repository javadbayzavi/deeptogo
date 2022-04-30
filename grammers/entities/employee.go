package entities

type Employee struct {
	id          int
	name        string
	family      string
	contactinfo contact
}
type contact struct {
	email       string
	phone       string
	address     string
	homeaddress int
}

//getters for employee struct
func (emp *Employee) getId() int {
	return emp.id
}
func (emp *Employee) getName() string {
	return emp.name
}
func (emp *Employee) getFamily() string {
	return emp.family
}

func (emp *Employee) getContactInfo() contact {
	return emp.contactinfo
}

//setters for employee struct
func (emp *Employee) setId(i int) {
	emp.id = i
}
func (emp *Employee) setName(nam string) {
	emp.name = nam
}
func (emp *Employee) setFamily(fam string) {
	emp.family = fam
}

func (emp *Employee) setContactInfo(cont contact) {
	emp.contactinfo = cont
}

//getters for contact struct
func (this *contact) getEmail() string {
	return this.email
}

func (this *contact) getPhone() string {
	return this.phone
}
func (this *contact) getAddress() string {
	return this.address
}
func (this *contact) getIsHomeAddress() int {
	return this.homeaddress
}

//setters for contact struct
func (this *contact) setEmail(emi string) {
	this.email = emi
}

func (this *contact) setPhone(pho string) {
	this.phone = pho
}
func (this *contact) setAddress(addres string) {
	this.address = addres
}
func (this *contact) setIsHomeAddress() int {
	return this.homeaddress
}
