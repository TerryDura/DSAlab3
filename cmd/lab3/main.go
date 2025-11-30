/////////////////////////////// THIS IS FOR DEMONSTRATION PURPOSES ONLY //////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"lab3/internal/core"
	"lab3/internal/io"
	"log"
	"strings"
)

func main() {
	mode := flag.String("mode", "C", "")
	probeType := flag.String("probe", "linear", "")
	hashType := flag.String("hash", "my", "")
	inFile := flag.String("in", "Words200D16", "")
	tableFile := flag.String("table", "table.bin", "")
	flag.Parse()

	keys, err := io.LoadKeys(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	keys = keys[:75] // first 75

	f, err := core.Open(*tableFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hashFunc := core.MyHash
	if *hashType == "yours" {
		hashFunc = core.YourHash
	}

	// INSERTION
	for _, key := range keys {
		orig := hashFunc(key, core.TableSize)
		final, probes := core.LinearProbe(orig, core.TableSize, func(idx int) bool {
			slot, err := core.ReadSlot(f, i)
			if err != nil {
				log.Fatal(err)
			}
			return slot.Filled == 0
		})
		if final == -1 {
			log.Fatal("table full")
		}

		var s core.Slot
		copy(s.Key[:], []byte(key))
		s.Orig = int32(orig)
		s.Final = int32(final)
		s.Probes = int32(probes)
		s.Filled = 1
		core.WriteSlot(f, final, s)
	}

	// SEARCHES – first 25
	type stats struct{ min, max, sum, count int }
	first25 := stats{min: 999}
	last25 := stats{min: 999}
	all75 := stats{min: 999}

	search := func(key string) int {
		orig := hashFunc(key, core.TableSize)
		probes := 1
		i := orig
		for {
			slot, err := core.ReadSlot(f, i)
			if err != nil || slot.Filled == 0 {
				log.Fatal("key not found")
			}
			if string(slot.Key[:]) == key {
				return probes
			}
			probes++
			i = (i + 1) % core.TableSize
		}
	}

	for _, key := range keys[:25] {
		p := search(key)
		if p < first25.min {
			first25.min = p
		}
		if p > first25.max {
			first25.max = p
		}
		first25.sum += p
		first25.count++
	}
	// repeat for last 25 and all 75... (same code)

	// TABLE DUMP
	for i := 0; i < core.TableSize; i++ {
		slot, _ := core.ReadSlot(f, i) // or readSlot depending on your func name
		if slot.Filled == 0 {
			fmt.Printf("%2d   (empty)   —   —   —\n", i)
		} else {
			keyStr := strings.ReplaceAll(string(slot.Key[:]), " ", "_")
			fmt.Printf("%2d   %-16s %4d %4d %4d\n", i, keyStr, slot.Orig, slot.Final, slot.Probes)
		}
	}

	// print your stats exactly as required
	fmt.Printf("First-25 min/max/avg: %d/%d/%.2f\n", first25.min, first25.max, float64(first25.sum)/float64(first25.count))
	// etc.
}
