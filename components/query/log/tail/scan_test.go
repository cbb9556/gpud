package tail

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	query_log_filter "github.com/leptonai/gpud/components/query/log/filter"

	"k8s.io/utils/ptr"
)

func TestScan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	tmpf, err := os.CreateTemp("", "test*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpf.Name())

	content := "line1\nline2\nline3\nline4\nline5\n"
	if _, err := tmpf.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := tmpf.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	t.Logf("wrote %q", tmpf.Name())

	largeTmpf, err := os.CreateTemp("", "large_test*.txt")
	if err != nil {
		t.Fatalf("failed to create large temp file: %v", err)
	}
	defer os.Remove(largeTmpf.Name())

	// Write 1000 lines to the large file
	for i := 1; i <= 1000; i++ {
		if _, err := largeTmpf.WriteString(fmt.Sprintf("line%d\n", i)); err != nil {
			t.Fatalf("failed to write to large temp file: %v", err)
		}
	}
	if err := largeTmpf.Close(); err != nil {
		t.Fatalf("failed to close large temp file: %v", err)
	}

	tests := []struct {
		name          string
		fileName      string
		commandArgs   []string
		n             int
		selectFilters []*query_log_filter.Filter
		want          []string
		wantError     bool
	}{
		{
			name:     "tail 3 lines",
			fileName: tmpf.Name(),
			n:        3,
			want:     []string{"line5", "line4", "line3"},
		},
		{
			name:     "tail more lines than file contains",
			fileName: tmpf.Name(),
			n:        10,
			want:     []string{"line5", "line4", "line3", "line2", "line1"},
		},
		{
			name:     "tail with match function",
			fileName: tmpf.Name(),
			n:        3,
			selectFilters: []*query_log_filter.Filter{
				{Regex: ptr.To("3")},
				{Regex: ptr.To("5")},
			},
			want: []string{"line5", "line3"},
		},
		{
			name:     "tail with match function",
			fileName: tmpf.Name(),
			n:        3,
			selectFilters: []*query_log_filter.Filter{
				{Substring: ptr.To("3")},
				{Substring: ptr.To("5")},
			},
			want: []string{"line5", "line3"},
		},
		{
			name:      "non-existent file",
			fileName:  "non-existent_file",
			n:         3,
			wantError: true,
		},

		{
			name:     "tail 100 lines from large file",
			fileName: largeTmpf.Name(),
			n:        100,
			want:     generateExpectedLines(1000, 100),
		},
		{
			name:        "tail 100 lines from large file but with cat",
			commandArgs: []string{"cat", largeTmpf.Name()},
			n:           100,
			want:        generateExpectedLines(1000, 100),
		},

		{
			name:     "tail 1000 lines from large file",
			fileName: largeTmpf.Name(),
			n:        1000,
			want:     generateExpectedLines(1000, 1000),
		},
		{
			name:        "tail 1000 lines from large file but with cat",
			commandArgs: []string{"cat", largeTmpf.Name()},
			n:           1000,
			want:        generateExpectedLines(1000, 1000),
		},

		{
			name:     "tail with regex filter on large file",
			fileName: largeTmpf.Name(),
			n:        1000,
			selectFilters: []*query_log_filter.Filter{
				{Regex: ptr.To("line(50|100|150)")},
			},
			want: []string{"line1000", "line509", "line508", "line507", "line506", "line505", "line504", "line503", "line502", "line501", "line500", "line150", "line100", "line50"},
		},
		{
			name:        "tail with regex filter on large file but with cat",
			commandArgs: []string{"cat", largeTmpf.Name()},
			n:           1000,
			selectFilters: []*query_log_filter.Filter{
				{Regex: ptr.To("line(50|100|150)")},
			},
			want: []string{"line1000", "line509", "line508", "line507", "line506", "line505", "line504", "line503", "line502", "line501", "line500", "line150", "line100", "line50"},
		},

		{
			name:     "tail kubelet.0.log",
			fileName: "testdata/kubelet.0.log",
			n:        5,
			want:     nil, // We'll check the length instead of exact content
		},
		{
			name:        "tail kubelet.0.log but with cat",
			commandArgs: []string{"cat", "testdata/kubelet.0.log"},
			n:           5,
			want:        nil, // We'll check the length instead of exact content
		},

		{
			name:     "tail kubelet.0.log with filter",
			fileName: "testdata/kubelet.0.log",
			n:        1000,
			selectFilters: []*query_log_filter.Filter{
				{Substring: ptr.To("error")},
			},
			want: nil, // We'll check the length instead of exact content
		},
		{
			name:        "tail kubelet.0.log with filter but with cat",
			commandArgs: []string{"cat", "testdata/kubelet.0.log"},
			n:           1000,
			selectFilters: []*query_log_filter.Filter{
				{Substring: ptr.To("error")},
			},
			want: nil, // We'll check the length instead of exact content
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			_, err := Scan(
				ctx,
				WithFile(tt.fileName),
				WithCommands([][]string{tt.commandArgs}),
				WithLinesToTail(tt.n),
				WithSelectFilter(tt.selectFilters...),
				WithParseTime(func(line []byte) (time.Time, error) {
					return time.Time{}, nil
				}),
				WithProcessMatched(func(line []byte, time time.Time, filter *query_log_filter.Filter) {
					got = append(got, string(line))
				}),
			)

			if (err != nil) != tt.wantError {
				t.Errorf("Scan = %v, wantError %v", err, tt.wantError)
				return
			}

			if tt.fileName == "testdata/kubelet.0.log" || strings.Contains(strings.Join(tt.commandArgs, " "), "testdata/kubelet.0.log") {
				// For kubelet.0.log, we'll just check if we got any results
				if len(got) == 0 {
					t.Errorf("Scan on kubelet.0.log returned no results")
				}
				if tt.selectFilters != nil && len(got) == 0 {
					t.Errorf("Scan on kubelet.0.log with filter returned no results")
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateExpectedLines(total, n int) []string {
	var result []string
	for i := total; i > total-n && i > 0; i-- {
		result = append(result, fmt.Sprintf("line%d", i))
	}
	return result
}
