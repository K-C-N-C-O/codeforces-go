// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[1,1,3,2]`, `[2,3]`, `[3]`, 
			`[3,2]`,
		},
		{
			`[3,1]`, `[2,3]`, `[1,2]`, 
			`[2,3,1]`,
		},
		{
			`[1,2,2]`, `[4,3,3]`, `[5]`, 
			`[]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, twoOutOfThree, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-262/problems/two-out-of-three/