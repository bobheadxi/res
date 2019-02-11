package base

import "testing"

func TestBaseResponse_Error(t *testing.T) {
	type fields struct {
		HTTPStatusCode int
		Message        string
		Err            string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"not an error",
			fields{200, "hi", ""},
			false},
		{"error with only message",
			fields{400, "hi", ""},
			true},
		{"error with message and error context",
			fields{400, "hi", "oh no"},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b = &Response{
				HTTPStatusCode: tt.fields.HTTPStatusCode,
				Message:        tt.fields.Message,
				Err:            tt.fields.Err,
			}
			if err := b.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Response.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
