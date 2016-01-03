package lang

import (
	"fmt" // test
	"testing"
)

func TestIPersistentVectorSatiesfiesInterfaces(t *testing.T) {

	var ipersistentvector IPersistentVector

	var associative Associative
	var indexed Indexed
	var ipersistentstack IPersistentStack
	var reversible Reversible
	var sequential Sequential

	// Check that IPersistentVector implements Associative
	associative = ipersistentvector

	// Check that IPersistentVector implements Indexed
	indexed = ipersistentvector

	// Check that IPersistentVector implements Reversible
	reversible = ipersistentvector

	// Check that IPersistentVector implements Sequential
	sequential = ipersistentvector

	if false {
		fmt.Println(
			associative,
			ipersistentstack,
			reversible,
			sequential,
			indexed,
		)
	}
}
