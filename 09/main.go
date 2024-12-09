package main

import (
	"fmt"
	"os"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <part> <input-file>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		part1()
	case "2":
		part2()
	default:
		fmt.Println("Invalid part")
		os.Exit(1)
	}
}

func part1() {
	diskMap := aoc.Map([]rune(aoc.StringFromFile(os.Args[2])), func(r rune) int64 {
		return aoc.Atoi(string(r))
	})
	aoc.Debug(diskMap)

	var seqs []*repeatingSequence
	var id int64
	for i, count := range diskMap {
		if i%2 == 0 {
			seqs = append(seqs, &repeatingSequence{id: id, count: count})
			id++
		} else {
			if count == 0 {
				continue
			}
			seqs = append(seqs, &repeatingSequence{id: -1, count: count})
		}
	}

	freeIndex := -1
	for freeIndex < len(seqs) {
		// Lol, when running the code for the actual input the `seqsToString` was making it extremely slow. Hence I'm
		// wrapping it in this condition, so it doesn't compute the debug strings for the actual input.
		if aoc.DebugMode {
			aoc.Debug(fmt.Sprintf("%02d", freeIndex), seqsToString(seqs))
		}

		freeIndex = findNextFreeIndex(seqs, freeIndex)
		if freeIndex < 0 {
			aoc.Debug("no free sequence found")
			break
		}

		lastIndex := len(seqs) - 1
		last := seqs[lastIndex]
		if last.id < 0 {
			seqs = seqs[:len(seqs)-1]
			continue
		}

		free := seqs[freeIndex]

		// replace free sequence
		if free.count == last.count {
			seqs[freeIndex] = last
			seqs = seqs[:len(seqs)-1]
			continue
		}

		// split free sequence
		if free.count > last.count {
			free.count -= last.count
			seqs = append(seqs[:freeIndex], append([]*repeatingSequence{last}, seqs[freeIndex:len(seqs)-1]...)...)
			continue
		}

		// split last sequence
		if last.count > free.count {
			newLast := &repeatingSequence{id: last.id, count: last.count - free.count}
			last.count = free.count
			seqs[freeIndex] = last
			seqs[lastIndex] = newLast
			continue
		}
	}

	var checksum int64
	var index int64
	for _, seq := range seqs {
		for j := int64(0); j < seq.count; j++ {
			checksum += index * seq.id
			index++
		}
	}

	fmt.Println(checksum)
}

func part2() {
	// part2_attempt1()
	part2_attempt2()
}

func part2_attempt2() {
	// This time I'm going to not track the empty space as sequences/blocks, but only keep track of the files.
	type file struct {
		index int64
		id    int64
		size  int64
	}

	diskMap := aoc.Map([]rune(aoc.StringFromFile(os.Args[2])), func(r rune) int64 {
		return aoc.Atoi(string(r))
	})

	files := map[int64]*file{} // files by id
	disk := map[int64]*file{}  // files by index
	var id int64
	var index, totalDiskSize int64

	diskAsString := func() string {
		if !aoc.DebugMode {
			return ""
		}

		output := ""
		for i := int64(0); i < totalDiskSize; i++ {
			if f, ok := disk[i]; ok {
				output += fmt.Sprintf("%d", f.id)
			} else {
				output += "."
			}
		}

		return output
	}

	for i, count := range diskMap {
		if i%2 == 0 {
			f := &file{
				index: index,
				id:    id,
				size:  count,
			}

			files[id] = f
			for j := index; j < index+count; j++ {
				disk[j] = f
			}

			id++
		}
		index += count
	}
	totalDiskSize = index

	for currentID := id - 1; currentID >= 0; currentID-- {
		aoc.Debug(fmt.Sprintf("%02d", currentID), diskAsString())

		f := files[currentID]

		lastFile := files[0]
		index := lastFile.size
		for {
			if index >= totalDiskSize || index >= f.index {
				break
			}
			nextFile := disk[index]
			if nextFile == nil {
				index++
				continue
			}

			startIndex := lastFile.index + lastFile.size
			endIndex := nextFile.index

			if f.size <= endIndex-startIndex {
				for i := f.index; i < f.index+f.size; i++ {
					delete(disk, i)
				}
				for i := startIndex; i < startIndex+f.size; i++ {
					disk[i] = f
				}
				f.index = startIndex
				break
			}

			lastFile = nextFile
			index = lastFile.index + lastFile.size
		}
	}

	var checksum int64

	for i := int64(0); i < totalDiskSize; i++ {
		if f, ok := disk[i]; ok {
			checksum += i * f.id
		}
	}

	// TOO HIGH
	fmt.Println(checksum)
}

