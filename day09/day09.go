package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

type Block struct {
	fid int
	free bool
}

func main() {

	var disk []Block
	var fid int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		for i,length := range line {
			l, _ := strconv.Atoi(string(length))
			if i%2 == 0 {  // file size
				for l > 0 {
					disk = append(disk,Block{fid,false})
					l--
				}
				fid++
			} else if i%2 == 1 { // free size
				for l > 0 {
					disk = append(disk,Block{-1,true})
					l--
				}
			}
		}		
	}
	odisk := make([]Block,len(disk))
	copy(odisk, disk)

	// part 1
	fileptr := len(disk)-1
	freeptr := 0
	for {
		fileptr := lastfile(disk, fileptr)
		freeptr := nxtfree(disk, freeptr)
		if fileptr < freeptr {
			break  // we are done
		}
		disk = swapblks(disk, nxtfree(disk, 0), lastfile(disk, len(disk)-1))
	}
	fmt.Println(chksum(disk))

	// part 2
	for fid := odisk[len(odisk)-1].fid; fid > -1; fid-- {
		fblks := fileblks(odisk, fid)
		freep := nxtfreesize(odisk,0, len(fblks))  // look for free segment large enough to hold file
		if freep > -1 {
			// found a space move 'em!
			for i, ifblk := range fblks {
				// only swap to the left
				if ifblk > freep {
					swapblks(odisk, ifblk, freep+i)
				}
			}
		} 
	}
	fmt.Println(chksum(odisk))
}

func chksum(disk []Block) int {
	var sum int
	for i, blk := range disk {
		if !blk.free {
			sum += i*blk.fid
		}
	}
	return sum
}

func swapblks(disk []Block, x, y int) []Block {
	disk[x], disk[y] = disk[y], disk[x]
	return disk
}

func nxtfree(disk []Block, freeptr int) int {
	for !disk[freeptr].free {
		freeptr++
	}
	return freeptr
}

func nxtfreesize(disk []Block, freeptr, fsize  int) int {
	for  {
		if freeptr > len(disk)-1 {
			// no free space large enough on disk
			return -1
		}
		if !disk[freeptr].free {
			//fmt.Println("not free at ", freeptr)
			freeptr++
			continue
		} else if freesize(disk, freeptr) < fsize {
			freeptr += freesize(disk, freeptr)
			continue
		}
		break
	}
	return freeptr
}

func freesize(disk []Block, fp int) int {
	size := 0
	for disk[fp].free {
		size++
		fp++
		if fp == len(disk) {
			break
		}
	}
	return size
}

func lastfile(disk []Block, fileptr int) int {
	for disk[fileptr].free {
		fileptr--
	}
	return fileptr 
}

func fileblks(disk []Block, fid int) []int {
	var ret []int
	for iblk, blk := range disk {
		if blk.fid == fid {
			ret = append(ret, iblk)
		}
	}
	return ret 
}
