## Concepts of Flushing
- Flushing means writing data from an in-memory buffer to a more permanent storage medium, such as a disk.
- fsync works on a lower level, it tells the OS to flush its buffers to the physical media. OSs heavily cache data you write to a file. If the OS enforced every write to hit the drive, things would be very slow. fsync (among other things) allows you to control when the data should hit the drive.
-  in-memory buffer is a temporary storage area in RAM that holds data
- atomicity  : All or noone occurs.

## OLAP and OLTP
- OLTP ->  handle large volumes of transactional data (users looking for flight data)
- OLAP -> large vol of data analyzed over broad spectrum.

## in memory vs disk data
for disk data structures-> B+tree and LSM-tree
updating them constantly while taking care of atomicity and durability is difficult.
- Lock Management: `lock` the data so that other multi threaded process xannot touch this.
- Mutex : mutually exclusive -> allowing only one thread of code to use.

## Go coroutine
- done by Go runtime scheduler - mapping multiple goroutines into smaller OS thread called  M scheduling, where M goroutines are multiplexed onto N OS threads.

- Fine-Grained Locking-> locking only parts of file.
- buffer cache (or just cache) is an area of memory used to store frequently accessed data temporarily. 