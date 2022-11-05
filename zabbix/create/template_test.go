package create

import (
	"testing"
	"zbxdb/zabbix/zbxtypes"
)

func Test_parseTemplateFile(t *testing.T) {
	type args struct {
		file zbxtypes.ZbxFile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Parsing Schema File",
			args:    args{file: zbxtypes.SchemaFile},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseTemplateFile(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("parseTemplateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
