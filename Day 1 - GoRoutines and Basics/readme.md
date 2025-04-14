## Concurrency vs. Parallelism
Concurrency and parallelism are often confused, but they are not the same thing. While both refer to the idea of things happening at the same time, they differ in important ways, especially when it comes to writing code and running programs.

Concurrency is about the structure of your code, while parallelism is about how your program is actually executed on the machine.

Concurrency refers to the idea of writing your code in a way that allows multiple tasks to make progress independently. But just because you write concurrent code doesn’t mean it will actually run in parallel (at the same time).

Parallelism happens when multiple tasks actually run at the same time on different processors or cores of the machine.

### Code vs. Runtime:

When we write concurrent code, we're creating opportunities for tasks to run at the same time. But we don’t control whether those tasks run at the same time (parallelism) or one after another (sequentially). That depends on the machine and how it runs the program.

For example, if you run your code on a computer with only one processor, tasks will seem to run at the same time because the computer switches between them quickly. But if the computer has more than one processor (like two or more cores), the tasks can actually run in parallel.

### Understanding the Role of Abstraction:

We don’t always need to worry about whether the tasks are running in parallel or not. The underlying system (the operating system, the machine, etc.) handles that for us. This abstraction allows us to focus on writing concurrent code and leave the details of parallel execution to the runtime.

### Context Matters:

Whether tasks are running concurrently or in parallel depends on the "context" we define. For example:

If the context is time, two tasks might seem to run at the same time if they finish within a specific window (like a second).

If the context is your computer, the tasks are considered concurrent if they don’t affect each other, even if they’re running on separate machines or separate threads.

### Example:
Imagine you and a friend are both doing simple calculations on your own computers. Each of you is running a "calculator process." The calculations don’t affect each other because you’re using separate machines (contexts). Even though your programs might be running at the same time, they are concurrent because they are isolated from each other, not parallel in the strict sense of using multiple processors