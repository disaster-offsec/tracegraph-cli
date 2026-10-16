package bundle

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// writeValidBundle creates a temp dir with all RequiredFiles filled with dummy data.
func writeValidBundle(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	for _, name := range RequiredFiles {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte("dummy"), 0644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}

func TestValidate(t *testing.T) {
	t.Run("valid folder", func(t *testing.T) {
		dir := writeValidBundle(t)
		if err := Validate(dir); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

	t.Run("missing folder", func(t *testing.T) {
		err := Validate("/nonexistent/path/xyz-12345")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("not a directory", func(t *testing.T) {
		tmp := filepath.Join(t.TempDir(), "file.txt")
		if err := os.WriteFile(tmp, []byte("x"), 0644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := Validate(tmp); err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		dir := writeValidBundle(t)
		if err := os.Remove(filepath.Join(dir, "alert.json")); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := Validate(dir); err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("empty file", func(t *testing.T) {
		dir := writeValidBundle(t)
		if err := os.WriteFile(filepath.Join(dir, "alert.json"), []byte{}, 0644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := Validate(dir); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestPack(t *testing.T) {
	data, err := Pack("../../testdata/case_001")
	if err != nil {
		t.Fatalf("Pack failed: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Pack returned empty data")
	}

	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		t.Fatalf("not a valid zip: %v", err)
	}

	if len(zr.File) != len(RequiredFiles) {
		t.Fatalf("expected %d files, got %d", len(RequiredFiles), len(zr.File))
	}

	names := make(map[string]bool)
	for _, f := range zr.File {
		names[f.Name] = true
	}
	for _, want := range RequiredFiles {
		if !names[want] {
			t.Errorf("missing %s in zip", want)
		}
	}
}
