package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// начало решения

type Employee struct {
	XMLName xml.Name `xml:"employee"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	City    string   `xml:"city"`
	Salary  int      `xml:"salary"`
}

type Department struct {
	XMLName   xml.Name    `xml:"department"`
	Code      string      `xml:"code"`
	Employees []*Employee `xml:"employees>employee"`
}

type Organization struct {
	XMLName     xml.Name      `xml:"organization"`
	Departments []*Department `xml:"department"`
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	//decode
	var org Organization

	decoder := xml.NewDecoder(inXML)

	writer := csv.NewWriter(outCSV)

	if err := decoder.Decode(&org); err != nil {
		return err
	}

	// Write CSV header
	header := []string{"id", "name", "city", "department", "salary"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, dept := range org.Departments {
		for _, emp := range dept.Employees {
			record := []string{
				fmt.Sprintf("%d", emp.Id),
				emp.Name,
				emp.City,
				dept.Code,
				fmt.Sprintf("%d", emp.Salary),
			}
			if err := writer.Write(record); err != nil {
				return err
			}
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil { // (2)
		return err
	}

	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
