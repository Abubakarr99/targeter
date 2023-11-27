package target

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"strings"
)

func ParseTerraformFile(filePath string) (*hcl.File, error) {
	parser := hclparse.NewParser()
	file, parseDiagnostics := parser.ParseHCLFile(filePath)
	if parseDiagnostics.HasErrors() {
		return nil, fmt.Errorf("error parsing HCL: %s", parseDiagnostics.Error())
	}
	return file, nil
}

func ExtractResources(file *hcl.File) ([]hcl.Block, error) {
	var schema = &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
			},
			{
				Type:       "module",
				LabelNames: []string{"name"},
			},
			{
				Type:       "provider",
				LabelNames: []string{"name"},
			},
			{
				Type:       "data",
				LabelNames: []string{"type", "name"},
			},
			{
				Type:       "locals",
				LabelNames: []string{},
			},
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	}
	var resourceBlocks []hcl.Block
	bodyContent, diags := file.Body.Content(schema)
	if diags.HasErrors() {
		return nil, fmt.Errorf("error %s", diags.Error())
	}
	// Iterate through the body blocks
	for _, block := range bodyContent.Blocks {
		// Check if the block is a resource block
		if block.Type == "resource" {
			resourceBlocks = append(resourceBlocks, *block)
		}
	}
	return resourceBlocks, nil
}

func StringOutput(blocks []hcl.Block) string {
	targets := ""
	for _, block := range blocks {
		if len(block.Labels) > 1 {
			targets += fmt.Sprintf("-target=\"%s.%s\" ", block.Labels[0], block.Labels[1])
		}
	}
	return strings.TrimSpace(targets)
}
