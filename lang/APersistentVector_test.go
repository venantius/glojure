package lang

import (
	"fmt"
	"testing"
)

func TestAPersistentVectorSatiesfiesInterfaces(t *testing.T) {

	var apersistentvector APersistentVector

	var ipersistentvector IPersistentVector

	var persistentvector PersistentVector

	/*
		var iterable Iterable
		var list List
		var randomaccess RandomAccess
		var comparable Comparable
		var serializable Serializable
		var ihasheq IHashEq
	*/

	// Check that PersistentVector implements APersistentVector
	apersistentvector = &persistentvector

	// Check that APersistentVector implements IPersistentVector
	ipersistentvector = apersistentvector

	if false {
		fmt.Println(
			ipersistentvector,
		)
	}
}
