package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Version struct {
	Major, Minor, Patch int
	PreRelease          string
}

func parseVersion(versionString string) Version {
	var v Version
	fmt.Sscanf(versionString, "v%d.%d.%d%s", &v.Major, &v.Minor, &v.Patch, &v.PreRelease)
	return v
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d%s", v.Major, v.Minor, v.Patch, v.PreRelease)
}

func main() {
	var versions []string

	// Read versions from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		versions = append(versions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Custom sorting function
	sort.Slice(versions, func(i, j int) bool {
		vi := parseVersion(versions[i])
		vj := parseVersion(versions[j])

		if vi.Major != vj.Major {
			return vi.Major < vj.Major
		}
		if vi.Minor != vj.Minor {
			return vi.Minor < vj.Minor
		}
		if vi.Patch != vj.Patch {
			return vi.Patch < vj.Patch
		}

		// Compare pre-release versions
		return strings.Compare(vi.PreRelease, vj.PreRelease) < 0
	})

	// Print sorted versions
	for _, version := range versions {
		fmt.Println(version)
	}
}

