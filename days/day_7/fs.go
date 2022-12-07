package main

import (
	"errors"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

func NewFile(name string, size int) *File {
	return &File{Name: name, Size: size}
}

type Dir struct {
	Name    string
	Parent  *Dir
	SubDirs []*Dir
	Files   []*File
}

func (d *Dir) Size() int {
	var total int
	d.Traverse(func(c *Dir) error {
		total += c.FlatSize()
		return nil
	})
	return total
}

func (d *Dir) FlatSize() int {
	var total int
	for _, f := range d.Files {
		total += f.Size
	}
	return total
}

func (d *Dir) Traverse(f func(c *Dir) error) error {
	err := f(d)
	if err != nil {
		return err
	}

	for _, subDir := range d.SubDirs {
		err := subDir.Traverse(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		Name:    name,
		Parent:  parent,
		SubDirs: make([]*Dir, 0),
		Files:   make([]*File, 0),
	}
}

// Returns nil if no SubDir can be found
func (d *Dir) FindSubDir(dir string) *Dir {
	for _, subdir := range d.SubDirs {
		if subdir.Name == dir {
			return subdir
		}
	}
	return nil
}

// Returns nil if no file can be found
func (d *Dir) FindFile(file string) *File {
	for _, f := range d.Files {
		if f.Name == file {
			return f
		}
	}
	return nil
}

func Parse(s []string) (*Dir, error) {
	root := NewDir("/", nil)
	currDir := root

	for _, out := range s {
		switch {
		case strings.HasPrefix(out, "$ cd"):
			newDirStr := strings.TrimPrefix(out, "$ cd ")
			if newDirStr == ".." {
				currDir = currDir.Parent
				continue
			}

			if newDirStr == "/" {
				currDir = root
				continue
			}

			prevDir := currDir // Store the curr dir again, as we need to reference it
			currDir = currDir.FindSubDir(newDirStr)
			if currDir == nil {
				prevDir.SubDirs = append(prevDir.SubDirs, NewDir(newDirStr, prevDir))
			}
		case strings.HasPrefix(out, "$ ls"):
			// Do nothing because the default case is an ls
			continue
		default:
			// Assume we just ran an ls
			if strings.HasPrefix(out, "dir ") {
				dirName := strings.TrimPrefix(out, "dir ")
				d := currDir.FindSubDir(dirName)
				if d == nil {
					currDir.SubDirs = append(currDir.SubDirs, NewDir(dirName, currDir))
				}
			} else {
				// Must be a file
				fileSize, err := strconv.Atoi(strings.Fields(out)[0])
				fileName := strings.Fields(out)[1]
				if err != nil {
					return nil, errors.New("invalid file size")
				}

				f := currDir.FindFile(fileName)
				if f == nil {
					currDir.Files = append(currDir.Files, NewFile(fileName, fileSize))
				}
			}
		}
	}
	return root, nil
}
