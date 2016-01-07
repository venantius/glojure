package lang

/*
	NOTE: This interface doesn't exist in JVM Clojure, but is needed here due
	to a need to check for the presence of custom equality methods.

	~ @venantius
*/

type IEquals interface {
	Equals(o interface{}) bool
}
