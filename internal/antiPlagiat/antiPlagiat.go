package antiplagiat

import (
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var index = make(map[uint32][]Fingerprint)

type Fingerprint struct {
	Hash     uint32
	Position int
	Source   string
}

type Match struct {
	Source   string
	Position int
	Hash     uint32
}

func ReadAndClean(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	clean := strings.ToLower(string(content))
	clean = regexp.MustCompile(`[^\p{L}\p{N}]+`).ReplaceAllString(clean, " ")
	return clean, nil
}

func GenerateFingerprints(text, source string, n, w int) []Fingerprint {
	var hashes []uint32
	for i := 0; i <= len(text)-n; i++ {
		gram := text[i : i+n]
		h := Hash(gram)
		hashes = append(hashes, h)
	}

	var fingerprints []Fingerprint
	var lastHash uint32 = 0
	for i := 0; i <= len(hashes)-w; i++ {
		window := hashes[i : i+w]
		minHash := window[0]
		minPos := i
		for j, h := range window {
			if h <= minHash {
				minHash = h
				minPos = i + j
			}
		}

		if minHash != lastHash {
			fingerprints = append(fingerprints, Fingerprint{
				Hash:     minHash,
				Position: minPos,
				Source:   source,
			})
			lastHash = minHash
		}
	}

	return fingerprints
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func IndexDocument(path string, n, w int) error {
	text, err := ReadAndClean(path)
	if err != nil {
		return err
	}

	fps := GenerateFingerprints(text, filepath.Base(path), n, w)
	for _, fp := range fps {
		index[fp.Hash] = append(index[fp.Hash], fp)
	}
	return nil
}

func CheckPlagiarism(newText string, n, w int) map[string][]Match {
	matches := make(map[string][]Match)

	fingerprints := GenerateFingerprints(newText, "NEW", n, w)
	for _, fp := range fingerprints {
		if entries, ok := index[fp.Hash]; ok {
			for _, e := range entries {
				matches[e.Source] = append(matches[e.Source], Match{
					Source:   e.Source,
					Position: e.Position,
					Hash:     fp.Hash,
				})
			}
		}
	}

	return matches
}

func ReportMatches(matches map[string][]Match) {
	for source, matchList := range matches {
		fmt.Printf("Совпадения с %s: %d\n", source, len(matchList))
		for _, m := range matchList {
			fmt.Printf("  Hash: %v, Позиция: %d\n", m.Hash, m.Position)
		}
	}
}
