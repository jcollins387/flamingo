package ldap

import (
	ber "github.com/atredispartners/flamingo/pkg/asn1-ber"
	log "github.com/sirupsen/logrus"
)

// debbuging type
//     - has a Printf method to write the debug output
type debugging bool

// write debug output
func (debug debugging) Printf(format string, args ...interface{}) {
	if debug {
		log.Debugf(format, args...)
	}
}

func (debug debugging) PrintPacket(packet *ber.Packet) {
	if debug {
		ber.PrintPacket(packet)
	}
}
