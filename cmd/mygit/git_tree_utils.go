package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type TreeNode struct {
	Mode string
	Name string
	Sha1 string
}

func getTreeName(treeNode TreeNode) string {
	return treeNode.Name
}

func getTreeInfo(treeNode TreeNode) string {
	return treeNode.Mode + " " + objectType(treeNode.Mode) + " " + treeNode.Sha1 + "   " + treeNode.Name
}

func objectType(mode string) string {
	if isDirectory(mode) {
		return "tree"
	} else {
		return "blob"
	}
}

func isDirectory(mode string) bool {
	return mode == "040000"
}

func getTreeNodes(content string) ([]TreeNode, error) {
	var treeNodes []TreeNode

	for len(content) > 0 {
		mode, rem, ok := strings.Cut(content, " ")
		if !ok {
			return []TreeNode{}, errors.New("failed to parse tree line while extracting mode")
		}

		name, rem, ok := strings.Cut(rem, "\x00")
		if !ok {
			return []TreeNode{}, errors.New("failed to parse tree line while extracting name and sha1")
		}

		sha1 := hex.EncodeToString([]byte(content[:20]))
		content = rem[20:]

		treeNode := TreeNode{Mode: mode, Name: name, Sha1: hex.EncodeToString([]byte(sha1))}
		treeNodes = append(treeNodes, treeNode)
	}

	return treeNodes, nil
}

func outputTreeNames(tree []TreeNode) string {
	names := Map(tree, getTreeName)

	return strings.Join(names, "\n")
}

func outputTreeInfo(tree []TreeNode) string {
	names := Map(tree, getTreeInfo)

	return strings.Join(names, "\n")
}

func treeNodesFromSha1(sha1 string) ([]TreeNode, error) {
	content, err := contentFromSha1(sha1)
	if err != nil {
		return []TreeNode{}, fmt.Errorf("failed to get tree nodes from sha1: %w", err)
	}

	nodes, err := getTreeNodes(content)
	if err != nil {
		return []TreeNode{}, fmt.Errorf("failed to get tree nodes from sha1: %w", err)
	}

	return nodes, nil
}

func outputTreeNamesFromSha1(sha1 string) (string, error) {
	treeNodes, err := treeNodesFromSha1(sha1)
	if err != nil {
		return "", fmt.Errorf("failed to get tree nodes from sha1: %w", err)
	}

	return outputTreeNames(treeNodes), nil
}

func outputTreeInfoFromSha1(sha1 string) (string, error) {
	treeNodes, err := treeNodesFromSha1(sha1)
	if err != nil {
		return "", fmt.Errorf("failed to get output tree info from sha1: %w", err)
	}
	return outputTreeInfo(treeNodes), nil
}

func sha1FromEntry(entry os.DirEntry, root string) (string, error) {
	fullPath := root + "/" + entry.Name()
	if entry.IsDir() {
		return WriteTree(fullPath)
	} else {
		return gitBlobSha1FromFile(fullPath)
	}
}

func unixMode(mode os.FileMode) string {
	if mode&os.ModeSymlink != 0 {
		return "120000"
	} else if mode&os.ModeDir != 0 {
		return "40000"
	} else {
		// Regular file
		perm := int64(mode.Perm())
		if perm == 0755 {
			return "100755"
		} else {
			return "100644"
		}
	}
}

func treeNodeFromEntry(entry os.DirEntry, path string) (TreeNode, error) {
	var node TreeNode
	node.Mode = unixMode(entry.Type())
	node.Name = entry.Name()
	sha1, err := sha1FromEntry(entry, path)
	node.Sha1 = sha1

	return node, err
}

func treeNodesFromEntries(entries []os.DirEntry, path string) ([]TreeNode, error) {
	var treeNodes []TreeNode
	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}

		treeNode, err := treeNodeFromEntry(entry, path)
		if err != nil {
			return []TreeNode{}, fmt.Errorf("failed to create tree nodes from entries: %w", err)
		}

		treeNodes = append(treeNodes, treeNode)
	}

	return treeNodes, nil
}

func treeNodesFromPath(path string) ([]TreeNode, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []TreeNode{}, err
	}

	return treeNodesFromEntries(entries, path)
}

func treeNodeSortFunc(treeNodes []TreeNode, i int, j int) bool {
	a := treeNodes[i]
	b := treeNodes[j]

	return a.Name < b.Name
}

func sortTreeNodes(treeNodes []TreeNode) {
	sort.Slice(treeNodes, func(i, j int) bool {
		return treeNodeSortFunc(treeNodes, i, j)
	})
}

func createGitTree(content string) string {
	return createGitObject("tree", content)
}

func writeTreeNodes(treeNodes []TreeNode) (string, error) {
	sortTreeNodes(treeNodes)

	var sb strings.Builder

	for _, node := range treeNodes {
		bytes, err := hex.DecodeString(node.Sha1)
		if err != nil {
			return "", err
		}
		sb.WriteString(node.Mode + " " + node.Name + "\x00" + string(bytes))
	}

	content := sb.String()
	sha1, err := WriteTreeFromContent(content)
	if err != nil {
		return "", err
	}

	return sha1, nil
}

func WriteTreeFromContent(content string) (string, error) {
	tree := createGitTree(content)
	return writeGitObject(tree)
}

func WriteTree(path string) (string, error) {
	treeNodes, err := treeNodesFromPath(path)
	if err != nil {
		return "", fmt.Errorf("failed to write tree: %w", err)
	}

	sha1, err := writeTreeNodes(treeNodes)
	if err != nil {
		return "", fmt.Errorf("failed to write tree: %w", err)
	}

	return sha1, nil
}
