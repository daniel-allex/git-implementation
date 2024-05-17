package main

import (
	"encoding/hex"
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

func getTreeNodes(content string) []TreeNode {
	var treeNodes []TreeNode

	for len(content) > 0 {
		mode, rem, ok := strings.Cut(content, " ")
		ExceptIfNotOk("Failed to parse tree line while extracting mode", ok)

		name, rem, ok := strings.Cut(rem, "\x00")
		ExceptIfNotOk("Failed to parse tree line while extracting name and sha1", ok)

		sha1 := hex.EncodeToString([]byte(content[:20]))
		content = rem[20:]

		treeNode := TreeNode{Mode: mode, Name: name, Sha1: hex.EncodeToString([]byte(sha1))}
		treeNodes = append(treeNodes, treeNode)
	}

	return treeNodes
}

func outputTreeNames(tree []TreeNode) string {
	names := Map(tree, getTreeName)

	return strings.Join(names, "\n")
}

func outputTreeInfo(tree []TreeNode) string {
	names := Map(tree, getTreeInfo)

	return strings.Join(names, "\n")
}

func treeNodesFromSha1(sha1 string) []TreeNode {
	return getTreeNodes(contentFromSha1(sha1))
}

func outputTreeNamesFromSha1(sha1 string) string {
	return outputTreeNames(treeNodesFromSha1(sha1))
}

func outputTreeInfoFromSha1(sha1 string) string {
	return outputTreeInfo(treeNodesFromSha1(sha1))
}
