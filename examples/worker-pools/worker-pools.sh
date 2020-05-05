# Program kita menunjukkan 5 pekerjaan dilakukan oleh
# berbagai worker-worker. Program hanya membutuhkan 
# sekitar 2 detik meskipun melakukan pekerjaan yang 
# membutuhka waktu  5 detik karena ada 3 worker
# bekerja secara _concurrent_.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
