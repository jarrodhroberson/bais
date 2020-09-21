package bais

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func stringToByteArrayPointer(s string) *[]byte {
	r := []byte(s)
	return &r
}

func Test_Encode(t *testing.T) {
	type args struct {
		ba                     *[]byte
		allowControlCharacters bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ABC",
			args: args{
				ba:                     stringToByteArrayPointer("ABC"),
				allowControlCharacters: false,
			},
			want: "ABC",
		},
		{
			name: "Cat\\b`@iE?tEB!CD",
			args: args{
				ba: func() *[]byte {
					content, err := ioutil.ReadFile("../testdata/test.jpg")
					if err != nil {
						t.Errorf("Could not read testdata/test.jpg")
					}
					return &content
				}(),
				allowControlCharacters: false,
			},
			want: "Cat\b`@iE?tEB!CD",
		},
		{
			name: "testdata/test.jpg",
			args: args{
				ba:                     &[]byte{},
				allowControlCharacters: false,
			},
			want: func() string {
				want, err := ioutil.ReadFile("../testdata/test.jpg.bais")
				if err != nil {
					t.Errorf("Could not read testdata/test.jpg.bais")
				}
				return string(want)
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.ba, tt.args.allowControlCharacters); got != tt.want {
				fmt.Println(string(got[:]))
				t.Errorf("ByteArrayInString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "ABC",
			args: args{
				s: "ABC",
			},
			want: []byte("ABC"),
		},
		{
			name: "Cat\\b`@iE?tEB!CD",
			args: args{
				s: "Cat\b`@iE?tEB!CD",
			},
			want: []byte{67, 97, 116, 128, 10, 69, 255, 65, 66, 67, 68},
		},
		{
			name: "test.jpg.bais",
			args: args{
				s: func() string {
					content, err := ioutil.ReadFile("../testdata/test.jpg.bais")
					if err != nil {
						t.Errorf("Could not read testdata/test.jpg.bais")
					}
					return string(content[:])
				}(),
			},
			want: func() []byte {
				want, err := ioutil.ReadFile("../testdata/test.jpg")
				if err != nil {
					t.Errorf("Could not read testdata/test.jpg")
				}
				return want
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Decode(tt.args.s); !bytes.Equal(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
