# Notes form the Big O Chapter

Big O complexity shows the rate of increase in time/space taken by algorithm in relation with the number of inputs. It describes how an algorithm performs with increasing the input.

1. State both average case and worst case scenarios if they aren't the same.
2. When thinking of Space complexity in a recursive function, consider the functions stored in the call stack as they exist simultaneously.
3. Drop za constants and the non-dominant terms.
4. O(1) < O(logN) < O(N) < O(NlogN) < O(N^2) < O(2^N) < O(N!) 
5. When you see a recursive function making multiple calls, the time complexity is O(branches^depth)
    e.g      (f4)
           /      \
         (f3)     (f3)          => branhes = 2 and depth = N
        /    \    /    \        => Time complexity = O(2^N)
      (f2)  (f2) (f2)  (f2)
     /   \
    (f1) (f1)
6. The base of the logarithm doesn't mater log2N => Log8N => logN
7. The base of the exponent matters 8^N != 2^N