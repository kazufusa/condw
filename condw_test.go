package condw

import (
	"bytes"
	"io"
	"testing"
)

func Test_CondWriter(t *testing.T) {
	bInfo := new(bytes.Buffer)
	bWarn := new(bytes.Buffer)
	bError := new(bytes.Buffer)

	lw := CondWriter(map[string]io.Writer{
		"[Info]":  bInfo,
		"[Warn]":  bWarn,
		"[Error]": bError,
	})

	var tests = []struct {
		name  string
		given string
		buf   *bytes.Buffer
	}{
		{"[Info]", "[Info] infomation message\n", bInfo},
		{"[Warn]", "[Warn] infomation message\n", bWarn},
		{"[Error]", "[Error] infomation message\n", bError},
		{"[???]", "[???] infomation message\n", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bInfo.Reset()
			bWarn.Reset()
			bError.Reset()

			n, err := lw.Write([]byte(tt.given))

			if tt.buf != nil {
				if err != nil {
					t.Errorf(
						`Output of CondWriter.Write = "%v", want nil.`,
						err,
					)
				}
				if tt.buf.String() != tt.given {
					t.Errorf(
						`Result(err) of CondWriter.Write = "%v", want "%v".`,
						tt.buf.String(),
						tt.given,
					)
				}
			} else {
				if n != len(tt.given) {
					t.Errorf(
						`Result(n) of CondWriter.Write() = "%v", want %v.`,
						n,
						len(tt.given),
					)
				}
				if err != nil {
					t.Errorf(
						`Result(err) of CondWriter.Write() = "%v", want nil.`,
						err,
					)
				}
				if bInfo.String() != "" {
					t.Errorf(
						`Output of CondWriter.Write = "%v", want "".`,
						bInfo.String(),
					)
				}
				if bWarn.String() != "" {
					t.Errorf(
						`Output of CondWriter.Write = "%v", want "".`,
						bWarn.String(),
					)
				}
				if bError.String() != "" {
					t.Errorf(
						`Output of CondWriter.Write = "%v", want "".`,
						bError.String(),
					)
				}
			}
		})
	}
}
