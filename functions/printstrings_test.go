package functions

import (
	"reflect"
	"testing"
)

func TestProcessString(t *testing.T) {

	tests := []struct {
		name string
		args []int
		filename string
		want []string
	}{
		{"h",
		[]int{650, 651 ,652 ,653, 654, 655, 656, 657},
		"standard.txt",
		[]string{"_", "       ", "|", " ", "|", "      ", "|", " ", "|", "|", "__", "    ", "|", "  ", "_", " \\ ", "   ", "|", " ", "|", " ", "|", "|", " ", "|", "|", " ", "|", "                   "},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessString(tt.args,tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessString() = %v, want %v", got, tt.want)
			}
		})
	}
}
