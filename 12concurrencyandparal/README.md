Before starting concurrency or parallelism, ensure you fully understand the difference between them.

to oversimplify it, in concurrency, tasks are running in overlapping time periods. this is because two tasks cannot run at the same time on a single core cpu, so multiple task will have to share the cpu, the cpu will switch from one task to another back and forth, very quickly, giving the illusion that it is doing them at once

in contrast, parallelism happens in multiple core cpu's, here, more that one task are running seperately and at the same time on different cores of the cpu.
