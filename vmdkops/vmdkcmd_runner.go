// An interface for sending Vmdk Commands to an ESX server.

package vmdkops

// VmdkCmdRunner interface
type VmdkCmdRunner interface {
	Run(cmd string, name string, opts map[string]string) ([]byte, error)
}