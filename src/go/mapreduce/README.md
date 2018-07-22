The mapreduce package provides a simple Map/Reduce library (in the mapreduce directory). Applications should normally call Distributed() [located in master.go] to start a job, but may instead call Sequential() [also in master.go] to get a sequential execution for debugging.

The code executes a job as follows:

1. The application provides a number of input files, a map function, a reduce function, and the number of reduce tasks (nReduce).
2. A master is created with this knowledge. It starts an RPC server (see master_rpc.go), and waits for workers to register (using the RPC call Register() [defined in master.go]). As tasks become available (in steps 4 and 5), schedule() [schedule.go] decides how to assign those tasks to workers, and how to handle worker failures.