func part2_attempt1() {
	diskMap := aoc.Map([]rune(aoc.StringFromFile(os.Args[2])), func(r rune) int64 {
		return aoc.Atoi(string(r))
	})
	aoc.Debug(diskMap)

	var seqs []*repeatingSequence
	var id int64
	for i, count := range diskMap {
		if i%2 == 0 {
			seqs = append(seqs, &repeatingSequence{id: id, count: count})
			id++
		} else {
			if count == 0 {
				continue
			}
			seqs = append(seqs, &repeatingSequence{id: -1, count: count})
		}
	}

	id -= 1
	if aoc.DebugMode {
		aoc.Debug(fmt.Sprintf("%02d", id), seqsToString(seqs))
	}
	for id >= 0 {
		seqIndex, seq := seqByID(seqs, id)
		if seq == nil {
			fmt.Println("sequence not found", id)
			os.Exit(1)
		}

		freeIndex := -1
		freeIndex = findNextFreeIndex(seqs, freeIndex)
		for freeIndex >= 0 && seqs[freeIndex].count < seq.count {
			freeIndex = findNextFreeIndex(seqs, freeIndex+1)
		}
		if freeIndex < 0 {
			id--
			continue
		}

		free := seqs[freeIndex]

		// replace free sequence
		if free.count == seq.count {
			aoc.Debug("replacing", id)
			seqs[freeIndex] = seq
			seqs[seqIndex] = free

			// merge free sequences
			if aoc.DebugMode {
				aoc.Debug(fmt.Sprintf("%02d", id), seqsToString(seqs))
			}
			seqs = mergeFreeAround(seqs, free, seqIndex)

			id--
			continue
		}

		// split free sequence
		if free.count > seq.count && freeIndex < seqIndex {
			aoc.Debug("splitting", id)
			free.count -= seq.count
			seqs = append(seqs[:freeIndex], append([]*repeatingSequence{seq}, seqs[freeIndex:]...)...)
			newFree := &repeatingSequence{id: -1, count: seq.count}
			seqs[seqIndex+1] = newFree

			// merge free sequences
			if aoc.DebugMode {
				aoc.Debug(fmt.Sprintf("%02d", id), seqsToString(seqs))
			}
			seqs = mergeFreeAround(seqs, newFree, seqIndex+1)

			id--
			continue
		}

		id--
		continue
	}

	var checksum int64
	var index int64
	for _, seq := range seqs {
		for j := int64(0); j < seq.count; j++ {
			if seq.id >= 0 {
				checksum += index * seq.id
			}
			index++
		}
	}

	// TOO HIGH
	fmt.Println(checksum)
}

type repeatingSequence struct {
	id    int64
	count int64
}

func findNextFreeIndex(seqs []*repeatingSequence, start int) int {
	if start < 0 {
		start = 0
	}
	for i := start; i < len(seqs); i++ {
		if seqs[i].id < 0 {
			return i
		}
	}

	return -1
}

func seqsToString(seqs []*repeatingSequence) string {
	output := ""
	for _, seq := range seqs {
		for i := int64(0); i < seq.count; i++ {
			if seq.id < 0 {
				output += "."
			} else {
				output += fmt.Sprintf("%d", seq.id)
			}
		}
	}
	return output
}

func seqByID(seqs []*repeatingSequence, id int64) (int, *repeatingSequence) {
	for i, seq := range seqs {
		if seq.id == id {
			return i, seq
		}
	}
	return -1, nil
}

func mergeFreeAround(seqs []*repeatingSequence, free *repeatingSequence, index int) []*repeatingSequence {
	startMergeIndex := index
	endMergeIndex := index
	if index > 0 && seqs[index-1].id == -1 {
		startMergeIndex--
	}
	if len(seqs) > index+1 && seqs[index+1].id == -1 {
		endMergeIndex++
	}
	mergedCount := int64(0)
	for i := startMergeIndex; i <= endMergeIndex; i++ {
		if i == index {
			mergedCount += free.count
		} else {
			mergedCount += seqs[i].count
		}
	}
	mergedSeq := &repeatingSequence{id: -1, count: mergedCount}

	aoc.Debug("  merging", seqs[startMergeIndex], "at", startMergeIndex, "until", seqs[endMergeIndex], "at", endMergeIndex, "to", mergedSeq)

	return append(seqs[:startMergeIndex], append([]*repeatingSequence{mergedSeq}, seqs[endMergeIndex+1:]...)...)
}
