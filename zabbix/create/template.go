package create

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"zbxdb/zabbix/zbxtypes"
)

func parseTemplateFile(file zbxtypes.ZbxFile) error {
	f, err := os.Open(string(file))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	scanner := bufio.NewScanner(f)

	template := new(zbxtypes.Template)

	for scanner.Scan() {
		switch {
		// skip line on comment
		case strings.HasPrefix(scanner.Text(), "--"):
			continue
		// skip blank lines
		case scanner.Text() == "":
			continue
		case strings.HasPrefix(scanner.Text(), "TABLE"):
			if err := parseTable(template, scanner.Text()); err != nil {
				return err
			}
		case strings.HasPrefix(scanner.Text(), "FIELD"):
			if err := parseField(template, scanner.Text()); err != nil {
				return err
			}
		}
	}

	fmt.Printf("%+v", template)

	return nil
}

func parseTable(template *zbxtypes.Template, table string) error {
	line := parseTemplateLine(table)

	// Quick check to see if the template format has changed for a table.
	if len(line) > 4 {
		return fmt.Errorf("table header has new unknown field")
	}

	template.Sections = append(template.Sections, zbxtypes.Section{
		Table: zbxtypes.Table{
			TableName: line[1],
			TableID:   line[2],
			TableType: line[3],
		},
		Fields:    nil,
		Unique:    zbxtypes.Unique{},
		ChangeLog: 0,
	})

	return nil
}

func parseField(template *zbxtypes.Template, table string) error {
	line := parseTemplateLine(table)

	currentSection := &template.Sections[len(template.Sections)-1]
	currentSection.Fields = append(currentSection.Fields, line)

	return nil
}

func parseTemplateLine(line string) []string {
	return strings.Split(line, "|")
}
