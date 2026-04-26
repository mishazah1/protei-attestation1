package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Company struct {
	XMLName   xml.Name   `xml:"company"`   // Корень документа
	Name      string     `xml:"name"`      // Простой элемент <name>
	City      string     `xml:"city,attr"` // Атрибут <company city="...">
	Employees []Employee `xml:"employee"`  // Список элементов <employee>
}

type Employee struct {
	ID    int    `xml:"id,attr"`       // Атрибут <employee id="...">
	Name  string `xml:"name"`          // Элемент <name>
	Role  string `xml:"role"`          // Элемент <role>
	Email string `xml:"contact>email"` // Вложенный элемент <contact><email>...
}

func main() {
	xmlData := `
    <company city="Moscow">
        <name>Protei Lab</name>
        <employee id="101">
            <name>Иван Иванов</name>
            <role>Разработчик</role>
            <contact>
                <email>ivan@example.com</email>
            </contact>
        </employee>
        <employee id="102">
            <name>Петр Петров</name>
            <role>Аналитик</role>
            <contact>
                <email>petr@example.com</email>
            </contact>
        </employee>
    </company>`
	var company Company
	err := xml.Unmarshal([]byte(xmlData), &company)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--- Данные из XML ---")
	fmt.Printf("Компания: %s, Город: %s\n", company.Name, company.City)
	for _, emp := range company.Employees {
		fmt.Printf("Сотрудник [%d]: %s, Роль: %s, Email: %s\n", emp.ID, emp.Name, emp.Role, emp.Email)
	}

	// --- MARSHALING (Struct -> XML) ---
	// Изменим данные, чтобы проверить запись
	company.Name = "Обновленный Протей"
	output, err := xml.MarshalIndent(company, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n--- Сгенерированный XML ---")
	fmt.Println(xml.Header + string(output))
}
