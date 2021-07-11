package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/miekg/dns"
)

func BenchmarkMessageRaw(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var m dns.Msg
		m.RecursionDesired = true
		m.Question = make([]dns.Question, 1)
		question := dns.Question{"", dns.TypeA, dns.ClassINET}
		m.Question[0] = question
		rando := rand.New(rand.NewSource(time.Now().Unix()))
		m.Id = uint16(rando.Uint32())
		// create a new lock free rand source for this goroutine
	}
}

func BenchmarkMsgWrapped(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var m dns.Msg
		m.SetQuestion("www.ucla.edu", dns.TypeA)
	}
}
