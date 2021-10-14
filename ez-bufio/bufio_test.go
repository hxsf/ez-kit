package ez_bufio

import (
	"bytes"
	"testing"
)

func TestReader_Seek(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		n      int
		expect byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				str: "0123456789abcdef",
			},
			args: args{
				n:      5,
				expect: '5',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewReaderSize(bytes.NewBufferString(tt.fields.str), 16)
			if err := b.Skip(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Seek() error = %s, wantErr %t", err, tt.wantErr)
			}
			c, _ := b.ReadByte()
			if c != tt.args.expect {
				t.Errorf("Seek() error = expect %b, got %b", tt.args.expect, c)
			}
		})
	}
}
