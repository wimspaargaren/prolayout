package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/wimspaargaren/prolayout/internal/analyzer"
	"github.com/wimspaargaren/prolayout/internal/model"
)

func TestPassingAnalysis(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	testdataPath := filepath.Join(wd, "testdata")

	tests := []struct {
		name  string
		input model.Root
	}{
		{
			name: "empty",
			input: model.Root{
				Module: "_" + testdataPath,
			},
		},
		{
			name: "allow all",
			input: model.Root{
				Module: "_" + testdataPath,
				Root: []*model.Dir{
					{
						Name: ".*",
					},
				},
			},
		},
		{
			name: "ensure matchin top to bottom",
			input: model.Root{
				Module: "_" + testdataPath,
				Root: []*model.Dir{
					{
						Name: "bar",
						Dirs: []*model.Dir{
							{
								Name: "baz",
							},
						},
					},
					{
						Name: ".*",
						Dirs: []*model.Dir{
							{
								Name: "other",
							},
						},
					},
				},
			},
		},
		{
			name: "extensive structure",
			input: model.Root{
				Module: "_" + testdataPath,
				Root: []*model.Dir{
					{
						Name: "cmd",
						Files: []*model.File{
							{
								Name: "main.go",
							},
						},
					},
					{
						Name: "bar",
						Files: []*model.File{
							{
								Name: "b.*r.go",
							},
						},
						Dirs: []*model.Dir{
							{
								Name: "^baz$",
								Files: []*model.File{
									{
										Name: "baz",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysistest.Run(t, testdataPath, analyzer.New(tt.input), "./...")
		})
	}

	analysistest.Run(t, testdataPath, analyzer.New(model.Root{
		Module: "_" + testdataPath,
		Root: []*model.Dir{
			{
				Name: "cmd",
				Files: []*model.File{
					{
						Name: "main.go",
					},
				},
			},
			{
				Name: "bar",
				Files: []*model.File{
					{
						Name: "b.*r.go",
					},
				},
				Dirs: []*model.Dir{
					{
						Name: "^baz$",
						Files: []*model.File{
							{
								Name: "baz",
							},
						},
					},
				},
			},
		},
	}), "./...")
}
