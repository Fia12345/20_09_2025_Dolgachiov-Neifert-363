package main

import (
    "fmt"
    "sync"
)

type MetricsCollector struct {
    successCounter int
    errorCounter   int
    lock           sync.Mutex
}

func NewMetricsCollector() *MetricsCollector {
    return &MetricsCollector{}
}

func (mc *MetricsCollector) IncrementSuccess() {
    mc.lock.Lock()
    defer mc.lock.Unlock()
    mc.successCounter++
}

func (mc *MetricsCollector) IncrementError() {
    mc.lock.Lock()
    defer mc.lock.Unlock()
    mc.errorCounter++
}

func (mc *MetricsCollector) Report() {
    mc.lock.Lock()
    defer mc.lock.Unlock()
    fmt.Printf("Количество успешных запросов: %d\nКоличество ошибок: %d\n", mc.successCounter, mc.errorCounter)
}

func main() {
    collector := NewMetricsCollector()
    collector.IncrementSuccess()
    collector.IncrementError()
    collector.Report()
}