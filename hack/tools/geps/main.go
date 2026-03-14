package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// GEPMetadata holds the fields we need from each metadata.yaml.
type GEPMetadata struct {
	Kind   string `yaml:"kind"`
	Number uint   `yaml:"number"`
	Name   string `yaml:"name"`
	Status string `yaml:"status"`
}

// excludedStatuses lists statuses that should not appear in the listing.
var excludedStatuses = map[string]bool{
	"Declined":  true,
	"Deferred":  true,
	"Withdrawn": true,
	"Completed": true,
	"Accepted":  true,
}

// statusOrder defines the display order for status tabs.
var statusOrder = []string{
	"Standard",
	"Memorandum",
	"Experimental",
	"Implementable",
	"Prototyping",
	"Provisional",
}

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	gepsDir := filepath.Join(root, "content", "en", "geps")
	indexPath := filepath.Join(gepsDir, "_index.md")

	// Read existing _index.md to preserve frontmatter.
	existing, err := os.ReadFile(indexPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", indexPath, err)
		os.Exit(1)
	}

	frontmatter, err := extractFrontmatter(string(existing))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing frontmatter: %v\n", err)
		os.Exit(1)
	}

	// Collect all GEP metadata.
	entries, err := filepath.Glob(filepath.Join(gepsDir, "gep-*", "metadata.yaml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error globbing: %v\n", err)
		os.Exit(1)
	}

	var geps []GEPMetadata
	for _, entry := range entries {
		data, err := os.ReadFile(entry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot read %s: %v\n", entry, err)
			continue
		}

		var m GEPMetadata
		if err := yaml.Unmarshal(data, &m); err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot parse %s: %v\n", entry, err)
			continue
		}

		if m.Kind != "GEPDetails" {
			continue
		}
		if m.Number == 0 || m.Name == "" || m.Status == "" {
			fmt.Fprintf(os.Stderr, "warning: skipping %s: missing required fields (number=%d, name=%q, status=%q)\n", entry, m.Number, m.Name, m.Status)
			continue
		}
		if excludedStatuses[m.Status] {
			continue
		}

		geps = append(geps, m)
	}

	// Group by status.
	grouped := make(map[string][]GEPMetadata)
	for _, g := range geps {
		grouped[g.Status] = append(grouped[g.Status], g)
	}

	// Sort each group by number.
	for status := range grouped {
		sort.Slice(grouped[status], func(i, j int) bool {
			return grouped[status][i].Number < grouped[status][j].Number
		})
	}

	// Generate output.
	var buf strings.Builder
	buf.WriteString(frontmatter)
	buf.WriteString("\n\n")
	buf.WriteString("{{< tabs >}}\n")

	for _, status := range statusOrder {
		items, ok := grouped[status]
		if !ok || len(items) == 0 {
			continue
		}

		fmt.Fprintf(&buf, "{{< tab name=%q >}}\n\n", status)
		for _, g := range items {
			fmt.Fprintf(&buf, "- [GEP-%d: %s]({{< ref \"/geps/gep-%d\" >}})\n", g.Number, g.Name, g.Number)
		}
		buf.WriteString("\n{{< /tab >}}\n")
	}

	buf.WriteString("{{< /tabs >}}\n")

	if err := os.WriteFile(indexPath, []byte(buf.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing %s: %v\n", indexPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s with %d GEPs\n", indexPath, len(geps))
}

// extractFrontmatter returns the full frontmatter block including the --- delimiters.
func extractFrontmatter(content string) (string, error) {
	const delimiter = "---"
	if !strings.HasPrefix(content, delimiter) {
		return "", fmt.Errorf("file does not start with frontmatter delimiter")
	}

	end := strings.Index(content[3:], delimiter)
	if end == -1 {
		return "", fmt.Errorf("no closing frontmatter delimiter found")
	}

	// Include both delimiters and the content between them.
	return content[:end+3+len(delimiter)], nil
}
