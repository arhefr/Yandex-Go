``` mermaid
graph TB
  subgraph Convert to post notation
  expr[1 * 2 + 3 * 4 + 5 * 6]
  exprPostnote1[1 2 * 3 4 * 5 6 * + +]
  expr --> exprPostnote1
  end
```
``` mermaid
graph TB
exprPostnote2[1 2 * 3 4 * 5 6 * + +]
exprPostnote4[2 42 +]
task1{1 2 *}
task2{3 4 *}
task3{5 6 *}
exprPostnote2 <--> task1
exprPostnote2 <--> task2
exprPostnote2 <--> task3

mes1[Give tasks to agents]
task1 --> mes1
task2 --> mes1
task3 --> mes1

exprPostnote3[2 12 30 + +]
exprPostnote2 ---> exprPostnote3
task4{12 30 +}
exprPostnote3 <--> task4
task4 <--> mes1


exprPostnote3 ---> exprPostnote4
task5{2 42 +}
exprPostnote4 <--> task5
task5 <--> mes1

exprPostnote5((44))
exprPostnote4 ---> exprPostnote5
```
