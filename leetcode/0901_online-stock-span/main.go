package main

type StockInfo struct {
	Price int
	Span  int
}

type StockSpanner struct {
	stack []StockInfo
	top   int
}

func Constructor() StockSpanner {
	return StockSpanner{[]StockInfo{}, -1}
}

func (this *StockSpanner) Next(price int) int {
	curStockInfo := StockInfo{price, 1}

	// pop top and sum
	for this.top >= 0 && this.stack[this.top].Price <= curStockInfo.Price {
		curStockInfo.Span += this.stack[this.top].Span
		this.stack = this.stack[:this.top]
		this.top--
	}

	// push
	this.stack = append(this.stack, curStockInfo)
	this.top++

	// return
	return this.stack[this.top].Span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {

}
