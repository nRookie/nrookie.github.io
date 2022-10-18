<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108105838265.png" alt="image-20211108105838265" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108105901671.png" alt="image-20211108105901671" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108105935917.png" alt="image-20211108105935917" style="zoom:50%;" />

They are independet to each other





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110019292.png" alt="image-20211108110019292" style="zoom:50%;" />



We make a special mark call it spawn, a spawn is either a function call or a procedure call.



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110210035.png" alt="image-20211108110210035" style="zoom:50%;" />



There is an dependent between a,b and return statement.





so we need a term, sync



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110259760.png" alt="image-20211108110259760" style="zoom:50%;" />





the sync waits for any spwan that has occurred so far within the same stack frame.

<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110421548.png" alt="image-20211108110421548" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110521799.png" alt="image-20211108110521799" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110546906.png" alt="image-20211108110546906" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110654563.png" alt="image-20211108110654563" style="zoom:50%;" />





even in this situation, the program is still wrong.



the sync is after the a + b, and the two spawn are only guaranteed to be complete at the sync.

 so the value of a and b are might not ready.





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108110919484.png" alt="image-20211108110919484" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108111026537.png" alt="image-20211108111026537" style="zoom:50%;" />



spwan creates a new path, and the path continues.



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108111232133.png" alt="image-20211108111232133" style="zoom:50%;" />

you reached to the sync, and the sync waits for the previous spwan to complete.



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108111545999.png" alt="image-20211108111545999" style="zoom:50%;" />





![image-20211108111755941](/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108111755941.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108112552469.png" alt="image-20211108112552469" style="zoom:50%;" />



if you eliminating the first spawn, you eliminate the concurrency.





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108113025733.png" alt="image-20211108113025733" style="zoom:50%;" />







## basic analysis of work and span



![image-20211108113656016](/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108113656016.png)





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108113748492.png" alt="image-20211108113748492" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108113859131.png" alt="image-20211108113859131" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108113917589.png" alt="image-20211108113917589" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/parallel_computing/image-20211108114048602.png" alt="image-20211108114048602" style="zoom:50%;" />

