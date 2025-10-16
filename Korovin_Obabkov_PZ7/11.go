package main

import "fmt"

func main() {
	cm := ContactManager{}
	cm.AddContact(Contact{"Иван", []ContactInfo{{"Телефон - ", "88005553535;", "Адрес - ", "Г.Курган ул.Гоголя 130;"}}})
	fmt.Println(cm.FindByName("Иван"))
}

type ContactInfo struct {
	Type   string
	Value  string
	Type1  string
	Value1 string
}

type Contact struct {
	Name  string
	Infos []ContactInfo
}

type ContactManager struct {
	Contacts []Contact
}

func (cm *ContactManager) AddContact(c Contact) {
	cm.Contacts = append(cm.Contacts, c)
}

func (cm *ContactManager) FindByName(name string) *Contact {
	for _, c := range cm.Contacts {
		if c.Name == name {
			return &c
		}
	}
	return nil
}
