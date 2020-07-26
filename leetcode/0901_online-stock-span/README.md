# 0901_online-stock-span

维护一个stack, stack内的元素存储两个值: stock price 和 这个位置的stock span

当输入一个数时, 初始化其stock span = 1:

- while stack内有元素 && 栈顶元素的stock price <= 当前stock price:
  - pop, 当前stock price += 被pop元素的stock price
  - 循环结束后push, 返回当前stock price
