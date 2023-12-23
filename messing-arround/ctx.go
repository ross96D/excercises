package main

// func Ctx() {
// 	ctx := context.TODO()
// 	doSomething(ctx)
// }

// func doSomething(ctx context.Context) {
// 	ctx, cancelCtx := context.WithCancel(ctx)

// 	printCh := make(chan int)
// 	go doAnother(ctx, printCh)

// 	for num := 1; num <= 3; num++ {
// 		printCh <- num
// 	}

// 	cancelCtx()

// 	time.Sleep(100 * time.Millisecond)

// 	fmt.Printf("doSomething: finished\n")
// }

// func doAnother(ctx context.Context, printCh <-chan int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			if err := ctx.Err(); err != nil {
// 				fmt.Printf("doAnother err: %s\n", err)
// 			}
// 			fmt.Printf("doAnother: finished\n")
// 			return
// 		case num := <-printCh:
// 			fmt.Printf("doAnother: %d\n", num)
// 		}
// 	}
// }
