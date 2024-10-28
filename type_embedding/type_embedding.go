package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) avgBand() int {
	sum := 0
	for _, element := range b.amount {
		sum = sum + int(element)
	}

	return sum / len(b.amount)
}

func (c *CpuTemp) avgCpu() int {
	sum := 0
	for _, element := range c.temp {
		sum = sum + int(element)
	}

	return sum / len(c.temp)
}

func (m *MemoryUsage) avgMem() int {
	sum := 0
	for _, element := range m.amount {
		sum = sum + int(element)
	}

	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 70, 72, 66, 65}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 830000, 800000}}

	dash := Dashboard{BandwidthUsage: bandwidth, CpuTemp: temp, MemoryUsage: memory}
	fmt.Printf("Average Bandwidth: %v\n", dash.avgBand())
	fmt.Printf("Average Cpu Temp: %v\n", dash.avgCpu())
	fmt.Printf("Average Memory: %v\n", dash.avgMem())
}
