package main

import (
    "fmt"
    "sync"
)

type Product struct {
    Name  string
    Stock int
    m     sync.Mutex
}

func (p *Product) Sell(qty int) error {
    p.m.Lock()
    defer p.m.Unlock()
    if qty > p.Stock {
        return fmt.Errorf("недостаточно товара '%s' на складе", p.Name)
    }
    p.Stock -= qty
    return nil
}

func (p *Product) Restock(qty int) {
    p.m.Lock()
    defer p.m.Unlock()
    p.Stock += qty
}

func main() {
    prod := &Product{Name: "Штаны", Stock: 300}
    err := prod.Sell(50)
    if err != nil {
        fmt.Println(err)
    }
    prod.Restock(60)
    fmt.Println(prod.Stock)
}