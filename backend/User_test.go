package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type User struct{
	Name string `valid:"required~Name cannot be blank"`
	Password string
	Email string `valid:"email~does not validate as email"`
}

func TestUserValidator (t *testing.T){
	g := NewGomegaWithT(t)

	u  := User{
		Name: "natthawat",
		Password: "1q23123",
		Email: "aha",
		}
	
	ok, err := govalidator.ValidateStruct(u)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err.Error()).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("does not validate as email"))

}

func TestName (t *testing.T){
	g  := NewGomegaWithT(t)

	t.Run("Name cannot blank", func (t*testing.T){
		u := User{
			Name: "",
			Password: "123",
			Email: "aha0037@gmail.com",
		}

		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))

	})
}