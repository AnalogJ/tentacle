package utils_test

import (

"path/filepath"
"testing"

"github.com/analogj/tentacle/pkg/utils"
"github.com/stretchr/testify/require"

)


func TestSimpleCmdExec_Date(t *testing.T) {
	t.Parallel()

	//test
	dateStr, cerr := utils.SimpleCmdExec("date", []string{}, "", nil, true)

	//assert
	require.NoError(t, cerr)
	require.NotEmpty(t, dateStr)
}

func TestSimpleCmdExec_Echo(t *testing.T) {
	t.Parallel()

	//test
	resp, cerr := utils.SimpleCmdExec("echo", []string{"hello", "world"}, "", nil, true)

	//assert
	require.NoError(t, cerr)
	require.Equal(t, "hello world", resp)
}

func TestSimpleCmdExec_Error(t *testing.T) {
	t.Parallel()

	//test
	_, cerr := utils.SimpleCmdExec("/bin/bash", []string{"exit", "1"}, "", nil, true)

	//assert
	require.Error(t, cerr)

}

func TestSimpleCmdExec_ErrorWorkingDirRelative(t *testing.T) {
	t.Parallel()

	//test
	_, cerr := utils.SimpleCmdExec("ls", []string{}, "testdata", nil, true)

	//assert
	require.Error(t, cerr)

}

func TestSimpleCmdExec_WorkingDirAbsolute(t *testing.T) {
	t.Parallel()

	//test
	absPath, aerr := filepath.Abs(".")
	_, cerr := utils.SimpleCmdExec("ls", []string{}, absPath, nil, true)

	//assert
	require.NoError(t, aerr)
	require.NoError(t, cerr)
}